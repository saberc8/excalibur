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
			<span style="font-weight: bold; font-size: 20px; color: #000"
				>列表</span
			>
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
		@close="close"
	/>
	<DialogForm 
		v-if="menuVisible"
		:title="'绑定菜单'"
		:visible="menuVisible"
		:formData="menuFormData"
		:renderForm="menuRenderForm"
		:formFunc="handleBindMenu"
		@close="closeMenuDialog"
	/>
</template>

<script setup>
defineOptions({
	name: 'Role',
})
import { ref, onMounted, createVNode, nextTick, computed, shallowRef } from 'vue'
import { getToken } from '@/utils/auth'
import {
	createButton,
	createSpaceGroup,
	createTag,
} from '@/utils/createElement'

import { roleList, updateRole, addRole, bindMenu, getRoleMenus, delRole } from '@/api/system/role'
import { getMenuTreeByRoleId, getMenuTree } from '@/api/system/menu'
import dayjs from 'dayjs'
import {
	ElButton,
	ElMessage,
	ElSpace,
	ElMessageBox,
	ElImage,
} from 'element-plus'
import DialogForm from '@/components/DialogForm/index.vue'
import ProTable from '@/components/ProTable/index.vue'
import router from '@/router'
const proTable = ref()
const visible = ref(false)
const title = ref('新增')
const getListFunc = roleList
const formFunc = ref()
const formData = ref({
  name: '',
  remark: '',
  status: 0
})
const renderForm = [
  {
    field: 'name',
    label: '角色名称',
    type: 'input',
    required: true,
    componentProps: {
      placeholder: '请输入角色名称',
      maxlength: 50
    }
  },
  {
    field: 'status',
    label: '状态',
    type: 'select',
    required: true,
    componentProps: {
      options: [
        { label: '正常', value: 0 },
        { label: '禁用', value: 1 }
      ]
    }
  },
  {
    field: 'remark',
    label: '备注',
    type: 'input',
    componentProps: {
      placeholder: '请输入备注信息',
      maxlength: 100,
      rows: 3
    }
  }
]
const columns = [
  {
    type: 'seq',
    width: 60,
    treeNode: false
  },
  { field: 'id', title: 'ID', width: 80 },
  { field: 'name', title: '角色名称', width: 150 },
  { 
    field: 'status', 
    title: '状态', 
    width: 100,
    formatter: (row) => {
      return row.row.status === 0 ? '正常' : '禁用'
    }
  },
  { field: 'remark', title: '备注' },
  {
    field: 'createAt',
    title: '创建时间',
    width: 180,
    formatter: (row) => {
      return dayjs(row.row.createdAt).format('YYYY-MM-DD HH:mm:ss')
    }
  },
  {
    title: '操作',
    width: 260, // 增加宽度以适应新按钮
    align: 'center',
    fixed: 'right',
    slots: {
      default: ({ row }) => {
        return createSpaceGroup([
          createButton('primary', 'small', '编辑', () => updateColumnData(row)),
          createButton('primary', 'small', '绑定菜单', () => bindMenuBtn(row)),
          createButton('danger', 'small', '删除', () => handleDelete(row))
        ])
      }
    }
  }
]
const showForm = true
// 搜索区域
const searchForm = [
  {
    label: '角色名称',
    field: 'name',
    type: 'input',
    componentProps: {
      placeholder: '请输入角色名称'
    }
  },
  {
    label: '状态',
    field: 'status',
    type: 'select',
    componentProps: {
      options: [
        { label: '全部', value: '' },
        { label: '正常', value: 0 },
        { label: '禁用', value: 1 }
      ]
    }
  }
]
const params = ref({
	pageNum: 1,
	pageSize: 10,
})
const add = async () => {
	title.value = '新增'
	visible.value = true
	formData.value = renderForm.reduce((acc, cur) => {
		acc[cur.field] = ''
		return acc
	}, {})
	formFunc.value = addRole
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
	formFunc.value = updateRole
	visible.value = true
}

// 修改查找节点的辅助函数，返回符合 setCheckedNodes 所需的格式
const findNodesByIds = (tree, ids) => {
  const nodes = []
  const traverse = (items) => {
    items.forEach(item => {
      if (ids.includes(item.id)) {
        nodes.push({
          id: item.id,
          label: item.name
        })
      }
      if (item.children) {
        traverse(item.children)
      }
    })
  }
  traverse(tree)
  return nodes
}

const bindMenuBtn = async (row) => {
  try {
    // 先获取菜单树
    const menuRes = await getMenuTree()
    console.log('菜单树数据:', menuRes)
    
    // 再获取角色已有菜单
    const roleMenuRes = await getRoleMenus(row.id)
    console.log('角色菜单数据:', roleMenuRes)
    
    // 先设置数据，后显示对话框
    menuTreeData.value = Array.isArray(menuRes) ? menuRes : [] // 修正这里的数据获取
    menuFormData.value = {
      menuIds: Array.isArray(roleMenuRes) ? roleMenuRes.map(item => item.id) : []
    }
    
    // 确保数据设置完成后再打开对话框
    await nextTick()
    currentRoleId.value = row.id
    menuVisible.value = true

    console.log('树形组件数据:', {
      treeData: menuTreeData.value,
      checkedKeys: menuFormData.value.menuIds
    })
  } catch (error) {
    console.error('获取菜单数据失败:', error)
    ElMessage.error('获取菜单数据失败')
  }
}

const handleBindMenu = async (formData) => {
  try {
    await bindMenu({
      roleId: currentRoleId.value,
      menuIds: formData.menuIds
    })
    ElMessage.success('菜单绑定成功')
    menuVisible.value = false
    return true
  } catch (error) {
    ElMessage.error('绑定菜单失败')
    return false
  }
}

const closeMenuDialog = () => {
  menuVisible.value = false
  menuFormData.value.menuIds = []
}

// 加载角色菜单
const loadRoleMenus = async (roleId) => {
  try {
    const [menuTree, roleMenus] = await Promise.all([
      getMenuTree(),
      getRoleMenus(roleId)
    ])
    menuTreeData.value = menuTree.data
    if(roleMenus?.data?.length) {
      menuTree.value?.setCheckedKeys(roleMenus.data.map(m => m.id))
    }
  } catch (error) {
    console.error('获取角色菜单失败:', error)
    ElMessage.error('获取角色菜单失败')
  }
}

const deleteFunc = async (id) => {
	ElMessageBox.confirm('确认删除吗？', '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	}).then(async () => {
		const res = await deleteQun(id)
		ElMessage.success('删除成功')
		refreshTable()
	})
}

const handleDelete = (row) => {
  ElMessageBox.confirm(`确定要删除角色"${row.name}"吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await delRole(row.id)
      ElMessage.success('删除成功')
      refreshTable()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  }).catch(() => {})
}

const menuVisible = ref(false)
const menuTreeData = shallowRef([])
const currentRoleId = ref(null)
const menuTree = ref(null)
const defaultProps = {
	children: 'children',
	label: 'name'
}

const menuFormData = ref({
  menuIds: []
})

const menuRenderForm = computed(() => {
  console.log('menuTreeData when computed:', menuTreeData.value)
  return [{
    field: 'menuIds',
    label: '',
    type: 'tree',
    required: true,
    componentProps: {
      data: menuTreeData.value || [],
      nodeKey: 'id',
      props: {
        children: 'children',
        label: 'name' // 确保这里使用 name 字段作为显示文本
      },
      showCheckbox: true,
      defaultExpandAll: true
    }
  }]
})

</script>
