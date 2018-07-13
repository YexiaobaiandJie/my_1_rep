<template>
<div class="container">
    <!--
    <div class="content">
        <div class="owner-info">{{ownername}} 的 {{coinnumber}} 硬币</div> 
        <div class="bubble">{{redPacket.message}}</div>
        -->
        <!--
        <div v-if="isExpired && record === null" class="slow-hint">
            硬币过期了，下手得早点儿
        </div>   
        <div v-else-if="!isAvailable && record === null" class="slow-hint">
            硬币被抢光了，下手得再快点
        </div>  
        <div v-else>
            <div class="balance-block" v-if="record">
               <balance-view :amount.sync="amountValue" :unit.sync="unit" :priceUsd.sync="priceUsd" text=""></balance-view>
                
            </div>
        </div>
        -->
        <!--
        <div v-if="!isNewUser" class="button-wrapper">
            
           <button @tap.user="tapBalanceButton">查看余额</button>
        </div>
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
            
            <div class="result-list">
                
                <div class="result-item">
                    <ul>
                        <li v-bind:v-for="result in pickResult">  
                            
                            <image class="result-item-avatar" src="{{result.avatarUrl}}" />  
                            <div class="result-item-left">
                                <div class="result-item-name">{{result.nickName}}</div>
                                <div class="result-item-time">{{result.time}}</div>   
                            </div>
                            <div class="result-item-right">
                                <div class="result-item-amount">{{result.amount}} {{unit}}</div>
                                <div wx:if="{{result.bestLuck}}" class="result-item-best-luck">👑 手气最佳</div>
                            </div>
                        </li>
                    </ul>
                </div>
                
            </div>
        </div>
    </div>
    
    <div class="back" v-bind:class="{'open':isOpen}" style="height:{{backHeight}}px">
        <div class="back-inner">
            <div class="back-border"></div>
            <div class="back-copyright"></div>
            
            <image class="back-img" mode="aspectFill" src="../assets/cover-back-img@2x.jpg" />
        </div>
    </div>
    
    
    <div class="cover" v-bind:class="{'open':isOpen}" style="height: {{coverHeight}}px">
        <div class="cover-inner">
            <image class="cover-img" src="../assets/cover-front-img@2x.png" />
            <div v-if="!isOpen" class="cover-summary">
               
            </div>
            <div class="open-view">
                <div v-if="isOpen" class="open-avatar">
                    <image mode="aspectFit" class="open-avatar-img" src="{{avatarUrl}}" />
                </div>
                <button v-else class="open-button" open-type="getUserInfo" bindgetuserinfo="bindGetUserInfo">
                    <image mode="aspectFit" src="../assets/open-button-label-img@2x.png" /> 
                </button>
            </div>
        </div>
    </div>
    -->

    <div class="back" v-bind:class="{'open':isOpen}" style="height:100%">
        <div class="back-inner">
            <div class="back-border"></div>
           <img class="back-img"  src="../assets/cover-back-img@2x.png" />
            
        </div>
    </div>
    
    <div class="cover" v-bind:class="{'open':isOpen}" style="height:63%">
        <div class="cover-inner">
            <img class="cover-img" src="../assets/cover-front-img@2x.png" />
            
            <div class="open-view">
                <div v-if="isOpen" class="open-avatar">
                   
                </div>
                <button v-else class="open-button">
                    <img src="../assets/open-button-label-img@2x.png" />
                </button>
            </div>
        </div>
    </div>
    
    
</div>
</template>

<script>
/*看不懂，先放着
        import wepy from 'wepy'
        import BalanceView from '@/components/balance-view'
        import LoginModal from '@/components/login-modal'
        import FxButton from '@/components/fx-button'
        import moment from 'moment'
*/
export default {
    /*
    config = {
      navigationBarTitleText: '幸运硬币',
      navigationBarTitleColor: '#FF6B49',
      backgroundColor: '#FF6B49'
    },
    components = {
      'login-modal': LoginModal,
      'fx-balance-button': FxButton,
      'fx-guide-button': FxButton,
      'balance-view': BalanceView
    },
    mixins = [],
    */
    data () {
      return {
            ownername:'helijie',
            coinnumber:'7',
            isOpen: false,
            coverHeight: 1000,
            pickResult: [
            {
         // avatarUrl:,
                nickName:'Tom',
                time:'2018-7-12 12:10:27',
                amount:'1',

            }],
      /*
      showGuideModal: false,
      showLoginModal: false,
      amountValue: '',
      unit: '',
      senderName: '',
      priceUsd: '',
      //pickResult
      pickResult: [],
      loginDoneFnName: 'getRecord',
      
      backHeight: 320,
      packetId: 0,
      uuid: 0,
      isExpired: false,
      isAvailable: false,
      //redPacket
      redPacket: null,
      //record
      record: null,
      asset: null,
      
      */
    }
  }
}
</script>


<style scoped>
    .cover {
        position: absolute;
        top:0;
        left: 0;
        right: 0;
        z-index: 100;
        height: 1000px;
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
        height: 125px;
        width: 125px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 50px;
        box-shadow: 0 1px 2px 0 rgba(172,105,18,0.50);
      }
      .open-avatar-img {
        border-radius: 99em;
        height: 100%;
      }


    .cover.open {
    transition: height 0.3s ease;
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
        height: 320px;
        z-index: 99;
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
    transition: height 0.3s ease;
    }


/*
  page {
    background: #FF6B49;
  }
  .container {
    background: #FFF;
  }
  .container.closed {
    overflow: hidden;
  }
  .content {
    padding: 170px 0 100px 0;
    width:100%;
    opacity: 0;
    display: flex;
    justify-content: center;
    flex-direction: column;
  }
  .content.open {
    opacity: 1;
  }
  .owner-info {
    text-align: center;
    font-size: 14px;
    margin-bottom: 10px;
  }
  //!!!!!
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
    &::before {
      content: " ";
      display: block;
      width: 0;
      height: 0;
      border-width: 6px;
      border-color: transparent transparent #eee transparent;
      border-style: solid;
      position: absolute;
      top: 0;
      left: 50%;
      margin-left: -6px;
      margin-top: -12px;
    }
  }
  
      //！！！这里放打开红包之后的样式 Open系列 没打开就用cover系列
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
        height: 88px;
        width: 88px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 50px;
        box-shadow: 0 1px 2px 0 rgba(172,105,18,0.50);
      }
      .open-avatar-img {
        border-radius: 99em;
        height: 100%;
      }
    }
  }
  .cover.open {
    transition: height 0.3s ease;
    .open-view {
      background: transparent;
      box-shadow: none;
      border: none;
    }
  }
  .back {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    height: 320px;
    z-index: 99;
    .back-inner {
      position: relative;
      height: 100%;
      width: 100%;
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
    }
  }
  .back.open {
    transition: height 0.3s ease;
  }
  .login-to-view-button {
    width: 200px;
    margin: 10px auto;
  }
  .result-list-meta {
    color: #aaa;
    font-size: 12px;
    padding: 0 0 0 16px;
    margin-bottom: 2px;
  }
  .result-list {
    padding: 0 16px;
    .result-item:first-child {
      border-top: 1px solid #aaa;
    }
    //！！！！！！！
    .result-item {
      display: flex;
      flex-direction: row;
      border-top: 1px dashed #aaa;
      padding: 10px 0;
      align-items: center;
      color: #000;
      .result-item-avatar {
        height: 32px;
        width: 32px;
        border-radius: 50%;
      }
      .result-item-name, .result-item-amount {
        align-items: center;
        display: flex;
      }
      .result-item-left {
        margin-left: 10px;
        flex: 1;
        .result-item-name {
          font-size: 16px;
        }
        .result-item-time {
          font-size: 12px;
          color: #aaa;
        }
      }
      .result-item-right {
        justify-content: flex-end;
        .result-item-amount {
          font-size: 16px;
          text-align: right;
          align-items: flex-end;
          justify-content: flex-end;
        }
        .result-item-best-luck {
          color: #F55C23;
          font-size: 12px;
          text-align: right;
        }
      }
    }
  }
  //!!!!!!!!!!!打开后的样式
  .slow-hint {
    text-align: center;
    color: #F55C23;
    margin: 10px auto 30px auto;
  }
  //!!!!!!!!!!!!
  .balance-block {
    margin-bottom: 0px;
  }
  .button-wrapper {
    margin: 0 auto 20px auto;
    width: 60%;
  }
  .guide-title {
    font-size: 20px;
    font-weight: bold;
  }
  .guide-text {
    font-size: 14px;
    padding: 10px 16px;
  }
  .guide-image-wrapper {
    margin: 0 auto 10px auto;
    width: 259px;
    height: 102px;
    image {
      width: 259px;
      height: 102px;
    }
  }
  .guide-button-wrapper {
    padding: 10px 16px;
    margin-bottom: 10px;
  }
  */
</style>
