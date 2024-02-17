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
        <el-button type="primary" size="mini" @click="$router.push({name:'DataReplicationEdit'})">创建数据复制
        </el-button>
        <el-row>
          <el-table
            :row-style="{height:'0'}"
            :cell-style="{padding:'7px'}"
            style="margin-top: 10px;height: 60vh;overflow:auto;font-weight: lighter;font-size: 13px"
            border
            :data="tableData"
          >
            <el-table-column
              label="任务名称"
              prop="taskName"
              align="center"
            />
            <el-table-column
              label="任务创建时间"
              prop="created_at"
              align="center"
            />
            <el-table-column
              label="源数据源"
              prop="sourceDataSource"
              align="center"
            >
              <template scope="scope">

              </template>
            </el-table-column>
            <el-table-column
              label="目标数据源"
              prop="targetDataSource"
              align="center"
            >
              <template scope="scope">
                <el-tag type="warning" size="small">{{ scope.row.targetDataSource.assetName }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column
              label="复制方式"
              prop="copyMethod"
              align="center"
            >
              <template scope="scope">
                <el-tag size="small">全量复制</el-tag>
              </template>
            </el-table-column>
            <el-table-column
              prop="date"
              label="操作"
              align="center"
            >
              <template scope="scope">
                <div v-if="scope.row.editingSteps !== 6">
                  <el-button
                    type="text"
                    size="mini"
                    @click="$router.push({name:'DataReplicationEdit',query:{stepsActive:scope.row.editingSteps,taskID:scope.row.id}})"
                  >
                    编 辑
                  </el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </el-row>
      </el-card>
    </el-main>
  </el-container>
</template>

<script>
import { dataReplicationTaskListGet } from '@/api/dataReplication'

export default {
  name: 'DataReplication',
  data() {
    return {
      searchForm: {
        taskName: '',
        size: 10,
        page: 1
      },
      tableData: []
    }
  },
  mounted() {
    this.searchForm = { ...this.searchForm, ...this.$route.query }
    this.showDataReplicationGet()
  },
  methods: {
    async showDataReplicationGet() {
      this.$router.push({ name: 'DataReplication', query: { ...this.searchForm } })
      const { result } = await dataReplicationTaskListGet(this.searchForm)
      this.tableData = result.data
    },
    searchClick() {
      this.searchForm.page = 1
      this.showDataReplicationGet()
    },
    // resetForm 重置表单
    resetForm(formName) {
      this.$refs[formName].resetFields()
    }
  }
}
</script>

<style></style>
