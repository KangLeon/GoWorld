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
		StartPos: &trippb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:  35,
			Longitude: 115,
		},
		PathLocations: []*trippb.Location{
			{
				Latitude:  31,
				Longitude: 119,
			},
			{
				Latitude:  32,
				Longitude: 118,
			},
		},
		Status: trippb.TripStatus_IN_PROGRESS,
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
