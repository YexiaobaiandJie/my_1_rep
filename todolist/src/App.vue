<template>
  <div id="app">
    <!--<img src="./assets/logo.png">-->
    <div id="notebook" v-show="!isopen">
      <h1>list of note</h1>
      <hr />
      <div v-for="noteitem in noteitems" v-on:click="opennote(noteitem)">
        <div class="box1" v-bind:class="{box2:noteitem.isFinished}">
          <img src="./assets/logo.png" style="width:145px;height:145px;">
          <br />
            {{noteitem.label}}
          <br />
          <div v-show="!noteitem.isFinished">
            {{noteitem.time}}
          </div>
          <div v-show="noteitem.isFinished" class="fin-w">
            completed 
          </div>
        </div>
      </div>
    </div>











    <div id="todolist" v-show="isopen">
    <h1>{{title}}</h1>
    <h2 class="back1" v-on:click="closenote"> < Back</h2>
    <input class="text1" type="text" v-model="newitem" v-on:keyup.enter="additem()"/>
    <div class="label1" v-for="item in items" v-bind:class="{FinC:item.isFinished}" v-on:click="finishitem(item)">
        {{item.label}}
        
    </div>
    <p class="del1" v-bind:class="{delp:items.length===0}" v-on:click="delall"> delete-all </p>
    </div>
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
      newitem:'',
      isopen:false,
      noteitems:[
        {
          label:'meeting',
          time:'2017-9-18',
          isFinished:true
        },
        {
          label:'go home',
          time:'2018-11-10',
          isFinished:false
        },
        {
          label:'visit',
          time:'2018-10-19',
          isFinished:true
        }
      ]
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
    },
    opennote:function(noteitem){
      this.isopen=true
      console.log(noteitem.label)
      this.items=Store.fetch(noteitem.label)
    },
    closenote:function(){
      this.isopen=false
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
.box1{
  background-color: #f1f1f1;
  width:180px;
  height:200px;
  float:left;
  border: none;
  border-radius: 0.3em;
  margin-left:15px;
  box-shadow:8px 8px 4px #c7c7c7;
}
.box2{

}
.fin-w{
  font-size:20px;
  color:#ff4500;
  font-weight:bold;
}
.back1{
  cursor:default;
}
</style>
