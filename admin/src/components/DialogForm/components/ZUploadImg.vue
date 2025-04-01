<template>
  <el-upload
    class="avatar-uploader"
    :action="action"
    :show-file-list="false"
    :on-success="handleSuccess"
    :before-upload="beforeUpload"
  >
    <img v-if="currentValue" :src="currentValue" class="avatar" />
    <svg-icon v-else icon-class="solar:add-square-bold-duotone" class="avatar-uploader-icon" />
  </el-upload>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import SvgIcon from '@/components/SvgIcon/index.vue'
const props = defineProps({
  value: {
    type: String,
    default: '',
  },
  placeholder: {
    type: String,
    default: '',
  },
  required: {
    type: Boolean,
    default: false,
  },
})

const action = import.meta.env.VITE_UPLOAD_URL
const currentValue = ref(props.value)

const emit = defineEmits(['imgChange'])

watch(
  () => props.value,
  (val) => {
    currentValue.value = val
  }
)

const handleSuccess = (response, uploadFile) => {
  console.log(response)
  currentValue.value = response.data
  emit('imgChange', currentValue.value)
}

const beforeUpload = (rawFile) => {
  console.log(rawFile.type)
  if (rawFile.type !== 'image/jpeg' && rawFile.type !== 'image/png') {
    ElMessage.error('Avatar picture must be JPG format!')
    return false
  } else if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('Avatar picture size can not exceed 2MB!')
    return false
  }
  return true
}
</script>

<style scoped lang="scss">
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}
.el-upload:hover {
  border-color: var(--el-color-primary);
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px !important;
  height: 178px !important;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
