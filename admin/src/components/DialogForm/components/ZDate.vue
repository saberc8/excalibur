<template>
	<el-date-picker
		v-model="currentValue"
		type="date"
		:placeholder="placeholder"
		@change="change"
		value-format="YYYY-MM-DD"
	/>
</template>

<script lang="ts" setup>
const props = defineProps({
	value: {
		type: String,
		default: '',
	},
	placeholder: {
		type: String,
		default: '',
	},
	required: {
		type: Boolean,
		default: false,
	},
})
const currentValue = ref(props.value)

const emit = defineEmits(['dateChange'])

watch(
	() => props.value,
	(val) => {
		currentValue.value = val
	}
)

const change = (val) => {
	emit('dateChange', val)
}
</script>
