<template>
	<a-modal :title="props.title" :on-before-ok="ok">
		<a-form validate ref="formRef" :model="form">
			<Unit_form label="规则名称" field="name"><a-input v-model="form.name"/></Unit_form>
			<Unit_form label="匹配模式" field="mode"><a-select v-model="form.mode" :options="DefineBigModel.modeOption"/></Unit_form>
			<Unit_form label="规则" field="rule"><a-textarea v-model="form.rule" :auto-size="{minRows:3,maxRows:5}"/></Unit_form>
			<Unit_form label="回复内容" field="rule_content"><a-textarea v-model="form.rule_content" :auto-size="{minRows:4,maxRows:5}"/></Unit_form>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {TypeUser} from "@/api/user_api";
import {reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {BaseLog} from "@/BAIYIN/log";
import {ApiBigModel, DefineBigModel, TypeBigModel} from "@/api/big_model_api";

type Props = { form: TypeBigModel.AutoReply; title: string };
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), {deep: true});

async function ok(): Promise<boolean> {
	if (await formRef.value.validate()) return false;
	let result = props.title === "创建自动回复" ?
		await ApiBigModel.AutoReplyCreate(form) :
		await ApiBigModel.AutoReplyUpdate(form);
	Object.assign(form, TypeBigModel.AutoReply())	// 重新赋值清空
	return BaseLog(result)
}
</script>

<style scoped lang="scss">

</style>
