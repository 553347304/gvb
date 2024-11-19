<template>
	<div class="web_article_list">
		<form_article_item v-for="item in data.list" :form="item"/>
		<div class="page">
			<a-pagination :total="data.total"
						  @change="getDate(PageInfo)"
						  v-model:current="PageInfo.page"
						  :default-page-size="PageInfo.limit"
						  :show-total="true"
						  show-jumper
			/>
		</div>
	</div>
</template>

<script setup lang="ts">
import form_article_item from "@/components/admin/form/form_article_item.vue";
import {reactive, watch} from "vue";
import {Api, TypeApi} from "@/api/api";
import {TypeArticle} from "@/api/article_api";
import {Base} from "@/BAIYIN/HTTP";
import {useRoute} from "vue-router";

const route = useRoute()
const PageInfo = reactive(TypeApi.PageInfo())
const data = reactive(Base.List<TypeArticle.Create>())

async function getDate(params?: TypeApi.PageInfo) {
	let result = await Api.List(Api.URL.Article, params)
	data.list = result.data.list
	data.total = result.data.total
}

defineExpose({getDate})

watch(() => route.query, () => {
	PageInfo.source = "like"
	if (typeof route.query.title === "string") {
		PageInfo.search = "title"
		PageInfo.key = route.query.title
	}
	if (typeof route.query.tag === "string") {
		PageInfo.search = "tags"
		PageInfo.key = route.query.tag
	}
	if (typeof route.query.date === "string") {
		PageInfo.search = "date"
		PageInfo.key = route.query.date
	}
	getDate(PageInfo)
}, {deep: true, immediate: true})
</script>

<style lang="scss">
.web_article_list {

	.gvb_article_item {
		background-color: var(--color-bg-1);
		margin-bottom: 20px;

		&:first-child {
			border-radius: 0 0 5px 5px;
		}

		.info {
			.abstract {
				margin-top: 10px;
			}
		}
	}

	.page {
		display: flex;
		justify-content: center;
		padding: 10px;
		border-radius: 5px;
		background-color: var(--color-bg-1);
	}
}
</style>