<template lang='pug'>
div
  h1 This is Test1
  h2 count is {{count}} double is {{doublecnt}}
  el-button( v-on:click="increment") +
  el-button( v-on:click="decrement") -
  el-button( v-on:click="asyncIncrement") Slow +
  el-Button( v-on:click="asyncDecrement") Slow -
  el-Button( v-on:click="logout") Logout

  p token is {{token}}
  icon( name="language" scale="3")
  tco(title='from test1--es6')
  el-radio.radio( v-model="radio", label="1") optionA
  el-radio.radio( v-model="radio", label="2") optionB
</template>

<script>
import Vue from 'vue'

import tco from './Tco.vue'
import { mapGetters, mapMutations, mapActions } from 'vuex'

export default {
  name: 'test1',
  data() {
    return {
      radio: '1'
    }
  },
  // usage 1
  computed: mapGetters(['count', 'doublecnt', 'token']),
  // usage 2
  /*
  computed: {
    ...mapGetters(['count', 'doublecnt'])
  },*/
  // usage 3
  /*
  computed: {
    count () {
      return this.$store.getters.count
    },
    doublecnt () {
      return this.$store.getters.doublecnt
    }
  },*/
  // usage 1
  //methods: mapMutations(['increment', 'decrement']),
  // usage 2 recommeded
  /*
  methods: {
    ...mapMutations(['increment', 'decrement'])
  },*/
  // usage 3
  /*
  methods: {
    increment () {
      this.$store.commit('increment')
    },
    decrement () {
      this.$store.commit('decrement')
    }
  },*/
  methods: {
    ...mapMutations(['increment', 'decrement']),
    ...mapActions({
      asyncIncrement: 'increment',
      asyncDecrement: 'decrement',
      logout: 'Logout'
    }),
    logout() {
      this.$store.dispatch('Logout')
      location.reload()
    }
  },
  components: {tco}
}
</script>

