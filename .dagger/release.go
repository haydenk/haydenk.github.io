package main

import (
	"context"
	"dagger/blog/internal/dagger"
	"fmt"
)

const (
	alpineImage      = "alpine:3.21"
	defaultReleaseRepo = "haydenk/haydenk.github.io"
)

// Package extracts the GitHub Pages artifact tarball and repacks the site as
// haydenk-blog-<tag>.tar.gz, ready to upload as a release asset.
func (m *Blog) Package(
	ctx context.Context,
	pagesArtifact *dagger.File,
	tag string,
) *dagger.File {
	name := fmt.Sprintf("haydenk-blog-%s.tar.gz", tag)
	return dag.Container().
		From(alpineImage).
		WithMountedFile("/in/artifact.tar", pagesArtifact).
		WithWorkdir("/out").
		WithExec([]string{
			"sh", "-c",
			`mkdir -p site && tar -xf /in/artifact.tar -C site && tar -czf "$1" -C site .`,
			"_", name,
		}).
		File("/out/" + name)
}

// Release packages the Pages artifact and creates a GitHub release with the
// bundle attached and auto-generated release notes. Replaces softprops/action-gh-release.
func (m *Blog) Release(
	ctx context.Context,
	pagesArtifact *dagger.File,
	tag string,
	token *dagger.Secret,
	// +optional
	// +default="haydenk/haydenk.github.io"
	repo string,
) (string, error) {
	if repo == "" {
		repo = defaultReleaseRepo
	}
	bundleName := fmt.Sprintf("haydenk-blog-%s.tar.gz", tag)
	bundle := m.Package(ctx, pagesArtifact, tag)

	return dag.Container().
		From(alpineImage).
		WithExec([]string{"apk", "add", "--no-cache", "github-cli"}).
		WithSecretVariable("GH_TOKEN", token).
		WithMountedFile("/work/"+bundleName, bundle).
		WithWorkdir("/work").
		WithExec([]string{
			"gh", "release", "create", tag,
			"--repo", repo,
			"--title", tag,
			"--generate-notes",
			bundleName,
		}).
		Stdout(ctx)
}
