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
sample.go:5: func main() {
```

### should be case insensitive

```execute
aux4 search content "HELLO" --path testdata --include "*.go"
```

```expect:partial
sample.go:6:*Hello*
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
code.js:1:*hello*
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
many.txt:1: line one match
many.txt:2: line two match
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
ctx.txt*2*line 2
ctx.txt:3: target line
ctx.txt*4*line 4
```
