package handler

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"time"
	"url-shortener/shared/database"
	"url-shortener/shared/utils"
)

type CreateUrlRequest struct {
	Url string `json:"url"`
}
type CreateUrlResponse struct {
	Code    int    `json:"code"`
	Url     string `json:"url"`
	Message string `json:"message"`
}
type GetUrlRequest struct {
	Url string `json:"url"`
}

func CreateUrlHandler(c *gin.Context) {
	//TODO VALIDATION
	//TODO ENCODE URL
	//TODO CONVERT TO BASE62 THEN SHA256
	//TODO GET 6 CHARACTER OF BEGINNING
	//TODO ADD RANDOM NUMBER TO 6 CHARACTER
	//TODO STORE TO REDIS
	//TODO RETURN 200 IF EVERYTHING IS OK
	var data CreateUrlRequest
	var hashedValue string
	if err := c.BindJSON(&data); err != nil {
		return
	}
	encoded := url.QueryEscape(data.Url)
	hashedValue = utils.DoCryptoAlgo(encoded)

	redis := database.GetInstance()

	val, _ := redis.Client.Get(database.Ctx, hashedValue).Result()
	if len(val) > 0 {
		hashedValue = utils.DoCryptoAlgo(val)
	}

	redis.Client.SetEX(database.Ctx, hashedValue, data.Url, time.Hour*24*30 *6) // expire after 180 days

	c.JSON(200, CreateUrlResponse{
		Url:     "http://localhost:8080/" + hashedValue,
		Code:    200,
		Message: "created",
	})

}
