<template>
	<div class="article_view">
		<web_nav/>
		<web_banner :data="{
			slogan: data.user_name,
			abstract: data.abstract,
			image: [data.banner_url],
		}"/>
		<main>
			<div class="container">
				<div class="article_container">
					<div class="head">
						<div class="title">{{ data.title }}</div>
						<div class="date">发布时间: {{ data.created_at }}</div>
						<div class="tag">
							<a-tag color="red" v-for="item in data.tags">{{ item }}</a-tag>
						</div>
					</div>

					<article>
						<MdPreview v-model="data.content" :editor-id="id"></MdPreview>
					</article>
					<!--					<gvb_card title="文章目录" class="article_dict">-->
					<!--						<MdCatalog :editor-id="id"-->
					<!--								   :scroll-element="scrollElement"-->
					<!--								   :offset-top="80"-->
					<!--								   :scroll-element-offset-top="80"-->
					<!--								   :theme="store.themeString"-->
					<!--						></MdCatalog>-->
					<!--					</gvb_card>-->

					<!--					<div class="next_prev">-->
					<!--						<span class="prev">上一篇:<a href="">xxx</a></span>-->
					<!--						<span class="next">下一篇:<a href="">xxx</a></span>-->
					<!--					</div>-->
					<gvb_comment_list :article-id="id"/>
				</div>
			</div>
		</main>
		<web_footer/>
	</div>
</template>

<script setup lang="ts">
import Web_nav from "@/components/web/web_nav.vue";
import Web_banner from "@/components/web/web_banner.vue";
import {useRoute} from "vue-router";
import Web_footer from "@/components/web/web_footer.vue";
import gvb_comment_list from "@/components/common/gvb_comment_list.vue";
import {Api} from "@/api/api";
import {nextTick, onMounted, reactive, ref} from "vue";
import {TypeArticle} from "@/api/article_api";
import {MdCatalog, MdPreview} from "md-editor-v3";
import Gvb_card from "@/components/admin/card/gvb_card.vue";
import {useStore} from "@/stores";


const scrollElement = document.documentElement	// 滚动条
const store = useStore()
const route = useRoute()
const id = ref("")
if (typeof route.params.id === 'string') {
	id.value = route.params.id
}

const data = reactive(TypeArticle.List())

async function getDate() {
	let result = await Api.List(Api.URL.Article, {
		search: "_id",
		key: id.value,
	})
	Object.assign(data, result.data.list[0])
}

getDate()

onMounted(() => {
	let hash = route.hash
	if (hash !== "") {
		nextTick(() => {
			let dom = document.querySelector(hash)
			console.log(dom)
		})
	}
})
</script>

<style scoped lang="scss">
.article_view {
	main {
		width: 100%;
		display: flex;
		justify-content: center;
		background-color: var(--bg);
		padding-top: 20px;
		padding-bottom: 20px;

		.container {
			width: var(--container_width);

			.article_container {

				.head {
					border-radius: 5px 5px 0 0;
					margin-bottom: 1px;
					background-color: var(--color-bg-1);
					padding: 20px;
					display: flex;
					flex-direction: column; // 设置轴向为从上到下
					align-items: center; // 水平居中

					.title {
						font-size: 20px;
						font-weight: 600;
						margin-bottom: 10px;
					}

					.date {
						margin-bottom: 5px;
						color: var(--color-text-2);
					}

					.tag {
						.arco-tag {
							margin-right: 10px;

							&:last-child {
								margin-right: 0;
							}
						}
					}
				}

				.article_dict {
					margin-top: 20px;
				}

				article {
					padding: 20px;
					margin-bottom: 20px;
					background-color: var(--color-bg-1);
				}

				.next_prev {
					background-color: var(--color-bg-1);
					border-radius: 0 0 5px 5px;
					padding: 20px;
					margin-top: 1px;
					margin-bottom: 20px;
					justify-content: space-between;

					a {
						color: var(--color-text-2);
						text-decoration: none; // 文本的装饰线
					}
				}
			}
		}
	}
}
</style>