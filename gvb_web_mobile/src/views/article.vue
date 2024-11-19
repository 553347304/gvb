<script setup lang="ts">
import {useRoute} from "vue-router";
import {Api, TypeApi} from "@/api/api";
import {reactive} from "vue";
import {TypeArticle} from "@/api/article_api";
import {MdPreview} from "md-editor-v3";
import 'md-editor-v3/lib/preview.css'

function goBack() {
	history.back()
}

const route = useRoute()
const params = reactive(TypeApi.PageInfo())
const data = reactive(TypeArticle.List())

async function articleList() {
	params.search = "_id"
	params.key = route.params.id as string
	let result = await Api.List(Api.URL.Article, params)
	Object.assign(data, result.data.list[0])
	console.log(data)
}

articleList()
</script>

<template>

	<div class="article_view">
		<div class="banner" :style="{backgroundImage: `url('${data.banner_url}')`}">
			<div class="head_top">
				<div class="left" @click="goBack">
					<i class="fa fa-angle-left"></i>
				</div>
				<div class="middle">文章详情页</div>
			</div>

			<div class="head_slogan">
				<div class="title">{{ data.title }}</div>
				<div class="abstract">{{ data.abstract }}</div>
			</div>


		</div>
		<div class="body">
			<article>
				<MdPreview v-model="data.content"></MdPreview>
			</article>
		</div>
	</div>
</template>

<style scoped lang="scss">
.article_view {
	height: 100vh;
	background-color: var(--bg);

	.banner {
		background-repeat: no-repeat;
		background-size: cover; // 图片自适应大小
		width: 100vw;
		height: 200px;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
		position: relative;

		.head_top {
			position: absolute;
			top: 0;
			color: white;
			display: flex;
			align-items: center;
			padding: 10px 20px;
			width: 100vw;

			.middle {
				position: absolute;
				top: 10px;
				left: 50%;
				transform: translateX(-50%);
			}
		}

		.head_slogan {
			display: flex;
			justify-content: center;
			align-items: center;
			flex-direction: column;
			font-size: 14px;
			color: white;
			padding-left: 20px;
			padding-right: 20px;

			.title {
				font-size: 18px;
				margin-bottom: 10px;
			}

			// 文本省略
			.abstract {
				display: -webkit-box;
				-webkit-box-orient: vertical;
				-webkit-line-clamp: 2;
				overflow: hidden;
				text-overflow: ellipsis;
			}
		}
	}


}

</style>