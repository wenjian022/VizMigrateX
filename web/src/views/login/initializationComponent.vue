<template>
  <div>
    <div style="padding-top: 180px;font-size: 30px;"><span style="font-weight:bolder">用户初始化</span></div>
    <div style="text-align: center;padding-top: 20px">
      <el-form
        ref="initializationUserForm"
        style="width: 80%;display: inline-block"
        :model="initializationUserForm"
        :rules="initializationUserRules"
        auto-complete="on"
        label-position="left"
      >
        <el-form-item prop="userName">
          <el-input v-model="initializationUserForm.userName" size="small" placeholder="账号"/>
        </el-form-item>

        <el-form-item prop="password">
          <el-input v-model="initializationUserForm.password" size="small" placeholder="密码" show-password/>

        </el-form-item>
        <el-form-item prop="confirmPassword">
          <el-input v-model="initializationUserForm.confirmPassword" size="small" placeholder="确认密码" show-password/>
        </el-form-item>
        <el-button
          type="primary"
          style="width:100%;margin-bottom:30px;background-color:rgb(44, 65, 180);height: 50px;font-size: 20px;box-shadow: 0 0 5px 1px #999"
          @click.native.prevent="createUser"
        >创建并登录
        </el-button>
      </el-form>
    </div>
  </div>
</template>
<script>
import { userInitializationPost } from '@/api/user'

export default {
  name: 'InitializationComponent',
  data() {
    const confirmPasswordRules = (rule, value, callback) => {
      if (this.initializationUserForm.password === value) {
        callback()
      } else {
        return callback(new Error('密码不一致'))
      }
    }
    const passwordRules = (rule, value, callback) => {
      const check = /^(?=.{8})(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d)(?=.*?[*?!&￥$%^#,./@";:><\[\]}{\-=+_\\|》《。，、？'‘“”~ `]).*$/
      if (check.test(value)) {
        callback()
      } else {
        return callback(new Error('密码规则: 强度不够,密码长度不得低于8位并,至少一个大小字母,一个特殊符号'))
      }
    }
    return {
      initializationUserForm: {
        userName: '',
        password: '',
        confirmPassword: ''
      },
      passwordType: 'password',
      initializationUserRules: {
        userName: [{ required: true, trigger: 'blur', message: '用户名不能为空' }],
        password: [
          { required: true, trigger: 'blur', message: '密码不能为空' },
          { min: 8, trigger: 'blur', message: '密码规则: 长度不能小于8位' },
          { trigger: 'blur', validator: passwordRules }],
        confirmPassword: [
          { required: true, trigger: 'blur', message: '确认密码不能为空' },
          { trigger: 'blur', validator: confirmPasswordRules }
        ]
      }
    }
  },
  methods: {
    createUser() {
      this.$refs.initializationUserForm.validate(async valid => {
        if (valid) {
          const { code } = await userInitializationPost(this.initializationUserForm)
          if (code === 0) {
            this.$store.dispatch('user/login', this.initializationUserForm).then(() => {
              this.$router.push('/dashboard')
              this.loading = false
            }).catch(() => {
              this.loading = false
            })
          }
        }
      })
    }
  }
}
</script>

<style>

</style>
