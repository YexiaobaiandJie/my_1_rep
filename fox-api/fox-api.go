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
	Userid 	 string `json:"userid"`
	Password string `json:"password"`
}

type getuser struct{
	Userid string    `json:"userid"`
	Password string	 `json:"password"`
	Id       bson.ObjectId `bson:"_id"`
}
//islogin :=false    
				//false表示未登录
				//true 表示登录

func main(){
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news",Index)
	router.HandleFunc("/register",RegisterPage)
	router.HandleFunc("/login",LoginPage)
	log.Fatal(http.ListenAndServe(":8080",router))
}


//   /news
func Index(w http.ResponseWriter,r *http.Request){
	var User []user
	//通过token取出对应的数据，否则报错
	temp :=r.URL.Query().Get("token")
	token:=bson.ObjectIdHex(temp)
	session1,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db2 :=session1.DB("userinfo") 
	cu :=db2.C("usertable")
	cu.Find(bson.M{"_id":token}).All(&User)
	if len(User) !=0{
		value1 :=r.URL.Query().Get("pagesize")
		value2,err:=strconv.Atoi(value1)
		if err !=nil{
			panic(err)
		}
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
// register
func RegisterPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var getUser []getuser
	newid :=r.URL.Query().Get("userid")
	newpwd :=r.URL.Query().Get("userpwd")
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db1 :=session.DB("userinfo")
	c1 :=db1.C("usertable")
	err=c1.Find(bson.M{"userid":newid}).All(&User)
	
	if len(User) != 0{
		fmt.Fprintln(w,"该userid已存在")
	}else{
		c1.Insert(&user{Userid:newid,Password:newpwd})
		c1.Find(bson.M{"userid":newid}).Select(bson.M{"_id":1,"userid":1,"password":1}).All(&getUser)
		fmt.Fprintln(w,"注册成功")
		fmt.Fprintln(w,getUser[0].Id)
	}
}

// LoginPage
func LoginPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var getUser []getuser
	
	userid  :=r.URL.Query().Get("userid")
	userpwd :=r.URL.Query().Get("userpwd")
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo")
	c :=db.C("usertable")
	c.Find(bson.M{"userid":userid}).All(&User)
	c.Find(bson.M{"userid":userid}).Select(bson.M{"_id":1,"userid":1,"password":1}).All(&getUser)
	if userpwd ==	User[0].Password{
		fmt.Fprintln(w,"身份验证成功")
		fmt.Fprintln(w,getUser[0].Id)
	}else{
		fmt.Fprintln(w,"身份验证出错")
	}
}