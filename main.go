package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bjmathias/go_barry_protodef"
	"github.com/golang/protobuf/proto"
)

func main() {

	httpListen()

}

// httpTestFunc handler for an http /test request
func httpPbTestFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Server response (PB Marshal/Unmarshal)\n")

	Messagetest := &go_barry_protodef.BarryTest{
		Messagetype: proto.Int32(1),
		Version:     proto.Int32(3),
		Timestamp:   proto.Int64(time.Now().Unix()),
		Value:       proto.Int64(3),
	}

	rawData, err := proto.Marshal(Messagetest)
	if err != nil {
		log.Fatal("serialisation failed: ", err)
	}
	NewMessageTest := &go_barry_protodef.BarryTest{}
	err = proto.Unmarshal(rawData, NewMessageTest)
	if err != nil {
		log.Fatal("deserialisation failed: ", err)
	}
	fmt.Fprintf(w, "MessageType: %d\n", NewMessageTest.GetMessagetype())
	fmt.Fprintf(w, "Version: %d\n", NewMessageTest.GetVersion())
	fmt.Fprintf(w, "Timestamp (time_t): %d\n", NewMessageTest.GetTimestamp())

	ts := time.Unix(NewMessageTest.GetTimestamp(), 0)
	ts.Format(time.UnixDate)
	fmt.Fprintf(w, "Timestamp formatted: %s\n", ts.Format(time.UnixDate))

}

// httpListen starts a basic http server on designated port
func httpListen() {

	httpPort := 8840
	fmt.Printf("Starting http server on port %d.......\n", httpPort)
	http.HandleFunc("/pbtest", httpPbTestFunc)
	http.ListenAndServe(":"+strconv.Itoa(httpPort), nil)
	return
}
