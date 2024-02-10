<template>
  <el-container class="mainClass">
    <el-header style="margin-top: 10px">
      <el-card shadow="never">
        <el-form ref="searchForm" :model="searchForm" @submit.native.prevent>
          <el-row :gutter="10">
            <el-col :span="6">
              <el-form-item prop="assetName" style="margin-bottom:0">
                <span style="font-size: 15px;font-weight: lighter">数据源名称: </span>
                <el-input
                  v-model="searchForm.dataSourceName"
                  style="width: 70%;"
                  size="small"
                  placeholder="请输入需要搜索的数据源名称"
                />
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item style="margin-bottom:0">
                <span style="font-size: 15px;font-weight: lighter">连接地址: </span>
                <el-input
                  v-model="searchForm.connectionAddress"
                  style="width: 70%;"
                  size="small"
                  placeholder="请输入需要搜索的数据库连接地址"
                />
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item style="margin-bottom:0">
                <span style="font-size: 15px;font-weight: lighter">数据库类型: </span>
                <el-select v-model="searchForm.databaseType" placeholder="请选择" size="small">
                  <template #prefix>
                    <span style="padding-left: 5px;">
                      <span style="float: left;"><svg-icon :icon-class="searchForm.databaseType" /></span>
                    </span>
                  </template>
                  <el-option
                    v-for="item in [{ value: 'all' }, { value: 'mysql' }]"
                    :key="item.value"
                    :label="item.value"
                    :value="item.value"
                  >
                    <span style="float: left;"><svg-icon :icon-class="item.value" /></span>
                    <span style="padding-left: 6px;">{{ item.value }}</span>
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item prop="hostAlias" style="margin-bottom:0">
                <el-button size="small" native-type="submit" type="primary" @click="searchClick">搜索</el-button>
                <el-button size="small" @click="resetForm('searchForm')">清空</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </el-header>
    <el-main style="padding-top: 30px">
      <el-card shadow="never">
        <el-button type="primary" size="mini" @click="addDatabaseDrawer = true">添加数据库</el-button>
        <el-row>
          <el-table
            :row-style="{height:'0'}"
            :cell-style="{padding:'7px'}"
            style="margin-top: 10px;height: 60vh;overflow:auto;font-weight: lighter;font-size: 13px"
            border
            :data="tableData"
            @selection-change="handleSelectionChange"
          >
            <el-table-column
              label="序号"
              type="index"
              width="70"
              align="center"
            />
            <el-table-column
              prop="dataSourceName"
              label="名称"
              align="center"
            />
            <el-table-column
              prop="created_at"
              label="创建日期"
              align="center"
            />
            <el-table-column
              prop="databaseType"
              label="数据库类型"
              align="center"
            >
              <template slot-scope="scope">
                <span><svg-icon
                  style="margin-right: 6px;"
                  :icon-class="scope.row.databaseType"
                />{{ scope.row.databaseType }} </span>
              </template>
            </el-table-column>
            <el-table-column
              prop="connectionAddress"
              label="连接地址"
              align="center"
            >
              <template slot-scope="scope">
                <span>{{ scope.row.connectionAddress }}:{{ scope.row.connectionPort }}</span>
              </template>
            </el-table-column>
            <el-table-column
              prop="updated_at"
              label="更新日期"
              align="center"
            />
            <el-table-column
              prop="date"
              label="操作"
              align="center"
            >
              <template scope="scope">
                <el-button type="text" size="mini" @click="upDataSource(scope.row.id)">编 辑</el-button>
                <span style="margin-right: 3px;margin-left: 3px">|</span>
                <el-button size="mini" type="text" style="color:#F56C6C" @click="delDataSource(scope.row)">删 除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </el-card>
    </el-main>
    <el-drawer
      :title="(drawerPattern) ? '添加数据库' :'编辑数据库' "
      :visible.sync="addDatabaseDrawer"
      size="40%"
      :before-close="handleClose"
    >
      <el-form
        ref="dataSourceForm"
        :rules="dataSourceRules"
        :model="dataSourceForm"
        style="margin-right: 40px"
        label-width="120px"
        class="labelClass"
      >
        <el-row :gutter="12">
          <el-col :span="24">
            <el-form-item label="数据源名称:" prop="dataSourceName">
              <el-input v-model="dataSourceForm.dataSourceName" placeholder="请输入数据源名称" size="small" />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="数据库类型:" prop="databaseType">

              <el-select
                v-model="dataSourceForm.databaseType"
                :disabled="!drawerPattern"
                placeholder="请选择"
                size="small"
                @change="databasesChange"
              >
                <template #prefix>
                  <span style="padding-left: 5px;">
                    <span style="float: left;"><svg-icon :icon-class="dataSourceForm.databaseType" /></span>
                  </span>
                </template>

                <el-option
                  v-for="item in databaseType"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value"
                >
                  <span style="float: left;"><svg-icon :icon-class="item.value" /></span>
                  <span style="padding-left: 6px;">{{ item.value }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-row>
              <el-col :span="12">
                <el-form-item label="连接地址:" prop="connectionAddress">
                  <el-input v-model="dataSourceForm.connectionAddress" size="small" placeholder="请输入连接地址" />
                </el-form-item>
              </el-col>
              <el-col :span="5" style="margin-left: 10px">
                <el-form-item label=":" label-width="20px" prop="databasesPort">
                  <el-input v-model="dataSourceForm.connectionPort" size="small" type="number" placeholder="3306" />
                </el-form-item>
              </el-col>
            </el-row>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数据库账号:" prop="databaseAccount">
              <el-input v-model="dataSourceForm.databaseAccount" size="small" placeholder="请输入数据库账号" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
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
            <el-form-item style="text-align: right">
              <el-link :underline="false" type="primary" @click="connectionTest">测试连接</el-link>
            </el-form-item>
            <el-form-item style="text-align: center">
              <el-button type="primary" size="small" @click="onSubmit">保 存</el-button>
              <el-button v-if="!drawerPattern" size="small" @click="$router.go(0)">取 消</el-button>
              <el-button v-else size="small" @click="resetForm('dataSourceForm')">重 置</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-drawer>
  </el-container>
</template>

<script>
import { validateIP, validatorPort } from '@/validated/validated'
import {
  dataSourcesConnectionTestPost,
  dataSourcesDatabasesCratePost,
  dataSourcesDatabasesGet,
  dataSourcesDatabasesPut, dataSourcesDelete, dataSourcesInfoGet
} from '@/api/dataSource'

export default {
  name: 'DataSource',
  data() {
    return {
      isOracle: false,
      drawerPattern: true,
      databaseType: [{ value: 'mysql' }],
      fileList: [],
      importDialogVisible: false,
      searchForm: {
        databaseType: 'all',
        dataSourceName: '',
        connectionAddress: '',
        size: 10,
        page: 1
      },
      total: 0,
      addDatabaseDrawer: false,
      tableData: [],
      multipleSelection: [],
      dataSourceForm: {
        databaseType: 'mysql',
        dataSourceName: '',
        connectionAddress: '',
        connectionPort: 3306,
        databaseAccount: '',
        databasePassword: '',
        additionalParameters: ''
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
        databaseAccount: [{ required: true, message: '请输入数据库账号', trigger: 'change' }],
        databasesPassword: [{ required: true, message: '请输入数据库密码', trigger: 'change' }],
        databaseType: [{ required: true, message: '请选择数据库类型', trigger: 'change' }]
      },
      assetDetails: {}
    }
  },
  mounted() {
    this.searchForm = { ...this.searchForm, ...this.$route.query }
    this.showDataSourceGet()
  },
  methods: {
    // showDataSourceGet
    async showDataSourceGet() {
      this.$router.push({ name: 'DataSource', query: { ...this.searchForm }})
      const { result } = await dataSourcesDatabasesGet(this.searchForm)
      this.tableData = result.data
    },
    // onSubmit
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
    // handleSelectionChange 表格多选事件
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    // handleClose 关闭添加数据库对话框
    handleClose(_) {
      this.$confirm('确认关闭？')
        .then(_ => {
          this.$router.go(0)
        })
        .catch(_ => {
        })
    },
    // resetForm 重置表单
    resetForm(formName) {
      this.$refs[formName].resetFields()
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
    // delDataSource 删除资产
    delDataSource(row) {
      this.$confirm('此操作将永久删除改资产信息包括当前的告警指标, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$loading(true)
        dataSourcesDelete(row.id).then(_ => {
          this.$message.success('删除成功!')
          this.$router.go(0)
        }).catch(_ => {
          this.$loading().close()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    // upDataSource 编辑数据库
    async upDataSource(id) {
      const { code, result } = await dataSourcesInfoGet(id)
      if (code === 0) {
        this.dataSourceForm = result
        this.drawerPattern = false
        this.addDatabaseDrawer = true
      }
    },
    searchClick() {
      this.searchForm.page = 1
      this.showDataSourceGet()
    },
    databasesChange(value) {
      this.isOracle = value === 'oracle'
    }
  }
}
</script>

<style scoped>
.labelClass /deep/ .el-form-item__label {
  font-weight: 100;
}
</style>
