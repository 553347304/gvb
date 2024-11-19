<template>
	<div class="message_list_view">
		<div class="user_list_menu">
			<Gvb_message_list :_data="data2.list" @_check="m.TalkList"/>
		</div>

		<div class="user_record_menu">
			<div class="record_list">

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

			<div class="message_record">
				<a-textarea v-model="sendMessage.content"
							placeholder="请输入聊天内容"
							:auto-size="{minRows:4,maxRows:5}"
				></a-textarea>
				<div class="button">
					<a-button type="primary" @click="m.TalkList(receiveUser)">刷新</a-button>
					<a-button type="primary" @click="m.Send" @keydown.enter.ctrl="m.Send">发送</a-button>
				</div>

			</div>
		</div>
	</div>
</template>
<script setup lang="ts">
import {ApiMessage, TypeMessage} from "@/api/message_api";
import Gvb_message_list from "@/components/message/gvb_message_list.vue";
import {nextTick, reactive} from "vue";
import {Api, TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import {useStore} from "@/stores";
import {BaseLog} from "@/BAIYIN/log";
import {Message} from "@arco-design/web-vue";

const store = useStore()
const PageInfo = reactive(TypeApi.PageInfo())

const data2 = reactive(Base.List<TypeMessage.Info>())
const talkDate = reactive(Base.List<TypeMessage.Talk>())

const receiveUser = reactive(TypeMessage.Info())
const sendMessage = reactive(TypeMessage.Create())



class _Message {
	constructor() {
	}

	private user_id = store.userInfo.user_id

	private AutoRoll = () => {
		// 自动滚动条
		nextTick(() => {
			setTimeout(() => {
				let doc = document.querySelector(".record_list")
				if (doc === null) return
				doc.scrollTo({
					top: doc.scrollHeight,
					behavior: "smooth", // 平滑效果
				})
			}, 0)
		})
	}

	UserList = async () => {
		let result = await ApiMessage.UserList(this.user_id, PageInfo)
		data2.list = result.data.list
		data2.total = result.data.total
	}

	TalkList = async (record: TypeMessage.Info) => {
		Object.assign(receiveUser, record)
		let result = await ApiMessage.List({
			PageInfo: PageInfo,
			user_id: this.user_id,
			receive_id: receiveUser.user_id,
		})
		const list = []
		result.data.list.forEach((item) => {
			item._is_me = item.send_user_id === this.user_id;
			list.push(item)
		})
		talkDate.list = result.data.list
		talkDate.total = result.data.total
		await this.AutoRoll()
	}

	Send = async () => {
		console.log(sendMessage)
		if (sendMessage.content === "") {
			Message.warning("消息不能为空")
			return
		}
		sendMessage.rev_user_id = receiveUser.user_id
		let result = await Api.Create(Api.URL.Message, sendMessage)
		if (!BaseLog(result)) return
		await this.TalkList(receiveUser)
		sendMessage.content = ""
		await this.AutoRoll()
	}
}

const m = new _Message()
m.UserList()

</script>

<style scoped lang="scss">
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
	}

	.user_menu {
		width: 280px;
		margin-left: 20px;
		padding: 20px 0;
	}

	.user_record_menu {
		width: calc(100% - 310px);
		margin-left: 20px;
		height: calc(100vh - 130px);

		.record_list {
			padding: 20px;
			height: calc(100% - 200px);
			overflow-y: auto;

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
							background-color: #95EC69;

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

		.message_record {
			top: 40px;
			border-top: 1px solid var(--bg);
			padding: 10px 20px;


			position: relative;

			.button {
				position: absolute;
				right: 30px;
				bottom: 20px;
				z-index: 1; // 层级
				> button {
					margin-right: 10px;

					&:last-child {
						margin-right: 0;
					}
				}
			}
		}
	}
}
</style>