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
	"time"
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
	Userid 		string    		`json:"userid"`
	Password 	string	 		`json:"password"`
	Id       	bson.ObjectId 	`bson:"_id"`
}

type posting struct{
	Title string 	`json:"title"`
	Author string 	`json:"author"`
	Content string 	`json:"content"`
	Date string 	`json:"date"`
}



func main(){
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news",Index)
	router.HandleFunc("/register",RegisterPage)
	router.HandleFunc("/login",LoginPage)
	router.HandleFunc("/publish",PublishPage)
	log.Fatal(http.ListenAndServe(":8080",router))
}


//   /news   输入新闻数量及token，若通过验证则返回新闻，否则报错
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
		fmt.Fprintln(w,"身份验证错误")
	}
}

// register 注册页，注册账号，注册成功之后会返回token
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

// LoginPage   登录后返回token
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


//PublishPage 发布帖子
func PublishPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var Posting posting
	temp :=r.URL.Query().Get("token")
	token:=bson.ObjectIdHex(temp)
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo") 
	c :=db.C("usertable")
	c.Find(bson.M{"_id":token}).All(&User)
	if len(User) !=0{
		//当token正确时，发布帖子
		fmt.Fprintln(w,"token正确，进入发布状态")
		//将帖子存入数据库中
		title :=r.URL.Query().Get("title")
		content :=r.URL.Query().Get("content")
		if(title=="" || content==""){
			fmt.Fprintln(w,"题目或内容未正确输入，发布失败")
		}else{
			dtime :=time.Now()
			date := dtime.Format("2006-01-02 15:04:05")
			Posting.Title=title
			Posting.Author=User[0].Userid
			Posting.Content=content
			Posting.Date=date
			session1,err1 :=mgo.Dial("localhost")
			if err1 !=nil{
				panic(err1)
			}
			db1 :=session1.DB("userinfo")
			c1 :=db1.C("postings")
			c1.Insert(Posting)
			fmt.Fprintln(w,"发布成功")
		}
	}else{
		//当token错误时，发布失败
		fmt.Fprintln(w,"token错误，发布失败")
	}
}