package main

import(
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
)


// register 注册页，注册账号，注册成功之后会返回token
func RegisterPage(cq *gin.Context){
	var User []user
	var getUser []getuser
	
	newid :=cq.Query("userid")
	
	newpwd :=cq.Query("userpwd")
	if newid=="" || newpwd=="" {
		cq.String(http.StatusOK,"userid or userpwd should not be empty")
	}else{
		session,err :=mgo.Dial("localhost")
		if err !=nil{
			panic(err)
		}
		db1 :=session.DB("userinfo")
		c1 :=db1.C("usertable")
		err=c1.Find(bson.M{"userid":newid}).All(&User)
		if len(User) != 0{
			cq.String(http.StatusOK,"该id已存在")
		}else{
			c1.Insert(&user{Userid:newid,Password:newpwd})
			c1.Find(bson.M{"userid":newid}).Select(bson.M{"_id":1,"userid":1,"password":1}).All(&getUser)
			cq.String(http.StatusOK,"注册成功")
			cq.JSON(200,getUser[0].Id)
		}
	}
}
	

// LoginPage   登录后返回token
func LoginPage(cq *gin.Context){
	var User []user
	var getUser []getuser
	userid :=cq.Query("userid")
	userpwd :=cq.Query("userpwd")
	if userid=="" || userpwd=="" {
		cq.String(http.StatusOK,"userid or userpwd should not be empty")
	}else{
		session,err :=mgo.Dial("localhost")
		if err !=nil{
			panic(err)
		}
		db :=session.DB("userinfo")
		c :=db.C("usertable")
		c.Find(bson.M{"userid":userid}).All(&User)
		c.Find(bson.M{"userid":userid}).Select(bson.M{"_id":1,"userid":1,"password":1}).All(&getUser)
		if userpwd ==	User[0].Password{
			cq.String(http.StatusOK,"身份验证成功")
			cq.JSON(200,getUser[0].Id)
		}else{
			cq.String(http.StatusOK,"身份验证出错")
		}
	}
}