<template>
<div>
    <div class="login_content" v-show="!islogin">
        <input type="text" id="userid" placeholder="userid" class="id_input" />
        <input type="password" id="password" placeholder="password" class="pwd_input" v-on:keyup.enter="login"/>
        <br />
        <div>
            <button v-on:click="login" class="login_button" v-bind:class="{onbutton:ison}" v-on:mouseover="onthebutton" v-on:mouseout="leavethebutton">login</button>
            <button v-on:click="register" class="regi_button" v-bind:class="{onbutton:ison2}" v-on:mouseover="onthebutton2" v-on:mouseout="leavethebutton2">register</button>
        </div>
    </div>
    <div v-show="islogin" class="loginsu">
        <div class="welcome">欢迎你,{{userid}}</div>
        <button class="logout_btn" v-on:click="logout" v-bind:class="{onbutton2:ison3}" v-on:mouseover="onthebutton3" v-on:mouseout="leavethebutton3">log-out</button>
    </div>
</div>
</template>


<script>
import Store from './store'
export default{
    data(){
        return{
            token:1,
            ison:false,
            ison2:false,
            ison3:false,
            islogin:false,
            userid:""
        }
    },
    mounted:function(){
            this.checktoken()
            this.checkstatus()
            this.userid=Store.getuserid()
    },
    methods:{
        logout:function(){
            this.islogin=false
            Store.savetoken()
            Store.saveuseid()
            Store.setlogin(false)
            this.$parent.change_login_title2()  //切换导航栏
        },
        checkstatus:function(){
            var status = Store.getstatus()
            if(status == "true")    //当浏览器内存中显示状态为登录时，从内存中取出userid显示在导航栏中
                this.islogin=true
        },
        checktoken:function(){
            this.token=Store.gettoken()
            
        },
        onthebutton:function(){
            this.ison=true
        },
        leavethebutton:function(){
            this.ison=false
        },
        onthebutton2:function(){
            this.ison2=true
        },
        leavethebutton2:function(){
            this.ison2=false
        },
        onthebutton3:function(){
            this.ison3=true
        },
        leavethebutton3:function(){
            this.ison3=false
        },
        login:function(){
            var idinput=document.getElementById("userid")
            var pwdinput=document.getElementById("password")
            var userid=idinput.value;
            var password=pwdinput.value;
            this.$http.post('http://localhost:3000/login',{"Userid":userid,"Password":password}).then(function(res){
                var data = res.body
                console.log(data)
                if(data=="userid or userpwd should not be empty"){
                    alert("userid or userpwd should not be empty")
                }else if(data=="身份验证出错"){
                    alert("身份验证出错")
                }else{
                    this.token=data.Token
                    Store.savetoken(data.Token)
                    Store.saveuseid(data.userid)
                    
                    Store.setlogin(true)
                    this.islogin=true
                    this.userid=data.userid
                    this.$parent.change_login_title()
                }
                
            },function(res){
                alert("username or password incorrect!")
            })
        },
        register:function(){
            this.$router.push({path:'/login/register'})
        }
    }
}
</script>


<style>
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
    margin-top:2%;
    width:250px;
    height:25px;
    margin-left:4px;
    border-style:none;
    border-bottom-style:solid;
    border-bottom-width:thin;
    color:black;
    font-size:20px;
    outline:none;
}
.login_content{
    margin-top:10%;
    padding-left:34%;
    padding-right:34%;
    text-align: center;
}
input::-webkit-input-placeholder{
    color:#cfcfcf;
}
.login_button{
    width:80px;
    height:30px;
    font-size:18px;
    margin-top:45px;
    margin-right:40px;
    background:none;
    border:none;
    outline:none;
}
.regi_button{
    width:80px;
    height:30px;
    /* margin-left:3px; */
    font-size:18px;
    background:none;
    border:none;
    outline:none;
}
.onbutton{
    border-radius: 0.3em;
    background-color:#ff4500;
    color:aliceblue;
    transition: 0.5s;
}
.onbutton2{  /*这里比较奇怪的是我没办法把背景颜色换掉*/
    color:#ff4500;
    transition: 0.5s;
}
.loginsu{
    text-align: center;
    margin-top:15%;
}
.welcome{
    font-size:30px;
    margin-bottom: 19px;
}
.logout_btn{
    font-size:20px;
    background:none;
    border:none;
    outline:none;
}
</style>