package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"musicRoom/models"
	"musicRoom/pkg/spotify"
	"net/http"
	"net/url"
	"time"
)

const (
	BASE_URL = "https://api.spotify.com/v1/me/"
)

func Spotify_is_authenticated(session_id string) bool {
	tokens := models.Spotify_get_user_tokens(session_id)
	if tokens.ID > 0 {
		expire := tokens.Expires_in
		if expire.Before(time.Now()) {
			Spotify_refresh_token(session_id)
		}
		return true
	}

	return false
}

func Spotify_refresh_token(session_id string) {
	refresh_token := models.Spotify_get_user_tokens(session_id).Refresh_token

	urlValues := url.Values{}
	urlValues.Add("grant_type", "refresh_token")
	urlValues.Add("refresh_token", refresh_token)
	urlValues.Add("client_id", spotify.ClIENT_ID)
	urlValues.Add("client_secret", spotify.CLIENT_SECRET)

	resp, err := http.PostForm("https://accounts.spotify.com/api/token", urlValues)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	str_body := string(body)
	fmt.Println("refresh token")
	fmt.Println(str_body)

	res := spotify.Callback_res{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	models.Spotify_Update_or_Create_user_tokens(session_id, res.Access_token, res.Token_type, res.Expires_in, res.Refresh_token)
}

func Spotify_execute_api_request(session_id string, endpoint string, post bool, put bool) string {
	tokens := models.Spotify_get_user_tokens(session_id)
	client := &http.Client{}
	if post {
		reqPost, _ := http.NewRequest(http.MethodPost, BASE_URL+endpoint, nil)
		reqPost.Header.Add("Content-type", "application/json")
		reqPost.Header.Add("Authorization", "Bearer "+tokens.Access_token)
		client.Do(reqPost)
	}

	if put {
		reqPut, _ := http.NewRequest(http.MethodPut, BASE_URL+endpoint, nil)
		reqPut.Header.Add("Content-type", "application/json")
		reqPut.Header.Add("Authorization", "Bearer "+tokens.Access_token)
		putRes, err := client.Do(reqPut)
		if err != nil {
			panic(err)
		}
		defer putRes.Body.Close()
		putBody, _ := ioutil.ReadAll(putRes.Body)
		fmt.Print("put Res")
		fmt.Println(string(putBody))
	}

	reqGet, _ := http.NewRequest(http.MethodGet, BASE_URL+endpoint, nil)
	reqGet.Header.Add("Content-type", "application/json")
	reqGet.Header.Add("Authorization", "Bearer "+tokens.Access_token)
	// fmt.Println("Bearer " + tokens.Access_token)
	res, err := client.Do(reqGet)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)

}

func Play_song(host string) {
	Spotify_execute_api_request(host, "player/play", false, true)
}

func Pause_song(host string) {
	Spotify_execute_api_request(host, "player/pause", false, true)
}
