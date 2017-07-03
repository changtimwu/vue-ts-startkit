const mutations = {
  plus(state, n) {
    state.count+=n
  },
  minus(state, n) {
    state.count-=n
  },
  increment: state => state.count++,
  decrement: state => state.count--
}

export default mutations
