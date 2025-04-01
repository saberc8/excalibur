<template>
  <div class="t-chart" v-bind="$attrs">
    <div
      v-show="!formatEmpty"
      class="t-chart-container"
      :id="id"
      ref="echartRef"
    />
    <slot v-if="formatEmpty" name="empty">
      <el-empty v-bind="$attrs" :description="description" />
    </slot>
    <slot></slot>
  </div>
</template>

<script setup name="TChart">
  import {
    onMounted,
    getCurrentInstance,
    ref,
    watch,
    nextTick,
    onBeforeUnmount,
    markRaw,
    useAttrs,
    computed
  } from 'vue'
  import { useResizeObserver, useDebounceFn } from '@vueuse/core'
  import { toLine } from '@/utils'

  const { proxy } = getCurrentInstance()
  const props = defineProps({
    options: {
      type: Object,
      default: () => ({}),
    },
    id: {
      type: String,
      default: () => Math.random().toString(36).substring(2, 8),
    },
    theme: {
      type: String,
      default: '',
    },
    isEmpty: {
      type: [Boolean, Function],
      default: false,
    },
    description: {
      type: String,
      default: '暂无数据',
    },
  })

  const echartRef = ref()
  const chart = ref()
  const emits = defineEmits()
  const events = Object.entries(useAttrs())

  // 图表初始化
  const renderChart = () => {
    chart.value = markRaw(proxy.$echarts.init(echartRef.value, props.theme))
    setOption(props.options)
    // 返回chart实例
    emits('chart', chart.value)

    // 监听图表事件
    events.forEach(([key, value]) => {
      if (key.startsWith('on') && !key.startsWith('onChart')) {
        const on = toLine(key).substring(3)
        chart.value.on(on, (...args) => emits(on, ...args))
      }
    })

    // 监听元素变化
    useResizeObserver(echartRef.value, resizeChart)
  }

  // 重绘图表函数
  const resizeChart = useDebounceFn(
    () => {
      chart.value?.resize()
    },
    300
  )

  // 设置图表函数
  const setOption = useDebounceFn(
    async (data) => {
      if (!chart.value) return
      chart.value.setOption(data, true, true)
      await nextTick()
      resizeChart()
    },
    300
  )

  const formatEmpty = computed(() => {
    if (typeof props.isEmpty === 'function') {
      return props.isEmpty(props.options)
    }
    return props.isEmpty
  })

  watch(
    () => props.options,
    async (nw) => {
      await nextTick()
      setOption(nw)
    },
    { deep: true },
  )

  watch(
    () => props.theme,
    async () => {
      chart.value.dispose()
      renderChart()
    },
  )

  onMounted(() => {
    renderChart()
  })
  onBeforeUnmount(() => {
    chart.value.dispose()
    chart.value = null
  })
</script>

<style scoped>
  .t-chart {
    position: relative;
    width: 100%;
    height: 100%;
  }
  .t-chart-container {
    width: 100%;
    height: 100%;
  }
</style>
