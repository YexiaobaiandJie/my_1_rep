package main

import(
	"github.com/gin-gonic/gin"
)



func main(){
	r :=gin.Default()
	r.GET("/news",Index)
	r.POST("/login",LoginPage)
	r.POST("/register",RegisterPage)
	r.POST("/publish",PublishPage)
	r.GET("/postings",PostingPage)
	r.GET("/detail",DetailPage)
	r.POST("/comment",CommentPage)
	r.Run(":3000")
}

