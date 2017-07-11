import Vue from 'vue'
import Vuex from 'vuex'
import counter from './counter'
import user from './user'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    counter,
    user
  }
})

export default store
