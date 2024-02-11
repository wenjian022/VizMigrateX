<template>
  <div>
    <el-row>
      <el-button
        type="primary"
        style="float:left;margin-right: 10px"
        size="small"
        @click="createEnvDialog = true"
      >创建环境
      </el-button>
      <el-form ref="searchForm" style="float:left;" :model="searchForm" @submit.native.prevent>
        <el-input v-model="searchForm.name" style="width: 300px;margin-right: 10px" size="small" placeholder="搜索环境名称">
          <el-button slot="append" icon="el-icon-search" @click="searchClick" />
        </el-input>
      </el-form>
    </el-row>

    <el-row>
      <el-table
        :data="envTableList"
        :row-style="{height:'0'}"
        :cell-style="{padding:'7px'}"
        style="margin-top: 10px;height: 60vh;overflow:auto;font-weight: lighter;font-size: 13px"
        border
      >
        <el-table-column
          prop="name"
          label="环境名称"
          align="center"
        />
        <el-table-column
          prop=""
          label="数据源数量"
          align="center"
        />
        <el-table-column
          prop="created_at"
          label="创建日期"
          align="center"
        />
        <el-table-column
          prop="creator"
          label="创建人"
          align="center"
        >
          <template slot-scope="scope">
            <div v-if="!scope.row.creator"><span>-</span></div>
            <div v-else />
          </template>
        </el-table-column>

        <el-table-column
          prop="date"
          label="操作"
          align="center"
        >
          <template scope="scope">
            <div v-if="scope.row.name==='生产'||scope.row.name ==='预发' ">
              <span>-</span>
            </div>
            <div v-else />
            <!--            <el-button type="text" size="mini" @click="upDataSource(scope.row.id)">编 辑</el-button>-->
            <!--            <span style="margin-right: 3px;margin-left: 3px">|</span>-->
            <!--            <el-button size="mini" type="text" style="color:#F56C6C" @click="delDataSource(scope.row)">删 除-->
            <!--            </el-button>-->
          </template>
        </el-table-column>
      </el-table>
    </el-row>

    <el-dialog width="20%" :visible.sync="createEnvDialog">
      <el-form ref="envForm" class="labelClass" :rules="envRules" :model="envForm">
        <el-form-item label="环境名称" prop="name">
          <el-input v-model="envForm.name" size="small" placeholder="请输入环境名称" />
        </el-form-item>
      </el-form>
      <div style="text-align: right">
        <el-button size="small" style="margin-left: 10px" @click="createEnvDialog = false">取 消</el-button>
        <el-button size="small" type="primary" @click="createEnv">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { dataSourcesEnvListGet, dataSourcesEnvPost } from '@/api/dataSource'

export default {
  name: '',
  data() {
    return {
      createEnvDialog: false,
      searchForm: {
        name: ''
      },
      envForm: {
        name: ''
      },
      envRules: { name: [{ required: true, message: '请输入环境名称', trigger: 'change' }] },
      envTableList: []
    }
  },
  mounted() {
    this.getEnv()
  },
  methods: {
    createEnv() {
      this.$refs['envForm'].validate(async(valid) => {
        if (valid) {
          const { code } = await dataSourcesEnvPost(this.envForm)
          if (code === 0) {
            console.log(code)
          }
        }
      })
    },
    async getEnv() {
      this.$router.push({ name: 'DataSource', query: { ...this.searchForm }})
      const { code, result } = await dataSourcesEnvListGet(this.searchForm)
      if (code === 0) {
        this.envTableList = result.data
      }
    },
    searchClick() {
      this.searchForm.page = 1
      this.getEnv()
    }
  }
}
</script>

<style scoped>
.labelClass /deep/ .el-form-item__label {
  font-weight: 100;
}

</style>
