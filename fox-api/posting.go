package main

import(
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"time"
	"strconv"
)

//PublishPage 发布帖子
func PublishPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var Posting posting
	temp :=r.URL.Query().Get("token")
	if temp ==""{
		fmt.Fprintln(w,"token should not be empty")
	}else{
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
				date := dtime.Unix()
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
	
}

//查看所有帖子 展示帖子概况
func PostingPage(w http.ResponseWriter,r *http.Request){
	var Postshort []postshort
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo")
	c :=db.C("postings")
	c.Find(nil).Sort("-date").Select(bson.M{"title":1,"author":1,"date":1}).All(&Postshort)
	if len(Postshort)!=0{
		w.Header().Set("Content-type","application/json")
		jsons,err:=json.Marshal(Postshort)
		if err !=nil{
			panic(err)
		}
		fmt.Fprintln(w,string(jsons))
	}else{
		fmt.Fprintln(w,"目前没有任何帖子")
		
	}
	
	
}

//展示帖子细节，包括帖子内容以及评论
func DetailPage(w http.ResponseWriter,r *http.Request){
	//根据时间和作者来确定帖子以及评论
	var Posting []posting
	var Comment []comment
	author :=r.URL.Query().Get("author")
	date :=r.URL.Query().Get("date")
	if author=="" || date==""{
		fmt.Fprintln(w,"author or date should not be empty")
	}else{
		date2,err:=strconv.ParseInt(date, 10, 64)   
		session,err :=mgo.Dial("localhost")
		if err !=nil{
			panic(err)
		}
		db :=session.DB("userinfo")
		c :=db.C("postings")
		c.Find(bson.M{"author":author,"date":date2}).All(&Posting)
		if len(Posting)!=0{
			c1 :=db.C("comments")
			c1.Find(bson.M{"author":author,"date":date2}).All(&Comment)
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
	
}


//发表评论
func CommentPage(w http.ResponseWriter,r *http.Request){
	var User user
	var Comment comment
	var Posting []posting
	//需要时间和作者来确定帖子
	author :=r.URL.Query().Get("author")
	date :=r.URL.Query().Get("date")
	temp :=r.URL.Query().Get("token")
	com :=r.URL.Query().Get("com")
	if(author=="" || date=="" || temp=="" || com==""){
		fmt.Fprintln(w,"author or date or temp or com should not be empty")
	}else{
		session2,err :=mgo.Dial("localhost")
		if err !=nil{
			panic(err)
		}
		db2 :=session2.DB("userinfo")
		c2 :=db2.C("postings")
		date2,err:=strconv.ParseInt(date, 10, 64)
		c2.Find(bson.M{"author":author,"date":date2}).All(&Posting)
		if len(Posting)!=0{
			  
			dtime :=time.Now()
			timenow := dtime.Unix()
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
			Comment.Date=date2
			Comment.Userid=User.Userid
			Comment.Com=com
			Comment.Time=timenow
			c1.Insert(Comment)
			fmt.Fprintln(w,"发表评论成功")
		}else{
			fmt.Fprintln(w,"不存在对应帖子，请确认作者和时间")
		}





		
	}
	
}