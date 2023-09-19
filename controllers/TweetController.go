package controllers

import (
	"crudtwitter/api/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TweetController struct {
	Tweets []entities.Tweet
}

func NewTweetController() *TweetController {
	controller := TweetController{}
	return &controller
}

func (t *TweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.Tweets)
}

func (t *TweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()
	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}
	t.Tweets = append(t.Tweets, *tweet)
	ctx.JSON(http.StatusCreated, t.Tweets)
}

func (t *TweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	for idx, tweet := range t.Tweets {
		if tweet.ID == id {
			t.Tweets = append(t.Tweets[0:idx], t.Tweets[idx+1:]...)
			return
		}

	}
	ctx.JSON(
		http.StatusNotFound,
		gin.H{"error": "Tweet not found"},
	)
}
