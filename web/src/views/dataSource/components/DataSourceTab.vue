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
          prop="data_source_name"
          label="数据源名称"
          align="center"
        >
          <template slot-scope="scope">
            <svg-icon style="margin-right: 3px" :icon-class="scope.row.database_type.split('_')[1]" />
            <el-link @click="$router.push({name:'DataSourceEdit',query:{dataSourceID :scope.row.id}})">
              {{ scope.row.data_source_name }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column
          prop="environment_name"
          label="所属环境"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag :style="{color: scope.row.environment_color}" style="background-color:#f8f8f8;" size="small">
              {{ scope.row.environment_name }}
            </el-tag>
          </template>
        </el-table-column>
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
            <span>{{ scope.row.connection_address }}:{{ scope.row.connection_port }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="database_account"
          label="数据库账号"
          align="center"
        />
        <el-table-column
          prop="label"
          label="标签"
          align="center"
        >
          <template slot-scope="scope">
            <el-tag
              v-for="(item,index) in scope.row.label"
              :key="index"
              style="margin-left: 3px"
              type="info"
              size="mini"
            >
              <span><svg-icon icon-class="tag" /> {{ item }}</span>
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          prop="creator_name"
          label="创建人"
          align="center"
        />

        <el-table-column
          prop="state"
          label="状态"
          align="center"
        >
          <template slot-scope="scope">
            <div v-if="scope.row.state === 0">
              <div class="status-point" style=" background-color:#67C23A" />
              正常
            </div>

            <div v-if="scope.row.state === 1">
              <div class="status-point" style="background-color:#ff1c1f" />
              异常
              <el-tooltip class="item" effect="dark" placement="bottom">
                <template slot="content">
                  <p style="max-width:200px;">{{ scope.row.connection_log }}</p>
                </template>
                <i class="el-icon-warning" />
              </el-tooltip>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          prop="date"
          label="操作"
          align="center"
        >
          <template scope="scope">
            <el-button type="text" size="mini" @click="connectionTest(scope.row.id)">检测可用性</el-button>
            <el-dropdown>
              <span class="el-dropdown-link">
                <i class="el-icon-arrow-down el-icon--right" />
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item @click="$router.push({name:'DataSourceEdit',query:{dataSourceID :scope.row.id}})">
                  编辑
                </el-dropdown-item>
                <el-dropdown-item @click="">详情</el-dropdown-item>
                <el-dropdown-item @click="">删除</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column>
      </el-table>
    </el-row>
  </div>
</template>
<script>
import {
  dataSourcesConnectionTestIDPost,
  dataSourcesDatabasesGet
} from '@/api/dataSource'

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
      this.DataSourceTableData = result.data
    },
    searchClick() {
      this.searchForm.page = 1
      this.showDataSourceGet()
    },
    async connectionTest(rowID) {
      await dataSourcesConnectionTestIDPost(rowID).then(_ => {
        this.$message.success('数据源可用,连接成功')
      }, (err) => {
        console.log(err)
      })
      await this.showDataSourceGet()
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
