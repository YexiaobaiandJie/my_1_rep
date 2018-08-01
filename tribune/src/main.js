// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import VueResource from 'vue-resource'
import VRouter from 'vue-router'
import News from './components/news'
import Postings from './components/postings'
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
