<template>
	<div class="chat_list_view">
		<div class="chat_list">

			<div class="head">
				<div class="title">白音在线聊天室</div>
				<div class="outline">在线人数: {{chatData.online}}</div>
			</div>

			<div class="record_list"><Gvb_chat_list :data="data"/></div>

			<div class="menu">
				<div class="icon" @click="ws.SendImage"><IconImage/></div>
				<div class="icon" @click="ws.SendFile"><IconFile/></div>
			</div>

			<div class="footer">

				<div class="inRoom" v-if="!socket">
					<a-button type="primary" @click="ws.Socket()">进入聊天室</a-button>
				</div>
				<MdEditor v-model="sendData.message"
						  :toolbars="[]"
						  :footers="[]"
						  :preview="false"
						  :on-upload-img="ApiImage.OnUploadImage"
						  :max-length="100"
						  show-word-limit
						  style="height: 100%"
						  @keydown.enter.prevent="ws.Send"
				/>
				<a-button type="primary" class="send_button"
						  @click="ws.Send">发送</a-button>
			</div>
		</div>
	</div>
</template>
<script setup lang="ts">
import Gvb_chat_list from "@/components/admin/gvb_chat_list.vue";
import {Message} from "@arco-design/web-vue";
import {nextTick, onUnmounted, reactive, ref} from "vue";
import {TypeChat} from "@/api/chat_api";
import {Base} from "@/BAIYIN/HTTP";
import "md-editor-v3/lib/style.css"
import {MdEditor} from "md-editor-v3";
import {ApiImage} from "@/api/image_api";

const data = reactive(Base.List<TypeChat.List>())
let chatData = reactive(TypeChat.Message())
let sendData = reactive(TypeChat.Send())
let index = 0
let socket = ref<WebSocket | null>()

class _ws {
	Socket = () => {
		socket.value = new WebSocket(`ws://${location.host}/ws/api/chat_group`)
		socket.value.onmessage = function (event){
			let _data = event.data
			let jsonData = JSON.parse(_data) as TypeChat.Message
			console.log(jsonData)
			// 第一条消息
			chatData.online = jsonData.online
			if (index === 0) {
				chatData.nick_name = jsonData.nick_name
				index++
				return
			}
			data.list.push({
				id: 0,
				addr: "",
				is_group:true,
				nick_name: jsonData.nick_name,
				avatar: jsonData.avatar,
				message: jsonData.message,
				type: jsonData.type,
				_is_me: chatData.nick_name === jsonData.nick_name,
				created_at: jsonData.created_at,
			})

			// 自动滚动条
			nextTick(() => {
				let doc = document.querySelector(".record_list")
				if (doc === null) return
				doc.scrollTo({
					top: doc.scrollHeight,
					behavior: "smooth", // 平滑效果
				})
			})
		}
		socket.value.onopen = function (e: any){
			Message.success("连接成功")
		}
		socket.value.onerror = function (e: any){
			Message.success(e)
			socket.value = null
		}
	}

	Send = (e: MouseEvent) => {
		// shift enter 换行    enter 发送
		if (e.ctrlKey) {
			sendData.message += "\n"
			return;
		}

		if (sendData.message === "") {
			Message.warning("消息不能为空")
			return
		}
		sendData.type = 1
		socket.value?.send(JSON.stringify(sendData))
		sendData.message = ""
	}

	SendImage = () => {
		Message.success("开发中")
	}
	SendFile = () => {
		Message.success("开发中")
	}
}
const ws = new _ws()
onUnmounted(()=>{
	socket.value?.close()
})
</script>

<style scoped lang="scss">
.chat_list_view {
	display: flex;
	height: calc(100vh - 130px);
	justify-content: space-between;


	.chat_list {
		width: 100%;
		height: 100%;
		background-color: var(--color-bg-1);
		border-radius: 5px;

		.head {
			border-bottom: 1px solid var(--bg);
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			height: 70px;

			.title {
				font-size: 16px;
				font-weight: 600;
			}

			.outline {
				margin-top: 5px;
				color: var(--color-text-2);
			}
		}

		.record_list {
			overflow-y: auto;
			padding: 20px;
			height: calc(100vh - 410px);
		}

		.menu{
			display: flex;
			border: 1px solid var(--bg);

			.icon{
				font-size: 18px;
				width: 30px;
				height: 30px;
				display: flex;
				align-items: center;
				cursor: pointer;
				margin-left: 10px;
				&:last-child{
					margin-left: 0;
				}
			}
		}

		.footer {
			height: 180px;
			border-top: 1px solid var(--bg);
			padding: 20px;
			position: relative;

			.inRoom{
				position: absolute;
				width: calc(100% - 40px);
				height: calc(100% - 40px);
				display: flex;
				justify-content: center;
				align-items: center;
				z-index: 2;
				background-color: var(--login_bg);
			}

			.send_button {
				position: absolute;
				left: 30px;
				bottom: 30px;
				z-index: 1;
			}
		}
	}
}
</style>