// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
//import App from './App'
import VueResource from 'vue-resource'
import Layout from './components/Layout'
import VueRouter from 'vue-router'
import IndexPage from './pages/index'
//import assets from './assets'
Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueResource)
let router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path:'/',
      component:IndexPage
    }
  ]
})

let _deviceId = ''
/* eslint-disable no-new */
/*
function generateDeviceId () {
  if (_deviceId) {
    return _deviceId
  }
  _deviceId = _generateDeviceId()
  console.log(_deviceId)
  return _deviceId
}

function _generateDeviceId () {
  const _0xcc12 = ['\x29\x92', '\x46\x4F\x58\x5F\x44\x45\x53\x4B\x54\x4F\x50\x5F\x44\x45\x56\x49\x43\x45\x5F\x49\x44\x5F\x4B\x45\x59\x5F\x33\x38\x31\x32\x30\x39\x32\x32\x61\x31\x30\x39\x34\x30\x37\x30\x34\x30\x62\x63\x32\x38\x30\x65\x31\x36\x35\x36\x64\x34\x30\x37\x38\x39\x33\x33\x63\x36\x38\x31\x33\x61\x62\x34\x34\x35\x66\x33\x36\x39\x33\x39\x65\x36\x66\x62\x35\x66\x64\x32\x65\x61\x36\x36']
  const sysInfo = wx.getSystemInfoSync()
  // const userInfo = wx.getStorageSync('user')
  const text = sysInfo.brand + sysInfo.model + ':'
  const hash = CryptoJS.HmacSHA256(text, _0xcc12[1])
  return hash.toString().toUpperCase().match(/.{1,2}/g).map((x) => x[0]).join('')
}
*/
new Vue({
  el: '#app',
  router,
  components: { Layout },
  template: '<Layout/>',
  

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
        console.log(response.data);
        redata=response.data;
    },response =>{
        console.log(error);
    });
  }
  
    
  
})
