<template>
	<div class="home_view">
		<Gvb_home_welcome/>
		<div class="main">
			<div class="left">
				<Gvb_card title="快捷入口" class="quick_entry">
					<div class="item" v-for="item in quickEntryList">
						<div class="icon" @click="goLink(item)">
							<component :is="item.font"/>
						</div>
						<div class="text">{{ item.text }}</div>
					</div>
				</Gvb_card>
				<echarts_statistics v-if="isShow" :data="data"/>
			</div>
			<div class="right">
				<Gvb_card title="更新日志" class="update_log">
					<div class="item" v-for="(item, index) in updateLogList">
						<div class="line1">
						    <span>
								 <span class="index">{{ index + 1 }}.</span>
								 <span class="content">
									 <a-typography-paragraph :ellipsis="{rows: 1,showTooltip: true,css: true}">
										 {{ item.title }}
									 </a-typography-paragraph>
								 </span>
							</span>
							<span class="date" :title="Time.Time(item.created_at).NowDay">
								{{ Time.Time(item.created_at).NowDay }}
							</span>
						</div>
						<ul class="line2" v-if="item.items">
							<li v-for="li in item.items">{{ li }}</li>
						</ul>
					</div>
				</Gvb_card>
				<!--				<a-button type="primary"></a-button>-->
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import {type Component, reactive, ref, watch} from "vue";
import {IconUser, IconMessage, IconBook, IconUserGroup, IconSettings} from "@arco-design/web-vue/es/icon";
import router from "@/router";
import Gvb_card from "@/components/admin/card/gvb_card.vue";
import Gvb_home_welcome from "@/components/admin/home/gvb_home_welcome.vue";
import echarts_statistics from "@/components/echarts/echarts_statistics.vue";
import {useStore} from "@/stores";
import {ApiDate, TypeDate} from "@/api/date_api";
import {Time} from "@/BAIYIN/time";
import {log} from "@/BAIYIN/log";

interface quickEntryType {
	font: Component
	text: string // 文字
	name?: string // 路由的名称
	link?: string // 可以跳外部链接
}

let quickEntryList: quickEntryType[] = [
	{font: IconSettings, text: "群聊管理", name: "chat_list"},
	{font: IconMessage, text: "评论列表", name: "comment_list"},
	{font: IconBook, text: "文章列表", name: "article_list"},
	{font: IconUserGroup, text: "用户列表", name: "user_list"},
	// font: defineComponent({
	// 	render: () => {
	// 		return h("i", {class: "icon ion-happy-outline"})
	// 	}
	// }),
]

function goLink(item: quickEntryType) {
	if (item.name) {
		router.push({name: item.name})
		return
	}
	if (item.link) {
		window.open(item.link, "_blank")
	}
}

interface updateLogType {
	id?: number
	title: string
	created_at: string
	items?: string[]
}

const updateLogList: updateLogType[] = [
	{title: " 博客", created_at: "2024-10-20", items: ["技术栈：gin vue3 TypeScript ArcoDesign",]},
]

// 数据统计监听主题切换
const store = useStore();
const isShow = ref(true)
watch(() => store.theme, () => {
	isShow.value = false
	setTimeout(() => isShow.value = true)
})

const data = reactive(TypeDate.List())

async function dataList() {
	isShow.value = false
	let result = await ApiDate.List()
	data.date_list = result.data.date_list
	data.login_data = result.data.login_data
	data.sign_data = result.data.sign_data
	isShow.value = true
}

dataList()
</script>

<style lang="scss">
.home_view {

	.main {
		display: flex;
		justify-content: space-between;
		margin-top: 20px;

		.left {
			width: 70%;

			.quick_entry {
				.body {
					display: flex;
					overflow: hidden;
				}

				.item {
					display: flex;
					flex-direction: column;
					justify-content: center;
					margin-right: 50px;
					align-items: center;

					&:last-child {
						margin-right: 0;
					}

					.icon {
						background-color: var(--color-fill-2);
						border-radius: 10px;
						width: 60px;
						height: 60px;
						display: flex;
						justify-content: center;
						align-items: center;
						font-size: 24px;
						transition: all .3s;
						cursor: pointer;

						svg {
							transition: all .3s;
						}

						&:hover {
							transform: scale(1.05);

							svg {
								transform: scale(1.20);
							}
						}
					}

					.text {
						margin-top: 5px;
						color: var(--color-text-2);
					}
				}
			}
		}

		.right {
			width: calc(30% - 20px);

			.update_log {
				.item {
					color: var(--color-text-2);
					margin-bottom: 15px;

					.line1 {
						display: flex;
						align-items: center;
						justify-content: space-between;

						> span {
							display: flex;
							align-items: center;
							white-space: nowrap;

							.index {
								margin-right: 10px;
							}
						}

						.content {
							display: flex;

							.arco-typography {
								margin-bottom: 0;
								display: inline-block;
							}
						}
					}

					.line2 {
						margin: 5px 0;
						line-height: 1.5rem;
					}

				}
			}
		}
	}
}


</style>