<template>
<div class="content">
    <div class="newtitle">
        <h2 class="backnews" v-on:click="loadnews">新闻>></h2>
    </div>
        <hr />
    <div>
        <h2 class="newstitle">{{title}}</h2>
        <h3 class="newsauthor">author:{{author}}</h3>
        <div class="newscontent">{{content}}</div>
        <div class="per" v-on:click="openh">查看全文</div>
        <div class="newsdate">{{date}}</div>
    </div>     
</div>  
</template>

<script>
export default{
    
    data(){
        return{
            message:"this is the newsdetail",
            title:"",
            author:"",
            content:"",
            date:"",
            newsurl:"",
            newsid:"12122"
        }
    },
    methods:{
        openh:function(){
            window.location.href=this.newsurl;
        },
        loadnews:function(){
            this.newsid=this.$route.query.id
             this.$http.get('http://localhost:3000/newsd?id='+this.newsid).then(function(res){
                var data = res.body
                console.log(data)
                this.title=data.title
                this.author=data.authorName
                this.date=data.publishDate
                this.content=data.summaryAuto
                this.newsurl=data.url
                },function(res){
                alert("返回新闻数据出错!")
            })
            
            
            // console.log(this.title)
        }
    },
    mounted:function(){
        this.loadnews()
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
.per{
    color:red;
    margin-left:10%;
    cursor: pointer;
    width:8%;
    float:left;
}
</style>