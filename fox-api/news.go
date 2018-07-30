package main

import(
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"strconv"
)


//   /news   输入新闻数量及token，若通过验证则返回新闻，否则报错
func Index(w http.ResponseWriter,r *http.Request){
	var User []user
	//通过token取出对应的数据，否则报错
	temp :=r.URL.Query().Get("token")
	if temp==""{
		fmt.Fprintln(w,"token should not be empty")
	}else{
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
			if value1==""{
				fmt.Fprintln(w,"pagesize should not be empty")
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
				w.Header().Set("Content-type","application/json")
				jsons,err :=json.Marshal(infos)
				if err !=nil{
					panic(err)
				}
				fmt.Fprintln(w,string(jsons))
			}
			
		}else{
			fmt.Fprintln(w,"身份验证错误")
		}
	}
}