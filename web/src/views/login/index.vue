<template style="height: 100%;margin: 0;padding: 0;">
  <div style="width: 100%;height: 100%;">
    <div style="width: 70%;height: 100%;text-align: center;float: left;background-color:#2c41b4;">
      <div style="padding-top: 80px;color:#ffffff;font-size: 50px">
        <svg-icon icon-class="logo" style="font-size: 70px"/>
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
      <login-component v-if="initialization"/>
      <initialization-component v-else/>
    </div>
  </div>
</template>

<script>
import settings from '@/settings'
import LoginComponent from '@/views/login/loginComponent.vue'
import InitializationComponent from '@/views/login/initializationComponent.vue'
import { userInitializationGet } from '@/api/user'

export default {
  name: 'Login',
  components: { InitializationComponent, LoginComponent },
  data() {
    return {
      title: settings.title,
      initialization: false,
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
  mounted() {
    // 判断是否初始化
    this.getPublicInitialization()
  },
  methods: {
    async getPublicInitialization() {
      const { code, result } = await userInitializationGet()
      if (code === 0) {
        this.initialization = result
      }
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
