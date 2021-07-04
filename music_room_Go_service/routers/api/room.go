package api

import (
	"fmt"
	"musicRoom/models"
	"musicRoom/pkg/e"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["list"] = models.GetRooms(maps)
	data["total"] = models.GetRoomTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type Room struct {
	Code string `json:"code"`
	// Host            string `json:"host"`
	Guest_can_pause bool `json:"guest_can_pause"`
	Votes_to_skip   int  `json:"votes_to_skip"`
}

func CreateRoom(c *gin.Context) {
	session := sessions.Default(c)
	Host := ""
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find host_key in session")
		Host = session.Get("host_key").(string)
	}
	json := Room{}
	data := &models.Room{}
	c.BindJSON(&json)

	code := e.INVALID_PARAMS
	if !models.ExistRoomByCode(json.Code) {
		code = e.SUCCESS
		data = models.CreateRoom(json.Code, Host, json.Guest_can_pause, json.Votes_to_skip)
	} else {
		code = e.ERROR_EXIST_CODE
	}

	session.Set("host_key", data.Host)
	session.Set("room_code", data.Code)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetRoomByCode(c *gin.Context) {
	session := sessions.Default(c)
	oldHost := ""
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find old host_key in session GetRoomByCode")
		oldHost = session.Get("host_key").(string)
	}
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["is_host"] = false

	roomCode := c.Query("roomCode")

	if len(roomCode) == 0 {
		code = e.INVALID_PARAMS
	} else if !models.ExistRoomByCode(roomCode) {
		code = e.NOT_FOUND
	} else {
		room := models.GetRoomByCode(roomCode)
		data["room"] = room
		data["is_host"] = room.Host == oldHost

	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func JoinRoom(c *gin.Context) {

	code := e.SUCCESS
	data := make(map[string]interface{})

	session := sessions.Default(c)
	if _, ok := session.Get("room_code").(string); ok {
		fmt.Println("find room_code in session")
	}
	json := Room{}
	c.BindJSON(&json)
	roomCode := json.Code
	if len(roomCode) == 0 {
		code = e.INVALID_PARAMS
	} else if !models.ExistRoomByCode(roomCode) {
		code = e.NOT_FOUND
	} else {
		room := models.GetRoomByCode(roomCode)
		session.Set("room_code", room.Code)
		session.Save()
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func UserInRoom(c *gin.Context) {

	code := e.SUCCESS
	data := make(map[string]interface{})
	data["code"] = ""

	session := sessions.Default(c)
	if _, ok := session.Get("room_code").(string); ok {
		fmt.Println("find room_code in session UserInRomm")
		data["code"] = session.Get("room_code").(string)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func LeaveRoom(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})

	session := sessions.Default(c)
	if _, ok := session.Get("room_code").(string); ok {
		fmt.Println("find room_code in session UserInRomm")
		session.Delete("room_code")
		session.Save()
	}

	oldHost := ""
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find old host_key in session GetRoomByCode")
		oldHost = session.Get("host_key").(string)
		session.Delete("host_key")
		session.Save()

	}

	if models.ExistRoomByHost(map[string]interface{}{"host": oldHost}) {
		models.DeleteRoomByHost(oldHost)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func UpdateRomm(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})
	session := sessions.Default(c)
	json := Room{}
	c.BindJSON(&json)

	roomCode := json.Code
	if !models.ExistRoomByCode(roomCode) {
		code = e.NOT_FOUND
	} else {
		oldHost := ""
		if _, ok := session.Get("host_key").(string); ok {
			fmt.Println("find old host_key in session UpdateRoom")
			oldHost = session.Get("host_key").(string)

		} else {
			panic("not host_key in session")
		}
		room := models.GetRoomByCode(roomCode)
		if room.Host != oldHost {
			code = e.NOT_HOST_OF_ROOM
		} else {
			data["room"] = models.UpdateRommByCode(roomCode, json.Guest_can_pause, json.Votes_to_skip)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
