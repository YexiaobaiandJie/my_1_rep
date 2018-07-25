package main

import(
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"strconv"
)
var newid string
var temp string

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
	fmt.Println("开始")
	//发送请求，接受请求
	client :=&http.Client{}
	url :="https://api.readhub.me/blockchain?pageSize=20"
	reqest,err :=http.NewRequest("GET",url,nil)
	if err !=nil{
		panic(err)
	}
	response,_ :=client.Do(reqest)
	body,err :=ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	tick := time.Tick(time.Minute*13)
    for _ = range tick {
            reptile(body)
	}
}

func reptile (body []byte){
	//读取最新数据的id
	fmt.Println("--------------------------时隔13分钟，重新抓取--------------------------")
	b, err := ioutil.ReadFile("newid.txt")
    if err != nil {
        fmt.Print(err)
    }
	newid= string(b)
	fmt.Println("当前最新新闻id为"+newid)
	var s Comp
	json.Unmarshal(body, &s)
	//将数据导入数据库
	session,err :=mgo.Dial("localhost")
	if err !=nil{
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)
	c:=session.DB("foxinfo").C("info")
	a:=0
	for a=0;a<20;a++{
		thisid:=strconv.Itoa(s.Data[a].Id)
		if a==0 {
			temp=thisid
		}
		if(thisid==newid){
			fmt.Println("此次抓取已停止")
			break
		}else{
				fmt.Println("a=",a)
					err = c.Insert(s.Data[a])
				fmt.Println("存入一条数据")
			}	
	}
	if a==20{
		fmt.Println("API返回数据全部为新数据，已全部存入数据库")
	}
	//将新的最新新闻id存入文件
	d1 :=[]byte(temp)
	ioutil.WriteFile("newid.txt",d1,0644)
	fmt.Println("---------------------》》最新新闻id为"+temp+"《《----------------------")
}
