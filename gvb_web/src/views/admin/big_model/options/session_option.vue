<script setup lang="ts">
import Gvb_title from "@/components/system/gvb_title.vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {reactive, ref} from "vue";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {MdPreview} from "md-editor-v3";
import {BaseLog} from "@/BAIYIN/log";

const formRef = ref()
const form = reactive(TypeBigModel.options())
async function getData() {
	let result = await ApiBigModel.SettingList();
	Object.assign(form, result.data)
}
getData()

async function update() {
	if (await formRef.value.validate()) return false
	let result = await ApiBigModel.SettingUpdate(form)
	return BaseLog(result)
}
</script>

<template>
	<div class="big_model_base_option">
		<div class="left">
			<a-alert>使用大模型必须得配置项</a-alert>
			<a-form ref="formRef" :model="form" :label-col-props="{span: 6}" :wrapper-col-props="{span: 18}">
				<unit_form label="对话积分消耗"><a-input-number v-model="form.session.chat_scope"/></unit_form>
				<unit_form label="会话积分消耗"><a-input-number v-model="form.session.session_score"/></unit_form>
				<unit_form label="每日积分领取"><a-input-number v-model="form.session.day_score"/></unit_form>
				<unit_form><a-button type="primary" @click="update">更新</a-button></unit_form>
			</a-form>
		</div>
		<div class="right">
			<gvb_title title="帮助"/>
			<md-preview :model-value="form.setting.help"/>
		</div>
	</div>
</template>

<style scoped lang="scss">
.big_model_base_option {
	display: flex;
	padding: 20px;

	.left {
		width: 50%;

		.arco-form {
			margin-top: 20px;
		}
	}

	.right {
		width: 40%;
		margin-left: 20px;
	}
}
</style>