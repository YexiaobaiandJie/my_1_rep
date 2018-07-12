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
    background:#f1f1f1;
}
.header{
    width:100%;
    height:50px;
    background-color: #ff4500;
    text-align: center;
    line-height: 50px;
    color:#fff;
}
.footer{
    width:100%;
    height:50px;
    background-color:#cfcfcf;
    text-align: center;
    line-height:50px;;
}
  
</style>
