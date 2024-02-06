<template>
  <div>
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
                <svg-icon style="font-size: 20px" :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"/>
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
</template>

<script>
export default {
  name: 'LoginComponent',
  data() {
    return {
      loginForm: {
        userName: '',
        password: ''
      },
      loginRules: {
        userName: [{ required: true, trigger: 'blur', message: '用户名不能为空' }],
        password: [{ required: true, trigger: 'blur', message: '密码不能为空' }]
      },
      loading: false,
      passwordType: 'password'
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
            this.$router.push('/dashboard')
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

<style scoped>

</style>
