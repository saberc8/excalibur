<template>
  <el-form ref="proFromRef" name="advanced_search" :model="formState">
    <el-row :gutter="24">
      <template v-for="(item, index) in searchForm" :key="index">
        <el-col v-show="expand || index < 3" :lg="8" :md="8" :sm="12" :xs="24">
          <el-form-item :prop="item.field" :label="item.label">
            <el-input
              v-model="formState[`${item.field}`]"
              :placeholder="item.componentProps.placeholder"
              clearable
            />
          </el-form-item>
        </el-col>
      </template>
      <el-col style="margin-left: auto" :span="8">
        <el-space style="float: right">
          <el-button @click="resetForm(proFromRef)">重置</el-button>
          <el-button type="primary" @click="submitForm(proFromRef)"
            >查询</el-button
          >
          <a v-if="searchForm.length > 2" style="font-size: 12px" @click="expand = !expand">
            <template v-if="expand">
              收起
              <svg-icon icon-class="solar:alt-arrow-up-bold-duotone" />
            </template>
            <template v-else>
              展开
              <svg-icon icon-class="solar:alt-arrow-down-bold-duotone" />
            </template>
          </a>
        </el-space>
      </el-col>
    </el-row>
  </el-form>
</template>
<script setup>
  import SvgIcon from '@/components/SvgIcon/index.vue'
  defineProps({
    searchForm: {
      type: Array,
      default: () => []
    }
  })
  const emit = defineEmits(['searchData', 'resetData'])
  const expand = ref(false)
  const proFromRef = ref()
  const formState = reactive({})
  const submitForm = async (formEl) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
      if (valid) {
        // 删除formSate的属性
        const data = JSON.parse(JSON.stringify(formState))
        emit('searchData', data)
      } else {
        console.log('error submit!', fields)
      }
    })
  }

  const resetForm = (formEl) => {
    console.log('formEl', formEl)
    if (!formEl) return
    formEl.resetFields()
    emit('resetData')
  }
</script>

<style scoped>
:deep(.icon-wrapper) {
  vertical-align: middle;
  width: 14px;
  height: 14px;
}
</style>
