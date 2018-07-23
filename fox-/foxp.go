package main

import(
	"fmt"
	
	"io/ioutil"
	"net/http"
	//"os"
	"encoding/json"
	//"gopkg.in/mgo.v2"
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

type Infoslice struct{
	Infos []Info
}

type Comp struct{
	Data 		Infoslice `json:"data"`
	PageSize 	int		  `json:"pageSize"`
	TotalItems 	int 	  `json:"totalitems"`
	TotalPages  int   	  `json:"totalpages"`
}

func main(){
	client :=&http.Client{}

	url :="https://api.readhub.me/blockchain?pageSize=20"

	reqest,err :=http.NewRequest("GET",url,nil)

	if err !=nil{
		panic(err)
	}
	var s Comp
	response,_ :=client.Do(reqest)
	body,err :=ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &s)
	//stdout :=os.Stdout
	//_,err = io.Copy(stdout,response.Body)
	fmt.Println(s.TotalItems)
	
}