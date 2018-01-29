### Build:
```
$ GOOS=linux go build -o main $ zip deployment.zip main
````

### Test payload (test event):
```
{
    "Body":"{\"Name\": \"John Doe\",\"Age\": 33,\"Address\": \"Somestrasse 99, 8001 Zürich\"}"
}
````
