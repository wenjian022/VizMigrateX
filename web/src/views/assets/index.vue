<template>
  <el-container class="mainClass">
    <el-header style="margin-top: 10px">
      <el-card shadow="never">
        <el-form ref="searchForm" :model="searchForm" @submit.native.prevent>
          <el-row :gutter="10">
            <el-col :span="6">
              <el-form-item prop="assetName" style="margin-bottom:0">
                <span style="font-size: 15px;font-weight: lighter">资产名称: </span>
                <el-input
                  v-model="searchForm.assetName"
                  style="width: 70%;"
                  size="small"
                  placeholder="杭州研发-246主库"
                />
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item prop="hostAlias" style="margin-bottom:0">
                <span style="font-size: 15px;font-weight: lighter">数据库地址: </span>
                <el-input
                  v-model="searchForm.databasesAddress"
                  style="width: 70%;"
                  size="small"
                  placeholder="192.168.11.246"
                />
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item prop="hostAlias" style="margin-bottom:0">
                <span style="font-size: 15px;font-weight: lighter">数据库类型: </span>
                <el-select v-model="searchForm.databasesType" placeholder="请选择" size="small">
                  <el-option
                    v-for="item in [{ value: 'all' }, { value: 'mysql' }, { value: 'postgresql' },{value: 'oracle'}]"
                    :key="item.value"
                    :label="item.value"
                    :value="item.value"
                  />
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
              prop="assetName"
              label="名称"
              align="center"
            />
            <el-table-column
              prop="databasesType"
              label="数据库类型"
              align="center"
            >
              <template slot-scope="scope">
                <el-tag type="info" size="small">{{ scope.row.databasesType }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column
              prop="databasesAddress"
              label="地址"
              align="center"
            />
            <el-table-column
              prop="databasesPort"
              label="端口"
              align="center"
            >
              <template scope="scope">
                <el-tag size="small">{{ scope.row.databasesPort }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column
              prop="updateTime"
              label="更新日期"
              align="center"
            />
            <el-table-column
              prop="date"
              label="操作"
              align="center"
            >
              <template scope="scope">
                <el-button type="text" size="mini" @click="upAssets(scope.row)">编 辑</el-button>
                <span style="margin-right: 3px;margin-left: 3px">|</span>
                <el-button size="mini" type="text" style="color:#F56C6C" @click="delAssets(scope.row)">删 除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </el-card>
    </el-main>
    <el-drawer
      :title="(drawerPattern) ? '添加数据库' :'编辑数据库' "
      :visible.sync="addDatabaseDrawer"
      size="50%"
      :before-close="handleClose"
    >
      <el-form
        ref="addMySQLForm"
        :rules="addMySQLRules"
        :model="addMySQLForm"
        style="margin-right: 40px"
        label-width="120px"
        class="labelClass"
      >
        <el-row :gutter="12">
          <el-col :span="12">
            <el-form-item label="资产名称" prop="assetName">
              <el-input v-model="addMySQLForm.assetName" placeholder="杭州研发-246主库"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数据库类型" prop="databasesType">
              <el-select
                v-model="addMySQLForm.databasesType"
                :disabled="!drawerPattern"
                placeholder="请选择"
                @change="databasesChange"
              >
                <el-option
                  v-for="item in databasesType"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数据库地址" prop="databasesAddress">
              <el-input v-model="addMySQLForm.databasesAddress" placeholder="192.168.11.246"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数据库端口" prop="databasesPort">
              <el-input-number v-model="addMySQLForm.databasesPort" type="number" placeholder="3306"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="用户名" prop="databasesUserName">
              <el-input v-model="addMySQLForm.databasesUserName" placeholder="请输入数据库用户名"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="密码" prop="databasesPassword">
              <el-input
                v-model="addMySQLForm.databasesPassword"
                show-password
                autocomplete=“new-password”
                placeholder="请输入数据库密码"
              />
            </el-form-item>
          </el-col>
          <el-col v-if="!isOracle" :span="24">
            <el-form-item label="默认库" prop="database">
              <el-input
                v-model="addMySQLForm.database"
                placeholder="请输入默认库"
              />
            </el-form-item>
          </el-col>
          <el-col v-else :span="24">
            <el-form-item label="Sid" prop="sid">
              <el-input
                v-model="addMySQLForm.sid"
                placeholder="请输入Sid库"
              />
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item style="text-align: right">
              <el-link :underline="false" type="primary" @click="connectionTest">测试连接</el-link>
            </el-form-item>
            <el-form-item style="text-align: center">
              <el-button type="primary" @click="onSubmit">保 存</el-button>
              <el-button v-if="!drawerPattern" @click="$router.go(0)">取 消</el-button>
              <el-button v-else @click="resetForm('addMySQLForm')">重 置</el-button>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
    </el-drawer>
  </el-container>
</template>

<script>
import {
  assetsConnectionTestPost, assetsDatabasesCratePost,
  assetsDatabasesGet, assetsDatabasesPut,
  assetsDelete
} from '@/api/assets'
import { validateIP, validatorPort } from '@/validated/validated'

export default {
  name: 'DatabasesAssets',
  data() {
    return {
      isOracle: false,
      drawerPattern: true,
      databasesType: [{ value: 'mysql' }],
      fileList: [],
      importDialogVisible: false,
      searchForm: {
        databasesType: 'all',
        assetName: '',
        databasesAddress: '',
        size: 10,
        page: 1
      },
      total: 0,
      addDatabaseDrawer: false,
      tableData: [],
      multipleSelection: [],
      addMySQLForm: {
        databasesType: 'mysql',
        assetName: '',
        databasesAddress: '',
        databasesPort: 3306,
        databasesUserName: '',
        databasesPassword: '',
        database: '/',
        sid: ''
      },
      addMySQLRules: {
        database: [{ required: true, message: '请输入一个默认库', trigger: 'blur' }],
        assetName: [{ required: true, message: '请输入数据库名称', trigger: 'blur' }, {
          max: 128,
          message: '长度不能超过 128 个字符',
          trigger: 'blur'
        }],
        databasesAddress: [{
          required: true,
          message: '请输入正确的数据库地址',
          trigger: 'change',
          validator: validateIP
        }],
        databasesPort: [{
          type: 'number',
          required: true,
          message: '请输入数据库端口',
          trigger: 'change',
          validator: validatorPort
        }],
        databasesUserName: [{ required: true, message: '请输入数据库用户名', trigger: 'change' }],
        databasesPassword: [{ required: true, message: '请输入数据库密码', trigger: 'change' }],
        databasesType: [{ required: true, message: '请选择数据库类型', trigger: 'change' }],
        sid: [{ required: true, message: '请输入sid', trigger: 'change' }]
      },
      assetDetails: {}
    }
  },
  mounted() {
    this.searchForm = { ...this.searchForm, ...this.$route.query }
    this.showAssetsGet()
  },
  methods: {
    // showAssetsGet 查询资产
    async showAssetsGet() {
      this.$router.push({ name: 'DatabasesAssets', query: { ...this.searchForm } })
      const { result } = await assetsDatabasesGet(this.searchForm)
      this.tableData = result.data
    },
    // onSubmit 创建MySQL数据库资产
    onSubmit() {
      this.$refs['addMySQLForm'].validate(async(valid) => {
        if (valid) {
          if (this.drawerPattern) {
            const { code } = await assetsDatabasesCratePost(this.addMySQLForm)
            if (code === 0) {
              this.$message.success('添加成功')
              this.$router.go(0)
            }
          } else {
            const { code } = await assetsDatabasesPut(this.addMySQLForm, this.addMySQLForm.id)
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
      this.$refs['addMySQLForm'].validate(async(valid) => {
        if (valid) {
          const { code } = await assetsConnectionTestPost(this.addMySQLForm)
          if (code === 0) {
            this.$message.success('连接可用')
          }
        }
      })
    },
    // delAssets 删除资产
    delAssets(row) {
      this.$confirm('此操作将永久删除改资产信息包括当前的告警指标, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$loading(true)
        assetsDelete(row.id).then(_ => {
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
    // upAssets 编辑数据库
    upAssets(row) {
      this.drawerPattern = false
      this.addDatabaseDrawer = true
      this.isOracle = row.databasesType === 'oracle'
      this.addMySQLForm = row
    },
    searchClick() {
      this.searchForm.page = 1
      this.showAssetsGet()
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
