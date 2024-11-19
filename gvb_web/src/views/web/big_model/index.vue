<script setup lang="ts">
import Web_nav from "@/components/web/web_nav.vue";
import {reactive, ref} from "vue";
import {TypeArticle} from "@/api/article_api";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {useStore} from "@/stores";
import {Message} from "@arco-design/web-vue";
import {useRoute} from "vue-router";

const data = reactive(TypeArticle.List())
const historyList = reactive<TypeBigModel.RoleSquare[]>([])
const store = useStore();
const route = useRoute()

async function roleHistory() {
	if (!store.isLogin) {
		Message.warning("未登录");
		return
	}
	let result = await ApiBigModel.RoleHistory();
	Object.assign(historyList, result.data.list)
}

roleHistory()

</script>

<template>
	<div class="big_model_view">
		<web_nav no-scroll/>
		<main>
			<div class="menu">
				<div class="head">
					<router-link :to="{name: 'role_square'}">角色广场</router-link>
				</div>
				<div class="role_list">
					<div :class="{item: true, active: Number(route.query.roleId) === item.id}"
						 v-for="item in historyList"
						 @click="ApiBigModel.CheckRole(item)">
						<img :src="item.icon" alt="" class="avatar">
						<div class="content">
							<a-typography-text class="title" :ellipsis="{rows: 1, css: true}">
								{{ item.name }}
							</a-typography-text>
							<a-typography-text class="abs" :ellipsis="{rows: 2, css: true}">{{ item.abstract }}
							</a-typography-text>
						</div>
					</div>
				</div>
			</div>
			<router-view class="view"/>
		</main>
	</div>
</template>


<style scoped lang="scss">
@import "@/global/mixin.scss";

.big_model_view {
	main {
		margin-top: 60px;
		height: calc(100vh - 60px);
		padding: 20px 40px;
		display: flex;

		.menu {
			width: 240px;
			background-color: var(--color-bg-1);
			height: 100%;
			border-radius: 5px;

			.head {
				padding: 20px 20px 10px 20px;
				border-bottom: 1px solid var(--bg);
				text-align: center;
			}

			.role_list {
				overflow-y: auto;
				height: calc(100% - 50px);

				.item {
					display: flex;
					align-items: center;
					cursor: pointer;
					padding: 20px;

					&.active {
						background-color: var(--color-fill-2);
					}

					&:hover {
						background-color: var(--color-fill-1);
					}

					&:last-child {
						margin-bottom: 0;
					}

					.content {
						margin-left: 10px;

						> .title {
							margin-bottom: 0;
						}

						> .abs {
							font-size: 12px;
							margin-bottom: 0;
							color: var(--color-text-2);
						}
					}

					@include image_avatar(40px, 14px)
				}
			}
		}

		a {
			text-decoration: none;
			color: var(--active);
		}

		.view {
			margin-left: 20px;
			width: 100%;
		}
	}
}
</style>