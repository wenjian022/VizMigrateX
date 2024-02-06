<template>
  <el-container class="mainClass">
    <el-header style="margin-top: 10px">
      <el-card shadow="never">
        <el-link @click="$router.push({name:'DataReplication'})"><i class="el-icon-back" style="font-size: 15px">
          创建数据复制</i></el-link>
      </el-card>
    </el-header>
    <el-main style="padding-top: 30px">
      <el-card shadow="never">
        <el-steps :active="stepsInfo.stepsActive" finish-status="success" simple>
          <el-step title="数据源与目标"></el-step>
          <el-step title="选择复制对象"></el-step>
          <el-step title="配置映射对象"></el-step>
          <el-step title="预检查"></el-step>
          <el-step title="启动任务"></el-step>
        </el-steps>
        <div style="margin-top: 20px;margin-left: 20px">
          <data-source-and-target-assembly
            :get-data-replication-task-info="dataReplicationTaskInfo"
            v-if="stepsInfo.stepsActive===0"
          />
          <select-copy-object-component
            :get-data-replication-task-info="dataReplicationTaskInfo"
            v-else-if="stepsInfo.stepsActive===1"
          />
          <configure-mapping-objects
            :get-data-replication-task-info="dataReplicationTaskInfo"
            v-else-if="stepsInfo.stepsActive===2"
          />
        </div>
      </el-card>
    </el-main>
  </el-container>
</template>

<script>
import dataSourceAndTargetAssembly from '@/views/dataReplication/replication/edit/dataSourceAndTarget'
import SelectCopyObjectComponent from '@/views/dataReplication/replication/edit/SelectCopyObjectComponent.vue'
import { dataReplicationTaskInfoGet } from '@/api/dataReplication'
import ConfigureMappingObjects from '@/views/dataReplication/replication/edit/configureMappingObjects.vue'

export default {
  name: 'DataReplicationEdit',
  components: { ConfigureMappingObjects, SelectCopyObjectComponent, dataSourceAndTargetAssembly },
  watch: {
    $route(to, from) {
      // 监听路由是否变化
      this.stepsInfo.stepsActive = Number(this.$route.query.stepsActive) || 0
      this.stepsInfo.taskID = Number(this.$route.query.taskID) || 0
    }
  },
  data() {
    return {
      stepsInfo: {
        stepsActive: 0,
        taskID: 0
      },
      dataReplicationTaskInfo: {
        taskName: '',
        sourceDataSource: '',
        targetDataSource: '',
        copyMethod: 0,
        copyType: [0, 1],
        editingSteps: 1
      }
    }
  },
  mounted() {
    this.stepsInfo.stepsActive = Number(this.$route.query.stepsActive) || 0
    this.stepsInfo.taskID = Number(this.$route.query.taskID) || 0
    if (this.stepsInfo.taskID !== 0) {
      this.getDataReplicationTaskInfo()
    }
  },
  methods: {
    async getDataReplicationTaskInfo() {
      const { code, result } = await dataReplicationTaskInfoGet(this.stepsInfo.taskID)
      if (code === 0) {
        result.copyType = JSON.parse(result.copyType)
        this.dataReplicationTaskInfo = result
      }
    }
  }
}
</script>

<style scoped lang="scss">

</style>
