package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
)
import trippb "Apple1/proto/gen/go"

func main() {
	trip := trippb.Trip{
		Start:    "abc",
		End:      "def",
		Duration: 3600,
		FreeCent: 10000,
	}
	fmt.Println(&trip)

	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)

	var trip2 trippb.Trip
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	b, err = json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
