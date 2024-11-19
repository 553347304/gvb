<template>
	<div class="message_list_view">
		<div class="user_list_menu">
			<div class="head">
				<a-input-search placeholder="搜索用户名"
								v-model="PageInfo.key"
								@search="getDate"
								@keydown.enter="getDate"
				></a-input-search>
			</div>

			<Gvb_message_list :_data="data1.list" @_check="check1"/>

			<div class="page">
				<a-pagination simple :total="data1.total"
							  @change="getDate"
							  v-model:current="PageInfo.page"
							  :page-size="PageInfo.limit"
				></a-pagination>
			</div>
		</div>

		<div class="user_menu">
			<Gvb_message_list :_data="data2.list" @_check="TalkList"/>
		</div>

		<div class="user_record_menu" v-if="talkDate.list && talkDate.list.length">
			<div v-for="item in talkDate.list" :key="item.id" :class="{message: true, isMe: item._is_me}">

				<img class="avatar" :src="item.send_user_avatar" alt="">

				<div class="message-main">
					<div class="message-user">{{ item.send_user_nick_name }}</div>
					<div class="message-context">
						<div class="context">
							<div class="txt-message">{{ item.content }}</div>
						</div>
					</div>
				</div>
			</div>
		</div>

	</div>
</template>
<script setup lang="ts">
import {ApiMessage, type TypeMessage} from "@/api/message_api";
import Gvb_message_list from "@/components/message/gvb_message_list.vue";
import {reactive, ref} from "vue";
import {TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";


const PageInfo = reactive(TypeApi.PageInfo())
const user1 = ref(0)
const user2 = ref(0)
const data1 = reactive(Base.List<TypeMessage.Info>())
const data2 = reactive(Base.List<TypeMessage.Info>())
const talkDate = reactive(Base.List<TypeMessage.Talk>())

async function getDate() {
	PageInfo.search = "send_user_nick_name"
	let result = await ApiMessage.UserList(0, PageInfo)
	data1.list = result.data.list
	data1.total = result.data.total
}

getDate()

async function check1(data: TypeMessage.Info) {
	let result = await ApiMessage.UserList(data.user_id, PageInfo)
	data2.list = result.data.list
	data2.total = result.data.total
	user1.value = data.user_id
}

async function TalkList(data: TypeMessage.Info) {
	user2.value = data.user_id
	let result = await ApiMessage.List({
		PageInfo: PageInfo,
		user_id: user1.value,
		receive_id: user2.value,
	})
	const list = []
	result.data.list.forEach((item) => {
		item._is_me = item.send_user_id === user1.value;
		list.push(item)
	})
	talkDate.list = result.data.list
	talkDate.total = result.data.total
}

</script>
<style lang="scss">
.message_list_view {
	display: flex;

	> div {
		background-color: var(--color-bg-1);
		height: calc(100vh - 130px);
		border-radius: 5px;
		color: var(--color-text-2);
		// 滚动条
		overflow-x: hidden;
		overflow-y: auto;
	}

	.user_list_menu {
		width: 280px;
		padding: 20px;

		.head {
			margin-bottom: 20px;
		}

		.page {
			display: flex;
			justify-content: center;
			margin-top: 10px;
		}
	}

	.user_menu {
		width: 280px;
		margin-left: 20px;
		padding: 20px 0;
	}

	.user_record_menu {
		width: calc(100% - 560px);
		margin-left: 20px;
		padding: 20px;

		.message {
			display: flex;
			margin-bottom: 20px;

			.avatar {
				width: 40px;
				height: 40px;
				border-radius: 50%;

			}

			.message-main {
				margin-left: 10px;

				.message-context {
					margin-top: 5px;

					.content {
						display: flex;
						margin-left: 5px;
					}
				}

				.txt-message {
					background-color: var(--color-fill-2);
					border-radius: 5px;
					padding: 10px;
					position: relative;
					width: fit-content;
					min-height: 41px;
					white-space: break-spaces;
					word-break: break-all;

					&:before {
						content: "";
						display: block;
						position: absolute;
						left: -20px;
						top: 6px;
						border-width: 5px 10px;
						border-style: solid;
						border-color: transparent var(--color-fill-2) transparent transparent;
					}
				}
			}

			&.isMe {
				justify-content: right; // 靠右
				flex-direction: row-reverse; // 反向

				.message-user {
					text-align: right;
				}

				.message-main {
					margin-left: 0;
					margin-right: 10px;

					.txt-message {
						&:before {
							content: "";
							display: block;
							position: absolute;
							right: -20px;
							left: inherit;
							top: 6px;
							border-width: 5px 10px;
							border-style: solid;
							border-color: transparent transparent transparent var(--color-fill-2);
						}
					}
				}
			}

		}
	}
}
</style>