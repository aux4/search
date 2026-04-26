# search count

## basic count

```beforeAll
mkdir -p testdata
printf 'TODO fix this\nTODO refactor\nall done\n' > testdata/code.js
printf 'no todos here\n' > testdata/clean.js
```

```afterAll
rm -rf testdata
```

### should count matches per file

```execute
aux4 search count "TODO" --path testdata --include "*.js"
```

```expect:partial
code.js:2
```

### should not show files with zero matches

```execute
aux4 search count "TODO" --path testdata --include "*.js"
```

```expect:partial
code.js*?
```

## no matches at all

```beforeAll
mkdir -p testdata
echo "hello" > testdata/file.txt
```

```afterAll
rm -rf testdata
```

### should output nothing when no matches

```execute
aux4 search count "nonexistent" --path testdata
```

```expect
```
