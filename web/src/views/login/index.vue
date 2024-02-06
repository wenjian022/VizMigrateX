<template style="height: 100%;margin: 0;padding: 0;">
  <div style="width: 100%;height: 100%;">
    <div style="width: 70%;height: 100%;text-align: center;float: left;background-color:#2c41b4;">
      <div style="padding-top: 80px;color:#ffffff;font-size: 50px">
        <svg-icon icon-class="logo" style="font-size: 70px" />
        <span style="padding-left: 20px;font-weight:bolder">Watchman</span></div>
      <img
        style="width: 50%;display:inline-block;margin-top: 50px"
        :alt="title"
        src="@/assets/login-box-bg.svg"
      >
      <div
        style="padding-right:200px;padding-top: 30px;color: #ffffff;font-size: 30px;font-weight:bolder;text-align: right"
      >运维管理平台
      </div>
    </div>
    <div style="width: 30%;height:100%;float: right;text-align: center">
      <div style="padding-top: 180px;font-size: 30px;"><span style="font-weight:bolder">登 录</span></div>
      <div style="text-align: center;padding-top: 20px">
        <el-form
          ref="loginForm"
          style="width: 80%;display: inline-block"
          :model="loginForm"
          :rules="loginRules"
          auto-complete="on"
          label-position="left"
        >
          <el-form-item prop="userName">
            <el-input
              ref="userName"
              v-model="loginForm.userName"
              class="s2"
              placeholder="账号"
              name="userName"
              type="text"
              tabindex="1"
              auto-complete="on"
            />
          </el-form-item>

          <el-form-item prop="password">
            <el-input
              :key="passwordType"
              ref="password"
              v-model="loginForm.password"
              class="s2"
              :type="passwordType"
              placeholder="密码"
              name="password"
              tabindex="2"
              size="medium"
              auto-complete="on"
              @keyup.enter.native="handleLogin"
            >
              <i slot="suffix">
                <span @click="showPwd">
                  <svg-icon style="font-size: 20px" :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
                </span>
              </i>
            </el-input>

          </el-form-item>

          <el-button
            :loading="loading"
            type="primary"
            style="width:100%;margin-bottom:30px;background-color:rgb(44, 65, 180);height: 50px;font-size: 20px;box-shadow: 0 0 5px 1px #999"
            @click.native.prevent="handleLogin"
          >登 录
          </el-button>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import settings from '@/settings'

export default {
  name: 'Login',
  data() {
    return {
      title: settings.title,
      loginForm: {
        userName: '',
        password: ''
      },
      loginRules: {
        userName: [{ required: true, trigger: 'blur', message: '用户名不能为空' }],
        password: [{ required: true, trigger: 'blur', message: '密码不能为空' }]
      },
      loading: false,
      passwordType: 'password',
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm).then(() => {
            this.$router.push('/')
            this.loading = false
          }).catch(() => {
            this.loading = false
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #283443;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  .el-input {
    display: inline-block;
    //height: 100%;
    //width: 100%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }
}
</style>

<style scoped>

.s2 { /*s2是我的输入框class*/

}

/deep/ .el-input__inner { /*或者 .s2>>>.el-input__inner  */
  height: 50px;
  font-size: 17px;
  border-radius: 1px; /*输入框圆角值*/
}

</style>
