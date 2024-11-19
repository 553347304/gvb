<template>
	<div class="search_view">
		<web_nav no-scroll/>
		<main>
			<div class="container">
				<div class="head">
					<a-input class="search_input" v-model="PageInfo.key as string"
							 @keydown.enter="s.Check({label: '标题', value: PageInfo.key as string}, 'title')"
					></a-input>
					<a-button type="primary"
							  @click="s.Check({label: '搜索', value: PageInfo.key as string}, 'title')">搜索
					</a-button>
				</div>
				<div class="action">
					<div class="item">
						<span :class="{active: search.sort === item.value}"
							  @click="s.Check(item, 'sort')"
							  v-for="item in DefineWeb.sortOption">
							{{ item.label }}</span>
					</div>
					<div class="item">
						<span :class="{active: search.category === item.value}"
							  @click="s.Check(item, 'category')"
							  v-for="item in DefineWeb.categoryOption">
							{{ item.label }}</span>
					</div>
					<div class="item">
						<span :class="{active: search.tags === item.value}"
							  @click="s.Check(item, 'tags')"
							  v-for="item in DefineWeb.tagOption">
							{{ item.label }}</span>
					</div>
				</div>

				<div class="source">
					<template v-if="data.list.length">
						<div class="article_list">
							<div class="item" v-for="item in data.list">
								<div class="top"><img :src="item.banner_url" alt=""></div>
								<div class="bottom">
									<div class="title">
										<a href="" v-html="item.title"></a>
									</div>
									<div class="abstract">
										<a-typography-paragraph :ellipsis="{rows: 2,css:true}">
											{{ item.abstract }}
										</a-typography-paragraph>
									</div>
									<div class="data">
										<div class="article_data_col">
											<span><icon-eye/>{{ Random.integer(0, 100) }}</span>
										</div>
										<div class="date">
											<IconClockCircle/>
											{{ Time.Time().NowDay }}
										</div>
									</div>
								</div>
							</div>
						</div>
						<div class="page">
							<a-pagination :total="data.total"
										  @change="s.Data"
										  v-model:current="PageInfo.page"
										  :default-page-size="PageInfo.limit"
										  :show-total="true"
										  show-jumper
							/>
						</div>
					</template>
					<a-empty v-else class="empty"/>
				</div>
			</div>
		</main>
	</div>
</template>

<script setup lang="ts">
import web_nav from "@/components/web/web_nav.vue";
import {reactive} from "vue";
import {Api, TypeApi} from "@/api/api";
import {DefineUtils} from "@/api/api_utils/type";
import {DefineWeb, TypeWeb} from "@/api/web_api";
import {TypeArticle} from "@/api/article_api";
import {Base} from "@/BAIYIN/HTTP";
import {Time} from "@/BAIYIN/time";
import {Random} from "mockjs";
import {Message} from "@arco-design/web-vue";

const PageInfo = reactive(TypeApi.PageInfo())
PageInfo.source = "like"
const search = reactive(TypeWeb.Search())
const data = reactive(Base.List<TypeArticle.List>())

class _search {
	Check = (item: DefineUtils.options, s: string) => {
		if (s === "sort") search.sort = item.value
		if (s === "category") search.category = item.value
		if (s === "tags") search.tags = item.value

		if (s === "sort") {
			PageInfo.sort = item.value as string
		} else {
			PageInfo.search = s
			PageInfo.key = item.value
		}

		this.Data()
		Message.success("查询成功")
	}

	Data = async () => {
		let result = await Api.List(Api.URL.Article, PageInfo)
		data.list = result.data.list
		data.total = result.data.total
	}
}

const s = new _search()
s.Data()
</script>

<style lang="scss">
.search_view {
	height: 100vh;

	main {
		width: 100%;
		display: flex;
		justify-content: center;
		padding-top: 80px;
		color: var(--color-text-2);

		.container {
			width: var(--container_width);

			.head {
				display: flex;
				justify-content: center;
				align-items: center;
				background-color: var(--color-bg-1);
				padding: 20px;

				.search_input {
					width: 400px;
					margin: 0 20px;
				}
			}

			.action {
				background-color: var(--color-bg-1);
				padding: 20px;
				border-radius: 5px;

				.item {
					margin-bottom: 25px;

					span {
						margin-right: 20px;
						cursor: pointer;
						padding: 5px 10px;
						border-right: 5px;


						&.active {
							background-color: var(--active);
							color: white;
						}
					}
				}
			}

			.source {
				margin-top: 20px;
				min-height: calc(100vh - 485px);

				.article_list {
					border-radius: 5px;
					display: grid;
					grid-template-columns: repeat(4, 1fr);
					grid-column-gap: 40px;
					grid-row-gap: 40px;

					.item {
						width: 100%;
						border-radius: 5px;
						overflow-y: hidden;

						.top {
							height: 120px;

							img {
								width: 100%;
								height: 100%;
								object-fit: cover; // 图片自适应大小
							}
						}
					}
				}

				a {
					color: var(--color-text-2);
					text-decoration: none; // 删除a标签下划线

				}

				.bottom {
					padding: 10px 20px;
					background-color: var(--color-bg-1);

					> div {
						margin-bottom: 5px;
					}

					.title {
						font-size: 16px;
						font-weight: 600;

						em {
							color: red;
						}
					}

					.abstract {
						.arco-typography {
							margin-bottom: 0;
							color: var(--color-text-2);
						}
					}

					.data {
						display: flex;
						justify-content: space-between;
					}

				}

				.page {
					margin-top: 20px;
					display: flex;
					justify-content: center;
					padding: 20px 0;
					border-radius: 5px;
					background-color: var(--color-bg-1);
				}

				.empty {
					width: 100%;
					min-height: calc(100vh - 485px);
					background-color: var(--color-bg-1);
					display: flex;
					justify-content: center;
					flex-direction: column; // 反向
					border-radius: 5px;
				}
			}

		}
	}
}
</style>