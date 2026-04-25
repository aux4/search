#### Description

The `files` command finds files matching a glob pattern by recursively walking the directory tree. It returns newline-separated file paths relative to the search path.

Supported glob patterns:

- `*.go` — match files by extension in any directory
- `test_*` — match files by prefix
- `**/*.md` — match files at any depth using double-star
- `config.?ml` — single character wildcard
- Plain text — case-insensitive substring match on filename

Directories listed in `--exclude` are skipped entirely, improving performance on large trees. Binary files are not filtered — this command only matches file names, not contents.

#### Usage

```bash
aux4 search files "<pattern>" [--path <dir>] [--exclude <patterns>] [--maxDepth <n>]
```

--pattern   Glob pattern to match (required, positional)
--path      Directory to search in (default: `.`)
--exclude   Comma-separated directory names to skip (default: `node_modules,.git`)
--maxDepth  Maximum recursion depth. 0 means only the target directory itself

#### Example

```bash
aux4 search files "*.go" --path src
```

```text
main.go
utils/helper.go
cmd/root.go
```

```bash
aux4 search files "*.md" --maxDepth 0
```

```text
README.md
CHANGELOG.md
```
