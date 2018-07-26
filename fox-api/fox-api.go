package main

import(
	"net/http"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	
	"encoding/json"
	"strconv"
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
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news",Index)
	log.Fatal(http.ListenAndServe(":8080",router))
}

func Index(w http.ResponseWriter,r *http.Request){
	value1 :=r.URL.Query().Get("pagesize")
	value2,err:=strconv.Atoi(value1)
	var infos []Info
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("foxinfo")
	c :=db.C("info")
	c.Find(nil).Sort("-publishdate").Limit(value2).All(&infos)
	w.Header().Set("Content-type","application/json")
	jsons,err :=json.Marshal(infos)
	if err !=nil{
		panic(err)
	}
	fmt.Fprintln(w,string(jsons))
}

