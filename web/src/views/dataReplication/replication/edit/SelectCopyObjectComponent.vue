<template>
  <div>
    <div class="edit_dev">
      <el-card style="width: 40%;" shadow="never">
        <div>
          <el-input
            size="small"
            style="width: 90%;"
            placeholder="输入关键字进行过滤"
            v-model="filterText"
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
          <span style="margin-left: 10px">源对象</span>
        </div>
        <el-tree
          style="margin-top: 10px"
          :data="targetSourceData"
          show-checkbox
          node-key="id"
          default-expand-all
          :expand-on-click-node="false"
          :filter-node-method="filterNode"
          ref="tree"
        >
          <span class="custom-tree-node" slot-scope="{ node, data }">
            <span>{{ node.label }}
              <el-tag size="mini" type="warning" v-if=" data.type === 0 "> 库对象 </el-tag>
              <el-tag size="mini" v-else> 表对象 </el-tag>
            </span>
          </span>
        </el-tree>
      </el-card>
    </div>
    <div style="margin-top: 10px">
      <el-button
        size="small"
        @click="$router.push({name:'DataReplicationEdit',query:{stepsActive:'0',taskID:$route.query.taskID}})"
      >
        上一步
      </el-button>
      <el-button size="small" type="primary" @click="nextStep2">下一步</el-button>
    </div>
  </div>

</template>

<script>
import { assetDatabasesInformationGet } from '@/api/dataSource'
import { dataReplicationTaskPut } from '@/api/dataReplication'

export default {
  name: 'SelectCopyObjectComponent',
  props: {
    getDataReplicationTaskInfo: {
      type: Object,
      require: true
    }
  },
  watch: {
    getDataReplicationTaskInfo(newVal) {
      this.dataReplicationTaskInfo = newVal
      this.getDatabaseOption()
    },
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
  data() {
    return {
      filterText: '',
      targetSourceData: [],
      editingSteps: 2,
      dataReplicationTaskInfo: {}
    }
  },
  mounted() {

  },
  methods: {
    async getDatabaseOption() {
      const { code, result } = await assetDatabasesInformationGet(this.dataReplicationTaskInfo.sourceDataSource)
      if (code === 0) {
        this.targetSourceData = result
      }
    },
    filterNode(value, data) {
      if (!value) return true
      return data.label.indexOf(value) !== -1
    },
    async nextStep2() {
      const data = this.$refs.tree.getCheckedNodes()
      const copyObjectFrom = {
        sourceObject: []
      }

      data.forEach((item) => {
        if (item.type !== 0) {
          copyObjectFrom.sourceObject.push({ schemaName: item.schemaName, tableName: item.label })
        }
      })
      //
      if (copyObjectFrom.sourceObject.length === 0 ){
        this.$message.warning("源对象不能为空")
        return false
      } else {
        const taskID = Number(this.$route.query.taskID)
        const { code } = await dataReplicationTaskPut(this.editingSteps, taskID, copyObjectFrom)
        if (code === 0) {
          this.$router.push({ name: 'DataReplicationEdit', query: { stepsActive: 2, taskID: taskID } })
        }
      }

    }
  }
}
</script>

<style scoped>
.edit_dev {
  display: flex;
  align-items: center;
  justify-content: center;
}

.edit_dev >>> .el-transfer-panel {
  width: 500px;

}
</style>
