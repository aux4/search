# aux4/search

File and content search utility with glob patterns and regex matching. Find files by name using glob patterns and search file contents using regular expressions.

## Installation

```bash
aux4 aux4 pkger install aux4/search
```

## Quick Start

```bash
# Find all Go files
aux4 search files "*.go"

# Find all markdown files under docs/
aux4 search files "*.md" --path docs

# Search for a function definition
aux4 search content "func main" --include "*.go"

# Search with context lines
aux4 search content "TODO" --context 2
```

## Commands

### search files

Find files matching a glob pattern recursively. Returns newline-separated file paths relative to the search path.

```bash
aux4 search files "<pattern>" [--path <dir>] [--exclude <patterns>] [--maxDepth <n>]
```

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pattern` | Glob pattern to match (e.g., `*.go`, `**/*.md`) | required |
| `--path` | Directory to search in | `.` |
| `--exclude` | Comma-separated directory names to skip | `node_modules,.git` |
| `--maxDepth` | Maximum recursion depth | unlimited |

**Supported patterns:**
- `*.go` — match files by extension
- `test_*` — match files by prefix
- `**/*.md` — match files at any depth
- `config.?ml` — single character wildcard

```bash
# Find all TypeScript files
aux4 search files "*.ts" --path src

# Find files, skip build directories
aux4 search files "*.js" --exclude "node_modules,.git,dist,build"

# Find files only in top-level directory
aux4 search files "*.md" --maxDepth 0
```

### search content

Search file contents using regex pattern matching (case-insensitive). Returns matches in `file:line: content` format.

```bash
aux4 search content "<pattern>" [--path <dir>] [--include <glob>] [--exclude <patterns>] [--limit <n>] [--context <n>]
```

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pattern` | Regex pattern to search for | required |
| `--path` | Directory to search in | `.` |
| `--include` | File glob filter (e.g., `*.go`, `*.md`) | all files |
| `--exclude` | Comma-separated directory names to skip | `node_modules,.git` |
| `--limit` | Maximum number of matches | `50` |
| `--context` | Lines to show before and after each match | `0` |

```bash
# Search for error handling patterns in Go files
aux4 search content "if err != nil" --include "*.go"

# Search with surrounding context
aux4 search content "TODO|FIXME" --context 3

# Limit results
aux4 search content "import" --include "*.js" --limit 10
```

### search replace

Find and replace text across multiple files. Uses grep to locate matches and sed to replace in-place.

```bash
aux4 search replace "<pattern>" --with "<replacement>" [--path <dir>] [--include <glob>] [--exclude <patterns>]
```

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pattern` | Text or regex pattern to find | required |
| `--with` | Replacement text | required |
| `--path` | Directory to search in | `.` |
| `--include` | File glob filter (e.g., `*.yaml`) | all files |
| `--exclude` | Comma-separated directory names to skip | `node_modules,.git` |

```bash
# Replace localhost with 127.0.0.1 in all yaml files
aux4 search replace "localhost" --with "127.0.0.1" --include "*.yaml"

# Rename a function across Go files
aux4 search replace "oldFunc" --with "newFunc" --include "*.go" --path src
```

### search count

Count matches per file. Shows `file:count` for files with at least one match.

```bash
aux4 search count "<pattern>" [--path <dir>] [--include <glob>] [--exclude <patterns>]
```

| Parameter | Description | Default |
|-----------|-------------|---------|
| `pattern` | Regex pattern to count | required |
| `--path` | Directory to search in | `.` |
| `--include` | File glob filter | all files |
| `--exclude` | Comma-separated directory names to skip | `node_modules,.git` |

```bash
# Count TODOs per file
aux4 search count "TODO" --include "*.js" --path src
```

## License

Apache-2.0
