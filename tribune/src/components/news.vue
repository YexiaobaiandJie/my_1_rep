<template>
    <div class="content">
        <div class="newtitle">
            <h2 v-on:click="backnews" class="backnews">新闻>></h2>
        </div>
        <hr />
        <div v-show="isShow">
            <div  v-for="newitem in newitems" class="newstyle">
                <div class="new-l"  v-on:click="clicknews(newitem)">{{newitem.title}}</div>
                <div class="new-r">{{newitem.publishDate}}</div>
            </div>
        </div>
        <div v-show="!isShow">
            <h2 class="newstitle">{{title}}</h2>
            <h3 class="newsauthor">{{author}}</h3>
            <div class="newscontent"><p>{{content}}</p></div>
            <div class="per" v-on:click="openh">查看全文</div>
            <div class="newsdate">{{date}}</div>
            
        </div>



    </div>
</template>


<script>
export default{
    mounted:function(){
        this.loadnews()
    },
    data(){
        return{
            message:"this is the news component",
            newitems:[],
            isShow:true,
            title:"title",
            author:"author",
            content:"content",
            date:"2018-08-01",
            newsurl:"http://www.bishijie.com/home/newsflashpc/detail?id=85995"
        }
    },
    methods:{
        loadnews:function(){
            this.$http.get('http://localhost:3000/news?token=5b5ff5794cf88aed639018e0&pagesize=30').then(response => {
                this.newitems=response.data;
            },response => {
                console.log("error");
            });
        },
        clicknews:function(newitem){
            this.isShow=!this.isShow
            this.title=newitem.title
            this.author=newitem.authorName
            this.date=newitem.publishDate
            this.content=newitem.summary
            this.newsurl=newitem.url
        },
        backnews:function(){
            this.isShow=true
        },
        openh:function(){
            window.location.href=this.newsurl;
        }
    }
}
</script>

<style scoped>
.content{
    width:90%;
    margin-left:5%;
    height:100%;
    
}
.newtitle{
    text-align:left;
    padding-left:10px;
    margin-bottom:-10px;
    margin-top:-2px;
   
}
.newstyle{
    width:100%;
    
}
.new-l{
   text-align: left;
   height:10%;
   width:70%;
   float: left;
   padding-left:1%;
   font-size:120%;
   padding-bottom:0.4%;
   padding-top:0.4%;
   color:black;
   text-decoration: none;
   cursor: pointer;
}
.new-r{
   text-align:right;
   height:10%;
   width:27%;
   float: left;
   font-size:120%;
   padding-right:15px;
   padding-bottom: 0.4%;
   padding-top:0.4%;
   cursor: default;
}
.newstitle{
    text-align:center;
}
.newsauthor{
    text-align: right;
    padding-right:8%;
}
.newscontent{
    width:80%;
    height:100%;
    margin-left:10%;
    font-size:18px;
}
.newsdate{
    padding-right:10%;
    width:50%;
    float:right;
    text-align: right;
}
.backnews{
    cursor: pointer;
}
.per{
    color:red;
    margin-left:10%;
    cursor: pointer;
    width:8%;
    float:left;
}
</style>