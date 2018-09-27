package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"./pb"
)

/*
https://developers.google.com/protocol-buffers/docs/gotutorial

https://godoc.org/github.com/golang/protobuf/proto

The complete guide to writing .proto files:
https://developers.google.com/protocol-buffers/docs/proto3

https://developers.google.com/protocol-buffers/docs/reference/go-generated
https://developers.google.com/protocol-buffers/docs/encoding
*/

func readJSONNotifs(filename string) *pb.Notifications {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	notifs := &pb.Notifications{}
	if err := json.Unmarshal([]byte(in), &notifs); err != nil {
		log.Fatalln("Failed to parse notifications:", err)
	}
	return notifs
}

func readJSONNotif(filename string) *pb.Notification {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	notif := &pb.Notification{}
	if err := json.Unmarshal([]byte(in), &notif); err != nil {
		log.Fatalln("Failed to parse notifications:", err)
	}
	return notif
}

func main() {
	// notifs := readNotifs("../notifications.pb")
	notif := readJSONNotif("../notif.json")
	// notifs := readJSONNotifs("../notifications.json")
	fmt.Println(">>>")
	fmt.Println(notif)
	fmt.Println("<<<")

	// j, err := json.Marshal(notif)
	// if err != nil {
	// 	fmt.Println("Err:", err)
	// }
	// fmt.Println(string(j))
}
