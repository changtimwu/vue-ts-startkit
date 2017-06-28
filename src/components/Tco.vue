<template lang='pug'>
div
  h2 {{ title }}
  h3 This is tco -- example typescript class based component
  input( v-model="message" placeholder="edit me")
  p
    | Message is: {{ message }}
    | Reverse Message is: {{ reverseMessage }}
  div
    div(className="block")
      span( className="demonstration") Default value
      el-slider( v-model="value1" show-input)
  el-button( type="primary", v-on:click="dopost()") Post Message

</template>

<script lang='ts'>
import Vue from 'vue'
import { Component, Inject, Model, Prop, Watch } from 'vue-property-decorator'
import axios from 'axios'

@Component
export default class tco extends Vue {
  @Prop()
  title: string
  message: string = ''
  value1: number=34
  dopost(): void {
    console.log('going to blah post', this.message)
  }
  get reverseMessage(): string {
    return this.message.split('').reverse().join('')
  }
}
Vue.component('tco', tco)
</script>

