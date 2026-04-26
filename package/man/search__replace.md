#### Description

The `replace` command finds and replaces text across multiple files. It first uses grep to locate files containing the pattern, then applies sed to perform the replacement in-place. Each modified file is reported.

The pattern is a regex compatible with both grep and sed. The replacement is a literal string (sed substitution syntax applies — use `\1` for capture groups).

#### Usage

```bash
aux4 search replace "<pattern>" --with "<replacement>" [--path <dir>] [--include <glob>] [--exclude <dirs>]
```

--pattern   Text or regex pattern to find (positional, required)
--with      Replacement text (required)
--path      Directory to search in (default: `.`)
--include   File glob filter (e.g., `*.js`, `*.yaml`)
--exclude   Comma-separated directory names to skip (default: `node_modules,.git`)

#### Example

```bash
aux4 search replace "localhost" --with "127.0.0.1" --include "*.yaml" --path config
```

```text
modified: config/dev.yaml
modified: config/test.yaml
```
