<script setup lang="ts">
import {ref, watch} from "vue";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {Time} from "@/BAIYIN/time";
import {MdPreview} from "md-editor-v3";
import {BaseLog} from "@/BAIYIN/log";

type Props = { id: number; title: string };
const props = defineProps<Props>()
const data = ref<TypeBigModel.ChatList[]>([])
const isLoading = ref(false)

async function getData() {
	let result = await ApiBigModel.ChatList({session_id: props.id,});
	data.value = result.data.list
}

watch(() => props.id, () => {
	isLoading.value = true
	getData()
	isLoading.value = false
}, {deep: true});


async function ChatAdminRemove(idList: number[]) {
	isLoading.value = true
	let result = await ApiBigModel.ChatAdminRemove(idList)
	await getData()
	isLoading.value = false
	return BaseLog(result)
}

</script>

<template>

	<a-drawer :placement="'right'" width="50%" :footer="false" :title="props.title">
		<a-spin :loading="isLoading" tip="加载中" style="width: 100%;height: 100%;">
			<div class="big_model_chat_list_view">
				<div class="item" v-for="item in data" @click="ChatAdminRemove([item.id])">
					<div class="right">
						<div class="avatar">
							<img :src="item.user_avatar" alt="">
						</div>
						<div class="container">
							<div class="date">{{ Time.Time(item.created_at).Now }}</div>
							<div class="content">{{ item.user_content }}</div>
						</div>
					</div>
					<div class="left">
						<div class="avatar">
							<img :src="item.bot_avatar" alt="">
						</div>
						<div class="container">
							<div class="date">{{ Time.Time(item.created_at).Now }}</div>
							<div class="content">
								<md-preview :model-value="item.bot_content"/>
							</div>
							<IconDelete @click="ChatAdminRemove([item.id])"/>
						</div>
					</div>
				</div>
			</div>
		</a-spin>

	</a-drawer>

</template>

<style scoped lang="scss">
@import "@/global/mixin.scss";

.big_model_chat_list_view {
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

			.data {
				font-size: 12px;
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
</style>

