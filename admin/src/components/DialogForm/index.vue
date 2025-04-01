<template>
	<el-dialog v-model="dialogFormVisible" :title="title" @close="close">
		<el-form :model="form">
			<el-form-item
				:label="item.label"
				:label-width="formLabelWidth"
				v-for="(item, index) in renderFormDeep"
				:key="index"
			>
				<ZInput
					v-if="item.type === 'input'"
					:value="form[`${item.field}`]"
					:placeholder="item.placeholder"
					:required="item.required"
					@inputChange="(val) => changeValue(val, item.field)"
				></ZInput>
				<ZInputNumber
					v-if="item.type === 'input-number'"
					:value="form[`${item.field}`]"
					:placeholder="item.placeholder"
					:required="item.required"
					@inputNumberChange="(val) => changeValue(val, item.field)"
				></ZInputNumber>
				<ZDate
					v-if="item.type === 'date'"
					:value="form[`${item.field}`]"
					:placeholder="item.placeholder"
					:required="item.required"
					@dateChange="(val) => changeValue(val, item.field)"
				></ZDate>
				<ZSelect
					v-if="item.type === 'select'"
					:value="form[`${item.field}`]"
					:placeholder="item.placeholder"
					:required="item.required"
					:options="item.componentProps.options"
					@selectChange="(val) => changeValue(val, item.field)"
				></ZSelect>
				<ZSelect
					v-if="item.type === 'select-dynamic'"
					:value="form[`${item.field}`]"
					:placeholder="item.placeholder"
					:required="item.required"
					:options="item.componentProps.options"
					:multiple="item.componentProps.multiple"
					@selectChange="(val) => changeValue(val, item.field)"
				></ZSelect>
				<ZUploadImg
					v-if="item.type === 'upload-img'"
					:value="form[`${item.field}`]"
					:placeholder="item.placeholder"
					:required="item.required"
					@imgChange="(val) => changeValue(val, item.field)"
				></ZUploadImg>
				<ZTree
					v-if="item.type === 'tree'"
					:value="form[`${item.field}`]"
					:data="item.componentProps.data"
					:props="item.componentProps.props"
					:node-key="item.componentProps.nodeKey"
					:default-expand-all="item.componentProps.defaultExpandAll"
					:required="item.required"
					@treeChange="(val) => changeValue(val, item.field)"
				/>
			</el-form-item>
		</el-form>
		<div class="link" v-if="showGeoLink">
			<a href="https://lbs.qq.com/getPoint/" target="_blank">地理坐标选取</a>
		</div>

		<template #footer>
			<span class="dialog-footer">
				<el-button @click="dialogFormVisible = false">取消</el-button>
				<el-button type="primary" @click="confirmEvent"> 确认 </el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus'
import ZInput from './components/ZInput.vue'
import ZInputNumber from './components/ZInputNumber.vue'
import ZDate from './components/ZDate.vue'
import ZSelect from './components/ZSelect.vue'
import ZUploadImg from './components/ZUploadImg.vue'
import ZTree from './components/ZTree.vue'
import { ref } from 'vue'
const props = defineProps({
	title: {
		type: String,
		default: '',
	},
	visible: {
		type: Boolean,
		default: false,
	},
	formData: {
		type: Object,
		default: () => ({}),
	},
	renderForm: {
		type: Array,
		default: () => [],
	},
	formFunc: {
		type: Function,
		default: () => {},
	},
	showGeoLink: {
		type: Boolean,
		default: false,
	},
})
const emit = defineEmits(['close'])
const dialogFormVisible = ref(props.visible)
const formLabelWidth = '140px'

const renderFormDeep = ref([])
let form = ref()
watch(
	() => props.visible,
	(val) => {
		console.log(val, '-------')
		dialogFormVisible.value = val
		form.value = JSON.parse(JSON.stringify(props.formData))
		console.log(props.renderForm, 'props.renderFormprops.renderForm')
		renderFormDeep.value = props.renderForm
		renderFormDeep.value.forEach(async (item) => {
			if (item.type === 'select-dynamic') {
				const res = await item.componentProps.renderFunc(
					item.componentProps.renderParams
				)
				const formatData = item.componentProps.formatData
				console.log(res.data)
				const listRowOrData = res.data ? res.data : res.data
				const optionList = []
				listRowOrData.forEach((item) => {
					optionList.push({
						label: item[formatData.label],
						value: item[formatData.value],
					})
				})
				item.componentProps.options = optionList
			}
		})
	},
	{
		immediate: true,
	}
)

// const show = computed({
//   get() {
//     return props.value
//   },
//   set(val: boolean) {
//     emit('update:value', val)
//   },
// })

const close = () => {
	emit('close', false)
}

const renderOptions = async (fn) => {
	const res = await fn()
	console.log(res, 'renderOptions-=-=-=-==-=')
	selectDyncmicOptions.value = res.data
}
// onMounted(() => {
// 	console.log('onMounted', props.renderForm);

// })
const confirmEvent = async () => {
	const res = await props.formFunc(form.value)
	console.log(res, 'confirmEvent')
	ElMessage.success(res.msg)
	emit('close', true)
}

const changeValue = (val, field) => {
	console.log(val, field, 'changeValue')
	form.value[field] = val
	console.log(form.value, 'changeValue')
}
</script>
<style lang="scss" scoped>
.link {
	color: #409eff;
	cursor: pointer;
	text-align: right;
	margin-bottom: 20px;
}
</style>
