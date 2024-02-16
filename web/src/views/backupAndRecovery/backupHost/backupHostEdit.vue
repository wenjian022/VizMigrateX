<template>

  <el-container class="mainClass">
    <el-main>
      <el-card shadow="never">
        <div>
          <el-link @click="$router.push({name:'BackupHostList'})">
            <i class="el-icon-back" style="font-size: 15px">
              <span style="margin-left: 3px">{{ backupHostID ? '编辑备份存储主机' : '创建备份存储主机' }}</span>
            </i>
          </el-link>
        </div>
        <el-row style="margin-top: 30px">
          <el-col :span="14">
            <el-form
              ref="backupHostForm"
              :rules="backupHostRules"
              :model="backupHostForm"
              style="margin-right: 40px"
              label-width="120px"
              class="labelClass backupHostForm"
            >
              <el-row :gutter="12">
                <el-col :span="24">
                  <el-form-item label="主机名称:" prop="hostName">
                    <el-input v-model="backupHostForm.hostName" placeholder="请输入备份存储主机名称" size="small" />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="连接方式" prop="lineType" size="small">
                    <el-radio-group v-model="backupHostForm.lineType" :disabled="backupHostID">
                      <el-radio v-for="(item) in connectionMethod" :key="item.value" :label="item.value">
                        <el-tag style="" :style="{color: item.color}" effect="plain" size="small">
                          {{ item.value.toUpperCase() }}
                        </el-tag>
                      </el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item style="float: left" label="连接地址:" prop="connectionAddress">
                    <el-input v-model="backupHostForm.connectionAddress" size="small" placeholder="请输入连接地址" />
                  </el-form-item>
                  <el-form-item style="float:left;margin-left: 10px" label=" :" label-width="20px" prop="databasesPort">
                    <el-input v-model="backupHostForm.connectionPort" size="small" type="number" placeholder="22" />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="登录账号:" prop="hostAccount">
                    <el-input v-model="backupHostForm.hostAccount" size="small" placeholder="请输入登录账号" />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="登录密码:" prop="hostPassword">
                    <el-input
                      v-model="backupHostForm.hostPassword"
                      show-password
                      autocomplete="off"
                      placeholder="请输入登录密码"
                      size="small"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="" prop="backupDirectory">
                    <template slot="label">
                      <span>备份目录</span>
                      <el-tooltip class="item" effect="dark" placement="right">
                        <i class="el-icon-question" style="font-size: 16px; vertical-align: middle;" />
                        <div slot="content">
                          备份文件存储的目录,需要这个目录存在并有登录账号的读写权限
                        </div>
                      </el-tooltip>
                    </template>
                    <el-input
                      v-model="backupHostForm.backupDirectory"
                      placeholder="请输入备份目录"
                      size="small"
                    />
                  </el-form-item>
                </el-col>
              </el-row>
            </el-form>
          </el-col>
        </el-row>

        <el-col :span="24" style="margin-bottom: 20px;text-align: center">
          <el-button type="primary" size="small" @click="onSubmit">
            {{ backupHostID ? '编辑备份存储主机' : '创建备份存储主机' }}
          </el-button>
          <el-button size="small" @click="connectionTest">测试连接</el-button>
        </el-col>
      </el-card>
    </el-main>
  </el-container>

</template>

<script>
import { validateIP, validatorPort } from '@/validated/validated'
import { backupHostCreatePost } from '@/api/backupAndRecovery'

export default {
  name: 'BackupHostEdit',
  data() {
    return {
      connectionMethod: [{ value: 'ssh', color: '#F56C6C' }, { value: 'ftp', color: '#E6A23C' }],
      backupHostID: 0,
      backupHostForm: {
        hostName: '',
        connectionAddress: '',
        connectionPort: 22,
        lineType: 'ssh',
        hostAccount: '',
        hostPassword: '',
        backupDirectory: ''
      },
      backupHostRules: {
        hostName: [{ required: true, message: '请输入备份存储主机名称', trigger: 'blur' }, {
          max: 128,
          message: '长度不能超过 128 个字符',
          trigger: 'blur'
        }],
        connectionAddress: [{
          required: true,
          message: '请输入正确的地址',
          trigger: 'change',
          validator: validateIP
        }],
        connectionPort: [{
          type: 'number',
          required: true,
          message: '请输入连接端口',
          trigger: 'change',
          validator: validatorPort
        }],
        lineType: [{ required: true, message: '请选择备份机存储主连接方式', trigger: 'blur' }],
        hostAccount: [{ required: true, message: '请输入登录账号', trigger: 'blur' }],
        hostPassword: [{ required: true, message: '请输入登录密码', trigger: 'blur' }],
        backupDirectory: [{ required: true, message: '请输入备份目录', trigger: 'blur' }]
      }
    }
  },
  mounted() {
    this.backupHostID = Number(this.$route.query['backupHostID']) || 0
  },
  methods: {
    onSubmit() {
      this.$refs['backupHostForm'].validate(async(valid) => {
        if (valid) {
          const { code } = await backupHostCreatePost(this.backupHostForm)
          if (code === 0) {
            this.$message.success('创建成功')
          }
        }
      })
    },
    connectionTest() {

    }
  }
}
</script>

<style scoped>
.labelClass /deep/ .el-form-item__label {
  font-weight: 100;
}

.backupHostForm {
  .el-input {
    width: 350px;
  }
}

</style>
