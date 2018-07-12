// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
//import App from './App'
import Layout from './components/Layout'
import VueRouter from 'vue-router'
import IndexPage from './pages/index'
//import assets from './assets'
Vue.config.productionTip = false
Vue.use(VueRouter)
let router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path:'/',
      component:IndexPage
    }
  ]
})
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { Layout },
  template: '<Layout/>'
})
