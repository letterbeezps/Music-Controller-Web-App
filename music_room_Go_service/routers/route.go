package routers

import (
	"musicRoom/pkg/middlewares"
	"musicRoom/routers/api"
	"musicRoom/routers/api/spotify"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(middlewares.Cors())

	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("sessionId", store))

	r.GET("/hello", api.TestGin)
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, ":3000")
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/rooms", api.GetRooms)
		apiV1.POST("/rooms", api.CreateRoom)
		apiV1.GET("/room", api.GetRoomByCode)
		apiV1.POST("/Join_room", api.JoinRoom)
		apiV1.GET("/User_in_room", api.UserInRoom)
		apiV1.POST("/Leave_room", api.LeaveRoom)
		apiV1.POST("/Update_room", api.UpdateRomm)
	}

	apiSpotify := r.Group("/api/spotify")
	{
		apiSpotify.GET("/Get_auth_url", spotify.AuthURL)
		apiSpotify.GET("/redirect", spotify.Spotify_Callback)
		apiSpotify.GET("/Is_authenticated", spotify.IsAuthenticated)
		apiSpotify.GET("/current_song", spotify.CurrentSong)
		apiSpotify.PUT("/pause", spotify.PauseSong)
		apiSpotify.PUT("/play", spotify.PlaySong)
	}

	return r
}
