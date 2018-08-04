<template>
<div class="content">
    <div class="post_title">
        <div class="title_p">{{title}}</div>
        <input id="title" type="text" class="title_input"/>    
    </div>
    <br />
    <div class="post_content">
        <div class="content_p">{{content}}</div>
        <textarea id="content" class="content_area"></textarea>
    </div>
    <div class="post_button">
    <button class="pub_btn" v-on:click="clickbtn" v-bind:class="{onbtn:isclick}">publish</button>
    </div>
</div>
</template>

<script>
import Store from './store.js'
export default{
    data(){
        return{
            title:"Title",
            content:"Content",
            isclick:false
        }
    },
    methods:{
        clickbtn:function(){
            this.isclick=true
            var token_=Store.gettoken()
            var title_post_input=document.getElementById("title")
            var content_post_input=document.getElementById("content")
            var title_post=title_post_input.value
            var content_post=content_post_input.value
            this.$http.post('http://localhost:3000/publish',{Token:token_,Title:title_post,Content:content_post}).then(function(res){
                var data=res.body
                alert(data)
                this.$router.push({path:'/postings'})
            },function(res){
                alert("publish failed")
            })
        }
    }
}
</script>

<style>
.title_p{
    font-size:20px;
    float:left;
    margin-right:10px;
}
.title_input{
    width:300px;
    height:25px;
    font-size:20px;
    margin-left:38px;
    border-style:none;
    border-bottom-style:solid;
    border-bottom-width:thin;
    outline:none;
}
.post_title{
    margin-left:13%;
    margin-top:50px;
}
.post_content{
    margin-left:13%;
    margin-top:12px;
}
.content_p{
    font-size:20px;
    float:left;
    
}
.content_area{
    margin-left:10px;
    width:80%;
    height:500px;
    font-size:19px;
    background-color: #c1c1c1;
    outline:none;
}
.post_button{
    margin-top:25px;
    width:90%;
    text-align: right;
}
.pub_btn{
    width:85px;
    height:33px;
    font-size:18px;
    background-color: #ff4500;
    color:aliceblue;
    border:none;
}
.onbtn{
    color:black;
    background-color:#fff;
    transition: 0.4s;
    outline:none;
}
.onbtn2{
    font-size:21px;
    transition:0.7s;
}
</style>