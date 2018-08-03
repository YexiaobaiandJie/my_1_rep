<template>
<div>
    <div class="login_content">
        <input type="text" id="userid" placeholder="userid" class="id_input" />
        <input type="text" id="password" placeholder="password" class="pwd_input" />
        <br />
        <div>
            <button v-on:click="login" class="login_button" v-bind:class="{onbutton:ison}" v-on:mouseover="onthebutton" v-on:mouseout="leavethebutton">login</button>
            <button class="regi_button" v-bind:class="{onbutton:ison2}" v-on:mouseover="onthebutton2" v-on:mouseout="leavethebutton2">register</button>
        </div>
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
            ison2:false
        }
    },
    methods:{
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
        login:function(){
            var idinput=document.getElementById("userid")
            var pwdinput=document.getElementById("password")
            var userid=idinput.value;
            var password=pwdinput.value;
            this.$http.post('http://localhost:3000/login',{"Userid":userid,"Password":password}).then(function(res){
                var date = res.body
                console.log(date)
                this.token=date
                Store.savetoken(date)
            },function(res){
                console.log("error")
            })
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
</style>