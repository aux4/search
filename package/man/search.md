#### Description

The `search` command provides file and content search utilities. It supports two subcommands:

- **files** — Find files matching a glob pattern recursively
- **content** — Search file contents using regex pattern matching

#### Usage

```bash
aux4 search <files|content> [options]
```

#### Example

```bash
aux4 search files "*.go"
aux4 search content "TODO" --include "*.js"
```
