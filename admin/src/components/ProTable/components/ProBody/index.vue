<template>
	<div>
		<vxe-grid ref="xGrid" v-bind="gridOptions">
			<template #toolbar_buttons>
				<slot name="toolbar_title"></slot>
			</template>

			<template #toolbar_tools>
				<el-space>
					<slot name="toolbar_buttons"></slot>
					<svg-icon
						icon-class="solar:refresh-circle-line-duotone"
						class="refresh-icon"
						@click="reloadData"
					/>
				</el-space>
			</template>

			<template #pager>
				<vxe-pager
					:layouts="[
						'Sizes',
						'PrevJump',
						'PrevPage',
						'Number',
						'NextPage',
						'NextJump',
						'FullJump',
						'Total',
					]"
					:total="total"
					v-model:current-page="tablePage.pageNum"
					v-model:pageNum-pageSize="tablePage.pageSize"
					@page-change="handlePageNumChange"
				/>
			</template>
		</vxe-grid>
	</div>
</template>

<script setup>
import SvgIcon from '@/components/SvgIcon/index.vue'
import { VxeGrid } from 'vxe-table'
const props = defineProps({
	dataSource: {
		type: Array,
		default: () => [],
	},
	columns: {
		type: Array,
		default: () => [],
	},
	total: {
		type: Number,
		default: 0,
	},
	gridOtherOptions: {
		type: Object,
		default: () => ({}),
	},
})
const emit = defineEmits(['reloadData', 'pageNumChange'])
const xGrid = ref()
defineExpose({
	xGrid,
})
let tablePage = ref({
	pageNum: 1,
	pageSize: 10,
})
let data = []
const gridOptions = reactive({
	border: true,
	treeConfig: {
		rowField: 'id',
		childrenField: 'children',
	},
	align: null,
	columnConfig: {
		resizable: true,
	},
	columns: props.columns,
	toolbarConfig: {
		slots: {
			buttons: 'toolbar_buttons',
			tools: 'toolbar_tools',
		},
	},
	pageConfig: {
		pageSize: tablePage.value.pageSize,
		pageNum: tablePage.value.pageNum,
		total: props.total,
	},
	data,
	...props.gridOtherOptions,
})
const handlePageNumChange = ({ currentPage }) => {
	console.log('handlePageNumChange - ', currentPage)
	const tablePageBak = {
		pageNum: currentPage,
		pageSize: tablePage.value.pageSize,
	}
	tablePage.value.pageNum = currentPage
	emit('pageNumChange', tablePageBak)
}
watch(
	() => props.dataSource,
	(val) => {
		console.log('val', val)
		data = val
		gridOptions.data = data
	},
	{ immediate: true }
)
watch(
	() => props.columns,
	(val) => {
		gridOptions.columns = val
	},
	{ immediate: true }
)

const reloadData = () => {
	console.log('redo')
	tablePage = {
		pageNum: 1,
		pageSize: 10,
	}
	emit('reloadData')
}
</script>

<style scoped>
.refresh-icon {
	cursor: pointer;
	width: 20px;
	height: 20px;
	color: #1677ff;
}
</style>
