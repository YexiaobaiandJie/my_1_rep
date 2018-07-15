<template>
<div class="container">
    
    <div class="content">
        <!--
        <div class="owner-info">{{redPacket_owner_fullname}} 的 {{coinid}} 硬币</div> 
        <div class="bubble">{{redPacket_message}}</div>
        -->
        <!--
        <div v-if="isExpired" class="slow-hint">
            硬币过期了，下手得早点儿
        </div>
        -->
       <!-- 
        <div v-else-if="!isAvailable && record === null" class="slow-hint">-->
          <!--
        <div v-if="isAvailable" class="slow-hint">
            硬币被抢光了，下手得再快点
        </div>
            -->

        <!--
        <div v-else>
            <div class="balance-block" v-if="record">
               <balance-view :amount.sync="amountValue" :unit.sync="unit" :priceUsd.sync="priceUsd" text=""></balance-view>
                
            </div>
        </div>
        -->
        
        <!--<div v-if="!isNewUser" class="button-wrapper">-->
        <div class="button-wrapper">
           <button class="browe">查看余额</button>
        </div>
        <!--
        <div v-if="isNewUser">
            
            <button v-if="isNewUser" class="login-to-view-button" open-type="getUserInfo" bindgetuserinfo="bindGetUserInfoToView">
                 登录查看领取记录
            </button>
        </div>

       
        <div v-else>
            
            <div v-if="redPacket" class="result-list-meta">
                
                <div v-if="redPacket.number !== 0">
                    共{{redPacket.total}}个，已领取{{redPacket.total-redPacket.number}}个
                </div>
                
                <div v-else>
                    硬币已抢光，用时{{during}}
                </div>
            </div>
            
            
        </div> 
         -->   
            <div class="result-list">  
                <div class="result-item">
                
                    <ol>
                        <li v-for="result in pickResults"class="li_c">  
                            <!--<image class="result-item-avatar" src="{{result.avatarUrl}}" />-->
                            <div class="result-item-left">
                                <div class="result-item-name">{{result.nickName}}</div>
                                <div class="result-item-time">{{result.time}}</div> 
                                
                            </div>
                            <div class="result-item-right">
                                <div class="result-item-amount">{{result.amount}}</div>
                                <div v-if="result.bestLuck" class="result-item-best-luck">👑 手气最佳</div>
                            </div>
                        </li>
                    </ol>
                  
                

                </div>
            </div>
            
            
            
            
    </div>

    
        <div class="back" v-show="!isOpen" v-bind:class="{'open':isOpen}" >
            <div class="back-inner">
                <div class="back-border"></div>
            <img class="back-img"  src="../assets/cover-back-img@2x.png" />
                <!--<div v-if="!isOpen" class="cover-summary">-->
            </div>
        </div>
        <div class="back_" v-show="isOpen" v-bind:class="{'open':isOpen}">
            <div class="back-inner">
                <div class="back-border"></div>
            <img class="back-img"  src="../assets/cover-back-img@2x.png" />
                <!--<div v-if="!isOpen" class="cover-summary">-->
            </div>
        </div>

    
    <div class="cover" v-show="!isOpen">
        <div class="cover-inner">
            <img class="cover-img" src="../assets/cover-front-img@2x.png" />
            <div v-if="!isOpen" class="cover-summary">
                <div class="cover-summary-text">来自{{senderName}}的{{unit}}硬币</div>
            </div>
            <div class="open-view">
                <div v-if="isOpen" class="open-avatar">
                   <!--这个可能是接收头像<image mode="aspectFit" class="open-avatar-img" src="{{avatarUrl}}" />-->
                </div>
                <button v-else class="open-button" v-on:click="openRedPacket(message)">
                    <img src="../assets/open-button-label-img@2x.png" />
                </button>
            </div>
        </div>
    </div>

    <div class="cover_" v-show="isOpen">
        <div class="cover-inner">
            <img class="cover-img" src="../assets/cover-front-img@2x.png" />
            <div v-if="!isOpen" class="cover-summary">
                <div class="cover-summary-text">来自{{senderName}}的{{unit}}硬币</div>
            </div>
            <div class="open-view">
                <!--
                <div v-if="isOpen" class="open-avatar">
                   <image mode="aspectFit" class="open-avatar-img" src="{{avatarUrl}}" />
                </div>
                
                <button v-else class="open-button_" v-on:click="openRedPacket(message)" >
                    <img src="../assets/open-button-label-img@2x.png" />
                </button>
                -->
                <button class="open-button_" v-on:click="openRedPacket(message)" >
                    
                </button>
            </div>
        </div>
    </div>
    
</div>
</template>

<script>

export default {
    
    data (){
        return{
            senderName:'',
            fullname:'',
            amountValue:'',
            unit:'',
            //priceUsd:'',
            isOpen: false,
            packetId:0,
            uuid:0,
            isExpired:false,  //是否过期
            isAvailable:false, //是否被抢光
            redPacket:null,
            record:null,
            asset:null,
            pickResults: [],
            openRedPacket:function(message){
                this.isOpen="true"
            }
    }
},
mounted: function(){
    //GET /someUrl
    
    this.$http.get( `https://dev.fox.one/api/luckycoin/detail/80`,{
      headers:{
        'content-type': 'application/json',
        'x-client-build': '101',
        'x-client-version': '2.0.1',
        'x-client-device-id': 'red_packet_web',
        'x-client-type': '2',
       // 'x-client-name': 'Mia'
      }
    }).then(response =>{
        
        this.senderName=response.data.data.redPacket.owner.fullname
        this.unit=response.data.data.redPacket.assetSymbol
        
    },response =>{
        console.log(error);
    });
  }

}
</script>


<style scoped>

    .open_{
        
    }
    .li_c{
        width:350px;
        margin-bottom:19px;
    }
    .container {
        background: #FFF;
    } 
    .cover {
        position: absolute;
        top:0;
        left: 0;
        right: 0;
        z-index: 100;
        max-height: 490px;
        max-width:450px;
        height:100%;
    }
    .cover_{
        position: absolute;
        top:0;
        left: 0;
        right: 0;
        z-index: 100;
        max-height: 490px;
        max-width:450px;
        height:35%;
    }
    .cover-inner {
         position: relative;
         height: 100%;
         width: 100%;
        }
    .cover-summary {
        position: absolute;
        bottom: -100px;
        left: 0;
        right: 0;
        display:flex;
        flex-direction:column;
        justify-content:center;
        align-items:center;
        }
    .cover-summary-text {
          margin-top:10px;
          text-align:center;
          color:white;
          background: rgba(255,255,255,0.2);
          padding:4px 8px;
          border-radius:4px;
          font-size:16px;
        }
      
    .cover-img {
        height: 100%;
        width: 100%;
      }

    .open-view {
        position: absolute;
        z-index: 110;
        bottom: 0;
        left: 50%;
        margin-left: -50px;
        margin-bottom: -50px;
        background: #FFD244;
        border: 1px solid #A22C13;
        box-shadow: 0 1px 2px 0 #A33720;
        color: #E0644A;
        border-radius: 99em;
        height: 100px;
        width: 100px;
        display: flex;
        align-items: center;
        justify-content: center;
      }
    .open-button, .open-avatar {
        background: #FFEFBE;
        border-radius: 99em;
        height:150px;
        width: 150px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 50px;
        border:10px solid #FFD244;
        box-shadow: 0 1px 2px 0 rgba(172,105,18,0.50);
    }
    .open-button_{
        background: #FFEFBE;
        border-radius: 99em;
        height:80px;
        width: 80px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 50px;
        border:5px solid #FFD244;
        box-shadow: 0 1px 2px 0 rgba(172,105,18,0.50);
    }
    .open-avatar-img {
        border-radius: 99em;
        height: 100%;
      }
    .cover.open {
    transition: height 0.1s ease;
     }

   .open-view {
      background: transparent;
      box-shadow: none;
      border: none;
    }
  
    .back {
        position: fixed;
        bottom: 0;
        left: 0;
        right: 0;
        max-height: 620px;
        z-index: 99;
        max-width:450px;
        margin-top:-60px;
        height:100%;
    }
    .back_{
        position: fixed;
        bottom: 0;
        left: 0;
        right: 0;
        max-height: 620px;
        z-index: 99;
        max-width:450px;
        margin-top:-60px;
        height:8%;
    }
    .back-inner {
      position: relative;
      height: 100%;
      width: 100%;
    }
    .back-border {
        position: absolute;
        top: -4px;
        border-top: 1px solid #DD6A37;
        border-bottom: 1px solid #DD6A37;
        height: 2px;
        left: 0;
        right: 0;
    }
    .back-img {
        height: 100%;
        width: 100%;
    }
    .back.open {
    transition: height 0.1s ease;

    }
    
    .bubble {
        border-radius:8px;
        margin: 2px auto 10px auto;
        text-align: center;
        font-size: 14px;
        padding: 4px 8px;
        background:#eee;
        color: #888;
        max-width:300px;
        display: inline-block;
        position: relative;
    }
    .owner-info {
        text-align: center;
        font-size: 19px;
        margin-bottom: 10px;
        color:#111
        
    }
    .container.closed {
        overflow: hidden;
    }
    .content {
        padding: 170px 0 100px 0;
        width:100%;
        opacity: 1;
        display: flex;
        justify-content: center;
        flex-direction: column;
    }
    .content.open {
        opacity: 1;
    }
    .slow-hint {
        text-align: center;
        color: #F55C23;
        margin: 10px auto 30px auto;
    }

    .result-list {
        padding: 0 16px;
    }
    .result-item:first-child {
      border-top: 1px solid #aaa;
    }
   
    .result-item {
      display:flex;
      flex-direction: row;
      border-top: 1px dashed #aaa;
      padding: 10px 0;
      align-items: center;
      color: #000;
    }
    .result-item-avatar {
        height: 32px;
        width: 32px;
        border-radius: 50%;
    }
    
    .result-item-name, .result-item-amount {
        align-items: center;
        display: flex;
        color:#000;
    }
    
    .result-item-name {
          font-size: 16px;

    }
    .result-item-time {
        
          font-size: 12px;
          color: #aaa;
    }
    .result-item-amount {
          font-size: 16px;
          text-align: right;
          align-items: flex-end;
          justify-content: flex-end;
    }
    .result-item-left {
        width:250px;
        margin-left: 10px;
        margin-bottom:5px;
        flex:1;
        float:left;
    }
    .result-item-right {
        width:90px;
        display:flex;
        /*justify-content: flex-end;*/  
        float:left;
    }
    
    .result-item-best-luck {
          color: #F55C23;
          font-size: 12px;
          text-align: right;
    }   
    .login-to-view-button {
        width: 200px;
        margin: 10px auto;
    }
    .browe{
        width:120px;
        height:40px;
        margin-left:36%;
        margin-bottom:5px;
        border:1px solid #ff4500;
        background: #F1F1F1;
        border-radius: 12em;
    }


</style>
