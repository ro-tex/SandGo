package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"./notif"
	"./pb"
	"github.com/golang/protobuf/proto"
)

/*
https://developers.google.com/protocol-buffers/docs/gotutorial

https://godoc.org/github.com/golang/protobuf/proto

The complete guide to writing .proto files:
https://developers.google.com/protocol-buffers/docs/proto3

https://developers.google.com/protocol-buffers/docs/reference/go-generated
https://developers.google.com/protocol-buffers/docs/encoding
*/

func readNotifsPB(filename string) *pb.Notifications {
	// Read the existing address book.
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	notifs := &pb.Notifications{}
	if err := proto.Unmarshal(in, notifs); err != nil {
		log.Fatalln("Failed to parse notifications:", err)
	}
	return notifs
}

func readNotifsJSON(filename string) *notif.Notifications {
	// Read the existing address book.
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	notifs := &notif.Notifications{}
	if err := json.Unmarshal([]byte(in), &notifs); err != nil {
		log.Fatalln("Failed to parse notifications:", err)
	}
	return notifs
}

func main() {
	fmt.Println(pb.NotificationStatus_name[1])

	// notifs := readNotifsPB("../notifications.pb")
	notifs := readNotifsJSON("../notifications.json")

	fmt.Println(notifs.Notifications[0].NotificationID)
}
