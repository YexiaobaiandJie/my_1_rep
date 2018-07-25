package main

import(
	"net/http"
	"fmt"
	"html"
	"log"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type Info struct{
	AuthorName 	string 	`json:"authorName"`
	Id 			int	   	`json:"id"`
	Language 	string	`json:"language"`
	MobileUrl 	string	`json:"mobileUrl"`
	PublishDate string	`json:"publishDate"`
	SiteName 	string	`json:"siteName"`
	Summary     string	`json:"summary"`
	SummaryAuto string	`json:"summaryAuto"`
	Title     	string	`json:"title"`
	Url  		string	`json:"url"`
}

func main(){
	//连接MongoDB服务器
	var infos []Info
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("foxinfo")
	c :=db.C("info")
	c.Find(nil).Sort("-publishdate").Limit(10).All(&infos)
	
	http.HandleFunc("/",func(writer http.ResponseWriter,request *http.Request){
		fmt.Fprintln(writer,infos,html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080",nil))
}