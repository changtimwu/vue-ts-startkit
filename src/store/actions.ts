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

export default actions
