<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
    <!-- <HelloWorld/> -->
    <div class="head">
      <div class="b" v-on:click="backtonews">{{head}}</div>
      <router-link to="/news" class="a">news</router-link>
      <router-link to="/postings" class="a">postings</router-link>
      <router-link to="/login" class="a">{{login}}</router-link>
    </div>
    <div class="underline"></div>
    <router-view></router-view>
    
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import News from './components/news'
import Postings from './components/postings'
import Store from './components/store'

export default {
  name: 'App',
  components: {
    HelloWorld,
    News,
    Postings
  },
  mounted:function(){
    var status = Store.getstatus()
    console.log(status)
    if(status == "true"){     //当浏览器内存中显示状态为登录时，从内存中取出userid显示在导航栏中
      var userid = Store.getuserid()
      this.login=userid

    }
  },
  data(){
    return{
      head:"论坛",
      login:"login/register"
    }
  },
  methods:{
    backtonews:function(){
      this.$router.push({path:'/news'})
    },
    change_login_title:function(){
      var userid = Store.getuserid()
      this.login=userid
    },
    change_login_title2:function(){
      this.login="login/register"
    }
  }
  }
</script>

<style>
#app {
  /* font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale; */
  /* text-align: center; */
  color: #2c3e50;
  min-width:1225px;
  /* margin-top: 60px; */
} 
.head{
  background-color: #555;
  padding-top:1%;
  padding-bottom:1%;
  font-size:130%;
}
/* .s1{
  color:aliceblue;
  margin-left:5%;
  margin-right:5%;
  cursor:default;
} */
.a{
  color:aliceblue;
  text-decoration:none;
  margin-left:2.5%;
  margin-right:2.5%;
  
}
.b{
  color:aliceblue;
  text-decoration:none;
  margin-left:5%;
  margin-right:55%;
  float:left;
  cursor: pointer;
}
.here{
  color:#ff4500;
  margin-left:5%;
  margin-right:5%;
  cursor: defalut;
  text-decoration: underline ;
}
.underline{
  width:100%;
  height:6px;
  background-color: #ff4500;
}
</style>
