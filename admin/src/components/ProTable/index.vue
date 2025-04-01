<template>
	<div class="table-container">
		<el-card
			><ProForm :searchForm="searchForm" @search-data="searchData"
		/></el-card>
		<el-card class="table-body">
			<ProBody
				ref="proBody"
				:columns="columns"
				:dataSource="data"
				:gridOtherOptions="gridOptions"
				:total="total"
				@reload-data="reloadData"
				@page-num-change="pageNumChange"
			>
				<template #toolbar_buttons>
					<slot name="toolbar_buttons"></slot>
          
				</template>
				<template #toolbar_title>
					<slot name="toolbar_title"></slot>
				</template>
			</ProBody>
		</el-card>
	</div>
</template>

<script setup>
import ProForm from './components/ProForm/index.vue'
import ProBody from './components/ProBody/index.vue'
const props = defineProps({
	dataSource: {
		type: Array,
		default: () => [],
	},
	columns: {
		type: Array,
		default: () => [],
	},
	searchForm: {
		type: Array,
		default: () => [],
	},
	showForm: {
		type: Boolean,
		default: false,
	},
	getListFunc: {
		type: Function,
		default: () => {},
	},
	params: {
		type: Object,
		default: () => {},
	},
	gridOptions: {
		type: Object,
		default: () => {},
	},
})
const data = ref([])
const proBody = ref()
let searchParams = {}
let total = ref(0)
const renderTable = (func, params) => {
	func(params).then((res) => {
		console.log(res, '请求的数据')
		data.value = res.data || res
		total.value = res.total
	})
}
if (typeof props.getListFunc === 'function') {
	renderTable(props.getListFunc, props.params)
}
const reloadData = () => {
	renderTable(props.getListFunc, props.params)
}
const pageNumChange = (e) => {
	searchData(e)
}

const searchData = (e) => {
	searchParams = e
	const params = {
		...props.params,
		...searchParams,
	}
	console.log(params, 'params')
	renderTable(props.getListFunc, params)
}
defineExpose({
	proBody,
	reloadData,
})
</script>

<style lang="scss" scoped>
.table-container {
	margin: 16px;
  .table-body {
	margin-top: 10px;
	min-width: 700px;
}
}

</style>
