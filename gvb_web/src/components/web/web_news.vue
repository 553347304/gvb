<template>
	<div class="web_news_view">
		<span class="index">{{ data.index }}</span>
		<a class="title" :href="data.link" target="_blank">{{ data.title }}</a>
		<span class="hot">{{ data.hotValue }}</span>
	</div>
</template>

<script setup lang="ts">
import {reactive} from "vue";

const data = reactive({
	index: "1",
	title: "户籍人口第一大省支棱起来了",
	hotValue: "190万",
	link: "",
})
</script>

<style scoped lang="scss">
.web_news_view {
	width: 100%;
	display: flex;
	justify-content: space-between;
	align-items: center;

	.title {
		color: var(--color-text-2);
		text-decoration: none;
		width: 93%;
	}
}
</style>