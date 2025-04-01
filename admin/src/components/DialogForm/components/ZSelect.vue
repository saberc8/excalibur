<template>
  <el-select
    v-model="currentValue"
    :placeholder="placeholder"
    :required="required"
    :multiple="multiple"
    style="width: 100%"
    :clearable="true"
    @change="change"
  >
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    >
    </el-option>
  </el-select>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'

const props = defineProps({
  value: {
    type: [String, Number, Array], // 支持字符串、数字和数组类型
    default: ''
  },
  placeholder: {
    type: String,
    default: ''
  },
  required: {
    type: Boolean,
    default: false
  },
  multiple: {
    type: Boolean,
    default: false
  },
  options: {
    type: Array,
    default: () => []
  }
})

const currentValue = ref(props.value)
const emit = defineEmits(['selectChange'])

// 优化 watch 处理
watch(
  () => props.value,
  (newVal) => {
    console.log('select value changed:', newVal)
    currentValue.value = newVal
  },
  { immediate: true }
)

const change = (val: any) => {
  console.log('select changed:', val)
  emit('selectChange', val)
}
</script>
