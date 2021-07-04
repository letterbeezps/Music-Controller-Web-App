package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type SpotifyToken struct {
	gorm.Model
	User          string `gorm:"size:50;unique"`
	Refresh_token string `gorm:"size:150"`
	Access_token  string `gorm:"size:150"`
	Expires_in    time.Time
	Token_type    string `gorm:"size:50"`
}

func Spotify_exist_by_user(session_id string) bool {
	var st SpotifyToken
	db.Select("id").Where("user = ?", session_id).First(&st)
	if st.ID > 0 {
		return true
	}
	return false
}

func Spotify_get_user_tokens(session_id string) (st SpotifyToken) {
	db.Where("user = ?", session_id).First(&st)
	return
}

func Spotify_Update_or_Create_user_tokens(
	session_id string,
	access_token string,
	token_type string,
	expires int,
	refresh_token string) {

	tokens := Spotify_get_user_tokens(session_id)
	expires_in := time.Now().Add(time.Second * time.Duration(expires))
	if tokens.ID > 0 {
		tokens.Access_token = access_token
		// tokens.Refresh_token = refresh_token
		tokens.Expires_in = expires_in
		tokens.Token_type = token_type
		db.Save(&tokens)
	} else {
		tokens = SpotifyToken{
			User:          session_id,
			Access_token:  access_token,
			Refresh_token: refresh_token,
			Token_type:    token_type,
			Expires_in:    expires_in,
		}

		db.Create(&tokens)
	}

}
