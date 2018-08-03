package main

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)



func main(){
	r :=gin.Default()
	conf := cors.DefaultConfig()
	conf.AllowOriginFunc = func(origin string) bool {
		return true
	}
	conf.AllowCredentials = true
	conf.ExposeHeaders =  []string{"Content-Length", "X-Server-Unique-ID", "Set-Cookie"}
	r.Use(cors.New(conf))
	r.GET("/news",Index)
	r.GET("/newsd",NewsdetailPage)
	r.POST("/login",LoginPage)
	r.POST("/register",RegisterPage)
	r.POST("/publish",PublishPage)
	r.GET("/postings",PostingPage)
	r.POST("/detail",DetailPage)
	r.POST("/comment",CommentPage)
	// r.POST("/getcomment",GetComPage)
	r.Run(":3000")
}

