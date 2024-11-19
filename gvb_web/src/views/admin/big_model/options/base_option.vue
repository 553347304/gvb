<script setup lang="ts">
import Gvb_title from "@/components/system/gvb_title.vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {reactive, ref} from "vue";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {MdPreview} from "md-editor-v3";
import {BaseLog} from "@/BAIYIN/log";
import {useStore} from "@/stores";

const formRef = ref()
const form = reactive(TypeBigModel.options())
const store = useStore();
Object.assign(form, store.bigModelInfo)

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
				<unit_form label="大模型选择"><a-select v-model="form.setting.name" :options="form.model_list"/></unit_form>
				<unit_form label="api_key"><a-input v-model="form.setting.api_key"/></unit_form>
				<unit_form label="api_secret"><a-input v-model="form.setting.api_secret"/></unit_form>
				<unit_form label="大模型名称"><a-input v-model="form.setting.title"/></unit_form>
				<unit_form label="大模型提示词"><a-textarea v-model="form.setting.prompt" :auto-size="{minRows: 4, maxRows: 5}"/></unit_form>
				<unit_form label="标语"><a-textarea v-model="form.setting.slogan" :auto-size="{minRows: 3, maxRows: 5}"/></unit_form>
				<unit_form label="菜单序号"><a-input-number v-model="form.setting.order"/></unit_form>
				<unit_form label="是否启用"><a-switch v-model="form.setting.enable"/></unit_form>
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