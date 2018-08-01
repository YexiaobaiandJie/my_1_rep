<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
    <!-- <HelloWorld/> -->
    <div>
      <h2>论坛</h2>
    </div>
    <div class="head">
      <span class="s1" to="/news" v-bind:class="{here:ishere1}" v-on:mouseover="trueishere1" v-on:mouseout="falseishere1">news</span>
      <span class="s1" v-bind:class="{here:ishere2}" v-on:mouseover="trueishere2" v-on:mouseout="falseishere2">postings</span>
      <span class="s1" v-bind:class="{here:ishere3}" v-on:mouseover="trueishere3" v-on:mouseout="falseishere3">login</span>
    </div>
    <div class="underline"></div>
    <!-- <News></News>
    <Postings></Postings> -->
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
      ishere1:false,
      ishere2:false,
      ishere3:false
    }
  },
  methods:{
    trueishere1:function(){
      this.ishere1=true
    },
    falseishere1:function(){
      this.ishere1=false
    },
    
    trueishere2:function(){
      this.ishere2=true
    },
    falseishere2:function(){
      this.ishere2=false
    },
    trueishere3:function(){
      this.ishere3=true
    },
    falseishere3:function(){
      this.ishere3=false
    }

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
  /* margin-top: 60px; */
} 
.head{
  background-color: #555;
  padding-top:1%;
  padding-bottom:1%;
  font-size:130%;
}
.s1{
  color:aliceblue;
  margin-left:5%;
  margin-right:5%;
  cursor:default;
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
