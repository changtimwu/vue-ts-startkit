<template lang='pug'>
  div.login-container
      el-form.card-box.login-form( autoComplete="on" :model="loginForm" :rules="loginRules" ref="loginForm" label-position="left" label-width="0px")
          h3.title System Login
          el-form-item( prop="uid")
              span.svg-container
                  icon( name="envelope" )
              el-input( name="uid" type="text" v-model="loginForm.uid" autoComplete="on" placeholder="User ID")
          el-form-item( prop="password")
              span.svg-container
                  icon( name="key")
              el-input( name="password" type="password" @keyup.enter.native="handleLogin" v-model="loginForm.password" autoComplete="on" placeholder="password")
          el-form-item
              el-button( type="primary" style="width:100%;" :loading="loading" @click.native.prevent="handleLogin") Login

</template>

<script>
    import { mapGetters } from 'vuex'
    export default {
      name: 'login',
      data() {
        const validateUid = (rule, value, callback) => {
          if (value.length === 0) {
            callback(new Error('Please input valid user ID'))
          } else {
            callback()
          }
        }
        const validatePass = (rule, value, callback) => {
          if (value.length < 5) {
            callback(new Error('password should not be shorter than 6'))
          } else {
            callback()
          }
        }
        return {
          loginForm: {
            uid: 'admin',
            password: ''
          },
          loginRules: {
            uid: [
                { required: true, trigger: 'blur', validator: validateUid }
            ],
            password: [
                { required: true, trigger: 'blur', validator: validatePass }
            ]
          },
          loading: false,
          showDialog: false
        }
      },
      computed: {
        ...mapGetters([])
      },
      methods: {
        handleLogin() {
          this.$refs.loginForm.validate(valid => {
            if (valid) {
                console.log('error submit!!')
                return false
            }
            this.loading= true
            this.$store.dispatch('Login', this.loginForm).then(() => {
              this.loading = false
              this.$router.push({path: '/'})
            })
              /*
              this.loading = true
              this.$store.dispatch('LoginByEmail', this.loginForm).then(() => {
                this.loading = false
                this.$router.push({ path: '/' })
                // this.showDialog = true
              }).catch(err => {
                this.$message.error(err)
                this.loading = false
              })*/
            } else {
            }
          })
        }
      },
      created() {
      },
      destroyed() {
      }
    }
</script>

<style lang="scss" scoped>
    @import "src/styles/mixin.scss";
    .tips {
      font-size: 14px;
      color: #fff;
      margin-bottom: 5px;
    }

    .login-container {
        @include relative;
        height: 100vh;
        background-color: #2d3a4b;

        input:-webkit-autofill {
            -webkit-box-shadow: 0 0 0px 1000px #293444 inset !important;
            -webkit-text-fill-color: #fff !important;
        }
        input {
            background: transparent;
            border: 0px;
            -webkit-appearance: none;
            border-radius: 0px;
            padding: 12px 5px 12px 15px;
            color: #eeeeee;
            height: 47px;
        }
        .el-input {
            display: inline-block;
            height: 47px;
            width: 85%;
        }
        .svg-container {
            padding: 6px 5px 6px 15px;
            color: #889aa4;
        }

        .title {
            font-size: 26px;
            font-weight: 400;
            color: #eeeeee;
            margin: 0px auto 40px auto;
            text-align: center;
            font-weight: bold;
        }

        .login-form {
            position: absolute;
            left: 0;
            right: 0;
            width: 400px;
            padding: 35px 35px 15px 35px;
            margin: 120px auto;
        }

        .el-form-item {
            border: 1px solid rgba(255, 255, 255, 0.1);
            background: rgba(0, 0, 0, 0.1);
            border-radius: 5px;
            color: #454545;
        }

        .forget-pwd {
            color: #fff;
        }
    }
</style>
