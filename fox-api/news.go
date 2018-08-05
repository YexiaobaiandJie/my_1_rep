package main

import(
	//"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"github.com/gin-gonic/gin"
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
	temp :=cq.Query("id")
	id,err:=strconv.Atoi(temp)
	
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	db:=session.DB("foxinfo")
	c:=db.C("info")
	c.Find(bson.M{"id":id}).One(&infos)
	cq.JSON(200,infos)
}