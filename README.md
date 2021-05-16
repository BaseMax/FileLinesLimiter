# File Lines Limiter

A tiny go program to limit lines of a input file and generate result in a new file. (You can specify custom range for needed lines)

## Using

```
$ go run limiter.go
input.txt
5
7
```

Finally, `limiter.out.txt` file will create.

### How build for Windows in Linux?

```
GOARCH=amd64 GOOS=windows go build
```

© Copyright 2021, Max Base
