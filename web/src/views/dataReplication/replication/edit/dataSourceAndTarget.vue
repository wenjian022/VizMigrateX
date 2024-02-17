<template>
  <div>
    <el-form
      ref="dataSourceAndTargetForm"
      :rules="dataSourceAndTargetRules"
      :model="dataSourceAndTargetForm"
      class="labelClass"
      label-width="160px"
    >
      <el-row :gutter="12">
        <el-col :span="24">
          <el-form-item label="任务名称:" prop="taskName">
            <el-input
              v-model="dataSourceAndTargetForm.taskName"
              size="small"
              style="width: 400px"
              placeholder="请输入任务名称"
            />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="源数据源:" prop="sourceDataSourceID">
            <el-select
              v-model="dataSourceSearchFrom.sourceDataSource.type"
              size="small"
              style="width: 150px;"
            >
              <el-option
                v-for="item in dataSourceTypeList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
                <span v-if="item.icon" style="float: left;margin-right: 5px"><svg-icon :icon-class="item.icon" /></span>
                <span>{{ item.label }}</span>
              </el-option>
            </el-select>
            <el-select
              v-model="dataSourceAndTargetForm.sourceDataSourceID"
              size="small"
              style="width: 250px"
              placeholder="请选择目标源数据源"
            >
              <el-option
                v-for="item in dataSourceOptionList"
                :key="item.id"
                :value="item.id"
                :label="item.data_source_name"
              >
                <el-row>
                  <el-col :span="20">
                    <div>
                      <span style="float: left;"><svg-icon :icon-class="item.database_type.split('_')[1]" /></span>
                      <span style="padding-left: 6px;">{{ item.data_source_name }}</span>
                    </div>
                  </el-col>
                  <el-col :span="4">
                    <el-tag style="background-color:#fff;" :style="{color:item.environment_color}" size="mini">{{ item.environment_name }}</el-tag>
                  </el-col>
                </el-row>

              </el-option>

              <div style="text-align: center">
                <span style="font-size: 12px;color: #97a8be">没有找到? </span>
                <el-button type="text">立即创建</el-button>
              </div>

              <template v-if="dataSourceOptionList.length === 0" slot="empty">
                <div style="text-align: center">
                  <div style="margin-top: 5px">
                    <i style="color: #97a8be;font-size: 25px" class="el-icon-files" />
                  </div>
                  <span style="color: #97a8be;font-size: 12px">暂无数据</span>
                </div>
                <div style="text-align: center;">
                  <span style="font-size: 12px;color: #97a8be">没有找到? </span>
                  <el-button type="text" @click="">立即创建</el-button>
                </div>
              </template>

            </el-select>
            <el-link class="el-icon-refresh" :underline="false" style="font-size: 15px;margin-left: 10px" />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="目标数据源:" prop="targetDataSourceID">
            <el-select
              v-model="dataSourceSearchFrom.targetDataSource.type"
              size="small"
              style="width: 150px;"
            >
              <el-option
                v-for="item in dataSourceTypeList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
                <span v-if="item.icon" style="float: left;margin-right: 5px"><svg-icon :icon-class="item.icon" /></span>
                <span>{{ item.label }}</span>
              </el-option>
            </el-select>
            <el-select
              v-model="dataSourceAndTargetForm.targetDataSourceID"
              size="small"
              style="width: 250px"
              placeholder="请选择目标源数据源"
            >
              <el-option
                v-for="item in dataSourceOptionList"
                :key="item.id"
                :value="item.id"
                :label="item.data_source_name"
              >
                <el-row>
                  <el-col :span="20">
                    <div>
                      <span style="float: left;"><svg-icon :icon-class="item.database_type.split('_')[1]" /></span>
                      <span style="padding-left: 6px;">{{ item.data_source_name }}</span>
                    </div>
                  </el-col>
                  <el-col :span="4">
                    <el-tag style="background-color:#fff;" :style="{color:item.environment_color}" size="mini">{{ item.environment_name }}</el-tag>
                  </el-col>
                </el-row>

              </el-option>

              <div style="text-align: center">
                <span style="font-size: 12px;color: #97a8be">没有找到? </span>
                <el-button type="text">立即创建</el-button>
              </div>

              <template v-if="dataSourceOptionList.length === 0" slot="empty">
                <div style="text-align: center">
                  <div style="margin-top: 5px">
                    <i style="color: #97a8be;font-size: 25px" class="el-icon-files" />
                  </div>
                  <span style="color: #97a8be;font-size: 12px">暂无数据</span>
                </div>
                <div style="text-align: center;">
                  <span style="font-size: 12px;color: #97a8be">没有找到? </span>
                  <el-button type="text" @click="">立即创建</el-button>
                </div>
              </template>

            </el-select>
            <el-link class="el-icon-refresh" :underline="false" style="font-size: 15px;margin-left: 10px" />
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="复制方式:" prop="copyMethod">
            <el-radio v-model="dataSourceAndTargetForm.copyMethod" :label="0">结构复制</el-radio>
            <el-radio v-model="dataSourceAndTargetForm.copyMethod" :label="1">全量复制</el-radio>
            <el-radio v-model="dataSourceAndTargetForm.copyMethod" :label="2">结构+全量复制</el-radio>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="目标库存在同名对象:" prop="sameNameOperation">
            <div v-for="item in sameNameOperationList" :key="item.value" style="margin-right: 10px;float: left">
              <el-radio v-model="dataSourceAndTargetForm.sameNameOperation" style="margin-right:10px;" :label="item.value">
                {{ item.label }}
              </el-radio>
              <el-tooltip class="item" effect="dark" placement="top-start">
                <template slot="content">
                  <p style="max-width:200px;">{{ item.explain }}</p>
                </template>
                <i class="el-icon-question" />
              </el-tooltip>
            </div>
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

import { dataReplicationTaskPost } from '@/api/dataReplication'
import { dataSourcesDatabasesGet } from '@/api/dataSource'

export default {
  name: 'DataSourceAndTargetAssembly',
  props: {
    getDataReplicationTaskInfo: {
      type: Object,
      require: true
    }
  },
  data() {
    const validatorDataSourceState = (rule, value, callback) => {
      this.dataSourceOptionList.forEach((item) => {
        if (value === item.id) {
          if (item.state !== 0) {
            callback(new Error(item.connection_log))
          }
        }
      })
      callback()
    }
    return {
      dataSourceTypeList: [
        { value: 'all', label: '全部类型', icon: '' },
        { value: 'idc_mysql', label: 'MySQL', icon: 'mysql' }
      ],
      sameNameOperationList: [
        { value: 0, label: '预检查报错并停止任务', explain: '预检查阶段，检测到同名对象，报错并停止任务' },
        { value: 1, label: '跳过并继续任务', explain: '预检查阶段，检测到同名对象，提示并不拦截任务。结构复制时，自动跳过同名对象，不复制' },
        { value: 2, label: '删除对象并重建', explain: '预检查阶段，检测到同名对象，提示并不拦截任务。结构复制时，自动删除同名对象，并重新创建' },
        { value: 3, label: '保留结构并清空数据，再覆盖写入', explain: '预检查阶段，检测到同名对象，提示并不拦截任务。保留结构，清空数据，再覆盖写入数据' }
      ],
      dataSourceSearchFrom: {
        sourceDataSource: {
          type: 'all'
        },
        targetDataSource: {
          type: 'all'
        }
      },
      dataSourceAndTargetForm: {
        taskName: '',
        sourceDataSourceID: '',
        targetDataSourceID: '',
        copyMethod: 0,
        sameNameOperation: 0
      },
      editingSteps: 1,
      dataSourceAndTargetRules: {
        taskName: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
        sourceDataSourceID: [
          { required: true, message: '请选择源数据源', trigger: 'blur' },
          { trigger: 'change', validator: validatorDataSourceState }
        ],
        targetDataSourceID: [
          { required: true, message: '请选择目标数据源', trigger: 'blur' },
          { trigger: 'change', validator: validatorDataSourceState }
        ],
        copyMethod: [{ required: true, message: '请选择复制方式', trigger: 'blur' }],
        sameNameOperation: [{ required: true, trigger: 'blur' }]
      },
      dataSourceOptionList: []
    }
  },
  watch: {
    getDataReplicationTaskInfo(newV) {
      this.dataSourceAndTargetForm = newV
    }
  },
  mounted() {
    this.getDatabaseOption()
    this.dataSourceAndTargetForm = { ...this.getDataReplicationTaskInfo }
  },
  methods: {
    async getDatabaseOption() {
      const { code, result } = await dataSourcesDatabasesGet()
      if (code === 0) {
        this.dataSourceOptionList = result.data
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
