# Security Policy

## Supported Versions

Only the latest deployment on the `master` branch is supported.

## Reporting a Vulnerability

If you discover a security vulnerability in this repository, please report it responsibly.

**Do not open a public issue.**

Instead, please use [GitHub's private vulnerability reporting](https://github.com/haydenk/haydenk.github.io/security/advisories/new) to submit your report.

You can expect an initial response within 7 days. If the vulnerability is accepted, a fix will be deployed as soon as possible. If it is declined, you will receive an explanation.

## Scope

This is a static Hugo site hosted on GitHub Pages. Relevant concerns include:

* Exposed secrets or credentials in the repository
* Cross-site scripting (XSS) in site content or templates
* Dependency vulnerabilities in Hugo modules or GitHub Actions
* Misconfigurations in GitHub Actions workflows

## Out of Scope

* Vulnerabilities in GitHub Pages infrastructure itself
* Vulnerabilities in Hugo upstream
* Social engineering attacks
* Denial of service
