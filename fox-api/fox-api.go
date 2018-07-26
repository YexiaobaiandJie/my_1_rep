package main

import(
	"net/http"
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type user struct{
	Userid int `json:"userid"`
	Password int `json:"password"`
}
func main(){
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news",Index)
	log.Fatal(http.ListenAndServe(":8080",router))
}

func Index(w http.ResponseWriter,r *http.Request){
	 var User []user
	  userid :=r.URL.Query().Get("userid")
	  userpwd :=r.URL.Query().Get("userpwd")
	 value1 :=r.URL.Query().Get("pagesize")
	userid2,err:=strconv.Atoi(userid)
	   if err !=nil{
	  	panic(err)
	   }
	 userpwd2,err:=strconv.Atoi(userpwd)
	 value2,err:=strconv.Atoi(value1)
	 session1,err :=mgo.Dial("localhost")
	 db2 :=session1.DB("userinfo") 
	 cu :=db2.C("usertable")
	 cu.Find(bson.M{"userid":userid2}).All(&User)
	
	if userpwd2 ==	User[0].Password{
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
	}else{
		fmt.Fprintln(w,"身份验证出错")
	}
	
}

