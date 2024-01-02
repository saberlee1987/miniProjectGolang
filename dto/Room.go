package dto

import (
	"encoding/json"
	"fmt"
)

type RoomType struct {
	Value int
	Title string
}

type StatusType struct {
	Value int
	Title string
}

var (
	Single   = RoomType{Value: 0, Title: "single"}
	Standard = RoomType{Value: 1, Title: "standard"}
	Suit     = RoomType{Value: 2, Title: "suit"}
	Double   = RoomType{Value: 3, Title: "double"}

	Reserved  = StatusType{Value: 0, Title: "reserved"}
	Available = StatusType{Value: 1, Title: "available"}
)

type Room struct {
	Id       int
	RoomType RoomType
	Status   StatusType
	BedCount int
	Price    int
}

func (r Room) String() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("{\"id\":%s,\"roomType\":\"%s\",\"status\":\"%s\",\"bedCount\":%d,\"price\":%d}",
			r.Id, r.RoomType.Title, r.Status.Title, r.BedCount, r.Price)
	}
	return string(marshal)
}
