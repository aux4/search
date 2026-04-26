# search files

## with glob pattern

```beforeAll
mkdir -p testdata/sub
echo "package main" > testdata/hello.go
echo "hello world" > testdata/world.txt
echo "package sub" > testdata/sub/nested.go
```

```afterAll
rm -rf testdata
```

### should find files matching extension

```execute
aux4 search files "*.go" --path testdata
```

```expect:partial
hello.go
```

### should find files in nested directories

```execute
aux4 search files "*.go" --path testdata
```

```expect:partial
sub/nested.go
```

### should not match non-matching files

```execute
aux4 search files "*.md" --path testdata
```

```expect
```

## with maxDepth

```beforeAll
mkdir -p testdata/deep
echo "top level" > testdata/top.txt
echo "deep level" > testdata/deep/bottom.txt
```

```afterAll
rm -rf testdata
```

### should limit depth

```execute
aux4 search files "*.txt" --path testdata --maxDepth 1
```

```expect
top.txt
```

## with exclude

```beforeAll
mkdir -p testdata/src testdata/dist
echo "code" > testdata/src/app.js
echo "built" > testdata/dist/app.js
```

```afterAll
rm -rf testdata
```

### should exclude directories

```execute
aux4 search files "*.js" --path testdata --exclude dist
```

```expect
src/app.js
```
