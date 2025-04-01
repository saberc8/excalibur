<template>
  <div class="tree-container">
    <el-tree
      ref="treeRef"
      v-if="data?.length"
      :data="data"
      :props="{
        children: 'children',
        label: 'name'  // 直接在这里指定使用 name 字段作为显示文本
      }"
      :node-key="nodeKey"
      show-checkbox
      check-strictly
      :default-expand-all="defaultExpandAll"
      :default-checked-keys="currentValue"
      @check="handleCheck"
    />
    <div v-else class="empty-text">暂无数据</div>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import type { ElTree } from 'element-plus'

const props = defineProps({
  value: {
    type: Array,
    default: () => []
  },
  data: {
    type: Array,
    default: () => []
  },
  nodeKey: {
    type: String,
    default: 'id'
  },
  defaultExpandAll: {
    type: Boolean,
    default: true
  },
  required: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['treeChange'])

const treeRef = ref<InstanceType<typeof ElTree>>()
const currentValue = ref(props.value)

watch(
  () => props.value,
  (val) => {
    currentValue.value = val
  }
)

// 获取选中的节点值
const handleCheck = (node: any, { checkedKeys }: any) => {
  // 直接使用选中的节点，不再需要处理半选状态
  emit('treeChange', checkedKeys)
}

// 添加更详细的调试日志
watch(
  () => props.data,
  (newData) => {
    console.log('Tree data structure:', {
      data: newData,
      firstItem: newData?.[0],
      label: newData?.[0]?.name
    })
  },
  { immediate: true, deep: true }
)

watch(
  () => props.value,
  (newVal) => {
    console.log('Tree value changed:', newVal)
    currentValue.value = newVal
  },
  { immediate: true }
)
</script>

<style scoped>
.tree-container {
  width: 100%;
  min-height: 200px;
}
.el-tree {
  margin-top: 8px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  padding: 10px;
}
.empty-text {
  text-align: center;
  color: #909399;
  padding: 20px 0;
}
</style>
