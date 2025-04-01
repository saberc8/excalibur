<template>
  <a-modal
    :visible="localVisible"
    title="创建空间"
    @ok="handleSubmit"
    @cancel="handleCancel"
    :modal-style="{ width: '500px' }"
  >
    <a-form :model="form" ref="formRef" :style="{ width: '100%' }">
      <a-form-item
        field="name"
        label="空间名称"
        required
        :rules="[{ required: true, message: '请输入空间名称' }]"
      >
        <a-input v-model="form.name" placeholder="请输入空间名称" />
      </a-form-item>
      <a-form-item field="description" label="描述">
        <a-textarea
          v-model="form.description"
          placeholder="请输入空间描述"
          :auto-size="{ minRows: 4, maxRows: 8 }"
        />
      </a-form-item>
      <a-form-item
        field="path"
        label="路径"
        required
        :rules="[
          { required: true, message: '请输入路径' },
          { 
            match: /^[a-zA-Z0-9-_]+$/, 
            message: '路径只能包含英文字母、数字、中划线和下划线' 
          }
        ]"
      >
        <a-input v-model="form.path" placeholder="请输入英文路径" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup>
import { ref, watch } from 'vue'
import { Message } from '@arco-design/web-vue'

const props = defineProps({
  visible: Boolean
})

const emit = defineEmits(['update:visible', 'success'])

const localVisible = ref(props.visible)

watch(
  () => props.visible,
  (val) => {
    localVisible.value = val
  }
)

watch(localVisible, (val) => {
  emit('update:visible', val)
})

const formRef = ref(null)
const form = ref({
  name: '',
  description: '',
  path: '',
  status: 1
})

const handleSubmit = async () => {
  const { validate, resetFields } = formRef.value
  try {
    await validate()
    emit('success', form.value)
    resetFields()
  } catch (err) {
    console.error(err)
  }
}

const handleCancel = () => {
  localVisible.value = false
  formRef.value?.resetFields()
}
</script>
