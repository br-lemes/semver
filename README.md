# semver

A lightweight, standalone CLI tool written in Go to automate semantic versioning
and the release process directly from your Git commit history.

The `semver` tool analyzes your local commits, calculates the next version,
updates the relevant files, generates tags, and creates a GitHub release,
automatically uploading the compiled build artifacts.

## Features

* **Automatic Versioning:** Calculates the next version (`vMajor.Minor.Patch`)
by analyzing Git history with Conventional Commits support.
* **Version State Management:** Synchronizes the current version by generating
and reading a local `.version` file.
* **Automated Git Flow:** Commits the version file, creates the corresponding
tag, and pushes changes to the remote repository.
* **GitHub CLI Integration:** Creates a GitHub release, generates release notes
automatically, and uploads multiple files simultaneously.
* **Cross-Compilation:** Ready to generate optimized binaries for Linux
(amd64/arm64) and Windows through native Makefile automation.

## How Version Calculation Works

The tool scans the commit history of the current branch from the oldest to the
newest and increments the version based on the following rules:

* **major (Breaking Change):** Commits starting with the `type!:` pattern (e.g.,
`feat!: breaking change`). Resets minor and patch counters.
* **minor (New Feature):** Commits starting with `feat` (e.g., `feat: add json
support`). Resets the patch counter.
* **patch (Fixes):** Commits starting with `fix` (e.g., `fix: memory leak`).
Increments the patch counter.

## Prerequisites

To use this tool in your development workflow, you only need:

* **Go** (version 1.26.3 or higher)
* **Git** installed and configured in the repository.
* **GitHub CLI (`gh`)** authenticated on the machine running the release.
* **GNU Make** (optional)

## Command Line Interface (CLI)

The executable has two main modes of operation:

### 1. Root Command (Version Generation)

Calculates the current version based on commits, prints the result, and updates
or creates the static `.version` file at your project root.

```sh
./semver
```

### 2. Release Command

Triggers the full tag publication pipeline and uploads binaries to GitHub.
The CLI syntax requires at least one mandatory asset file but accepts multiple
parameters dynamically.

```sh
./semver release semver-linux-amd64 semver-linux-arm64 semver-windows-amd64.exe
```

## License

This project is licensed under the BSD Zero Clause License (0BSD). You may
copy, modify, and/or distribute this software for any purpose with or without
fee.
