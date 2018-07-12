<template>
  <div class="content2">
     
        
          <button class="btn1">
              <image class="imgc" src="./assets/logo.png"/>
              <div class="zm">
                  <div class="zm1">{{reccoin}}</div>
                  <div class="zm2">{{fromf}}</div>
              </div>
          </button>   
          <br/>
          <button class="btn1">
              <image class="imgc" src="./assets/logo.png"/>
              <div class="zm">
                  <div class="zm1">{{sendcoin}}</div>
                  <div class="zm2">{{tof}}</div>
              </div>
          </button> 
        

        
          <div >
             <button class="btn2">我要发幸运硬币</button>
          </div>
          <div >
            <button class="btn3" >我的余额账户</button>
          </div>
        

        
  </div>

  
</template>

<script>
export default {
 // name: 'HelloWorld',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      reccoin:'收到的硬币',
      fromf:'来自朋友的馈赠',
      sendcoin:'送出的硬币',
      tof:'给朋友的祝福'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>

.btn1{
    background-color:#fff;
    margin:20px;
    border:none;
    width:300px;
    height:100px;
    margin-top:5px;
    border-radius: 10px;
    box-shadow: 0 2px 10px rgba(0,0,0,0.05);
}
.btn2{
    width:300px;
    margin:5px;
    height:30px;
    background-color:#ffcc00;
    border:none;
    color:#ff6600;
    border-radius: 8px;
}
.btn3{
    width:300px;
    margin:5px;
    height:30px;
    border:1px solid#ff4500;
    color:#ff4500;
    border-radius: 8px;
}
.imgc{
width:50px;
height:50px;
}
.zm{
    width:200px;
    height:100px;
}
.zm1{
    font-size:20px;
    margin:3px;
}
.zm2{

}






















.content2{
    text-align: center;
    margin-top:30px;
}
.row-action {
    display: flex;
    /*align-items: center;*/
    border-radius: 8px;
    background: #fff;
    padding: 10px 20px;
    margin-bottom: 10px;
    box-shadow: 0 2px 10px rgba(0,0,0,0.05);
}
    .row-action-image {
      flex-basis: 70px;
      height: 70px;
      width: 70px;
    }
    .row-action-content {
      padding: 10px 0;
      margin-left: 20px;
      text-align: left;
      display: flex;
      flex-direction: column;
     /* flex: 1;*/
    }
    .row-action-title {
      font-size: 20px;
      flex: 1;
    }
    .row-action-text {
      flex: 0;
      font-size: 14px;
    }
    .form {
    padding: 20px 16px;
    }
    .form-row{
        width:200px;
        height:30px;
       
        background: #fff;
       /* display: flex;*/
        align-items: center;
        
       /* margin-bottom: 10px;*/
        /*padding: 10px 20px;*/
        
        box-shadow: 0 2px 10px rgba(0,0,0,0.05);
        border-radius: 8px;
       
    }
  
</style>
