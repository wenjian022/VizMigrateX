<template>
  <div>
    <el-row>
      <el-button
        type="primary"
        style="float:left;margin-right: 10px"
        size="small"
        @click="$router.push({name:'DataSourceEdit'})"
      >创建数据源
      </el-button>
      <el-form ref="searchForm" style="float:left;" :model="searchForm" @submit.native.prevent>
        <el-input style="width: 300px;margin-right: 10px" size="small" placeholder="搜索名称/IP">
          <el-button slot="append" icon="el-icon-search" @click="searchClick" />
        </el-input>
        <el-select
          v-model="searchForm.label"
          multiple
          filterable
          allow-create
          default-first-option
          placeholder="搜索标签"
          size="small"
          style="margin-right: 10px"
        >
          <el-option
            v-for="item in labelList"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
        <el-select v-model="searchForm.state" size="small" style="width: 100px;margin-right: 10px">
          <el-option
            v-for="item in stateList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select v-model="searchForm.databaseType" size="small" style="width: 150px;margin-right: 10px">
          <el-option
            v-for="item in databaseTypeList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select v-model="searchForm.environment" size="small" style="width: 150px;margin-right: 10px">
          <el-option
            v-for="item in environmentList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-form>
    </el-row>
    <el-row>
      <el-table
        :row-style="{height:'0'}"
        :cell-style="{padding:'7px'}"
        style="margin-top: 10px;height: 60vh;overflow:auto;font-weight: lighter;font-size: 13px"
        border
        :data="DataSourceTableData"
      >
        <el-table-column
          prop="dataSourceName"
          label="数据源名称"
          align="center"
        />
        <el-table-column
          prop="dataSourceName"
          label="所属环境"
          align="center"
        />
        <el-table-column
          prop="created_at"
          label="创建日期"
          align="center"
        />
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
          prop="connectionAddress"
          label="数据库账号"
          align="center"
        >
          <template slot-scope="scope">
            <span>{{ scope.row.connectionAddress }}:{{ scope.row.connectionPort }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="label"
          label="标签"
          align="center"
        >
          <template slot-scope="scope">
            <span>{{ scope.row.connectionAddress }}:{{ scope.row.connectionPort }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="creator"
          label="创建人"
          align="center"
        />

        <el-table-column
          prop="state"
          label="状态"
          align="center"
        >
          <template slot-scope="scope">
            <div v-if="scope.row.status === 0">
              <div class="status-point" style=" background-color:#67C23A" />
              正 常
            </div>

            <div v-if="scope.row.status === 1">
              <div class="status-point" style="background-color:#ff1c1f" />
              异 常
            </div>
          </template>
        </el-table-column>
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
  </div>
</template>
<script>
import { dataSourcesDatabasesGet } from '@/api/dataSource'

export default {
  name: 'DataSourceTab',
  data() {
    return {
      labelList: ['杭州', '上海'],
      environmentList: [{ value: 'all', label: '全部环境' }],
      stateList: [{ value: 2, label: '全部状态' }, { value: 0, label: '正常' }, { value: 1, label: '异常' }],
      databaseTypeList: [{ value: 'all', label: '全部类型' }, { value: 'mysql', label: 'mysql' }],
      searchForm: {
        state: 2,
        databaseType: 'all',
        environment: 'all',
        label: '',
        dataSourceName: '',
        connectionAddress: '',
        size: 10,
        page: 1
      },
      total: 0,
      DataSourceTableData: []
    }
  },
  mounted() {
    this.showDataSourceGet()
  },
  methods: {
    async showDataSourceGet() {
      this.$router.push({ name: 'DataSource', query: { ...this.searchForm }})
      const { result } = await dataSourcesDatabasesGet(this.searchForm)
      this.tableData = result.data
    },
    searchClick() {
      this.searchForm.page = 1
      this.showDataSourceGet()
    }
  }
}
</script>

<style scoped>

.status-point {
  display: inline-block;
  width: 6px;
  height: 6px;
  border-radius: 50%;
}
</style>
