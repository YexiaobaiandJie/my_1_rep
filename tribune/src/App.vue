<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
    <!-- <HelloWorld/> -->
    <div>
      <h2>论坛</h2>
    </div>
    <div class="head">
      <router-link to="/news" class="a">news</router-link>
      <router-link to="/postings" class="a">postings</router-link>
      <router-link to="/news" class="a">login</router-link>
    </div>
    <div class="underline"></div>
    <router-view></router-view>
    
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import News from './components/news'
import Postings from './components/postings'

export default {
  name: 'App',
  components: {
    HelloWorld,
    News,
    Postings
  },
  data(){
    return{
      
    }
  },
  methods:{
    

  }
  }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  min-width:944px;
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
  margin-left:5%;
  margin-right:5%;
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
