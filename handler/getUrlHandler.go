package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"url-shortener/shared/database"
)

func GetUrlHandler(c *gin.Context){
	key:=c.Params.ByName("key")

	redis:=database.GetInstance()
	val,err:=redis.Client.Get(database.Ctx,key).Result()

	fmt.Println("val",val)
	if err != nil {
		fmt.Println(err)
	}
	//c.JSON(200,gin.H{
	//	"test":val,
	//})
	c.Redirect(302, val)

}
