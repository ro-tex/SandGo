package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"./pb"
	"github.com/golang/protobuf/proto"
)

func readNotifsFromJSON(filename string) *pb.Notifications {
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

func readNotifsFromPB(filename string) *pb.Notifications {
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

func writeNotifsToJSON(filename string, notifs *pb.Notifications) {
	j, err := json.Marshal(notifs)
	if err != nil {
		fmt.Println("Err:", err)
	}
	err = ioutil.WriteFile(filename, []byte(string(j)), 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}

func writeNotifsToPB(filename string, notifs *pb.Notifications) {
	pb, err := proto.Marshal(notifs)
	if err != nil {
		fmt.Println("Failed to encode notifications:", err)
	}
	err = ioutil.WriteFile(filename, pb, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}

func main() {
	// notifs := readNotifsFromPB("notifications.pb")
	notifs := readNotifsFromJSON("notifications.json")

	writeNotifsToPB("notifications.pb", notifs)
}
