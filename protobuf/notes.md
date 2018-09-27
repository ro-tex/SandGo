## Notes

-   no inheritance between message types
-   ids under 16 take one less byte to encode, so use those for fields you encounter often
-   repeated items encode their id with each item, so make sure they use an id under 16 as well

## Resources

<https://developers.google.com/protocol-buffers/docs/gotutorial>

<https://godoc.org/github.com/golang/protobuf/proto>

The complete guide to writing .proto files:
<https://developers.google.com/protocol-buffers/docs/proto3>

<https://developers.google.com/protocol-buffers/docs/reference/go-generated>
<https://developers.google.com/protocol-buffers/docs/encoding>

## Compile the .proto

`cd ~/source/SandGo/protobuf/notifications`
`protoc --go_out=./ notification.proto`
