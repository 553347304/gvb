<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {Time} from "@/BAIYIN/time";
import {MdPreview} from "md-editor-v3";
import {BaseLog} from "@/BAIYIN/log";
import {useRoute} from "vue-router";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";
import {SSE} from "@/api/api_index/sse";
import {Base} from "@/BAIYIN/HTTP";
import {Unit} from "@/api/api_utils/unit";
import {EventSource, type MessageEvent} from "event-source-polyfill";

const store = useStore();
const route = useRoute()
type Props = { session_id: number; info: TypeBigModel.RoleDetail };
const props = defineProps<Props>()
const isLoading = ref(false)
const dataList = ref<TypeBigModel.ChatList[]>([])

async function getData() {
	isLoading.value = true
	let result = await ApiBigModel.ChatList({session_id: props.session_id,});
	dataList.value = result.data.list
	Unit.Scroll(".chat_list")
	isLoading.value = false
}

watch(() => props.session_id, () => {
	getData()
}, {immediate: true})

async function ChatAdminRemove(idList: number[]) {
	let result = await ApiBigModel.ChatAdminRemove(idList)
	await getData()
	return BaseLog(result)
}

const chat = reactive(TypeBigModel.ChatSse())

function userChat(e: any) {

	// shift enter 换行    enter 发送
	if (e.ctrlKey) {
		chat.content += "\n"
		return;
	}

	chat.session_id = Number(route.query.sessionId)
	if (chat.content.trim() === "") {
		Message.warning("内容不能为空")
		return
	}
	const chatItem = reactive<TypeBigModel.ChatList>({
		id: 0,
		created_at: "",
		user_content: chat.content,
		user_avatar: store.userInfo.avatar,
		bot_content: "",
		bot_avatar: props.info.icon,
		status: false,
	})
	dataList.value.push(chatItem)

	const eventSource = SSE.get("/api/big_model/chat_sse", chat)
	eventSource.onopen = function () {
		chat.content = ""
		chatItem.status = true
	}
	eventSource.onmessage = function (ev) {
		const result: Base.Response<string> = JSON.parse(ev.data)
		if (result.code) {
			Message.error(result.message)
			return
		}
		chatItem.bot_content += result.data
		Unit.Scroll(".chat_list")
	}
}

</script>

<template>
	<a-spin :loading="isLoading" tip="加载中" style="width: 100%;height: 100%;">
		<div class="chat_message_region_view">

			<div class="chat_list">

				<div class="item" v-for="item in dataList">
					<div class="right">
						<div class="avatar"><img :src="item.user_avatar" alt=""></div>
						<div class="container">
							<div class="date">{{ Time.Time(item.created_at).Now }}</div>
							<div class="content">{{ item.user_content }}</div>
						</div>
					</div>
					<div class="left">
						<div class="avatar"><img :src="item.bot_avatar" alt=""></div>
						<div class="container">
							<div class="date">{{ Time.Time(item.created_at).Now }}</div>
							<div class="content">
								<md-preview :model-value="item.bot_content" :code-foldable="false"/>
							</div>
							<IconDelete @click="ChatAdminRemove([item.id])"/>
						</div>
					</div>
				</div>
			</div>
			<div class="bottom">
				<div>
					<div class="input">
						<IconSend @click="userChat"/>
						<a-textarea placeholder="请输入您的问题" :auto-size="{minRows: 1, maxRows: 5}"
									v-model="chat.content"
									@keydown.enter.prevent="userChat"
						></a-textarea>
					</div>
				</div>
			</div>

		</div>
	</a-spin>

</template>

<style lang="scss">
@import "@/global/mixin.scss";

.chat_message_region_view {
	width: 100%;
	height: 100%;


	.chat_list {
		width: 100%;
		height: calc(100% - 70px);
		overflow-y: auto;
		padding: 20px;

		.item {
			margin-bottom: 20px;

			.left {
				margin-bottom: 10px;
				display: flex;


				.container {
					margin-left: 10px;

					.content {
						background-color: var(--bg);
					}
				}

			}

			.right {
				flex-direction: row-reverse; // 翻转
				display: flex;


				.container {
					text-align: right;

					.content {
						margin-left: auto;
						background-color: var(--chat_bg);
					}
				}
			}


			@include image_avatar(50px);

			.container {
				width: calc(100% - 100px);

				.date {
					font-size: 12px;
					color: var(--color-text-4);
				}

				.content {
					border-radius: 5px;
					padding: 10px;
					margin-top: 5px;
					width: fit-content; // 根据内容排列宽度

					@include md-editor();
				}
			}
		}

	}


	.bottom {
		padding: 20px;

		> div {
			display: flex;
			align-items: center;
			width: 100%;
			position: relative;

			.input {
				width: 100%;
				align-items: center;
				position: absolute;
				bottom: -30px;

				svg {
					font-size: 22px;
					position: absolute;
					right: 10px;
					top: 40%;
					transform: translateY(-50%) rotate(-45deg);
					color: var(--color-text-2);
					z-index: 2;
				}

				.arco-textarea-wrapper {
					border-radius: 5px;
					background-color: transparent;
					border: 1px solid #CCCCCC;
				}
			}
		}

	}


}
</style>

