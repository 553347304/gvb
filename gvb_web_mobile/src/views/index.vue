<script setup lang="ts">
import {Search, Tab, Tabs, Icon, TextEllipsis} from 'vant';
import {Random} from "mockjs";
import {reactive, ref} from "vue";
import {Api, TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import {TypeArticle} from "@/api/article_api";
import {Time} from "@/BAIYIN/time";
import Gvb_tabbar from "@/components/gvb_tabbar.vue";
import {DefineUtils} from "@/api/api_utils/type";

const params = reactive(TypeApi.PageInfo())

const tags = reactive(Base.List<string>())
const data = reactive(Base.List<TypeArticle.List>())
const tab = ref()

class _Article {
	List = async () => {
		if (params.key === "全部") params.key = ""
		let result = await Api.List(Api.URL.Article, params)
		Object.assign(data, result.data)
		console.log(data)
	}
	Tags = async () => {
		params.search = "tags"
		let result = await Api.Search(Api.URL.Article, params)
		tags.total = result.data.total
		tags.list = ["全部", ...result.data.list]
	}

	search = async () => {
		params.search = "tags"
		await this.List()
	}
}

const a = new _Article()
a.Tags()
a.search()
</script>

<template>
	<div class="base_view index_view">
		<div class="head">
			<Search v-model="params.key" :search="a.search"/>
			<Tabs v-model:active="params.key" @click="a.search">
				<Tab v-for="item in tags.list" :title="item" :name="item"/>
			</Tabs>
		</div>
		<div class="body">
			<div class="article_list">
				<div class="item" v-for="item in data.list">
					<div class="title"><a :href="`/article/${item.id}`">{{ item.title }}</a></div>
					<div class="abstract">
						<TextEllipsis :content="item.abstract"/>
					</div>
					<div class="info">
						<span class="date"><Icon name="clock"/>
							{{ Time.Time(item.created_at).Now }}
						</span>
						<span class="tags">
							<i class="fa fa-tags"/>
							<span v-for="i in item.tags">{{ i }}</span>
						</span>
						<span class="look"><Icon name="eye"/>{{ Random.integer(0, 100) }}</span>
					</div>
				</div>
			</div>
			<gvb_tabbar/>
		</div>
	</div>
</template>

<style scoped lang="scss">
.index_view {

	.head {
		position: fixed;
		top: 0;
		background-color: white;
		z-index: 1;

		.van-search {
			width: 100vw;
		}

		.van-tabs {
			width: 100vw;
		}
	}


	.body {
		padding-top: 110px;
		padding-bottom: 80px;
		min-height: 100vh;
		background-color: var(--bg);
	}


	.article_list {
		margin-top: 10px;

		.item {
			width: 100vw;
			overflow: hidden;
			background-color: white;
			padding: 10px;
			margin-top: 10px;

			.title {
				font-size: 18px;

				a {
					color: var(--active);
				}

			}

			.abstract {
				color: #555;
				margin: 5px 0;
				font-size: 14px;
			}

			.info {

				span {
					i {
						margin-right: 2px;
					}

					color: #444;
				}

				.tags {
					margin-left: 10px;

					> span {
						margin-left: 10px;
					}
				}

				.look {
					margin-left: 10px;
				}
			}

		}

	}
}
</style>
