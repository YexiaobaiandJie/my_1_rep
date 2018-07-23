package main

import(
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	//"os"
	"encoding/json"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
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

type Infos []Info

type Comp struct{
	Data 		Infos 	  `json:"data"`
	PageSize 	int		  `json:"pageSize"`
	TotalItems 	int 	  `json:"totalitems"`
	TotalPages  int   	  `json:"totalpages"`
}

func main(){
	tick := time.Tick(1 * time.Minute)
      for _ = range tick {
            reptile()
      }
}

func reptile (){
	//发送请求，接受请求
	fmt.Println("执行")
	client :=&http.Client{}

	url :="https://api.readhub.me/blockchain?pageSize=20"

	reqest,err :=http.NewRequest("GET",url,nil)

	if err !=nil{
		panic(err)
	}
	//将数据放入结构体
	var s Comp
	response,_ :=client.Do(reqest)
	body,err :=ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &s)
	//stdout :=os.Stdout
	//_,err = io.Copy(stdout,response.Body)
	//fmt.Println(s.TotalItems)
	//fmt.Println(s.Data[0].Summary)
	//将数据导入数据库
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)
	c:=session.DB("foxinfo").C("info")
	for a:=0;a<20;a++{
		err = c.Insert(&Info{
			s.Data[a].AuthorName,
			s.Data[a].Id,
			s.Data[a].Language,
			s.Data[a].MobileUrl,
			s.Data[a].PublishDate,
			s.Data[a].SiteName,
			s.Data[a].Summary,
			s.Data[a].SummaryAuto,
			s.Data[a].Title,
			s.Data[a].Url,
		})
	}
}