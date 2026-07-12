// Pipelines for haydenk.github.io.
//
// Build, lint, and validate the Hugo site, and notify the profile repo of
// publishes. Invoked locally via mise wrappers and from GitHub Actions.
package main

import (
	"context"
	"dagger/blog/internal/dagger"

	"golang.org/x/sync/errgroup"
)

const (
	hugoImage         = "hugomods/hugo:0.157.0"
	markdownlintImage = "davidanson/markdownlint-cli2:v0.18.1"
	taploImage        = "tamasfe/taplo:0.10.0"
	htmltestImage     = "wjdp/htmltest:v0.17.0"
	lycheeImage       = "lycheeverse/lychee:sha-b990100-alpine"
	curlImage         = "curlimages/curl:8.16.0"
)

type Blog struct{}

// Build the Hugo site. Returns the generated dist/ directory.
func (m *Blog) Build(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=[".dagger", "dist", "public", "resources", ".git", "node_modules", "sermon-notes"]
	source *dagger.Directory,
	// +optional
	baseURL string,
) *dagger.Directory {
	args := []string{
		"hugo",
		"--gc",
		"--minify",
		"--noBuildLock",
		"--cleanDestinationDir",
		"--buildDrafts=false",
		"--destination", "/dist",
	}
	if baseURL != "" {
		args = append(args, "--baseURL", baseURL)
	}

	return dag.Container().
		From(hugoImage).
		WithEnvVariable("HUGO_ENVIRONMENT", "production").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args).
		Directory("/dist")
}

// Lint runs markdownlint over content and taplo over TOML config.
func (m *Blog) Lint(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=[".dagger", "dist", "public", "resources", ".git", "node_modules", "sermon-notes"]
	source *dagger.Directory,
) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		_, err := dag.Container().
			From(markdownlintImage).
			WithUser("0:0").
			WithMountedDirectory("/workdir", source).
			WithWorkdir("/workdir").
			WithExec(
				[]string{"content/**/*.md", "README.md"},
				dagger.ContainerWithExecOpts{UseEntrypoint: true},
			).
			Sync(ctx)
		return err
	})

	g.Go(func() error {
		_, err := dag.Container().
			From(taploImage).
			WithUser("0:0").
			WithMountedDirectory("/workdir", source).
			WithWorkdir("/workdir").
			WithExec(
				[]string{"check"},
				dagger.ContainerWithExecOpts{UseEntrypoint: true},
			).
			Sync(ctx)
		return err
	})

	return g.Wait()
}

// HtmlTest builds the site and validates the rendered HTML.
func (m *Blog) HtmlTest(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=[".dagger", "dist", "public", "resources", ".git", "node_modules", "sermon-notes"]
	source *dagger.Directory,
) error {
	dist := m.Build(ctx, source, "")

	_, err := dag.Container().
		From(htmltestImage).
		WithMountedDirectory("/test", dist).
		WithMountedFile("/htmltest.yml", source.File(".htmltest.yml")).
		WithWorkdir("/test").
		WithExec(
			[]string{"-c", "/htmltest.yml", "-s"},
			dagger.ContainerWithExecOpts{UseEntrypoint: true},
		).
		Sync(ctx)
	return err
}

// LinkCheck builds the site and runs lychee against the output.
func (m *Blog) LinkCheck(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=[".dagger", "dist", "public", "resources", ".git", "node_modules", "sermon-notes"]
	source *dagger.Directory,
) error {
	dist := m.Build(ctx, source, "")

	_, err := dag.Container().
		From(lycheeImage).
		WithMountedDirectory("/site", dist).
		WithWorkdir("/site").
		WithExec(
			[]string{
				"--no-progress",
				"--base-url", "/site",
				"--max-retries", "2",
				"--retry-wait-time", "5",
				// 202 = some sites (e.g. ballotpedia) return Accepted to bot clients;
				// 403/429 = treat anti-bot rejections as OK rather than dead links.
				"--accept", "200,202,204,403,429",
				// Own domain is verified via htmltest; skip externally so build state isn't
				// gated on deployed content.
				"--exclude", `^https?://(www\.)?haydenk\.github\.io`,
				// LinkedIn returns 999 to non-browser clients; checking it adds noise.
				"--exclude", `^https?://(www\.)?linkedin\.com`,
				"--exclude", `^https?://localhost`,
				// GoatCounter attribution in the footer; intermittent connection refused
				// from CI runners means every page reports a false positive.
				"--exclude", `^https?://(www\.)?goatcounter\.com`,
				// Social share endpoints — these URLs are constructed for logged-in
				// browser clicks, not for bot crawls, and routinely return 4xx/5xx.
				"--exclude", `^https?://news\.ycombinator\.com/submitlink`,
				"--exclude", `^https?://(www\.)?reddit\.com/submit`,
				"--exclude", `^https?://(www\.)?facebook\.com/sharer/`,
				"--exclude", `^https?://(twitter\.com|x\.com)/intent/`,
				"--exclude", `^https?://api\.whatsapp\.com/send`,
				"--exclude", `^https?://(t\.me|telegram\.me)/share`,
				// congress.gov sits behind bot protection: browsers get 403 while the
				// CI runner's connection is reset outright, so 403 in --accept can't help.
				"--exclude", `^https?://(www\.)?congress\.gov`,
				// Known dead external links; remove from this list if/when content is updated.
				"--exclude", `^https://www\.youtube\.com/embed/MSm3w9JO9GQ$`,
				"--exclude", `^https://www\.britannica\.com/biography/Ken_Paxton$`,
				"--exclude", `^https://www\.houstonpublicmedia\.org/articles/news/politics/election-2026/2026/03/03/texas-senate-republican-primary-election-results-cornyn-paxton-hunt/$`,
				"--exclude", `^https://www\.topbug\.net/blog/2013/04/14/install-and-use-gnu-command-line-tools-in-mac-os-x/$`,
				"**/*.html",
			},
			dagger.ContainerWithExecOpts{UseEntrypoint: true},
		).
		Sync(ctx)
	return err
}

// Test runs Lint, HtmlTest, and LinkCheck in parallel.
func (m *Blog) Test(
	ctx context.Context,
	// +defaultPath="/"
	// +ignore=[".dagger", "dist", "public", "resources", ".git", "node_modules", "sermon-notes"]
	source *dagger.Directory,
) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return m.Lint(ctx, source) })
	g.Go(func() error { return m.HtmlTest(ctx, source) })
	g.Go(func() error { return m.LinkCheck(ctx, source) })
	return g.Wait()
}
