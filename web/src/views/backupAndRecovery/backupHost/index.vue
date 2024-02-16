<template>
  <el-container class="mainClass">
    <el-main style="padding-top: 10px">
      <el-card shadow="never">
        <el-row>
          <el-button
            type="primary"
            style="float:left;margin-right: 10px"
            size="small"
            @click="$router.push({name:'BackupHostEdit'})"
          >创建备份主机
          </el-button>
          <el-form ref="searchForm" style="float:left;" :model="searchForm" @submit.native.prevent>
            <el-input style="width: 300px;margin-right: 10px" size="small" placeholder="搜索名称/IP">
              <el-button slot="append" icon="el-icon-search" @click="searchClick" />
            </el-input>

            <el-select v-model="searchForm.state" size="small" style="width: 100px;margin-right: 10px">
              <el-option
                v-for="item in stateList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form>
        </el-row>

        <el-table :data="backupHostList" />
      </el-card>
    </el-main>
  </el-container>

</template>

<script>

export default {
  name: 'BackupHostList',
  data() {
    return {
      stateList: [{ value: 2, label: '全部状态' }, { value: 0, label: '正常' }, { value: 1, label: '异常' }],
      searchForm: {
        state: 2
      },
      backupHostList: []
    }
  },
  methods: {
    searchClick() {

    }
  }
}
</script>
<style scoped>

</style>
