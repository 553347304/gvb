<script setup lang="ts">
import {reactive, ref} from "vue";
import {IconDelete} from "@arco-design/web-vue/es/icon";
import {Api, TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import {TypeArticle} from "@/api/article_api";
import Gvb_comment_list from "@/components/common/gvb_comment_list.vue";
import router from "@/router";
import {useRoute} from "vue-router";

const route = useRoute()
const activeId = ref("")
if (typeof route.query.id === "string") {
	activeId.value = route.query.id;
}

const PageInfo = reactive(TypeApi.PageInfo())
const articleList = reactive(Base.List<TypeArticle.List>())


class Article {
	constructor() {
	}

	List = async () => {
		let result = await Api.List(Api.URL.Article, PageInfo)
		articleList.list = result.data.list
		articleList.total = result.data.total
	}
	Check = async (record: TypeArticle.List) => {
		activeId.value = record.id
		await router.push({query: {id: record.id}})
	}
}

const a = new Article()
a.List()


</script>
<template>
	<div class="comment_list_view">
		<div class="article">
			<a-input-search class="head" placeholder="搜索文章标题"/>
			<template v-for="item in articleList.list">
				<div class="article_list">
					<div :class="{item: true, active: activeId === item.id}"
						 @click="a.Check(item)"
					>
						<div class="title">
							<a-typography-paragraph :ellipsis="{rows:1,showTooltip:true,css:true}">
								{{ item.title }}
							</a-typography-paragraph>
						</div>
						<div class="count">{{ articleList.total }}</div>
						<div class="action">
							<IconDelete/>
						</div>
					</div>
					<div class="page">
						<a-pagination simple :total="articleList.total"/>
					</div>
				</div>
			</template>
		</div>

		<div class="comment">
			<Gvb_comment_list :article-id="activeId"/>
		</div>
	</div>
</template>
<style lang="scss">
.comment_list_view {
	display: flex;

	> div {
		background-color: var(--color-bg-1);
		height: calc(100vh - 130px);
		border-radius: 5px;
		color: var(--color-text-2);
		overflow-x: hidden;
		overflow-y: auto;
	}

	.article {
		width: 280px;
		padding: 20px;
		position: relative;

		.head {
			margin-bottom: 10px;
		}

		.page {
			display: flex;
			justify-content: center;
			position: absolute;
			bottom: 10px;
			left: 50%;
			transform: translateX(-50%);
		}

		.item {
			width: 100%;
			height: 60px;
			padding: 10px 20px;
			display: flex;
			justify-content: space-between;
			align-items: center;
			cursor: pointer;

			.title {
				width: 8rem;

				.arco-typography {
					margin-bottom: 0;
				}
			}

			&.active {
				background-color: var(--color-fill-2);
				border-radius: 5px;
			}
		}
	}

	.comment {
		width: calc(100% - 300px);
		margin-left: 20px;
	}
}
</style>