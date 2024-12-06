<template>
	<div class="index_view">
		<web_nav/>
		<web_banner :data="DefineWeb.bannerIndex"/>
		<main>
			<div class="container">
				<div class="left">
					<gvb_card title="今日热搜">
						<template #head-right>
							<router-link :to="{name: 'news'}">
								<icon-double-right/>更多
							</router-link>
						</template>
						<web_news/>
					</gvb_card>
					<gvb_card title="文章日历">
						<echarts_calendar v-if="isShow" :data="data" :theme="store.theme"/>
					</gvb_card>
					<gvb_card title="文章列表" class="article_card">
						<template #head-right>
							<a-input-search v-model="key" placeholder="搜索文章"
											@search="search" @keydown.enter="search"/>
						</template>
						<web_article_list ref="webArticleList"/>
					</gvb_card>
				</div>
				<div class="right">
					<gvb_card title="独家推广"><web_promotion/></gvb_card>
					<gvb_card title="标签云"><web_tags/></gvb_card>
					<gvb_card title="个人名片"><web_user_card/></gvb_card>
					<gvb_card title="意见反馈"><web_feed_back/></gvb_card>
				</div>
			</div>
		</main>
		<web_footer/>
	</div>
</template>

<script setup lang="ts">
import web_nav from "@/components/web/web_nav.vue";
import Web_banner from "@/components/web/web_banner.vue";
import Gvb_card from "@/components/admin/card/gvb_card.vue";
import Web_footer from "@/components/web/web_footer.vue";
import Web_news from "@/components/web/web_news.vue";
import Web_promotion from "@/components/web/web_promotion.vue";
import Web_user_card from "@/components/web/web_user_card.vue";
import Web_feed_back from "@/components/web/web_feed_back.vue";
import Echarts_calendar from "@/components/echarts/echarts_calendar.vue";
import Web_tags from "@/components/web/web_tags.vue";
import web_article_list from "@/components/web/web_article_list.vue";

import {reactive, ref, watch} from "vue";
import {ApiArticle, TypeArticle} from "@/api/article_api";
import {useStore} from "@/stores";
import {useRouter} from "vue-router";
import {DefineWeb} from "@/api/web_api";

const router = useRouter()
const store = useStore();
const data = reactive<TypeArticle.Calendar>(TypeArticle.Calendar())
const isShow = ref(false)
watch(() => store.theme, () => {
	isShow.value = false
	setTimeout(() => isShow.value = true)
})

async function getDate() {
	isShow.value = false
	let result = await ApiArticle.Calendar()
	let length = result.data.list.length
	data.date = [result.data.list[0][0], result.data.list[length - 1][0]]
	data.total = result.data.list
	isShow.value = true
}

getDate()

const webArticleList = ref()
const key = ref("")

function search() {
	router.push({query: {title: key.value}})
}
</script>

<style lang="scss">
.index_view {

	main {
		width: 100%;
		display: flex;
		justify-content: center;
		margin-top: 20px;

		.container {
			width: var(--container_width);
			display: flex;
			justify-content: space-between;

			.left {
				width: calc(100% - 400px);

				.article_card {
					background: none;

					.head {
						background-color: var(--color-bg-1);
					}

					.body {
						padding: 0;
					}
				}
			}

			> .right {
				width: 380px;
			}

			.gvb_card {
				margin-bottom: 20px;
			}

		}
	}
}


</style>
