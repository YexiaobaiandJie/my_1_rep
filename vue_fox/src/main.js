// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
//import App from './App'
import VueResource from 'vue-resource'
import Layout from './components/Layout'
import VueRouter from 'vue-router'
import IndexPage from './pages/index'
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

new Vue({
  el: '#app',
  router,
  components: { Layout },
  template: '<Layout/>',
  
/*
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
        
    },response =>{
        console.log(error);
    });
  }
  */
    
  
})
