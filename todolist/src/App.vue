<template>
  <div id="app">
    <!--<img src="./assets/logo.png">-->
    <h1>{{title}}</h1>
    <input class="text1" type="text" v-model="newitem" v-on:keyup.enter="additem()"/>
    <div class="label1" v-for="item in items" v-bind:class="{FinC:item.isFinished}" v-on:click="finishitem(item)">
        {{item.label}}
        
    </div>
    <p class="del1" v-bind:class="{delp:items.length===0}" v-on:click="delall"> >delete-all< </p>
    
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import Store from './components/store'
export default {
  name: 'App',
  components: {
    HelloWorld
  },
  data(){
    return{
      title:'TODOLIST',
      value:'',
      items:Store.fetch(),
      newitem:''
    }
  },
  methods:{
    finishitem:function(item){
      item.isFinished=!item.isFinished
    },
    additem:function(){
      this.items.push({
        label:this.newitem,
        isFinished:false
      })
      this.newitem=''
    },
    delall:function(){
      Store.clear()
      this.items=[]
    }
  },
  watch:{
    items:{
      handler:function(items){
        Store.save(items)
      },
      deep:true
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
  margin-top: 60px;
}
.FinC{
  text-decoration:line-through;
  opacity: 0.4;
  
}
.text1{
  width:50%;
  height:30px;
  font-size:22px;
}
.ul1{
  text-align:center;
  width:100%;
  float:left;
}
.label1{
  font-size:22px;
  width:50.3%;
  margin-left:24.9%;
  text-align: center;
  background:#f1f1f1;
  cursor:pointer;
}
.delp{
  opacity: 0;
  cursor: default;
}
.del1{
  cursor:default;
  color:#111;
}
</style>
