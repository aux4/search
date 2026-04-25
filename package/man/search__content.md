#### Description

The `content` command searches file contents using regex pattern matching. Searches are case-insensitive by default. Results are returned in `file:line: content` format, similar to grep output.

Features:

- **Regex patterns** — full Go regex syntax supported (e.g., `func\s+\w+`, `TODO|FIXME`)
- **File filtering** — use `--include` to restrict search to specific file types
- **Context lines** — show surrounding lines with `--context` for better understanding of matches
- **Result limiting** — cap output with `--limit` to avoid overwhelming results
- **Binary skip** — automatically skips binary files (images, archives, executables, etc.)
- **Large file support** — handles lines up to 1MB

When `--context` is used, non-matching context lines use `-` separator (e.g., `file-10- line`) while matching lines use `:` (e.g., `file:10: line`). Groups of matches are separated by `--`.

#### Usage

```bash
aux4 search content "<pattern>" [--path <dir>] [--include <glob>] [--exclude <patterns>] [--limit <n>] [--context <n>]
```

--pattern   Regex pattern to search for (required, positional)
--path      Directory to search in (default: `.`)
--include   File glob filter, e.g., `*.go`, `*.md` (default: all files)
--exclude   Comma-separated directory names to skip (default: `node_modules,.git`)
--limit     Maximum number of matches to return (default: `50`)
--context   Number of lines to show before and after each match (default: `0`)

#### Example

```bash
aux4 search content "func main" --include "*.go"
```

```text
main.go:10: func main() {
cmd/root.go:15: func main() {
```

```bash
aux4 search content "TODO" --context 1 --limit 5
```

```text
src/handler.go-14- // handle the request
src/handler.go:15: // TODO: add error handling
src/handler.go-16- func handleRequest() {
```
