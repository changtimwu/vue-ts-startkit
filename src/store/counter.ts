import Cookies from 'js-cookie'

function setcount(state, n) {
  state.count = n
  Cookies.set('count', n)
}

const mutations = {
  plus(state, n) {
    setcount(state, state.count+n)
  },
  minus(state, n) {
    setcount(state, state.count-n)
  },
  increment(state) {
    setcount(state, state.count+1)
  },
  decrement(state) {
    setcount(state, state.count-1)
  }
}

const actions = {
  increment ({commit}) {
    setTimeout(() => {
      commit('increment')
    }, 1000)
  },
  decrement ({commit}) {
    setTimeout(() => {
      commit('decrement')
    }, 1000)
  },
  plus({commit}, amount) {
    let tid = setInterval(() => {
      commit('plus', amount)
    }, 1000)
    setTimeout(() => {
      clearInterval(tid)
    }, 10*1000)
  },
  minus({commit}, amount) {
    setTimeout(() => {
      commit('minus', amount)
    }, 1000)
  }
}

const getters ={
  count: state => state.count,
  doublecnt: state => state.count*2
}

const counter = {
  state: {
    count: Cookies.get('count') || 0
  },
  actions,
  mutations,
  getters
}

export default counter

