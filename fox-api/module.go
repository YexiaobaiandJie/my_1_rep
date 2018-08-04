package main

import(
	"gopkg.in/mgo.v2/bson"
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

type user struct{
	Userid 	 string `json:"userid"`
	Password string `json:"password"`
}

type getuser struct{
	Userid 		string    		`json:"userid"`
	Password 	string	 		`json:"password"`
	Id       	bson.ObjectId 	`bson:"_id"`
}

type posting struct{
	Title   string 	`json:"title"`
	Author  string 	`json:"author"`
	Content string 	`json:"content"`
	Date    int64 	`json:"date"`
}

type postshort struct{
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	Date 	int64  `json:"date"`
} 

type comment struct{
	Author string `json:"author"`
	Date   int64 `json:"date"`
	Userid string `json:"userid"`
	Com    string `json:"com"`
	Time   int64 `json:"time"`
}

type detailcom struct{
	Userid string `json:"userid"`
	Com    string `json:"com"`
	Time   int64  `json:"time"`
}

type detail struct{
	Title   string 	`json:"title"`
	Author  string 	`json:"author"`
	Content string 	`json:"content"`
	Date    int64 	`json:"date"`
	Com 	[]detailcom `json:"com"`
	Comcount int    `json:"comcount"`
}

type detailpost struct{
	Author string `json:"author"`
	Date   int64  `json:"date"`
}

type tokenid struct{
	Userid string `json:"userid"`
	Token  bson.ObjectId 	`bson:"_id"`
}

type token_title_con struct{
	Token bson.ObjectId `bson:"_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}