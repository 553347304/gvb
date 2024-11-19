<script setup lang="ts">
import {nextTick, reactive, ref} from "vue";
import {Api, TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import {TypeChat} from "@/api/chat_api";
import {MdPreview} from "md-editor-v3";

const PageInfo = reactive(TypeApi.PageInfo())
type Props = { data: Base.List<TypeChat.List> }
const props = defineProps<Props>()
const show = ref(false)

class _Chat {
	constructor() {
	}


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

	List = async () => {
		show.value = false
		let result = await Api.List(Api.URL.Chat, PageInfo)
		props.data.list = result.data.list
		props.data.total = result.data.total
		await this.AutoRoll()
		show.value = true
	}
}

const c = new _Chat()
c.List()

</script>

<template>
	<template v-if="show" v-for="item in props.data.list" :key="item.id">

		<!--		<div class="date">{{ Time.Time(item.created_at).Now }}</div>-->

		<div v-if="item.type === 0 || item.type === -1"
			 :class="'system' + ' system_' + item.type">
			<div class="txt-message">{{ item.message }}</div>
		</div>


		<div v-else-if="item.type === 1" :class="{message: true, isMe: item._is_me}">
			<img class="avatar" :src="item.avatar" alt="">
			<div class="message-main">
				<div class="message-user">{{ item.nick_name }}</div>
				<div class="message-context">
					<div class="context">
						<div class="txt-message">

							<MdPreview v-model="item.message" :editor-id="'md__'+item.id"></MdPreview>


							<!--							{{ item.message }}-->

						</div>
					</div>
				</div>
			</div>
		</div>
	</template>
</template>

<style lang="scss">
.system {
	display: flex;
	justify-content: center;
	margin-bottom: 20px;
	border-radius: 5px;

	&:hover {
		background-color: var(--color-fill-1);
		border-radius: 5px;
	}

	.txt-message {
		padding: 5px 10px;
		border-radius: 5px;
		background-color: var(--color-fill-2);
		color: var(--color-text-2);
	}
}

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

			.txt-message {
				background-color: var(--color-fill-2);
				border-radius: 5px;
				padding: 10px;
				position: relative;
				width: fit-content;
				min-height: 41px;
				//white-space: break-spaces;
				//word-break: break-all;

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

				.md-editor {
					background-color: inherit;

					.md-editor-preview-wrapper {
						padding: 0;

						.md-editor-preview {
							img {
								padding: 0;
								border: none;
							}

							p {
								margin: 0;
							}
						}
					}
				}
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
					border-color: transparent transparent transparent #95EC69;
				}
			}
		}
	}
}
</style>

