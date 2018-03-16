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

## Further compress the binary:

```
upx --brute main
```

Be advised that compressing a binary in this way:

1. Takes some time.
2. Incurs a small delay when starting the binary (~150ms, ballpark value).
