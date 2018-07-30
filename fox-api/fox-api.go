package main

import(
	// "net/http"
	// "log"
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)



func main(){
	// router :=mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/news",Index)
	// router.HandleFunc("/register",RegisterPage)
	// router.HandleFunc("/login",LoginPage)
	// router.HandleFunc("/publish",PublishPage)
	// router.HandleFunc("/postings",PostingPage)
	// router.HandleFunc("/detail",DetailPage)
	// router.HandleFunc("/comment",CommentPage)
	// log.Fatal(http.ListenAndServe(":8080",router))
	r :=gin.Default()
	r.GET("/news",Index)
	r.GET("/login",LoginPage)
	r.POST("/register",RegisterPage)
	r.Run(":8080")
}

