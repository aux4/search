#### Description

The `count` command counts how many times a pattern appears in each file. It uses grep with the `-c` flag and filters out files with zero matches. Output format is `file:count`.

#### Usage

```bash
aux4 search count "<pattern>" [--path <dir>] [--include <glob>] [--exclude <dirs>]
```

--pattern   Regex pattern to count (positional, required)
--path      Directory to search in (default: `.`)
--include   File glob filter (e.g., `*.go`, `*.md`)
--exclude   Comma-separated directory names to skip (default: `node_modules,.git`)

#### Example

```bash
aux4 search count "TODO" --include "*.js" --path src
```

```text
src/handler.js:3
src/utils.js:1
```
