<template>
	<a-modal title="全文搜索" body-class="web_search_body"
			 v-model:visible="visible"
			 :footer="false"
	>
		<div class="head">
			<a-input-search ref="refSearch"
							v-model="PageInfo.key"
							@keydown.enter="search"
							@search="search"
							search-button
			/>
		</div>
		<template v-if="data.list.length">
			<div class="body">
				<div class="item" v-for="item in data.list" @click="goTo(item)">
					<div class="title" v-html="item.title"/>
					<div class="content">{{ item.content }}</div>
				</div>
			</div>
			<div class="footer">
				共搜索到{{ data.total }}条数据
			</div>
		</template>

	</a-modal>
	<IconSearch @click="showSearch"/>
</template>

<script setup lang="ts">
import {nextTick, reactive, ref} from "vue";
import {ApiSearch, TypeSearch} from "@/api/search_api";
import {TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";

const refSearch = ref()
const visible = ref(false)
const PageInfo = reactive(TypeApi.PageInfo())
PageInfo.search = "title"
const data = reactive(Base.List<TypeSearch.Search>())

function showSearch() {
	visible.value = true
	nextTick(() => {
		refSearch.value.focus()	// 触发焦点
	})
}

async function search() {
	let result = await ApiSearch.Search(PageInfo)
	Object.assign(data, result.data)
}

function goTo(item: TypeSearch.Search) {
	window.open(`/article/${item.slug}`, "_black")
}
</script>

<style lang="scss">
.web_search_body {
	.body {
		max-height: 50vh;
		overflow-y: auto;
		color: var(--color-text-2);
		padding: 20px 0;

		.item {
			padding: 10px 20px;
			cursor: pointer;
			// 鼠标效果
			&:hover {
				background-color: var(--color-fill-2);
			}

			.title {
				color: var(--color-text-1);
				font-size: 16px;
				margin-bottom: 5px;

				em {
					color: red;
				}
			}
		}
	}

	.footer {
		text-align: center;
		padding: 10px 0;
		color: var(--color-text-2);
	}
}
</style>