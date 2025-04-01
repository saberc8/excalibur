<template>
  <el-input-number
    v-model="currentValue"
    :placeholder="placeholder"
    :required="required"
    @change="change"
    style="width: 60%;"
  >
  </el-input-number>
</template>

<script lang="ts" setup>
  const props = defineProps<{
    value: string
    placeholder: string
    required: boolean
  }>()
  const currentValue = ref(props.value)

  const emit = defineEmits(['inputNumberChange'])

  watch(
    () => props.value,
    (val) => {
      currentValue.value = val
    },
  )

  const change = (val: string) => {
    emit('inputNumberChange', val)
  }
</script>
