package main

import (
	"context"
	"dagger/blog/internal/dagger"
	"encoding/json"
	"fmt"
	"path"
	"sort"
	"strings"
	"time"
)

const (
	hugoBaseURL  = "https://haydenk.github.io"
	startMarker  = "<!-- ARTICLES:START -->"
	endMarker    = "<!-- ARTICLES:END -->"
	dispatchRepo = "haydenk/haydenk"
)

type frontmatter struct {
	Title string
	Slug  string
	Date  time.Time
	Draft bool
}

type post struct {
	date    time.Time
	display string
	title   string
	url     string
}

func parseFrontmatter(body string) (*frontmatter, bool) {
	const fence = "+++"
	body = strings.TrimLeft(body, "\ufeff \t\r\n")
	if !strings.HasPrefix(body, fence) {
		return nil, false
	}
	rest := strings.TrimPrefix(body, fence)
	rest = strings.TrimLeft(rest, "\r\n")
	end := strings.Index(rest, fence)
	if end == -1 {
		return nil, false
	}
	fm := rest[:end]

	out := &frontmatter{}
	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}
	for _, line := range strings.Split(fm, "\n") {
		eq := strings.Index(line, "=")
		if eq == -1 {
			continue
		}
		key := strings.TrimSpace(line[:eq])
		val := strings.TrimSpace(line[eq+1:])
		val = strings.Trim(val, `"'`)
		switch key {
		case "title":
			out.Title = val
		case "slug":
			out.Slug = val
		case "draft":
			out.Draft = val == "true"
		case "date":
			for _, layout := range layouts {
				if t, err := time.Parse(layout, val); err == nil {
					out.Date = t
					break
				}
			}
		}
	}
	return out, true
}

func buildTable(posts []post) string {
	var sb strings.Builder
	sb.WriteString("| **Date** | **Title** |\n")
	sb.WriteString("|:---------|:----------|\n")
	for _, p := range posts {
		fmt.Fprintf(&sb, "| %s | [%s](%s) |\n", p.display, p.title, p.url)
	}
	return sb.String()
}

// RecentPosts returns a markdown table of the N most recent published posts.
// Walks content/ excluding notes/ and _index.md files.
func (m *Blog) RecentPosts(
	ctx context.Context,
	// +defaultPath="/content"
	content *dagger.Directory,
	// +optional
	// +default=5
	count int,
) (string, error) {
	if count <= 0 {
		count = 5
	}

	entries, err := content.Glob(ctx, "**/*.md")
	if err != nil {
		return "", err
	}

	var posts []post
	for _, entry := range entries {
		if strings.HasPrefix(entry, "notes/") || strings.Contains(entry, "/notes/") {
			continue
		}
		if path.Base(entry) == "_index.md" {
			continue
		}

		body, err := content.File(entry).Contents(ctx)
		if err != nil {
			return "", err
		}
		fm, ok := parseFrontmatter(body)
		if !ok || fm.Draft || fm.Title == "" || fm.Slug == "" || fm.Date.IsZero() {
			continue
		}

		section := strings.SplitN(entry, "/", 2)[0]
		// Match Hugo's permalink layout: /<section>/YYYY/MM/DD/<slug>/
		url := fmt.Sprintf("%s/%s/%04d/%02d/%02d/%s/",
			hugoBaseURL, section,
			fm.Date.Year(), int(fm.Date.Month()), fm.Date.Day(),
			fm.Slug)

		posts = append(posts, post{
			date:    fm.Date,
			display: fm.Date.Format("2006-01-02"),
			title:   strings.ReplaceAll(fm.Title, "|", `\|`),
			url:     url,
		})
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].date.After(posts[j].date) })
	if len(posts) > count {
		posts = posts[:count]
	}

	return buildTable(posts), nil
}

// UpdateReadme splices the recent-posts table into README.md between the
// ARTICLES:START / ARTICLES:END markers and returns the updated README file.
func (m *Blog) UpdateReadme(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=[".dagger", "dist", "public", "resources", ".git", "node_modules", "sermon-notes", "themes", "layouts", "static", "assets", "config"]
	source *dagger.Directory,
) (*dagger.File, error) {
	table, err := m.RecentPosts(ctx, source.Directory("content"), 5)
	if err != nil {
		return nil, err
	}

	readme, err := source.File("README.md").Contents(ctx)
	if err != nil {
		return nil, err
	}

	si := strings.Index(readme, startMarker)
	ei := strings.Index(readme, endMarker)
	if si == -1 || ei == -1 || ei < si {
		return nil, fmt.Errorf("ARTICLES markers not found in README.md")
	}

	updated := readme[:si+len(startMarker)] + "\n" + table + readme[ei:]

	return dag.Directory().
		WithNewFile("README.md", updated).
		File("README.md"), nil
}

// NotifyProfile dispatches a blog-updated event to haydenk/haydenk with the
// current recent-posts table in the client_payload.
func (m *Blog) NotifyProfile(
	ctx context.Context,
	// +defaultPath="/content"
	content *dagger.Directory,
	token *dagger.Secret,
) (string, error) {
	table, err := m.RecentPosts(ctx, content, 5)
	if err != nil {
		return "", err
	}

	payload := map[string]any{
		"event_type": "blog-updated",
		"client_payload": map[string]any{
			"articles_markdown": table,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/dispatches", dispatchRepo)
	payloadFile := dag.Directory().WithNewFile("payload.json", string(body)).File("payload.json")

	script := fmt.Sprintf(`set -eu
code=$(curl -sS -o /tmp/resp -w "%%{http_code}" \
  --max-time 30 --retry 2 --retry-connrefused \
  -X POST \
  -H "Authorization: Bearer $GH_PAT" \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  -d @/payload.json \
  %q)
if [ "$code" != "204" ]; then
  echo "Dispatch to %s failed: HTTP $code" >&2
  cat /tmp/resp >&2
  exit 1
fi
echo "Dispatched blog-updated to %s (HTTP $code)"
`, url, dispatchRepo, dispatchRepo)

	return dag.Container().
		From(curlImage).
		WithSecretVariable("GH_PAT", token).
		WithMountedFile("/payload.json", payloadFile).
		WithExec([]string{"sh", "-c", script}).
		Stdout(ctx)
}
