package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"./pb"
	"github.com/golang/protobuf/proto"
)

func readNotifsFromJSON(filename string) (*pb.Notifications, time.Duration) {
	start := time.Now()

	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	notifs := &pb.Notifications{}

	err = json.Unmarshal([]byte(in), &notifs)
	if err != nil {
		log.Fatalln("Failed to parse notifications:", err)
	}

	return notifs, time.Since(start)
}

func readNotifsFromPB(filename string) (*pb.Notifications, time.Duration) {
	start := time.Now()

	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	notifs := &pb.Notifications{}

	err = proto.Unmarshal(in, notifs)
	if err != nil {
		log.Fatalln("Failed to parse notifications:", err)
	}

	return notifs, time.Since(start)
}

func writeNotifsToJSON(filename string, notifs *pb.Notifications) time.Duration {
	start := time.Now()

	// I'm using MarshalIndent instead of Marshal here because that better reflects the everyday use case.
	// This *does* affect the JSON's size -> it explodes from 14MB to 22MB in my example.
	j, err := json.MarshalIndent(notifs, "", "  ")
	if err != nil {
		fmt.Println("Err:", err)
	}

	err = ioutil.WriteFile(filename, []byte(string(j)), 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}

	return time.Since(start)
}

func writeNotifsToPB(filename string, notifs *pb.Notifications) time.Duration {
	start := time.Now()

	pb, err := proto.Marshal(notifs)
	if err != nil {
		fmt.Println("Failed to encode notifications:", err)
	}

	err = ioutil.WriteFile(filename, pb, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}

	return time.Since(start)
}

func prt(msg string, dur time.Duration, div int64) {
	fmt.Println(msg, float64(int64(dur)/div)/1000/1000, "ms")
}

func main() {
	var runtime time.Duration
	var notifsPB, notifsJSON *pb.Notifications
	var durations struct {
		readPB    time.Duration
		readJSON  time.Duration
		writePB   time.Duration
		writeJSON time.Duration
	}

	var i int64
	for i = 0; i < 100; i++ {
		notifsPB, runtime = readNotifsFromPB("out.pb")
		durations.readPB += runtime
		notifsJSON, runtime = readNotifsFromJSON("out.json")
		durations.readJSON += runtime

		runtime = writeNotifsToPB("out.pb", notifsPB)
		durations.writePB += runtime
		runtime = writeNotifsToJSON("out.json", notifsJSON)
		durations.writeJSON += runtime
	}

	// readJSON		445.677192 ms
	// readPB			 59.701868 ms
	// writeJSON	371.340576 ms
	// writePB		 62.022134 ms

	prt("readJSON\t", durations.readJSON, i)
	prt("readPB\t\t", durations.readPB, i)
	prt("writeJSON\t", durations.writeJSON, i)
	prt("writePB\t\t", durations.writePB, i)
}
