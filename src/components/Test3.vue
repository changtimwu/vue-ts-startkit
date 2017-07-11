<template lang='pug'>
.hello
  h1 This is Test3
  h2 count is {{count}} doublecnt is {{doublecnt}}
  el-button( v-on:click="plus(2)") plus 2
  el-button( v-on:click="minus(2)") minus 2
  el-button( v-on:click="asyncPlus(2)") slow plus 2
  el-button( v-on:click="asyncMinus(2)") slow minus 2
  el-button( v-on:click="doLogin") login ok
  el-button( v-on:click="doLoginFail") login fail
  el-button( v-on:click="doLogout") logout

  br
  el-button( v-on:click="doGetCompany") get company
  input( v-model="companyName" placeholder="company name")
  el-button( v-on:click="doSetCompany") set company
  br
  icon( name="envelope" spin scale=3)
  tco(title='from test3 -- ts')
  el-checkbox( v-model="checked") Option
</template>

<script lang="ts">
import Vue from 'vue'
import { Component } from 'vue-property-decorator'
import { Getter, Mutation, Action } from 'vuex-class'
import tco from './Tco.vue'
import * as api from '../api'

@Component({
  components: { tco }
})

export default class test3 extends Vue {
  checked: boolean=false
  companyName: string='superdry'
  // conjunct with vuex's getter -- way 1
  get count() {
    return this.$store.state.count
  }
  // conjunct with vuex's getter -- way 2
  @Getter('doublecnt') doublecnt
  // or with type
  /*
  @Getter('doublecnt')
  doublecnt: number*/

  // mutations to methods  -- way 1
  /*
  plus (amount) {
    store.commit('plus', amount)
  }
  minus (amount) {
    store.commit('minus', amount)
  }*/
  // mutations to methods  -- way 2
  @Mutation('plus') plus
  @Mutation('minus') minus
  @Action('plus') asyncPlus
  @Action('minus') asyncMinus
  async doSetCompany() {
    console.log('companyName=', this.companyName)
    let rcompany = await api.SetCompany({Name: this.companyName, City: 'Taipei', Other: 'blah'})
  }
  async doGetCompany() {
    let rcompany = await api.GetCompany('intrising')
    console.log('rcompany=', rcompany)
  }
  async doLogin() {
    let ret = await api.LoginWithAccount({Id: 'admin', Password: 'admin'})
    console.log('login ret=', ret)
  }
  async doLoginFail() {
    let ret = await api.LoginWithAccount({Id: 'admin', Password: 'xxxn'})
    console.log('loginfail ret=', ret)
  }
  async doLogout() {
    let ret = await api.Logout()
    console.log('logout ret=', ret)
  }

  // or with type
  /*
  @Mutation('plus')
  plus: (amount:number) => void
  @Mutation('minus')
  minus: (amount:number) => void
  */
}
</script>

