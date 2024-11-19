<template>
	<div class="web_feed_back">
		<a-form ref="formRef" :model="form" :label-col-props="{span: 0}" :wrapper-col-props="{span:24}">
			<unit_form label="邮箱" field="email">
				<a-input v-model="form.email" placeholder="邮箱"/>
			</unit_form>
			<unit_form label="反馈内容" field="content">
				<a-textarea v-model="form.content" placeholder="反馈内容"
							:auto-size="{minRows: 4, maxRows: 4}"
							:max-length="160" show-word-limit/>
			</unit_form>
			<unit_form>
				<a-button type="primary" @click="ok" style="width: 100%">发布</a-button>
			</unit_form>
		</a-form>
		<div class="fade_back_list">
			<div class="title">反馈列表</div>
			<div class="list">
				<div class="item" v-for="item in List">
					<div class="content">{{ item.content }}</div>
					<div class="date">{{ item.created_at }}</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {ApiSetting} from "@/api/setting_api";
import {Message} from "@arco-design/web-vue";

const formRef = ref()
const form = reactive({email: "553347304@qq.com", content: "",})
const List = ref([{content: "反馈内容xxxxxx", created_at: "2024-10-22",}])

async function ok() {
	if (await formRef.value.validate()) return false
	await ApiSetting.Send({receive: form.email, body: form.content})
	Message.success("发布成功")
}

</script>

<style scoped lang="scss">
.web_feed_back {
	.fade_back_list {
		color: var(--color-text-2);

		.title {
			font-size: 12px;
			display: flex;
			align-items: center;
			justify-content: space-between;

			&::before {
				display: inline-flex;
				width: 110px;
				height: 1px;
				content: "";
				background-color: var(--color-text-3);
			}

			&::after {
				display: inline-flex;
				width: 110px;
				height: 1px;
				content: "";
				background-color: var(--color-text-3);
			}
		}

		.list {
			margin-top: 20px;

			.item {
				margin-bottom: 20px;
				padding: 20px;
				border-radius: 5px;
				background-color: var(--color-fill-2);
				display: flex;
				justify-content: space-between;
				font-size: 12px;
			}

			&:last-child {
				margin-right: 0;
			}

			.date {
				white-space: nowrap;
				display: flex;
				align-items: center;
			}
		}
	}
}
</style>