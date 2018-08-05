package main

import(
	
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
)


// register 注册页，注册账号，注册成功之后会返回token
func RegisterPage(cq *gin.Context){
	var User []user
	var getUser []getuser
	var user2 user
	// newid :=cq.Query("userid")
	// newpwd :=cq.Query("userpwd")
	if cq.ShouldBind(&user2) == nil{
		newid :=user2.Userid
		newpwd :=user2.Password
		if newid=="" || newpwd=="" {
			cq.JSON(400,"userid or userpwd should not be empty")
		}else{
			session,err :=mgo.Dial("localhost")
			if err !=nil{
				panic(err)
			}
			db1 :=session.DB("userinfo")
			c1 :=db1.C("usertable")
			err=c1.Find(bson.M{"userid":newid}).All(&User)
			if len(User) != 0{
				cq.JSON(400,"该id已存在")
			}else{
				c1.Insert(&user{Userid:newid,Password:newpwd})
				c1.Find(bson.M{"userid":newid}).Select(bson.M{"_id":1,"userid":1,"password":1}).All(&getUser)
				cq.JSON(200,"注册成功")
				
			}
		}
	}
	
}
	

// LoginPage   登录后返回token
func LoginPage(cq *gin.Context){
	var User []user
	var getUser []getuser
	var user2 user
	var Tokenid tokenid
	//cq.Header("Access-Control-Allow-Origin", "*")
	if cq.ShouldBind(&user2) == nil{
		userid :=user2.Userid
		userpwd :=user2.Password
		if userid=="" || userpwd=="" {
			cq.JSON(400,"userid or userpwd should not be empty")
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
				// cq.String(http.StatusOK,"身份验证成功")
				Tokenid.Userid=getUser[0].Userid
				Tokenid.Token=getUser[0].Id
				cq.JSON(200,Tokenid)
			}else{
				cq.JSON(400,"身份验证出错")
			}
		}
	}
	
}