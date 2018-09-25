package main

import (
	"fmt"

	"./notifications"
)

/*
https://developers.google.com/protocol-buffers/docs/gotutorial

https://godoc.org/github.com/golang/protobuf/proto

The complete guide to writing .proto files:
https://developers.google.com/protocol-buffers/docs/proto3

https://developers.google.com/protocol-buffers/docs/reference/go-generated
https://developers.google.com/protocol-buffers/docs/encoding
*/

func main() {
	fmt.Println(notifications.NotificationStatus_name[1])
}
