<template>
  <div>
    <el-form
      ref="dataSourceAndTargetForm"
      :rules="dataSourceAndTargetRules"
      :model="dataSourceAndTargetForm"
      class="labelClass"
      label-position="top"
    >
      <el-row :gutter="12">
        <el-col :span="24">
          <el-form-item label="任务名称" prop="taskName">
            <el-input
              size="small"
              style="width: 400px"
              v-model="dataSourceAndTargetForm.taskName"
              placeholder="请输入任务名称"
            />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="源数据源" prop="sourceDataSource">
            <el-select size="small" style="width: 400px" v-model="dataSourceAndTargetForm.sourceDataSource"
                       placeholder="请选择"
            >
              <el-option
                v-for="item in dataSourceOptionList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
                <span style="float: left;"><svg-icon :icon-class="item.databaseType"></svg-icon></span>
                <span style="padding-left: 6px;">{{ item.label }}</span>
              </el-option>
            </el-select>

          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="目标数据源" prop="targetDataSource">
            <el-select size="small" style="width: 400px" v-model="dataSourceAndTargetForm.targetDataSource"
                       placeholder="请选择"
            >
              <el-option
                v-for="item in dataSourceOptionList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
                <span style="float: left;"><svg-icon :icon-class="item.databaseType"></svg-icon></span>
                <span style="padding-left: 6px;">{{ item.label }}</span>
              </el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="复制方式" prop="copyMethod">
            <el-radio v-model="dataSourceAndTargetForm.copyMethod" :label="0">单项复制</el-radio>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="复制类型" prop="copyType">
            <el-checkbox-group v-model="dataSourceAndTargetForm.copyType">
              <el-checkbox :label="0">结构复制</el-checkbox>
              <el-checkbox :label="1">全量复制</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-button type="primary" size="small" @click="nextStep">下一步</el-button>
        </el-col>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import { assetsDatabasesSelectGet } from '@/api/assets'
import { dataReplicationTaskPost } from '@/api/dataReplication'

export default {
  name: 'dataSourceAndTargetAssembly',
  props: {
    getDataReplicationTaskInfo: {
      type: Object,
      require: true
    }
  },
  watch: {
    getDataReplicationTaskInfo(newV) {
      this.dataSourceAndTargetForm = newV
    }
  },
  data() {
    return {
      dataSourceAndTargetForm: {
        taskName: '',
        sourceDataSource: '',
        targetDataSource: '',
        copyMethod: 0,
        copyType: [0, 1]
      },
      editingSteps: 1,
      dataSourceAndTargetRules: {
        taskName: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
        sourceDataSource: [{ required: true, message: '请选择源数据源', trigger: 'blur' }],
        targetDataSource: [{ required: true, message: '请选择目标数据源', trigger: 'blur' }],
        copyMethod: [{ required: true, message: '请选择复制方式', trigger: 'blur' }],
        copyType: [{ required: true, message: '至少选择一种复制类型', trigger: 'blur' }]
      },
      dataSourceOptionList: []
    }
  },
  mounted() {
    this.getDatabaseOption()
    this.dataSourceAndTargetForm = { ...this.getDataReplicationTaskInfo }
  },
  methods: {
    async getDatabaseOption() {
      const { code, result } = await assetsDatabasesSelectGet()
      if (code === 0) {
        this.dataSourceOptionList = result
      }
    },
    async nextStep() {
      await this.$refs['dataSourceAndTargetForm'].validate((valid) => {
        if (!valid) {
          return false
        } else {
          // 判断2个数据源是不是一样
          if (this.dataSourceAndTargetForm.sourceDataSource === this.dataSourceAndTargetForm.targetDataSource) {
            this.$confirm('源数据源 与 目标数据源 是一致的,是否继续?', '警告', {
              confirmButtonText: '继续',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(async _ => {
              // 提交操作
              const { code, result } = await dataReplicationTaskPost(this.dataSourceAndTargetForm)
              if (code === 0) {
                // 添加成功后跳转到对应的流程
                this.$router.push({
                  name: 'DataReplicationEdit',
                  query: { stepsActive: this.editingSteps, taskID: result.rowID }
                })
              }
            }).catch(_ => {
              return false
            })
          }
        }
      })
    }
  }
}
</script>

<style scoped>
.labelClass /deep/ .el-form-item__label {
  font-weight: 100;
}
</style>
