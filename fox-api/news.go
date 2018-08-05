package main

import(
	//"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)


//   /news   输入新闻数量及token，若通过验证则返回新闻，否则报错
func Index(cq *gin.Context){
	var User []user
	//通过token取出对应的数据，否则报错
	//cq.Header("Access-Control-Allow-Origin", "*")
	temp :=cq.Query("token")
	if temp==""{
		cq.JSON(400,"token should not be empty")
	}else if len(temp)==24{
		token:=bson.ObjectIdHex(temp)
		session1,err :=mgo.Dial("localhost")
		if err !=nil{
			panic(err)
		}
		db2 :=session1.DB("userinfo") 
		cu :=db2.C("usertable")
		cu.Find(bson.M{"_id":token}).All(&User)
		if len(User) !=0{
			value1 :=cq.Query("pagesize")
			if value1==""{
				cq.JSON(400,"pagesize should not be empty")
			}else{
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
				if err !=nil{
					panic(err)
				}
				cq.JSON(200,infos)
			}
		}else{
			cq.JSON(400,"身份验证错误")
		}
	}else{
		cq.JSON(400,"token 残缺")
	}
}

func NewsdetailPage(cq *gin.Context){
	var infos Info
	var News_comment []news_comment_save
	var Newsdetail newsdetail 
	temp :=cq.Query("id")
	id,err:=strconv.Atoi(temp)
	
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db:=session.DB("foxinfo")
	c:=db.C("info")
	c.Find(bson.M{"id":id}).One(&infos)
	// cq.JSON(200,infos)
	
	//新增返回新闻评论
	c2 := db.C("newscomment")
	c2.Find(bson.M{"newsid":id}).Sort("-date").All(&News_comment)
	Newsdetail.Newsinfo=infos
	Newsdetail.NewsComment=News_comment
	cq.JSON(200,Newsdetail)
}

//提交对新闻评论
func NewsComment (cq *gin.Context) {
	var User []user
	var newscomment news_comment
	if err10 :=cq.ShouldBind(&newscomment);err10 == nil{
		newsid := newscomment.Newsid
		newscom := newscomment.Com
		token := newscomment.Token
		//根据token获得评论用户的用户id
		session,err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		db := session.DB("userinfo")
		c := db.C("usertable")
		c.Find(bson.M{"_id":token}).All(&User)
		if len(User) == 0 {  //如果没有找到对应用户
			cq.JSON(400,"token is wrong")
		} else {
			userid := User[0].Userid
			db2 := session.DB("foxinfo")
			c2  := db2.C("newscomment")
			//获得当前时间，为防止之后在前端无法转换回通用格式
			//时间使用字符串形式存入数据库
			date:=time.Now().Format("2006-01-02 15:04:05")
			c2.Insert(&news_comment_save{Newsid:newsid,Com:newscom,Userid:userid,Date:date})
			cq.JSON(200,"新闻评论存储成功")
		}

	}else{
		fmt.Println(err10)
		fmt.Println("未获得数据")
	}
}