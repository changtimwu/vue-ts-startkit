<template lang='pug'>
.hello
  h1 This is Test3
  tco(title='from test3 -- ts')
  el-checkbox( v-model="checked") Option
</template>

<script lang="ts">
import Vue from 'vue'
import { Component } from 'vue-property-decorator'
import tco from './Tco.vue'

@Component({
  components: { tco }
})
export default class test3 extends Vue {
  checked: boolean=false
}
</script>

