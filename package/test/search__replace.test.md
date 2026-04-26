# search replace

## basic replacement

```beforeAll
mkdir -p testdata
echo "host: localhost" > testdata/config.yaml
echo "url: localhost:3000" > testdata/app.yaml
```

```afterAll
rm -rf testdata
```

### should replace text in matching files

```execute
aux4 search replace "localhost" --with "127.0.0.1" --path testdata
cat testdata/config.yaml
```

```expect:partial
127.0.0.1
```

### should report modified files

```execute
echo "host: localhost" > testdata/config.yaml
aux4 search replace "localhost" --with "myhost" --path testdata --include "*.yaml"
```

```expect:partial
modified:*?
```

## no matches

```beforeAll
mkdir -p testdata
echo "hello world" > testdata/test.txt
```

```afterAll
rm -rf testdata
```

### should report no matches

```execute
aux4 search replace "nonexistent" --with "replaced" --path testdata
```

```expect
No matches found
```
