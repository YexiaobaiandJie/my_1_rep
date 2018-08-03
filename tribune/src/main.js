// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueResource from 'vue-resource'
import VRouter from 'vue-router'
import News from './components/news'
import Postings from './components/postings'
import Login from './components/login'
import Newsdetail from './components/newsdetail'
Vue.config.productionTip = false
Vue.use(VueResource)
Vue.use(VRouter)
let router = new VRouter({
  mode:'history',
  routes:[
    {
      path:'/news',
      component: News
    },
    {
      path:'/postings',
      component: Postings
    },
    {
      path:'/login',
      component:  Login
    },
    {
      path:'/news/newsdetail',
      component: Newsdetail
    }
  ]
})
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
