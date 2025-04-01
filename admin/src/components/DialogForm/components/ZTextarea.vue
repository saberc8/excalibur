<template>
  <el-input
    v-model="currentValue"
    :placeholder="placeholder"
    :required="required"
    @change="change"
  >
  </el-input>
</template>

<script lang="ts" setup>
  const props = defineProps<{
    value: string
    placeholder: string
    required: boolean
  }>()
  const currentValue = ref(props.value)

  const emit = defineEmits(['inputChange'])

  watch(
    () => props.value,
    (val) => {
      currentValue.value = val
    },
  )

  const change = (val: string) => {
    emit('inputChange', val)
  }
</script>
