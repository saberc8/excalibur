<template>
	<ProTable
		ref="proTable"
		:columns="columns"
		:params="params"
		:searchForm="searchForm"
		:showForm="showForm"
		:getListFunc="getListFunc"
	>
		<template #toolbar_title>
			<span style="font-weight: bold; font-size: 20px; color: #000">列表</span>
		</template>
		<template #toolbar_buttons>
			<el-button type="primary" @click="add"> 新增 </el-button>
		</template>
	</ProTable>
	<DialogForm
		v-if="visible"
		:title="title"
		:visible="visible"
		:formData="formData"
		:renderForm="renderForm"
		:formFunc="formFunc"
		:showGeoLink="false"
		@close="close"
	></DialogForm>
</template>

<script setup>
defineOptions({
	name: 'Menu',
})
import { ref } from 'vue'
import { createButton, createSpaceGroup } from '@/utils/createElement'

import {
	getMenuTree,
	updateMenu,
	addMenu,
	delMenu,
} from '@/api/system/menu'
import dayjs from 'dayjs'
import { ElButton, ElMessage, ElMessageBox } from 'element-plus'
import DialogForm from '@/components/DialogForm/index.vue'
import ProTable from '@/components/ProTable/index.vue'
const proTable = ref()
const visible = ref(false)
const title = ref('新增')
const getListFunc = getMenuTree
const formFunc = ref()
const formData = ref({
	name: '', // 菜单名称
	parentId: 0, // 顶级菜单使用 0
	orderNum: 0, // 显示排序
	path: '', // 路由地址，目录可为空
	component: '', // 组件路径，目录可为空
	query: '', // 路由参数
	isFrame: 0, // 是否外链
	menuType: 'M', // M-目录 C-菜单 F-按钮
	isCatch: 0, // 是否缓存
	isHidden: 0, // 是否隐藏
	perms: '', // 权限标识
	icon: '', // 菜单图标
	status: 0, // 0-正常 1-停用
	remark: '', // 备注信息
})

const renderForm = [
	{
		field: 'menuType',
		label: '菜单类型',
		type: 'select',
		required: true,
		componentProps: {
			options: [
				{ label: '目录', value: 'C' },
				{ label: '菜单', value: 'M' },
				{ label: '按钮', value: 'B' },
			],
		},
	},
	{
		field: 'name',
		label: '菜单名称',
		type: 'input',
		required: true,
		placeholder: '请输入菜单名称',
	},
	{
		field: 'parentId',
		label: '上级菜单',
		type: 'input-number',
		placeholder: '顶级菜单请输入0',
	},
	{
		field: 'orderNum',
		label: '显示排序',
		type: 'input-number',
		required: true,
		placeholder: '请输入显示排序',
	},
	{
		field: 'path',
		label: '路由地址',
		type: 'input',
		placeholder: '目录可为空',
	},
	{
		field: 'component',
		label: '组件路径',
		type: 'input',
		placeholder: '目录可为空',
	},
	{
		field: 'perms',
		label: '权限标识',
		type: 'input',
		placeholder: '请输入权限标识',
	},
	{
		field: 'icon',
		label: '菜单图标',
		type: 'input',
		placeholder: '请输入菜单图标',
	},
	{
		field: 'isFrame',
		label: '是否外链',
		type: 'select',
		componentProps: {
			options: [
				{ label: '否', value: 0 },
				{ label: '是', value: 1 },
			],
		},
	},
	{
		field: 'isCatch',
		label: '是否缓存',
		type: 'select',
		componentProps: {
			options: [
				{ label: '否', value: 0 },
				{ label: '是', value: 1 },
			],
		},
	},
	{
		field: 'isHidden',
		label: '是否隐藏',
		type: 'select',
		componentProps: {
			options: [
				{ label: '显示', value: 0 },
				{ label: '隐藏', value: 1 },
			],
		},
	},
	{
		field: 'status',
		label: '菜单状态',
		type: 'select',
		componentProps: {
			options: [
				{ label: '正常', value: 0 },
				{ label: '停用', value: 1 },
			],
		},
	},
	{
		field: 'remark',
		label: '备注',
		type: 'input',
		placeholder: '请输入备注信息',
	},
]

const columns = [
	{
		type: 'seq',
		width: 60,
	},
	{ field: 'id', title: 'ID', width: 80, treeNode: true },
	{ field: 'name', title: '菜单名称', width: 160 },
	{ field: 'path', title: '路由地址', width: 160 },
	{ field: 'component', title: '组件路径', width: 160 },
	{
		field: 'menuType',
		title: '菜单类型',
		width: 100,
		formatter: (row) => {
			const types = {
				C: '目录',
				M: '菜单',
				B: '按钮',
			}
			return types[row.row.menuType] || row.row.menuType
		},
	},
	{
		field: 'status',
		title: '状态',
		width: 100,
		formatter: (row) => {
			return row.row.status === 0 ? '正常' : '停用'
		},
	},
	{ field: 'orderNum', title: '排序', width: 80 },
	{ field: 'perms', title: '权限标识', width: 160 },
	{
		title: '操作',
		width: 280,
		align: 'center',
		fixed: 'right',
		slots: {
			default: ({ row }) => {
				const buttons = []
				
				// 只有目录类型(C)显示新增子菜单按钮
				if (row.menuType === 'C') {
					buttons.push(
						createButton('success', 'small', '新增子菜单', () => addSubMenu(row))
					)
				}
				
				// 添加编辑和删除按钮
				buttons.push(
					createButton('primary', 'small', '编辑', () => updateColumnData(row)),
					createButton('danger', 'small', '删除', () => deleteMenu(row.id))
				)
				
				return createSpaceGroup(buttons)
			},
		},
	},
]
const showForm = true
// 搜索区域
const searchForm = [
	{
		label: '菜单名称',
		field: 'name',
		type: 'input',
		componentProps: {
			placeholder: '请输入',
		},
	},
	{
		label: '昵称',
		field: 'nickname',
		type: 'input',
		componentProps: {
			placeholder: '请输入',
		},
	},
]
const params = ref({
	pageNum: 1,
	pageSize: 10,
})
const add = async () => {
	title.value = '新增菜单'
	visible.value = true
	formData.value = {
		name: '',
		parentId: null,
		orderNum: 0,
		path: '',
		component: '',
		query: '',
		isFrame: 0,
		menuType: 'M',
		isCatch: 0,
		isHidden: 0,
		perms: '',
		icon: '',
		status: 0,
		remark: '',
	}
	formFunc.value = addMenu
}

const close = (e) => {
	visible.value = false
	if (e) {
		refreshTable()
	}
}

const refreshTable = () => {
	setTimeout(() => {
		proTable.value.reloadData()
	}, 500)
}

const updateColumnData = (row) => {
	const rowBak = JSON.parse(JSON.stringify(row))
	// console.log(rowBak)
	delete rowBak._X_ROW_KEY
	title.value = '编辑'
	formData.value = rowBak
	formFunc.value = updateMenu
	visible.value = true
}

// 修改删除处理函数
const deleteMenu = async (id) => {
  ElMessageBox.confirm('删除菜单后，所有角色对应的菜单权限也将被删除，确认要删除吗？', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await delMenu(id)
      ElMessage.success('删除成功')
      refreshTable()
    } catch (error) {
      console.error('删除菜单失败:', error)
      ElMessage.error('删除失败')
    }
  }).catch(() => {
    // 取消删除操作
  })
}

// 新增处理子菜单的函数
const addSubMenu = (parentMenu) => {
    title.value = '新增子菜单'
    visible.value = true
    formData.value = {
        name: '',
        parentId: parentMenu.id, // 设置父菜单ID
        orderNum: 0,
        path: '',
        component: '',
        query: '',
        isFrame: 0,
        menuType: 'M', // 默认为菜单类型
        isCatch: 0,
        isHidden: 0,
        perms: '',
        icon: '',
        status: 0,
        remark: '',
    }
    formFunc.value = addMenu
}
</script>
