<template>
	<ProTable
		ref="proTable"
		:dataSource="dataSource"
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
		:showGeoLink="false"
		@close="close"
	></DialogForm>
	<DialogForm
		v-if="roleVisible"
		title="绑定角色"
		:visible="roleVisible"
		:formData="roleFormData"
		:renderForm="roleRenderForm"
		:formFunc="bindUserRole"
		@close="closeRole"
	></DialogForm>
</template>

<script setup>
defineOptions({
	name: 'User',
})
import { ref, onMounted, createVNode } from 'vue'
import { getToken } from '@/utils/auth'
import {
	createButton,
	createSpaceGroup,
	createTag,
} from '@/utils/createElement'

import { userList, updateUser, addUser, changePassword, delUser, bindUserRole, getUserRoles } from '@/api/system/user'
import { roleList } from '@/api/system/role'

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
const dataSource = ref([])
const proTable = ref()
const visible = ref(false)
const title = ref('新增')
const getListFunc = userList
const formFunc = ref()
const formData = ref({
	username: '',
	nickname: '',
})
const renderForm = [
  {
    field: 'username',
    label: '账号',
    type: 'input',
    required: true,
    componentProps: {
      placeholder: '请输入账号',
      maxlength: 50
    }
  },
  {
    field: 'nickname',
    label: '昵称',
    type: 'input',
    required: true,
    componentProps: {
      placeholder: '请输入昵称',
      maxlength: 50
    }
  },
  {
    field: 'avatar',
    label: '头像',
    type: 'upload-img',
    componentProps: {
      placeholder: '请上传头像'
    }
  },
  {
    field: 'email',
    label: '邮箱',
    type: 'input',
    componentProps: {
      placeholder: '请输入邮箱',
      maxlength: 50
    }
  },
  {
    field: 'mobile',
    label: '手机号',
    type: 'input',
    componentProps: {
      placeholder: '请输入手机号',
      maxlength: 30
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
    field: 'sort',
    label: '排序',
    type: 'input-number',
    componentProps: {
      placeholder: '请输入排序号',
      min: 0
    }
  }
]
const columns = [
	{
		type: 'seq',
		width: 60,
		treeNode: false, // 开启树图表
	},
	{ field: 'id', title: 'ID', width: 80 },
	{
		field: 'avatar',
		title: '头像',
		width: 140,
		slots: {
			default: ({ row }) => {
				return createVNode(
					ElImage,
					{
						src: row.avatar,
						fit: 'contain',
						style: 'width: 60px; height: 60px',
						previewSrcList: [row.avatar],
					},
					{
						default: () => '暂无图片',
					}
				)
			},
		},
	},
	{ field: 'username', title: '登录账号', width: 140 },
	{ field: 'nickname', title: '昵称', width: 140 },
	{
		field: 'createAt',
		title: '创建时间',
		width: 180,
		formatter: (row) => {
			return dayjs(row.row.createdAt).format('YYYY-MM-DD HH:mm:ss')
		},
	},
	{
		field: 'updateAt',
		title: '更新时间',
		width: 180,
		formatter: (row) => {
			return row.row.updatedAt
				? dayjs(row.row.updatedAt).format('YYYY-MM-DD HH:mm:ss')
				: dayjs(row.row.createAt).format('YYYY-MM-DD HH:mm:ss')
		},
	},
  { 
    field: 'status', 
    title: '状态', 
    width: 100,
    formatter: (row) => {
      return row.row.status === 0 ? '正常' : '禁用'
    }
  },
	{
		title: '操作',
		width: 380,
		align: 'center',
		fixed: 'right',
		slots: {
			default: ({ row }) => {
				return createSpaceGroup([
					createButton('primary', 'small', '编辑', () => updateColumnData(row)),
					createButton('success', 'small', '绑定角色', () => showBindRole(row)),
					createButton(row.status === 0 ? 'warning' : 'success', 'small', 
            row.status === 0 ? '禁用' : '启用', 
            () => toggleUserStatus(row)
          ),
					createButton('warning', 'small', '改密码', () => changePasswordModel(row)),
					createButton('danger', 'small', '删除', () => handleDelete(row))
				])
			},
		},
	},
]

const showForm = true
// 搜索区域
const searchForm = [
	{
		label: '登录名',
		field: 'username',
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
  title.value = '新增用户'
  visible.value = true
  formData.value = {
    username: '',
    nickname: '',
    password: '123456', // 默认密码
    avatar: '',
    email: '',
    mobile: '',
    status: 0,
    sort: 1,
    userType: 0
  }
  formFunc.value = addUser
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
	formFunc.value = updateUser
	visible.value = true
}

const changePasswordModel = (row) => {
  title.value = '修改密码'
  visible.value = true
  formData.value = {
    id: row.id,
    password: '',
    confirmPassword: ''
  }
  renderForm.value = [
    {
      field: 'password',
      label: '新密码',
      type: 'input',
      required: true,
      componentProps: {
        type: 'password',
        placeholder: '请输入新密码'
      }
    },
    {
      field: 'confirmPassword',
      label: '确认密码',
      type: 'input',
      required: true,
      componentProps: {
        type: 'password',
        placeholder: '请再次输入新密码'
      }
    }
  ]
  formFunc.value = changePassword
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

const bindData = (row) => {
	router.push({
		path: `/qun/qunBindData/${row.qunId}`,
		query: {
			name: row.name,
		},
	})
}

const detailData = (row) => {
	router.push({
		path: `/qun/qunDetailData/${row.qunId}`,
		query: {
			name: row.name,
		},
	})
}

const roleVisible = ref(false)
const roleFormData = ref({
  userId: '',
  roleIds: []
})

const roleRenderForm = [
  {
    field: 'roleIds',
    label: '角色',
    type: 'select-dynamic',
    placeholder: '请选择角色',
    required: true,
    componentProps: {
      multiple: true,
      renderFunc: roleList,
      renderParams: {},
      formatData: {
        label: 'name',
        value: 'id'
      }
    }
  }
]

const showBindRole = async (row) => {
  try {
    // 先获取用户当前的角色
    const userRoles = await getUserRoles(row.id)
    console.log('用户当前角色:', userRoles)
    
    // 设置表单数据
    roleVisible.value = true
    roleFormData.value = {
      userId: row.id,
      roleIds: userRoles.map(role => role.id) || []
    }
    
    console.log('角色表单数据:', roleFormData.value)
  } catch (error) {
    console.error('获取用户角色失败:', error)
    ElMessage.error('获取用户角色失败')
  }
}

const bindRole = async (formData) => {
  try {
    await bindUserRole({
      userId: formData.userId,
      roleIds: formData.roleIds
    })
    ElMessage.success('角色绑定成功')
    roleVisible.value = false
    return true
  } catch (error) {
    ElMessage.error('角色绑定失败')
    return false
  }
}

const closeRole = (refresh) => {
  roleVisible.value = false
  roleFormData.value = { userId: '', roleIds: [] }  // 重置表单数据
  if(refresh) {
    refreshTable()
  }
}

const handleDelete = (row) => {
  ElMessageBox.confirm(`确认删除用户"${row.nickname}"吗？`, '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await delUser(row.id)
      ElMessage.success('删除成功')
      refreshTable()
    } catch (error) {
      ElMessage.error('删除失败')
    }
  }).catch(() => {})
}

// 切换用户状态
const toggleUserStatus = async (row) => {
  try {
    await updateUser({
      id: row.id,
      status: row.status === 0 ? 1 : 0
    })
    ElMessage.success(`${row.status === 0 ? '禁用' : '启用'}成功`)
    refreshTable()
  } catch (error) {
    ElMessage.error(`${row.status === 0 ? '禁用' : '启用'}失败`)
  }
}

</script>
