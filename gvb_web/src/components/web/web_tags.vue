<template>
	<div class="web_tags">
		<div v-for="item in list" :class="{item: true, active: active === item.tag}">
			<a href="javascript:void(0)" @click="checkTag(item)">
				<span>{{ item.tag }}</span>
<!--				<span>{{ item.total }}</span>-->
			</a>
		</div>
	</div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {Random} from "mockjs";
import router from "@/router";
import {useRoute} from "vue-router";

interface Tag {
	tag: string
	total: number
}

const list = ref<Tag[]>([
	{tag: "golang", total: Random.integer(0, 100)},
	{tag: "python", total: Random.integer(0, 100)},
	{tag: "redis", total: Random.integer(0, 100)},
	{tag: "gorm", total: Random.integer(0, 100)},
	{tag: "elasticsearch", total: Random.integer(0, 100)},
	{tag: "mysql", total: Random.integer(0, 100)},
])
const route = useRoute()
const active = ref(route.query.tag as string)

function checkTag(item: Tag) {
	if (active.value === item.tag) {
		active.value = ""
	} else {
		active.value = item.tag
	}
	router.push({query: {tag: active.value}})
}
</script>


<style scoped lang="scss">
@import "@/global/mixin.scss";

.web_tags {
	position: relative;

	&::after {
		display: block;
		position: absolute;
		width: 20px;
		height: 20px;
		background-color: var(--active);
		content: "";

		@include animation_box(20px);
	}

	.item {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		color: var(--color-text-2);
		background-color: var(--color-fill-1);
		width: 50%;
		height: 40px;

		&:nth-child(4n + 1), &:nth-child(4n + 2) {
			background-color: var(--color-fill-2);
		}

		&:nth-child(4n + 1), &:nth-child(4n + 3) {
			border-right: 1px solid var(--bg);
		}

		a {
			color: var(--color-text-2);
			text-decoration: none;
		}

		&.active {
			a {
				color: var(--active);
			}
		}

		span:nth-child(1) {
			margin-right: 5px;
		}
	}
}
</style>