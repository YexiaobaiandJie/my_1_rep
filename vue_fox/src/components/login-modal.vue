<!--<template>
  <fx-login-modal :show.sync="show">
    <view slot="title">登录 Fox.One 帐号</view>
    <view slot="content">
      <view wx:if="{{!codeSent}}" class="form">
        <view class="form-row">
          <fx-phone-input inputType="number" prefix="+86" suffix="" placeholder="手机号" :value.sync="phoneValue"></fx-phone-input>
        </view>
        <view class="form-row">
          <text class="form-hint">暂时只支持中国大陆手机号</text>
        </view>
        <view class="form-row">
          <fx-code-button label="发送验证码" buttonType="primary" @tap.user="tapCodeButton"></fx-code-button>
        </view>
      </view>
      <view wx:else class="form">
        <view class="form-row">
          <text class="form-hint">验证码已发送到 {{phoneValue}}</text>
        </view>
        <view class="form-row">
            <fx-code-input inputType="number" prefix="" suffix="" placeholder="验证码" :value.sync="codeValue"></fx-code-input>
        </view>
        <view class="form-row">
          <fx-login-button label="验证" buttonType="primary" @tap.user="tapLoginButton"></fx-login-button>
        </view>
      </view>
    </view>
  </fx-login-modal>
</template>
<script>
  import wepy from 'wepy'
  import FXModal from '@/components/fx-modal'
  import FXInput from '@/components/fx-input'
  import FXButton from '@/components/fx-button'

  export default class LoginModal extends wepy.component {
    components = {
      'fx-login-modal': FXModal,
      'fx-phone-input': FXInput,
      'fx-code-input': FXInput,
      'fx-login-button': FXButton,
      'fx-code-button': FXButton
    }
    props = {
      show: {
        type: Boolean,
        default: false,
        twoWay: true
      }
    }
    data = {
      phoneValue: '',
      codeValue: '',
      codeSent: false
    }
    methods = {
      tapCodeButton (e) {
        wx.showLoading()
        console.log(this.$parent)
        this.$parent.$parent.globalData.client.fxLogin({
          phone: '86 ' + this.phoneValue
        })
          .then((res) => {
            console.log('code', res)
            this.codeSent = true
            this.$apply()
            wx.hideLoading()
          })
          .catch((err) => {
            wx.hideLoading()
            if (err.code === 1585) {
              wx.showToast({icon: 'none', title: `你的操作太频繁，请稍后再试(${err.code})`})
              this.codeSent = true
              this.$apply()
            }
            console.log('code err', err)
          })
      },
      tapLoginButton (e) {
        wx.showLoading()
        console.log('xxxxx', this.$parent.$parent.globalData)
        this.$parent.$parent.globalData.client.fxLoginCode({
          phone: '86 ' + this.phoneValue,
          code: this.codeValue,
          nickName: this.$parent.$parent.globalData.userInfo.nickName,
          avatar: this.$parent.$parent.globalData.userInfo.avatarUrl
        })
          .then((res) => {
            console.log('login', res)
            this.$parent.$parent.globalData.fxUserInfo = res.data.user
            this.$apply()
            wx.hideLoading()
            this.$emit('done', e)
          })
          .catch((err) => {
            wx.hideLoading()
            console.log('login err', err)
          })
      }
    }
  }
</script>

<style type="less" scoped>
</style>
-->