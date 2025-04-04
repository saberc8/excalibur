<template>
	<div v-if="!item.isHidden">
		<template
			v-if="
				hasOneShowingChild(item.children, item) &&
				(!onlyOneChild.children || onlyOneChild.noShowingChildren) &&
				!item.alwaysShow
			"
		>
			<app-link
				v-if="onlyOneChild.meta"
				:to="resolvePath(onlyOneChild.path, onlyOneChild.query)"
			>
				<el-menu-item
					:index="resolvePath(onlyOneChild.path)"
					:class="{ 'submenu-title-noDropdown': !isNest }"
				>
					<svg-icon
						:icon-class="onlyOneChild.meta.icon || (item.meta && item.meta.icon)"
					/>
					<template #title>
						<TooltipOver
							:content="onlyOneChild.meta.title"
							placement="right"
						></TooltipOver>
					</template>
				</el-menu-item>
			</app-link>
		</template>

		<el-sub-menu
			v-else
			ref="subMenu"
			:index="resolvePath(item.path)"
			teleported
		>
			<template v-if="item.meta" #title>
				<svg-icon :icon-class="item.meta && item.meta.icon" />
				<TooltipOver :content="item.meta.title" placement="right"></TooltipOver>
			</template>

			<sidebar-item
				v-for="child in item.children"
				:key="child.path"
				:is-nest="true"
				:item="child"
				:base-path="resolvePath(child.path)"
				class="nest-menu"
			/>
		</el-sub-menu>
	</div>
</template>

<script setup>
import { isExternal } from '@/utils/validate'
import AppLink from './Link'
import { getNormalPath } from '@/utils/tools'
import TooltipOver from '@/components/TooltipOver/index.vue'

const props = defineProps({
	// route object
	item: {
		type: Object,
		required: true,
	},
	isNest: {
		type: Boolean,
		default: false,
	},
	basePath: {
		type: String,
		default: '',
	},
})

const onlyOneChild = ref({})

function hasOneShowingChild(children = [], parent) {
	if (!children) {
		children = []
	}
	const showingChildren = children.filter((item) => {
		if (item.isHidden) {
			return false
		} else {
			// Temp set(will be used if only has one showing child)
			onlyOneChild.value = item
			return true
		}
	})

	// When there is only one child router, the child router is displayed by default
	if (showingChildren.length === 1) {
		return true
	}

	// Show parent if there are no child router to display
	if (showingChildren.length === 0) {
		onlyOneChild.value = { ...parent, path: '', noShowingChildren: true }
		return true
	}

	return false
}

function resolvePath(routePath, routeQuery) {
	if (isExternal(routePath)) {
		return routePath
	}
	if (isExternal(props.basePath)) {
		return props.basePath
	}
	if (routeQuery) {
		let query = JSON.parse(routeQuery)
		return {
			path: getNormalPath(props.basePath + '/' + routePath),
			query: query,
		}
	}
	return getNormalPath(props.basePath + '/' + routePath)
}
</script>

<style lang="scss" scoped>
:deep(.icon-wrapper) {
  margin-right: 8px;
  width: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.el-menu-item,
.menu-title {
  overflow: isHidden !important;
  text-overflow: ellipsis !important;
  white-space: nowrap !important;
}

.submenu-title-noDropdown {
  border-radius: 6px;
  &:not(.is-active):hover {
    background: var(--el-menu-active-color) !important;
    color: #fff;
    border-radius: 6px;
    opacity: 0.4;
  }
}

.el-menu-item {
  margin: 0 10px;
  padding-left: 10px !important;
  display: flex;
  align-items: center;
  border-radius: 6px;
  position: relative;

  :deep(.svg-icon) {
    margin-right: 8px;
    width: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  &.is-active {
    background: var(--el-menu-active-color) !important;
    color: #fff;
  }

  &:not(.is-active):hover {
    background: var(--el-menu-active-color) !important;
    color: #fff;
    opacity: 0.4;
  }
}

.el-sub-menu {
  :deep(.el-sub-menu__title) {
    display: flex;
    align-items: center;
    
    .svg-icon {
      margin-right: 8px;
      width: 24px;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    &:hover {
      background-color: rgba(0, 0, 0, 0.06) !important;
    }
  }
}

.nest-menu {
  .el-menu-item {
    margin: 0 10px;
    padding: 0 30px;
    position: relative;

    &::before {
      content: '';
      width: 24px;
      display: inline-block;
    }

    :deep(.svg-icon) {
      position: absolute;
      left: 30px;
      margin-right: 0;
    }
  }
}

:deep(.el-menu--collapse) {
  .el-menu-item,
  .el-sub-menu__title {
    padding: 0 !important;
    text-align: center;

    .svg-icon {
      margin: 0 !important;
      width: 54px !important;
      position: static !important;
    }

    span, i, .title-content {
      display: none !important;
    }
  }

  .el-menu-item.is-active {
    .svg-icon {
      color: var(--el-menu-active-color);
    }
  }

  .submenu-title-noDropdown {
    .svg-icon {
      position: static !important;
      margin: 0 !important;
      width: 54px !important;
    }
  }
}

// 调整子菜单样式，确保不受折叠影响
.el-menu--popup-container {
  :deep(.el-menu-item),
  :deep(.el-sub-menu__title) {
    .svg-icon {
      width: 24px !important;
      margin-right: 8px !important;
      position: static !important;
    }
  }
}
</style>