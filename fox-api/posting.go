package main

import(
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"time"
)

//PublishPage 发布帖子
func PublishPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var Posting posting
	var Postshort postshort
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
	if author=="" || date==""{
		fmt.Fprintln(w,"author or date should not be empty")
	}else{
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
	if(author=="" || date=="" || temp=="" || com==""){
		fmt.Fprintln(w,"author or date or temp or com should not be empty")
	}else{
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
	
}