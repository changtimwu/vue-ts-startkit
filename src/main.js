// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import store from './store'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'

import 'vue-awesome/icons/envelope'
import 'vue-awesome/icons/key'
import 'vue-awesome/icons/language'
import 'vue-awesome/icons/bug'

import icon from 'vue-awesome/components/Icon'
Vue.component('icon', icon)

Vue.use(ElementUI)
Vue.config.productionTip = false

const whiteList = ['/login', '/authredirect', '/reset', '/sendpwd']
router.beforeEach((to, from, next) => {
  const rstore = store.getters
  if (rstore.token) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      next()
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next('/login')
    }
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {
    App
  }
})
