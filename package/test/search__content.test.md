# search content

## basic regex search

```beforeAll
mkdir -p testdata
cat > testdata/sample.go << 'GOEOF'
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func helper() {
	fmt.Println("helper")
}
GOEOF
```

```afterAll
rm -rf testdata
```

### should find matching lines

```execute
aux4 search content "func main" --path testdata --include "*.go"
```

```expect:partial
sample.go*func main*
```

### should be case insensitive

```execute
aux4 search content "HELLO" --path testdata --include "*.go"
```

```expect:partial
sample.go*Hello*
```

## with include filter

```beforeAll
mkdir -p testdata
echo 'function hello() { return "hi"; }' > testdata/code.js
echo 'func hello() string { return "hi" }' > testdata/code.go
```

```afterAll
rm -rf testdata
```

### should filter by file extension

```execute
aux4 search content "hello" --path testdata --include "*.js"
```

```expect:partial
code.js*hello*
```

## with limit

```beforeAll
mkdir -p testdata
printf 'line one match\nline two match\nline three match\nline four match\nline five match\n' > testdata/many.txt
```

```afterAll
rm -rf testdata
```

### should limit results

```execute
aux4 search content "match" --path testdata --include "*.txt" --limit 2
```

```expect:partial
many.txt*line one match
**many.txt*line two match
```

## with context lines

```beforeAll
mkdir -p testdata
printf 'line 1\nline 2\ntarget line\nline 4\nline 5\n' > testdata/ctx.txt
```

```afterAll
rm -rf testdata
```

### should show context around match

```execute
aux4 search content "target" --path testdata --include "*.txt" --context 1
```

```expect:partial
**line 2
**target line
**line 4
```
