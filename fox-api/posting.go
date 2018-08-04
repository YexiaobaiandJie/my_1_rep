package main

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
)

//PublishPage 发布帖子
func PublishPage(cq *gin.Context){
	var User []user
	var Posting posting
	var Token_title_con token_title_con
	if cq.ShouldBind(&Token_title_con) == nil{
		token :=Token_title_con.Token
			// token:=bson.ObjectIdHex(temp)
			session,err :=mgo.Dial("localhost")
			if err !=nil{
				panic(err)
			}
			db :=session.DB("userinfo") 
			c :=db.C("usertable")
			c.Find(bson.M{"_id":token}).All(&User)
			if len(User) !=0{
				//当token正确时，发布帖子
				cq.String(200,"token正确，进入发布状态->")
				//将帖子存入数据库中
				// title :=cq.Query("title")
				// content :=cq.Query("content")
				title:=Token_title_con.Title
				content :=Token_title_con.Content
				if(title=="" || content==""){
					cq.String(200,"题目或内容未正确输入，发布失败")
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
					cq.String(200,"发布成功")
				}
			}else{
				//当token错误时，发布失败
				cq.String(200,"token错误，发布失败")
			}
	}
}

//查看所有帖子 展示帖子概况
func PostingPage(cq *gin.Context){
	var Postshort []postshort
	//cq.Header("Access-Control-Allow-Origin", "*")
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db :=session.DB("userinfo")
	c :=db.C("postings")
	c.Find(nil).Sort("-date").Select(bson.M{"title":1,"author":1,"date":1}).All(&Postshort)
	if len(Postshort)!=0{
			cq.JSON(200,Postshort)
	}else{
		cq.String(200,"目前没有任何帖子")
	}
}

//展示帖子细节，包括帖子内容以及评论
func DetailPage(cq *gin.Context){
	//根据时间和作者来确定帖子以及评论
	var Posting []posting
	var Comment []detailcom
	var Detail  detail
	var Detailpost detailpost
	if cq.ShouldBind(&Detailpost) == nil{
		author :=Detailpost.Author
		date :=Detailpost.Date
		if author=="" || date==0{
			cq.String(200,"author or date should not be empty")
		}else{
			date2:=date
			session,err :=mgo.Dial("localhost")
			if err !=nil{
				panic(err)
			}
			db :=session.DB("userinfo")
			c :=db.C("postings")
			c.Find(bson.M{"author":author,"date":date2}).All(&Posting)
			if len(Posting)!=0{
				c1 :=db.C("comments")
				c1.Find(bson.M{"author":author,"date":date2}).Select(bson.M{"userid":1,"com":1,"time":1}).All(&Comment)
				Detail.Title=Posting[0].Title
				Detail.Author=Posting[0].Author
				Detail.Content=Posting[0].Content
				Detail.Date=Posting[0].Date
				Detail.Com=Comment
				Detail.Comcount=len(Comment)
				cq.JSON(200,Detail)
			}else{
				cq.String(200,"找不到相应帖子，请确认作者和时间")
			}
		}
	}
}


//发表评论
func CommentPage(cq *gin.Context){
	var User user
	var fake user
	var Comment comment
	var Posting []posting
	//需要时间和作者来确定帖子
	author :=cq.Query("author")
	date :=cq.Query("date")
	temp :=cq.Query("token")
	com :=cq.Query("com")
	if(author=="" || date=="" || temp=="" || com==""){
		cq.String(200,"author or date or temp or com should not be empty")
	}else if len(temp)==24{
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
			if User==fake{
				cq.String(200,"token错误")
			}else{
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
					cq.String(200,"发表评论成功")
				}
		}else{
			cq.String(200,"不存在对应帖子，请确认作者和时间")
		}		
	}else{
		cq.String(200,"token残缺")
	}
	
}

// func GetComPage(cq *gin.Context){
// 	var Posting []posting
// 	var Comment []detailcom
// 	var Detail  detail
// 	author :=cq.Query("author")
// 	date :=cq.Query("date")
// 	if author=="" || date==""{
// 		cq.String(200,"author or date should not be empty")
// 	}else{
// 		date2,err:=strconv.ParseInt(date, 10, 64)   
// 		session,err :=mgo.Dial("localhost")
// 		if err !=nil{
// 			panic(err)
// 		}
// 		db :=session.DB("userinfo")
// 		c :=db.C("postings")
// 		c.Find(bson.M{"author":author,"date":date2}).All(&Posting)
// 		if len(Posting)!=0{
// 			c1 :=db.C("comments")
// 			c1.Find(bson.M{"author":author,"date":date2}).Select(bson.M{"userid":1,"com":1,"time":1}).All(&Comment)
// 			Detail.Title=Posting[0].Title
// 			Detail.Author=Posting[0].Author
// 			Detail.Content=Posting[0].Content
// 			Detail.Date=Posting[0].Date
// 			Detail.Com=Comment
// 			cq.JSON(200,Detail)
// 		}else{
// 			cq.String(200,"找不到相应帖子，请确认作者和时间")
// 		}
// 	}
// }