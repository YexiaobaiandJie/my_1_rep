<template>
  <div id="app">
    <!--<img src="./assets/logo.png">-->
    <div id="notebook" v-show="!isopen"> <!--笔记本展示页-->
      <h1>list of note</h1>
      <button v-on:click="addnotefunc" class="newnotep">new</button>
      <hr />
      <div style="margin-left:2.8%">
        <div v-for="noteitem in noteitems" v-on:click="opennote(noteitem)">
          <div class="box1" v-bind:class="{box2:noteitem.isFinished}">
            <img src="./assets/logo.png" style="width:90%;height:70%;">
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
      <!--添加笔记本-->
      <!--
      <div class="addnote" v-on:click="addnotefunc">
        <p class="addnotep">new</p>
      </div>
      -->
    </div>
    <transition name="fade">
      <div v-show="createnote">
        <Addnote v-on:getNoteInfo="getNoteInfo" v-bind:createnote="createnote"  v-bind:parnoteitems="noteitems"></Addnote>
       <!--<Addnote></Addnote>-->
      </div>
    </transition>









    <div id="todolist" v-show="isopen">  <!--清单页-->
    <h1>{{title}} of {{thisnoteitem}}</h1>
    <h2 class="back1" v-on:click="closenote"> < Back</h2>
    <input class="text1" type="text" v-model="newitem" v-on:keyup.enter="additem()"/>
    <div class="label1" v-for="item in items" v-bind:class="{FinC:item.isFinished}" v-on:click="finishitem(item)">
        {{item.label}}
        
    </div>
    <p class="del1" v-bind:class="{delp:items.length===0}" v-on:click="delall"> delete-note </p>
    </div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld'
import Store from './components/store'
import Addnote from './components/addnote'
export default {
  name: 'App',
  components: {
    HelloWorld,Addnote
  },
  data(){
    return{
      title:'list',
      value:'',
      items:[],
      newitem:'',
      newnoteinfo:{
        label:'',
        time:''
      },
      isopen:false,
      createnote:false,
      noteitems:Store.fetchnote(),
      thisnoteitem:'',
      thisnote:''
    }
  },
  methods:{
    finishitem:function(item){   //当点击事件时触发事件条目状态更改
      item.isFinished=!item.isFinished
    },
    additem:function(){          //当按下回车触发增加事件条目
      this.items.push({
        label:this.newitem,
        isFinished:false
      })
      this.newitem=''
    },
    crNewNote:function(){       //将笔记数据放入笔记数组中，使其能够在笔记展示显示出来
      this.noteitems.push({
        label:this.newnoteinfo.label,
        time:this.newnoteinfo.time,
        isFinished:false
      })
      this.createnote=false
    },
    delall:function(){                 //点击删除按钮，删除当前笔记里全部事件条目
      Store.remove(this.thisnote)
      this.items=[]
      Store.removenote(this.thisnote.label)
      this.isopen=false
      this.noteitems=Store.fetchnote()
    },
    opennote:function(noteitem){       //点击笔记，进入笔记调出当前笔记所有内容
      this.isopen=true
      this.thisnote=noteitem
      this.thisnoteitem=noteitem.label
      this.items=Store.fetch(noteitem.label)
    },
    closenote:function(){              //点击返回退出笔记，回到笔记展览状态
      this.isopen=false
    },
    addnotefunc:function(){           //点击增加笔记，弹出输入框
      this.createnote=true
     },
    getNoteInfo:function(noteinfo){    //从子组件那获得新创建的笔记的信息
      this.newnoteinfo.label=noteinfo.label
      this.newnoteinfo.time=noteinfo.time
      this.crNewNote()
    },
    ChangeCreatenoteStatue:function(){       //改变新增笔记框的状态 隐藏新增笔记框
      this.createnote=false
    }
  },
   
  watch:{
    items:{
      handler:function(items){          //监视条目变化，将条目变化存入storage
        Store.save(items,this.thisnoteitem)
      },
      deep:true
    },
    noteitems:{
      handler:function(noteitems){
        Store.savenote(noteitems)
      },
      deep:true
    }
  }
}
</script>

<style scoped>

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.FinC{        /*当事件条目是完成状态时，此时条目的样式*/
  text-decoration:line-through;
  opacity: 0.4;
  
}
.text1{             /*事件输入框样式*/
  width:50%;
  height:30px;
  font-size:22px;
}
.label1{         /*事件条目默认样式*/
  font-size:22px;
  width:50.3%;
  height:30px;
  line-height:30px;
  margin-left:24.9%;
  text-align: center;
  background:#f1f1f1;
  cursor:pointer;
}
.delp{       /*删除标签当没有事件条目时隐藏的样式*/
  opacity: 0;
  cursor: default;
}
.del1{      /*删除标签显现的样式*/
  cursor:default;
  color:#111;
}
.box1{      /*笔记框样式*/
  background-color: #f1f1f1;
  width:43%;
  height:23%;
  max-width:160px;
  max-height:190px;
  float:left;
  border: none;
  border-radius: 0.3em;
  margin-left:17px;
  margin-top:8px;
  box-shadow: 8px 8px 60px 1px #c7c7c7;
  
}
.box2{   /*当笔记内事件全部完成时的样式*/

}
.fin-w{   /*完成提示显示出来的样式*/
  font-size:20px;
  color:#ff4500;
  font-weight:bold;
}
.back1{    /*返回 笔记展示模块 标签的样式*/
  cursor:default;
}

.fade-enter-active, .fade-leave-active{
  transition: opacity 0.6s
}
.fade-enter, .fade-leave-to{
  opacity: 0
}
.newnotep{
  font-size:21px;
  cursor:default;
  background-color:#cfcfcf;
  border:none;
}
</style>
