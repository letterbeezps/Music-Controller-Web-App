package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

type Room struct {
	gorm.Model
	Code            string `gorm:"size:8;default:'';unique" json:"code"`
	Host            string `gorm:"size:25;unique" json:"host"`
	Guest_can_pause bool   `gorm:"not null;default:false" json:"guest_can_pause"`
	Votes_to_skip   int    `gorm:"not null;default:1" json:"votes_to_skip"`
}

func ExistRoomByCode(code string) bool {
	var room Room
	db.Select("id").Where("code = ?", code).First(&room)
	if room.ID > 0 {
		return true
	}
	return false
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Generate_unique_code(l int) string {
	rand.Seed(time.Now().UnixNano())

	var res string
	for {
		res = RandStringBytes(l)
		if !ExistRoomByCode(res) {
			break
		}
	}

	return res
}

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// CRUD

// select all rooms
func GetRooms(maps interface{}) (rooms []Room) {
	db.Where(maps).Find(&rooms)
	return
}

func GetRoomByCode(code string) (room Room) {
	db.Where("code = ?", code).First(&room)
	return
}

func GetRoomTotal(maps interface{}) (count int) {
	db.Model(&Room{}).Where(maps).Count(&count)
	return
}

func ExistRoomByHost(maps interface{}) bool {
	var room Room
	db.Where(maps).First(&room)
	if room.ID > 0 {
		return true
	}
	return false
}

// create rooms
func CreateRoom(code string, host string, guest_can_pause bool, votes_to_skip int) *Room {

	// check if host is exist
	newRoom := &Room{}
	db.Where(map[string]interface{}{"Host": host}).First(newRoom)
	if newRoom.ID > 0 {
		fmt.Println("update old Room")
		newRoom.Guest_can_pause = guest_can_pause
		newRoom.Votes_to_skip = votes_to_skip
		db.Save(newRoom)

	} else {
		fmt.Println("create new room")
		if len(code) == 0 {
			fmt.Println("no code")
			code = Generate_unique_code(6)
			fmt.Println("new code", code)
		}

		host = Generate_unique_code(12)

		newRoom = &Room{
			Code:            code,
			Host:            host,
			Guest_can_pause: guest_can_pause,
			Votes_to_skip:   votes_to_skip,
		}

		db.Create(newRoom)

	}
	return newRoom

}

func DeleteRoomByHost(host string) {
	db.Where("host = ?", host).Delete(&Room{})
}

func UpdateRommByCode(code string, guest_can_pause bool, votes_to_skip int) (room Room) {
	room = GetRoomByCode(code)
	room.Guest_can_pause = guest_can_pause
	room.Votes_to_skip = votes_to_skip
	db.Save(&room)
	return
}
