<template>
	<el-tooltip
		effect="dark"
		:disabled="isShowTooltip"
		:content="content"
		:placement="placement"
	>
		<div class="menu-title" :class="className" @mouseover="onMouseOver">
			<span class="title-content" ref="refName">{{ content }}</span>
		</div>
	</el-tooltip>
</template>
<script setup name="TooltipOver">
import { ref } from 'vue'
const props = defineProps({
	// 显示的文字内容
	content: {
		type: String,
		default: () => {
			return ''
		},
	},
	// 显示文字方向
	placement: {
		type: String,
		default: 'top',
	},
	// 外层框的样式，在传入的这个类名中设置文字显示的宽度
	className: {
		type: String,
		default: () => {
			return ''
		},
	},
})
const isShowTooltip = ref(true)
const refName = ref(null)
function onMouseOver() {
	let parentWidth = refName.value.parentNode.offsetWidth
	let contentWidth = refName.value.offsetWidth
	if (contentWidth > parentWidth) {
		isShowTooltip.value = false
	} else {
		isShowTooltip.value = true
	}
}
</script>
<style lang="scss" scoped></style>
