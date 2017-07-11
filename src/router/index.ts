// /// <reference path="../../node_modules/@types/webpack-env/index.d.ts" />

import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  // base: __dirname,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  },
  routes: [{
    path: '/',
    name: 'Hello',
    component: resolve => require(['@/components/Hello.vue'], resolve)
  }, {
    path: '/test1',
    name: 'test1',
    component: resolve => require(['@/components/Test1.vue'], resolve)
  }, {
    path: '/test3',
    name: 'test3',
    component: resolve => require(['@/components/Test3.vue'], resolve)
  }, {
    path: '/login',
    name: 'login',
    component: resolve => require(['@/components/login'], resolve)
  }]
})
