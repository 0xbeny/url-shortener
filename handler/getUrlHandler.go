package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"url-shortener/shared/database"
)

type ErrorResponse struct {
	Code    int64
	Message string
}

func GetUrlHandler(c *gin.Context) {
	key := c.Params.ByName("key")
	length := os.Getenv("URL_LENGTH")

	s, err := strconv.Atoi(length)
	if len(key) != s {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "key's length is not enough",
		})
	}
	redis := database.GetInstance()
	val, err := redis.Client.Get(database.Ctx, key).Result()

	if err != nil {
		fmt.Println(err)
	}

	c.Redirect(302, val)

}
