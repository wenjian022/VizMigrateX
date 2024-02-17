<template>
  <el-container class="mainClass">
    <el-header style="margin-top: 10px">
      <el-card shadow="never">
        <el-link @click="$router.push({name:'DataSource'})">
          <i class="el-icon-back" style="font-size: 15px">
            <span style="margin-left: 3px">{{ dataSourceID ? '编辑数据源' : '创建数据源' }}</span>
          </i>
        </el-link>
      </el-card>
    </el-header>

    <el-main style="padding-top: 30px">
      <el-card shadow="never">
        <el-row>
          <el-col :span="14">
            <el-form
              ref="dataSourceForm"
              :rules="dataSourceRules"
              :model="dataSourceForm"
              style="margin-right: 40px"
              label-width="120px"
              class="labelClass dataSourceForm"
            >
              <el-row :gutter="12">
                <el-col :span="24">
                  <el-form-item label="数据源名称:" prop="dataSourceName">
                    <el-input v-model="dataSourceForm.dataSourceName" placeholder="请输入数据源名称" size="small" />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="数据库类型:" prop="databaseType" size="small">
                    <el-cascader
                      v-model="dataSourceForm.databaseType"
                      :disabled="Boolean(dataSourceID)"
                      :options="cascaderOptions"
                      style="width: 350px;"
                      @change="cascaderChange"
                    >
                      <template slot-scope="{node,data}">
                        <svg-icon v-if="data.icon !== ''" style="margin-right: 3px" :icon-class="data.icon" />
                        <span>{{ data.label }}</span>
                      </template>
                    </el-cascader>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item style="float: left" label="连接地址:" prop="connectionAddress">
                    <el-input v-model="dataSourceForm.connectionAddress" size="small" placeholder="请输入连接地址" />
                  </el-form-item>
                  <el-form-item style="float:left;margin-left: 10px" label=" :" label-width="20px" prop="databasesPort">
                    <el-input v-model="dataSourceForm.connectionPort" size="small" type="number" placeholder="3306" />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="数据库账号:" prop="databaseAccount">
                    <el-input v-model="dataSourceForm.databaseAccount" size="small" placeholder="请输入数据库账号" />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="数据库密码:" prop="databasePassword">
                    <el-input
                      v-model="dataSourceForm.databasePassword"
                      show-password
                      autocomplete="off"
                      placeholder="请输入数据库密码"
                      size="small"
                    />
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="环境:" prop="environment">
                    <el-radio-group v-model="dataSourceForm.environment">
                      <el-radio v-for="(item,index) in environmentList" :key="index" :label="item.id">
                        <el-tag style="background-color:#f8f8f8;" :style="{color: item.color,}" size="small">{{
                          item.environment_name
                        }}
                        </el-tag>
                      </el-radio>
                    </el-radio-group>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item label="标签:">
                    <el-checkbox-group v-model="dataSourceForm.label">
                      <el-checkbox v-for="(item,index) in labelList" :key="index" :label="item.label_name">
                        <span><svg-icon icon-class="tag" /> {{ item.label_name }}</span>
                      </el-checkbox>
                    </el-checkbox-group>
                  </el-form-item>
                </el-col>
                <el-col :span="24">
                  <el-form-item prop="additionalParameters">
                    <template slot="label">
                      <span>额外参数</span>
                      <el-tooltip class="item" effect="dark" placement="right">
                        <i class="el-icon-question" style="font-size: 16px; vertical-align: middle;" />
                        <div slot="content">
                          额外参数指的是连接数据时?后面带的参数,如: ?useUnicode=true&useSSL=true
                        </div>
                      </el-tooltip>
                    </template>
                    <el-input
                      v-model="dataSourceForm.additionalParameters"
                      size="small"
                      placeholder="请输入数据库账号"
                    />
                  </el-form-item>
                </el-col>
              </el-row>
            </el-form>
          </el-col>
          <el-col v-if="dataSourceID" :span="10">
            <el-card class="box-card" shadow="never" style="max-height:400px;overflow:auto;">
              <div slot="header" class="clearfix">
                <span>受影响的任务(0) <i class="el-icon-warning" style="color:#E6A23C;" /> <span style="font-size: 12px">修改数据源相关信息，可能会影响以下任务</span></span>
              </div>
              <el-button round size="small">结构同步(1)</el-button>
              <el-button round size="small">数据复制(2)</el-button>
              <el-table :data="taskList">
                <el-table-column
                  prop="task_name"
                  label="任务名称"
                  align="center"
                />
                <el-table-column
                  prop="state"
                  label="状态"
                  align="center"
                />
                <el-table-column
                  prop="creator_name"
                  label="创建人"
                  align="center"
                />
                <el-table-column
                  prop="created_at"
                  label="创建日期"
                  align="center"
                />
              </el-table>
            </el-card>
          </el-col>
        </el-row>

        <el-col :span="24" style="margin-bottom: 20px;text-align: center">
          <el-button type="primary" size="small" @click="onSubmit">{{
            dataSourceID ? '更新数据源' : '创建数据源'
          }}
          </el-button>
          <el-button size="small" @click="connectionTest">测试连接</el-button>
        </el-col>
      </el-card>
    </el-main>
  </el-container>

</template>

<script>
import {
  dataSourcesConnectionTestPost,
  dataSourcesDatabasesCratePost, dataSourcesDatabasesPut,
  dataSourcesEnvListGet, dataSourcesInfoGet, dataSourcesLabelListGet
} from '@/api/dataSource'
import { validateIP, validatorPort } from '@/validated/validated'

export default {
  name: 'DataSourceEdit',
  data() {
    return {
      labelList: [],
      dataSourceID: 0,
      taskList: [{ task_name: '生产同步表结构', state: 0, creator_name: 'admin', created_at: '2024-02-13 19:33:03' }],
      environmentList: [],
      cascaderOptions: [
        {
          label: '自建数据库',
          value: 'idc',
          icon: 'buildByOneself',
          children: [{
            label: '关系型数据库',
            value: 'relational',
            icon: '',
            children: [{
              label: 'MySQL',
              value: 'idc_mysql',
              icon: 'mysql'
            }
              // {
              //   label: 'Oracle',
              //   value: 'idc_oracle',
              //   icon: 'oracle'
              // }
            ]

          }]
        }],
      dataSourceForm: {
        databaseType: 'idc_mysql',
        dataSourceName: '',
        connectionAddress: '',
        connectionPort: 3306,
        databaseAccount: '',
        databasePassword: '',
        additionalParameters: '',
        environment: 0,
        label: []
      },
      dataSourceRules: {
        dataSourceName: [{ required: true, message: '请输入数据源名称', trigger: 'blur' }, {
          max: 128,
          message: '长度不能超过 128 个字符',
          trigger: 'blur'
        }],
        connectionAddress: [{
          required: true,
          message: '请输入正确的数据源地址',
          trigger: 'change',
          validator: validateIP
        }],
        connectionPort: [{
          type: 'number',
          required: true,
          message: '请输入数据库端口',
          trigger: 'change',
          validator: validatorPort
        }],
        environment: [{ required: true, message: '请选择环境', trigger: 'change' }],
        databaseAccount: [{ required: true, message: '请输入数据库账号', trigger: 'change' }],
        databasePassword: [{ required: true, message: '请输入数据库密码', trigger: 'change' }],
        databaseType: [{ required: true, message: '请选择数据库类型', trigger: 'change' }]
      }
    }
  },
  mounted() {
    this.dataSourceID = Number(this.$route.query['dataSourceID']) || 0
    if (this.dataSourceID) {
      this.editDataSources()
    }
    this.initEnv()
    this.initLabel()
  },
  methods: {
    async editDataSources() {
      const { code, result } = await dataSourcesInfoGet(this.dataSourceID)
      if (code === 0) {
        this.dataSourceForm = result
      }
    },
    async initEnv() {
      const { code, result } = await dataSourcesEnvListGet({ size: 100, page: 1 })
      if (code === 0) {
        this.environmentList = result.data
        this.environmentList.forEach((item) => {
          if (item.environment_name === '生产') {
            this.dataSourceForm.environment = item.id
          }
        })
      }
    },
    async initLabel() {
      const { code, result } = await dataSourcesLabelListGet({ size: 100, page: 1 })
      if (code === 0) {
        this.labelList = result.data
      }
    },
    onSubmit() {
      this.$refs['dataSourceForm'].validate(async(valid) => {
        if (valid) {
          const loading = this.$loading({
            lock: true,
            text: '数据加载中，请稍后……',
            background: 'rgba(0, 0, 0, 0.5)'
          })

          const submit = this.dataSourceID ? dataSourcesDatabasesPut(this.dataSourceForm, this.dataSourceForm.id) : dataSourcesDatabasesCratePost()

          submit.then(_ => {
            setTimeout(() => {
              loading.close()
              this.$message.success(this.dataSourceID ? '修改成功' : '创建成功')
              this.$router.push({ name: 'DataSource' })
            }, 1500)
          }).catch(_ => {
            loading.close()
          })
        }
      })
    },
    // connectionTest 数据库测试
    connectionTest() {
      this.$refs['dataSourceForm'].validate(async(valid) => {
        if (valid) {
          const { code } = await dataSourcesConnectionTestPost(this.dataSourceForm)
          if (code === 0) {
            this.$message.success('连接可用')
          }
        }
      })
    },
    cascaderChange(value) {
      this.dataSourceForm.databaseType = value[value.length - 1]
    }
  }
}
</script>

<style scoped>
.labelClass /deep/ .el-form-item__label {
  font-weight: 100;
}

.dataSourceForm {
  .el-input {
    width: 350px;
  }
}
</style>
