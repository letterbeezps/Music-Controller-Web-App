package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"musicRoom/models"
	"musicRoom/pkg/e"
	"musicRoom/pkg/spotify"
	pkgSpotify "musicRoom/pkg/spotify"
	utilSpotify "musicRoom/utils/spotify"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthURL(c *gin.Context) {
	scopes := "user-read-playback-state user-modify-playback-state user-read-currently-playing"
	s_params := url.Values{}
	Url, err := url.Parse("https://accounts.spotify.com/authorize")
	if err != nil {
		panic(err)
	}

	s_params.Set("scope", scopes)
	s_params.Set("response_type", "code")
	s_params.Set("redirect_uri", pkgSpotify.REDIRECT_URI)
	s_params.Set("client_id", pkgSpotify.ClIENT_ID)
	Url.RawQuery = s_params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)

	code := e.SUCCESS
	data := make(map[string]interface{})
	data["url"] = urlPath

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func Spotify_Callback(c *gin.Context) {
	s_code := c.Query("code")
	// s_error := c.Query("error")

	urlValues := url.Values{}
	urlValues.Add("grant_type", "authorization_code")
	urlValues.Add("code", s_code)
	urlValues.Add("redirect_uri", spotify.REDIRECT_URI)
	urlValues.Add("client_id", spotify.ClIENT_ID)
	urlValues.Add("client_secret", spotify.CLIENT_SECRET)

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", urlValues)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	str_body := string(body)
	fmt.Println(str_body)

	res := spotify.Callback_res{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	session := sessions.Default(c)
	oldHost := ""
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find host_key in session Spotify_Callback")
		oldHost = session.Get("host_key").(string)
	}

	models.Spotify_Update_or_Create_user_tokens(oldHost, res.Access_token, res.Token_type, res.Expires_in, res.Refresh_token)
	c.Redirect(http.StatusMovedPermanently, "http://192.168.199.133:3000")

}

func IsAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	oldHost := ""
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find host_key in session IsAuthenticated")
		oldHost = session.Get("host_key").(string)
	} else {
		fmt.Println("not find host_key in session IsAuthenticated")
	}
	is_au := utilSpotify.Spotify_is_authenticated(oldHost)

	code := e.SUCCESS
	data := make(map[string]interface{})
	data["status"] = true
	if !is_au {
		code = e.ERROR_AUTH
		data["status"] = false
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func CurrentSong(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})
	session := sessions.Default(c)
	room_code := ""
	if _, ok := session.Get("room_code").(string); ok {
		// fmt.Println("find room_code in session")
		room_code = session.Get("room_code").(string)
	}
	room := models.GetRoomByCode(room_code)
	if room.ID > 0 {
		host := room.Host
		endpoint := "player/currently-playing"
		response := utilSpotify.Spotify_execute_api_request(host, endpoint, false, false)
		// fmt.Println("currentSong")
		// fmt.Println(response)

		if strings.Contains(response, "error") || !strings.Contains(response, "item") {
			code = e.NOT_CONTENT
		} else {
			var res pkgSpotify.CurrentSong_res
			err := json.Unmarshal([]byte(response), &res)
			if err != nil {
				panic(err)
			}

			item := res.Item
			duration := item.DurationMs
			progress := res.ProgressMs
			ablum_cover := item.Album.Images[0].URL
			is_playing := res.IsPlaying
			song_id := item.ID

			artist_string := ""
			for i, artist := range item.Artists {
				if i > 0 {
					artist_string += ", "
				}
				name := artist.Name
				artist_string += name
			}

			data["title"] = item.Name
			data["artist"] = artist_string
			data["duration"] = duration
			data["time"] = progress
			data["image_url"] = ablum_cover
			data["is_playing"] = is_playing
			data["votes"] = 0
			data["id"] = song_id
		}

	} else {
		code = e.NOT_FOUND
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func PauseSong(c *gin.Context) {
	fmt.Println("get PauseSong")
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["code"] = ""
	room_code := ""
	host := ""

	session := sessions.Default(c)
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find host_key in session PauseSong")
		host = session.Get("host_key").(string)
	}

	if _, ok := session.Get("room_code").(string); ok {
		fmt.Println("find room_code in session PauseSong")
		room_code = session.Get("room_code").(string)
	}
	room := models.GetRoomByCode(room_code)

	if room.ID > 0 {
		if room.Host == host || room.Guest_can_pause {
			utilSpotify.Pause_song(room.Host)
		} else {
			code = e.FORBIDDEN_PAUSE_OR_PLAY_SONG
		}
	} else {
		code = e.NOT_ROOMCODE_OF_ROOM
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func PlaySong(c *gin.Context) {
	fmt.Println("get playSong")
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["code"] = ""
	room_code := ""
	host := ""

	session := sessions.Default(c)
	if _, ok := session.Get("host_key").(string); ok {
		fmt.Println("find host_key in session PlaySong")
		host = session.Get("host_key").(string)
	}

	if _, ok := session.Get("room_code").(string); ok {
		fmt.Println("find room_code in session PlaySong")
		room_code = session.Get("room_code").(string)
	}
	room := models.GetRoomByCode(room_code)

	if room.ID > 0 {
		if room.Host == host || room.Guest_can_pause {
			utilSpotify.Play_song(room.Host)
		} else {
			code = e.FORBIDDEN_PAUSE_OR_PLAY_SONG
			fmt.Println("forbidden")
		}
	} else {
		code = e.NOT_ROOMCODE_OF_ROOM
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
