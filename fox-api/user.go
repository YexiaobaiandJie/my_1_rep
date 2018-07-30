package main

import(
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


// register 注册页，注册账号，注册成功之后会返回token
func RegisterPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var getUser []getuser
	newid :=r.URL.Query().Get("userid")
	newpwd :=r.URL.Query().Get("userpwd")
	if newid=="" || newpwd=="" {
		fmt.Fprintln(w,"userid or userpwd should not be empty")
	}else{
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
}
	

// LoginPage   登录后返回token
func LoginPage(w http.ResponseWriter,r *http.Request){
	var User []user
	var getUser []getuser
	userid  :=r.URL.Query().Get("userid")
	userpwd :=r.URL.Query().Get("userpwd")
	if userid=="" || userpwd=="" {
		fmt.Fprintln(w,"userid or userpwd should not be empty")
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
			fmt.Fprintln(w,"身份验证成功")
			fmt.Fprintln(w,getUser[0].Id)
		}else{
			fmt.Fprintln(w,"身份验证出错")
		}
	}
}