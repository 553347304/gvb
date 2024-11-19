<template>
	<div class="gvb_tabs">
		<!-- 鼠标中键删除	@click.middle="closeTab(item)" -->
		<span :class="{gvb_tab:true, active: route.name === item.name}"
			  @click="clickTab(item)"
			  @click.middle="closeTab(item)"
			  v-for="item in tabList" :key="item.name">
			{{ item.title }}
			<IconClose
				@click.stop="closeTab(item)"
				v-if="item.name !== 'home'">
			</IconClose>
		</span>
		<span class="gvb_tab close_all_tab" @click="closeAllTab">全部关闭</span>
	</div>
</template>

<script setup lang="ts">
import {IconClose} from "@arco-design/web-vue/es/icon";
import {type Ref, ref, watch} from "vue";
import {useRoute, useRouter} from "vue-router";

const route = useRoute()
const router = useRouter()

interface tabType {
	name: string
	title: string
}

const tabList: Ref<tabType[]> = ref([{name: "home", title: "首页"}])

function clickTab(item: tabType) {
	router.push({name: item.name})
}

// 监听列表
watch(() => route.name, () => {
	if (!inList(route.name as string)) {
		// console.log(route.name)
		tabList.value.push({
			name: route.name as string,
			title: route.meta.title as string,
		})
		// 当前路由不在列列表
	}
}, {immediate: true})	// 一开始执行一次

// 判断值是否在列表中
function inList(name: string): boolean {
	for (const tab of tabList.value) {
		if (tab.name === name) {
			return true
		}
	}
	return false
}

// 关闭全部
function closeAllTab() {
	tabList.value = [{name: "home", title: "首页"}]
	router.push({name: "home"})
}

// 关闭单个
function closeTab(item: tabType) {
	if (item.name === "home") {
		return
	}
	let index = tabList.value.findIndex((tab) => item.name === tab.name)	// 获取索引
	tabList.value.splice(index, 1)	// 根据索引删  个数
	// 删除后向左切换一个
	if (route.name === item.name) {
		let beforeIndex = tabList.value[index - 1]
		clickTab(beforeIndex)
	}
	console.log(index)
}


// 持久化存储
watch(() => tabList.value.length, () => {
	localStorage.setItem("tabList", JSON.stringify(tabList.value))
})

function loading() {
	let val = localStorage.getItem("tabList")
	if (val === null) {
		return
	}
	let tab = []
	try {
		tab = JSON.parse(val)
	} catch (e) {
		return;
	}
	tabList.value = tab
}

loading()

</script>

<style lang="scss">
.gvb_tabs {
	height: 30px;
	width: 100%;
	border: 1px solid var(--bg);
	padding: 0 20px;
	display: flex;
	align-items: center;
	position: relative;
	background-color: var(--color-bg-1);

	.gvb_tab {
		border-radius: 5px;
		border: 1px solid var(--bg);
		padding: 2px 6px;
		margin-right: 10px;
		cursor: pointer;

		&.active {
			background-color: var(--active);
			color: #FFF;
			border: none;

			svg:hover {
				background-color: rgb(var(--arcoblue-7));
			}
		}

		svg {
			font-size: 12px;
			border-radius: 50%;

			&:hover {
				background-color: var(--bg);
			}
		}
	}

	.close_all_tab {
		position: absolute;
		right: 10px;
	}
}
</style>