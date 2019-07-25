# Notes

## Compiling for AWS Lambda:

With debug info:

```
$ GOOS=linux go build -o main
$ zip deployment.zip main
```

Without debug info:

```
$ GOOS=linux go build -ldflags="-s -w" -o main
$ zip deployment.zip main
```

## Deploying to AWS Lambda:

```
aws lambda update-function-code \
  --function-name YourLambdaName \
  --zip-file fileb://$PWD/deployment.zip \
  --publish \
  --no-dry-run
```

## Reducing the size of the binary:

See "[Shrink Your Go Binaries With This One Weird Trick](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)".

Strip all debugging information:

`$ go build -ldflags="-s -w" -o main`

Compress the binary (it will auto-decompress on run):

`$ upx --brute main`

Be advised that compressing a binary in this way:

1. Takes some time.
2. Incurs a small delay when starting the binary (~150ms, ballpark value).
