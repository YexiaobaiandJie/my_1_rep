<template>
<div class="content">
    <div class="post_title">
        <div class="title_p">{{title}}</div>
        <input type="text" class="title_input"/>    
    </div>
    <br />
    <div class="post_content">
        <div class="content_p">{{content}}</div>
        <textarea class="content_area"></textarea>
    </div>
</div>
</template>

<script>
export default{
    data(){
        return{
            title:"Title",
            content:"Content"
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
    border-style:none;
    border-bottom-style:solid;
    border-bottom-width:thin;
}
.post_title{
    margin-left:20%;
    margin-top:50px;
}
.post_content{
    margin-left:20%;
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
    font-size:15px;
    background-color: #cfcfcf;
}
</style>