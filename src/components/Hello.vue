<template lang='pug'>
.hello
  h1 {{ msg }}
  h2 Essential Links xxxx

  ul
    li
      router-link( :to="'test1'") test 1
    li
      router-link( :to="'test3'") test 3
    li
      router-link( :to="'login'") Login
    li
      a(href='https://vuejs.org', target='_blank') Core Docs
    li
      a(href='https://forum.vuejs.org', target='_blank') Forum
    li
      a(href='https://gitter.im/vuejs/vue', target='_blank') Gitter Chat
    li
      a(href='https://twitter.com/vuejs', target='_blank') Twitter
    br
    li
      a(href='http://vuejs-templates.github.io/webpack/', target='_blank') Docs for This Template
  h2 Ecosystem
  ul
    li
      a(href='http://router.vuejs.org/', target='_blank') vue-router
    li
      a(href='http://vuex.vuejs.org/', target='_blank') vuex
    li
      a(href='http://vue-loader.vuejs.org/', target='_blank') vue-loader
    li
      a(href='https://github.com/vuejs/awesome-vue', target='_blank') awesome-vue
</template>

<script lang="ts">
import Vue from 'vue'
import { Component } from 'vue-property-decorator'

@Component
export default class Hello extends Vue {
  msg: string = 'Welcome to Your Vue.js App'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b913;
}
</style>
