<template>
  <div>
    <div>
      <el-row :gutter="12">
        <el-col :span="11" style="width: 400px;">
          <el-input size="small" v-model="sourceDataSource" disabled>
            <template slot="prepend">源端</template>
            <i slot="prefix">
              <svg-icon icon-class="mysql" style="width: 20px;height: 30px"></svg-icon>
            </i>
          </el-input>
        </el-col>
        <el-col :span="2" style="height: 30px;width: 30px;" class="edit_dev">
          <div>
            <span style="font-size: 15px"> -> </span>
          </div>
        </el-col>
        <el-col :span="11" style="width: 400px;">
          <el-input size="small" v-model="targetDataSource" disabled >
            <template slot="prepend">目标端</template>
            <i slot="prefix">
              <svg-icon icon-class="mysql" style="width: 20px;height: 30px"></svg-icon>
            </i>
          </el-input>
        </el-col>
      </el-row>
    </div>
    <el-row :gutter="12" style="margin-top: 20px">
      <el-row>
        <el-col :span="20">
          <el-tabs v-model="activeName">
            <el-tab-pane  name="first">
            <span  slot="label">
              <span>库</span>
              <el-badge :value="12" class="item" size="mini"></el-badge>
            </span>
            </el-tab-pane>
            <el-tab-pane  name="second">
            <span slot="label">
              <span>表</span>
              <el-badge :value="12" class="item" size="mini"></el-badge>
            </span>
            </el-tab-pane>
          </el-tabs>
        </el-col>
        <el-col :span="4">
          <el-input size="small" v-model="search"  placeholder="搜索名称">
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>
      </el-row>
      <el-col :span="24">
        <el-table v-if="activeName === 'first'" :data="sourceDataTable">
          <el-table-column
            prop="name"
            label="源码库名称"
          >
          </el-table-column>
          <el-table-column
            prop="name"
            label="目标表名称"
          >
          </el-table-column>
        </el-table>
        <el-table v-else-if="activeName === 'second'" :data="sourceDataTable">
          <el-table-column
            prop="name"
            label="源码库名称"
          >
          </el-table-column>
          <el-table-column
            prop="name"
            label="源表名称"
          >
          </el-table-column>
          <el-table-column
            prop="name"
            label="目标库名称"
          >
          </el-table-column>
          <el-table-column
            prop="name"
            label="目标表名称"
          >
          </el-table-column>
        </el-table>
      </el-col>
    </el-row>
    <div style="margin-top: 10px">
      <el-button
        size="small"
        @click="$router.push({name:'DataReplicationEdit',query:{stepsActive:'1',taskID:$route.query.taskID}})"
      >
        上一步
      </el-button>
      <el-button size="small" type="primary" @click="">下一步</el-button>
    </div>
  </div>

</template>
<script>
export default {
  name: 'configureMappingObjects',
  props: {
    getDataReplicationTaskInfo: {
      type: Object,
      require: true
    }
  },
  data() {
    return {
      sourceDataSource: '129  192.168.111.129:3306',
      targetDataSource: '129  192.168.111.129:3306',
      sourceDataTable: [],
      activeName: 'first',
      search: ""
    }
  },
  mounted() {

  },
  methods: {

  }
}

</script>
<style scoped>
.edit_dev {
  display: flex;
  align-items: center;
  justify-content: center;
}

/deep/ .el-input.is-disabled .el-input__inner {
  background-color: #ffffff;
  color: #000000;
}

/deep/ .el-menu-item, .el-submenu__title {
  height: 30px;
  line-height: 30px;
}
::v-deep .el-tabs__nav-wrap::after {
  position: static !important;
}
</style>
