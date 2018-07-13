<template>
  <div>
      <div class="header">幸运硬币</div>
      <div class="content">
          <keep-alive>
          <router-view></router-view>
          </keep-alive>
      </div>
      <div class="footer">footer</div>
  </div>
</template>

<script>
export default {
 // name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
.body{
    sbackground:#f1f1f1;
    
      
}
.header{
    max-width:442px;
    width:100%;
    height:50px;
    background-color: #ff4500;
    text-align: center;
    line-height: 50px;
    color:#fff;
}
.content{
    color:#fff;
    max-width:442px;
    height:400px;
    width:100%;
}
.footer{
    max-width:442px;
    width:100%;
    height:50px;
    background-color:#cfcfcf;
    text-align: center;
    line-height:50px;
    margin-top:230px;
}
  
</style>
