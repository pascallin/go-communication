package protocol

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Message string `json:"message"`
}

func Encode(e Event) []byte {
	str, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	return str
}

func Decode(b []byte) Event {
	// data := Event{}
	// err := json.Unmarshal(b, &data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return data
	return Event{
		Message: string(b),
	}
}
