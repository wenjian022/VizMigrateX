<template>
  <el-container class="mainClass">
    <el-header style="margin-top: 10px">
      <el-card shadow="never">
        <el-link @click="$router.push({name:'DataSource'})">
          <i class="el-icon-back" style="font-size: 15px">
            <span style="margin-left: 3px">{{ $route.query.dataSourceID ? '编辑数据源':'创建数据源' }}</span>
          </i>
        </el-link>
      </el-card>
    </el-header>

    <el-main style="padding-top: 30px">
      <el-card shadow="never">
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
                  <el-radio v-for="(item,index) in environmentList" :key="index" :label="item.value">
                    <el-tag style="background-color:#fff;" :style="{color: item.color,}" size="small">{{ item.value }}</el-tag>
                  </el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="标签:">
                <el-checkbox-group v-model="dataSourceForm.label">
                  <el-checkbox v-for="(item,index) in labelList" :key="index" :label="item.value">
                    <el-tag size="small" effect="plain" type="info">{{ item.value }}</el-tag>
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
                <el-input v-model="dataSourceForm.additionalParameters" size="small" placeholder="请输入数据库账号" />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item style="text-align: center">
                <el-button type="primary" size="small" @click="onSubmit">创建数据源</el-button>
                <el-button size="small" @click="connectionTest">测试连接</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </el-main>
  </el-container>

</template>

<script>
import { dataSourcesConnectionTestPost, dataSourcesDatabasesCratePost, dataSourcesDatabasesPut } from '@/api/dataSource'
import { validateIP, validatorPort } from '@/validated/validated'

export default {
  name: 'DataSourceEdit',
  data() {
    return {
      labelList: [{ value: '杭州' }],
      environmentList: [{ value: '生产', color: 'red' }, { value: '预发', color: '' }],
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
        environment: '生产',
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
  methods: {
    onSubmit() {
      this.$refs['dataSourceForm'].validate(async(valid) => {
        if (valid) {
          if (this.drawerPattern) {
            const { code } = await dataSourcesDatabasesCratePost(this.dataSourceForm)
            if (code === 0) {
              this.$message.success('添加成功')
              this.$router.go(0)
            }
          } else {
            const { code } = await dataSourcesDatabasesPut(this.dataSourceForm, this.dataSourceForm.id)
            if (code === 0) {
              this.$message.success('修改成功')
              this.$router.go(0)
            }
          }
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
