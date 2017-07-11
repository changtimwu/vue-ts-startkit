import Cookies from 'js-cookie'
const user = {
  state: {
    token: Cookies.get('Admin-Token'),
    logined: false
  },
  mutations: {
    LOGIN(state, userInfo) {
      state.token = userInfo.uid + userInfo.password
      state.logined = true
      Cookies.set('Admin-Token', state.token)
    },
    LOGOUT(state) {
      state.token = ''
      state.logined = false
      Cookies.set('Admin-Token', '')
    }
  },
  actions: {
    Login({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        commit('LOGIN', {uid: userInfo.uid, password: userInfo.password})
        resolve()
      })
    },
    Logout({ commit }) {
      return new Promise((resolve, reject) => {
        commit('LOGOUT')
        resolve()
      })
    }
  },
  getters: {
    logined: state => state.logined,
    token: state => state.token
  }
}

export default user
