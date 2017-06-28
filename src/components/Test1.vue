<template lang='pug'>
div
  h1 This is Test1
  tco(title='from test1--es6')
  el-radio.radio( v-model="radio", label="1") optionA
  el-radio.radio( v-model="radio", label="2") optionB
</template>

<script>
import tco from './Tco.vue'
export default {
  name: 'test1',
  data() {
    return {
      radio: '1'
    }
  },
  methods: {
  },
  components: {tco}
}
</script>

