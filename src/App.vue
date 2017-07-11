<template lang="pug">
  #app
    router-view
</template>

<script lang="ts">
import Vue from 'vue'
import {
  Component
} from 'vue-property-decorator'

@Component
export default class App extends Vue {}
</script>

<style lang="scss">
#app {
    @import './styles/index.scss';
}
</style>
