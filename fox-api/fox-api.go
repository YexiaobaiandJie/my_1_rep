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

type postshort struct{
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	Date 	string `json:"date"`
} 

type comment struct{
	Author string `json:"author"`
	Date   string `json:"date"`
	Userid string `json:"userid"`
	Com    string `json:"com"`
	Time   string `json:"time"`
}



func main(){
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/news",Index)
	router.HandleFunc("/register",RegisterPage)
	router.HandleFunc("/login",LoginPage)
	router.HandleFunc("/publish",PublishPage)
	router.HandleFunc("/postings",PostingPage)
	router.HandleFunc("/detail",DetailPage)
	router.HandleFunc("/comment",CommentPage)
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
	var Postshort postshort
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
			Postshort.Title=title
			Postshort.Author=User[0].Userid
			Postshort.Date=date
			session1,err1 :=mgo.Dial("localhost")
			if err1 !=nil{
				panic(err1)
			}
			db1 :=session1.DB("userinfo")
			c1 :=db1.C("postings")
			c2 :=db1.C("postshort")
			c1.Insert(Posting)
			c2.Insert(Postshort)
			fmt.Fprintln(w,"发布成功")
		}
	}else{
		//当token错误时，发布失败
		fmt.Fprintln(w,"token错误，发布失败")
	}
}

//查看所有帖子 展示帖子概况
func PostingPage(w http.ResponseWriter,r *http.Request){
	var postshort []postshort
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo")
	c :=db.C("postings")
	c.Find(nil).Sort("-date").All(&postshort)
	w.Header().Set("Content-type","application/json")
	jsons,err:=json.Marshal(postshort)
	if err !=nil{
		panic(err)
	}
	fmt.Fprintln(w,string(jsons))
}

//展示帖子细节，包括帖子内容以及评论
func DetailPage(w http.ResponseWriter,r *http.Request){
	//根据时间和作者来确定帖子以及评论
	var Posting []posting
	var Comment []comment
	author :=r.URL.Query().Get("author")
	date :=r.URL.Query().Get("date")
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo")
	c :=db.C("postings")
	c.Find(bson.M{"author":author,"date":date}).All(&Posting)
	if len(Posting)!=0{
		//session1,err :=mgo.Dial("hostlocal")
		//db1 :=session1.DB("userinfo")
		c1 :=db.C("comments")
		c1.Find(bson.M{"author":author,"date":date}).All(&Comment)
		json1,err :=json.Marshal(Posting)
		if err!=nil{
			panic(err)
		}
		json2,err :=json.Marshal(Comment)
		if err !=nil{
			panic(err)
		}
		fmt.Fprintln(w,string(json1))
		
		fmt.Fprintln(w,string(json2))
	}else{
		fmt.Fprintln(w,"找不到对应帖子，请确认作者和时间")
	}
}


//发表评论
func CommentPage(w http.ResponseWriter,r *http.Request){
	var User user
	var Comment comment
	//需要时间和作者来确定帖子
	author :=r.URL.Query().Get("author")
	date :=r.URL.Query().Get("date")
	temp :=r.URL.Query().Get("token")
	com :=r.URL.Query().Get("com")
	dtime :=time.Now()
	timenow := dtime.Format("2006-01-02 15:04:05")
	token :=bson.ObjectIdHex(temp)
	//根据token查找用户名
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo")
	c  :=db.C("usertable")
	c.Find(bson.M{"_id":token}).One(&User)
	//将评论存入数据库
	session1,err1 :=mgo.Dial("localhost")
	if err1 !=nil{
		panic(err)
	}
	db1 :=session1.DB("userinfo")
	c1  :=db1.C("comments")
	Comment.Author=author
	Comment.Date=date
	Comment.Userid=User.Userid
	Comment.Com=com
	Comment.Time=timenow
	c1.Insert(Comment)
	fmt.Fprintln(w,"发表评论成功")
}