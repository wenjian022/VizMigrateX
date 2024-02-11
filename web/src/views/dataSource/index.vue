<template>
  <el-container class="mainClass">
    <el-main style="padding-top: 10px">
      <el-card shadow="never">
        <el-tabs v-model="activeName" @tab-click="handleClick">
          <el-tab-pane label="数据源" name="dataSource">
            <data-source-tab v-if="activeName === 'dataSource'" />
          </el-tab-pane>
          <el-tab-pane label="环境" name="environment">
            <environment-tab v-if="activeName === 'environment'" />
          </el-tab-pane>
          <el-tab-pane label="标签" name="label">
            <label-tab v-if="activeName === 'label'" />
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-main>
  </el-container>
</template>

<script>
import DataSourceTab from '@/views/dataSource/components/DataSourceTab.vue'
import EnvironmentTab from '@/views/dataSource/components/EnvironmentTab.vue'
import LabelTab from '@/views/dataSource/components/LabelTab.vue'

export default {
  name: 'DataSource',
  components: { LabelTab, EnvironmentTab, DataSourceTab },
  data() {
    return {
      activeName: 'dataSource'
    }
  },
  mounted() {
    const name = sessionStorage.getItem('currentTab')
    // 判断是否存在currentTab，即tab页之前是否被点击切换到别的页面
    if (name) {
      this.activeName = name
    }
  },
  methods: {
    handleClick(tab, event) {
      sessionStorage.setItem('currentTab', tab.name)
    }
  }
}
</script>

<style scoped>

</style>
