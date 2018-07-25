package main

import(
	"net/http"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"encoding/json"
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
	router.HandleFunc("/",Index)
	router.HandleFunc("/10",firstIndex)
	router.HandleFunc("/20",secondIndex)
	log.Fatal(http.ListenAndServe(":8080",router))
}

func Index(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json;charset=utf-8")
	fmt.Fprintf(w,"welcome!")
}

func firstIndex(w http.ResponseWriter,r *http.Request){
	var infos []Info
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("foxinfo")
	c :=db.C("info")
	c.Find(nil).Sort("-publishdate").Limit(10).All(&infos)
	w.Header().Set("Content-type","application/json")
	jsons,err :=json.Marshal(infos)
	if err !=nil{
		panic(err)
	}
	fmt.Fprintln(w,string(jsons))
}

func secondIndex(w http.ResponseWriter,r *http.Request){
	var infos []Info
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("foxinfo")
	c :=db.C("info")
	c.Find(nil).Sort("-publishdate").Limit(10).Skip(10).All(&infos)
	w.Header().Set("Content-type","application/json")
	jsons,err :=json.Marshal(infos)
	if err !=nil{
		panic(err)
	}
	fmt.Fprintln(w,string(jsons))
}
