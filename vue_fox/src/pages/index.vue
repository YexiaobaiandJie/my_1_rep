<template>
<!-- 首页界面
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
              <image class="imgc" src="../assets/logo.png"/>
              <div class="zm">
                  <div class="zm1">{{sendcoin}}</div>
                  <div class="zm2">{{tof}}</div>
              </div>
          </button> 
          <div >
             <button class="btn2">{{iwant}}</button>
          </div>
          <div >
            <button class="btn3" >{{mycount}}</button>
          </div>
  </div>
-->
<div class="container">
    <div class="content">
        <div class="owner-info">{{redPacket.owner.fullname}} 的 {{unit}} 硬币</div> 
        <div class="bubble">{{redPacket.message}}</div>
        <div v-if="isExpired && record === null" class="slow-hint">
            硬币过期了，下手得早点儿
        </div>   
        <div v-else-if="!isAvailable && record === null" class="slow-hint">
            硬币被抢光了，下手得再快点
        </div>  
        <div v-else>
            <div class="balance-block" v-if="record">
              <!--  <balance-view :amount.sync="amountValue" :unit.sync="unit" :priceUsd.sync="priceUsd" text=""></balance-view>
                -->
            </div>
        </div>
        <!--如果是新用户-->
        <div v-if="!isNewUser" class="button-wrapper">
            <!--给他一个查看余额的按钮-->
            <button @tap.user="tapBalanceButton">查看余额</button>
        </div>
        <div v-if="isNewUser">
            <!--也给一个查看领取记录的按钮-->
            <button v-if="isNewUser" class="login-to-view-button" open-type="getUserInfo" bindgetuserinfo="bindGetUserInfoToView">
                 登录查看领取记录
            </button>
        </div>
        <!--如果不是新用户-->
        <div v-else>
            <!--直接显示红包领取结果-->
            <div v-if="redPacket" class="result-list-meta">
                <!--如果红包还有剩余，则显示抢到了-->
                <div v-if="redPacket.number !== 0">
                    共{{redPacket.total}}个，已领取{{redPacket.total-redPacket.number}}个
                </div>
                <!--如果没有剩余，则没抢到-->
                <div v-else>
                    硬币已抢光，用时{{during}}
                </div>
            </div>
            <!--展示抢红包结果列表-->
            <div class="result-list">
                <!--使用索引和循环输出结果列表-->
                <div class="result-item">
                    <ul>
                        <li v-bind:v-for="result in pickResult">  <!--&&&&&&&&&&这里注意一下&&&&&&&&&&&-->
                            <!--这里是相关样式-->
                            <image class="result-item-avatar" src="{{result.avatarUrl}}" />  <!--这里放一张不知道哪来的图片，大概是头像啥的，循环的-->
                            <div class="result-item-left">
                                <div class="result-item-name">{{result.nickName}}</div><!--这里是列表里的用户名-->
                                <div class="result-item-time">{{result.time}}</div>   <!--这里是抢红包的时间-->
                            </div>
                            <div class="result-item-right">
                                <div class="result-item-amount">{{result.amount}} {{unit}}</div><!--这里应该放该用户抢到了多少-->
                                <div wx:if="{{result.bestLuck}}" class="result-item-best-luck">👑 手气最佳</div><!--如果他满足最幸运的条件，给他一个称号-->
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
    <!--这里似乎是未点击打开红包所呈现的界面，就是打开链接一开始可以看到的-->
    
</div>
</template>

<script>

export default {
 // name: 'HelloWorld',
  data () {
    return {
    /*
      msg: 'Welcome to Your Vue.js App',
      reccoin:'收到的硬币',
      fromf:'来自朋友的馈赠',
      sendcoin:'送出的硬币',
      tof:'给朋友的祝福',
      iwant:'我要发幸运硬币',
      mycount:'我的余额账户'
      */
      pickResult: [
      {
         // avatarUrl:,
          nickName:'Tom',
          time:'2018-7-12 12:10:27',
          amount:'1',

      }]
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
/*
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

.content2{
    text-align: center;
    margin-top:30px;
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
      
        align-items: center;
        
       
        
        box-shadow: 0 2px 10px rgba(0,0,0,0.05);
        border-radius: 8px;
       
    }
  */





</style>
