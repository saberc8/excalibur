<template>
  <Icon v-if="iconClass" :icon="iconClass" :width="size" :height="size" class="icon-wrapper" />
</template>

<script setup>
import { Icon } from '@iconify/vue'

defineProps({
  iconClass: {
    type: String,
    required: true
  },
  size: {
    type: [String, Number],
    default: '16'
  }
})
</script>
