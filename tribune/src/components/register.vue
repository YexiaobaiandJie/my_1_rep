<template>
<div class="content">
    <div class="title">Create a New Account</div>
    <input type="text" id="userid" placeholder=">userid<" class="id_input"/>
    <input type="password" id="password" placeholder=">password<" class="pwd_input"/>
    <br />
    <button class="cr_btn" v-bind:class="{onbtn:ison}" v-on:click="register" v-on:mouseover="moveover_btn" v-on:mouseout="moveout_btn">CREATE</button>
</div>
</template>

<script>
export default{
    data(){
        return{
            message:"this is the register",
            ison:false
        }
    },
    methods:{
        moveover_btn:function(){
            this.ison=true
        },
        moveout_btn:function(){
            this.ison=false
        },
        register:function(){
            var idinput=document.getElementById("userid")
            var pwdinput=document.getElementById("password")
            var userid=idinput.value;
            var password=pwdinput.value;
            this.$http.post('http://localhost:3000/register',{Userid:userid,Password:password}).then(function(res){
                var data=res.body
                alert(data)
                if(data == "注册成功"){
                    this.$router.push({path:'/login'})
                }
            },function(res){
                alert("register failed")
            })
        }
    }
}
</script>

<style>
.content{
    text-align: center;
    margin-top:10%;
    padding-left:34%;
    padding-right:34%;
}
.id_input{
    width:250px;
    height:25px;
    border-style:none;
    border-bottom-style:solid;
    border-bottom-width:thin;
    color:black;
    font-size:20px;
    outline:none;
}
.pwd_input{
    margin-top:10px;
    margin-left:4px;
    width:250px;
    height:25px;
    border-style:none;
    border-bottom-style:solid;
    border-bottom-width:thin;
    color:black;
    font-size:20px;
    outline:none;
}
input::-webkit-input-placeholder{
    color:#cfcfcf;
}
.title{
    font-size:32px;
    margin-bottom: 16px;
}
.cr_btn{
    margin-top:40px;
    width:100px;
    height:35px;
    background-color: #fff;
    border:none;
    color:#cfcfcf;
    border-radius: 0.6em;
    font-size:15px;
}
.onbtn{
    color:black;
    background-color:#fff;
    font-size:20px;
    transition: 0.5s;
    font-family:sans-serif;
}
</style>