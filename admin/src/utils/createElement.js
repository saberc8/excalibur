import { createVNode } from 'vue'
import { ElButton, ElSpace, ElTag } from 'element-plus'

export const createButton = (type, size, text, onClick) =>
	createVNode(ElButton, { type, size, onClick }, () => text)

export const createTag = (type, text) =>
	createVNode(ElTag, { type, size: 'small' }, () => text)

export const createSpaceGroup = (children) => {
	return createVNode(
		ElSpace,
		{ size: 10, wrap: true },
		{
			default: () => children,
		}
	)
}
