package main

import (
	"fmt"
	"log"
	"miniProjectGolang/dto"
	"os"
)

var rooms []dto.Room = GenerateRooms()

func main() {
	fmt.Println("***** Welcome To MyProject *****")
	startProject()
}
func startProject() {
	var operationTypeCode int
	fmt.Println("Enter a operation : ")
	fmt.Println("1: Room List ")
	fmt.Println("2: Add Room ")
	fmt.Println("3: Reserve Room ")
	fmt.Println("4: Exit")
	fmt.Print("your operation : ")
	_, err := fmt.Scanf("%d\n", &operationTypeCode)
	if err != nil {
		return
	}
	switch operationTypeCode {
	case 1:
		showRoomList()
	case 2:
		addRoom()
	case 3:
		reserveRoom()
	case 4:
		os.Exit(0)
	default:
		fmt.Println("wrong Operation")
	}
	startProject()

}

func reserveRoom() {
	room := GetRoom()
	var nights int
	var personCount int
	if room.Status == dto.Reserved {
		fmt.Println("room already reserved .......")
		return
	}
	fmt.Print("enter your nights : ")
	_, err := fmt.Scanf("%d\n", &nights)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("enter your personCount : ")
	_, err = fmt.Scanf("%d\n", &personCount)
	if err != nil {
		log.Fatal(err)
	}
	price, tax, discountAmount, finalPrice := calculateRoomPrice(*room, nights, personCount)
	fmt.Printf("Room price : %.2f, tax : %.2f, discountAmount : %.2f, final Price : %.2f\n", price, tax, discountAmount, finalPrice)
	room.Status = dto.Reserved
}

func GetRoom() *dto.Room {
	var room *dto.Room
	var roomId int
	fmt.Print("enter your bed roomId : ")
	_, err := fmt.Scanf("%d\n", &roomId)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Id == roomId {
			room = &rooms[i]
			break
		}
	}
	if room == nil {
		fmt.Println("invalid roomId")
		fmt.Println("please try again")
		startProject()
	}
	return room
}

func addRoom() {
	room := GetRoomFromInput()
	rooms = append(rooms, room)
	fmt.Println("your room information added ........")
}

func GetRoomFromInput() dto.Room {
	var operationTypeCode int
	var price int
	var roomType dto.RoomType
	roomId := getLastIdFromRooms() + 1000
	var room = dto.Room{Id: roomId, Status: dto.Available}
	fmt.Println("enter your room information  ")
	fmt.Print("enter your bed count : ")
	_, err := fmt.Scanf("%d\n", &room.BedCount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("choose your roomType")
	fmt.Println("0 : single  price(200)")
	fmt.Println("1 : standard  price(400) ")
	fmt.Println("2 : suit  price(600) ")
	fmt.Println("3 : double  price(800)   ")
	_, err = fmt.Scanf("%d\n", &operationTypeCode)
	if err != nil {
		log.Fatal(err)
	}

	if operationTypeCode == dto.Single.Value {
		roomType = dto.Single
		price = 200
	} else if operationTypeCode == dto.Standard.Value {
		roomType = dto.Standard
		price = 400
	} else if operationTypeCode == dto.Suit.Value {
		roomType = dto.Suit
		price = 600
	} else if operationTypeCode == dto.Double.Value {
		roomType = dto.Double
		price = 800
	} else {
		fmt.Println("wrong Operation")
		fmt.Println("please try again .........")
		startProject()
	}
	room.RoomType = roomType
	room.Price = price
	return room
}
func calculateRoomPrice(room dto.Room, nights int, personCount int) (roomPrice float64, tax float64, discountAmount float64, finalPrice float64) {
	var priceRate float64
	discountPercentage := 0.0
	if nights >= 7 && nights <= 15 {
		discountPercentage = 0.1
	} else if nights > 15 && nights <= 30 {
		discountPercentage = 0.15
	} else if nights > 30 {
		discountPercentage = 0.2
	}

	switch room.RoomType {
	case dto.Single:
		priceRate = 1.0
	case dto.Double:
		priceRate = 1.2
	case dto.Standard:
		priceRate = 1.4
	case dto.Suit:
		priceRate = 1.6
	}
	roomPrice = float64(nights*room.Price*personCount) * priceRate
	tax = roomPrice * 0.09
	discountAmount = roomPrice * discountPercentage
	finalPrice = roomPrice + tax - discountAmount
	return
}

func showRoomList() {
	for _, room := range rooms {
		fmt.Println(room)
	}
}
func getLastIdFromRooms() int {
	var roomId int
	room := rooms[len(rooms)-1]
	roomId = room.Id
	return roomId
}

func GenerateRooms() []dto.Room {
	var rooms []dto.Room
	rooms = append(rooms, dto.Room{Id: 1000, RoomType: dto.Single, Status: dto.Available, BedCount: 1, Price: 200})
	rooms = append(rooms, dto.Room{Id: 2000, RoomType: dto.Suit, Status: dto.Available, BedCount: 3, Price: 600})
	rooms = append(rooms, dto.Room{Id: 3000, RoomType: dto.Standard, Status: dto.Available, BedCount: 2, Price: 400})
	rooms = append(rooms, dto.Room{Id: 4000, RoomType: dto.Double, Status: dto.Available, BedCount: 4, Price: 800})

	rooms = append(rooms, dto.Room{Id: 5000, RoomType: dto.Single, Status: dto.Available, BedCount: 1, Price: 200})
	rooms = append(rooms, dto.Room{Id: 6000, RoomType: dto.Suit, Status: dto.Available, BedCount: 3, Price: 600})
	rooms = append(rooms, dto.Room{Id: 7000, RoomType: dto.Standard, Status: dto.Available, BedCount: 2, Price: 400})
	rooms = append(rooms, dto.Room{Id: 8000, RoomType: dto.Double, Status: dto.Available, BedCount: 4, Price: 800})

	rooms = append(rooms, dto.Room{Id: 9000, RoomType: dto.Single, Status: dto.Available, BedCount: 1, Price: 200})
	rooms = append(rooms, dto.Room{Id: 10000, RoomType: dto.Suit, Status: dto.Available, BedCount: 3, Price: 600})
	rooms = append(rooms, dto.Room{Id: 11000, RoomType: dto.Standard, Status: dto.Available, BedCount: 2, Price: 400})
	rooms = append(rooms, dto.Room{Id: 12000, RoomType: dto.Double, Status: dto.Available, BedCount: 4, Price: 800})

	return rooms
}
