package service

import (
	"github.com/alscaldeira/twitter/repository"
)

const (
	urlTweets = "https://api.twitter.com/2/tweets"
	POST      = "POST"
)

func Post(username string, password string, content string) {
	web := Init(username, password)
	PostContent(content, web)
}

func PostUser(username string, password string) {
	if username != "" && password != "" {
		repository.UserPost(username, password)
	}
}

func Login(username string) {
	user := repository.UserGet(username)
	Init(username, user.Password)
}
