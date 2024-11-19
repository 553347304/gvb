<template>
	<div class="gvb_title">
		{{ title }}
	</div>
</template>

<script setup lang="ts">
defineProps<{ title: string }>()

</script>

<style scoped lang="scss">
.gvb_title {
	font-size: 16px;
	font-weight: 700;
	display: flex;
	align-items: center;

	// 条图标
	&::before {
		width: 4px;
		height: 1.2rem; // 1rem=1个文字高度
		margin-right: 5px;
		content: "";
		display: inline-block;
		background-color: var(--active);
	}
}
</style>