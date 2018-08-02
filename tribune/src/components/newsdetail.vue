<template>
<div>
    <h2 class="title">{{title}}</h2>
    <h3 class="author">{{author}}</h3>
    <div class="content"><p>{{content}}</p></div>
    <div class="date">{{date}}</div>
</div>
</template>


<script>
export default{
    data(){
        return{
            message:"this is the detail of news",
            title:"这是一个题目",
            author:"作者",
            content:"这是新闻内容",  
            date:"2018-01-12"
        }
    }
}
</script>


<style>
.title{
    text-align:center;
}
.author{
    text-align: right;
    padding-right:8%;
}
.content{
    width:80%;
    height:100%;
    margin-left:10%;
    font-size:18px;
}
.date{
    text-align: right;
    padding-right:8%;
}
</style>