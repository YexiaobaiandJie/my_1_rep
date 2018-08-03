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
import Bus from './bus.js'
export default{
    props:['message'],
    mounted:function(){
        // console.log("B is running")
        
    },
    data(){
        return{
            message:"",
            title:"",
            author:"",
            content:"",
            date:"",
            newsurl:""

        }
    },
    methods:{
        openh:function(){
            window.location.href=this.newsurl;
        },
        loadnews:function(){
            // console.log("newsdetail running")
            // Bus.$on('val',(data)=>{
            // console.log("bus is running")
            // console.log(data)
            // this.title=data.title
            // this.author=data.authorName
            // this.date=data.publishDate
            // this.content=data.summary
            // this.newsurl=data.url
            // console.log(this.title)
            console.log(this.message)
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