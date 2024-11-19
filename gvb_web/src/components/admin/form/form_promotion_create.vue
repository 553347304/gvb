<template>
	<a-modal :title="props.title" :on-before-ok="ok">
		<a-form validate ref="formRef" :model="form">
			<Unit_form label="广告名称" field="title"><a-input v-model="form.title"/></Unit_form>
			<Unit_form label="跳转链接" field="href"><a-input v-model="form.href"/></Unit_form>
			<Unit_form label="图片链接" field="images"><Gvb_upload_image v-model="form.images"/></Unit_form>
			<Unit_form label="是否显示" field="is_show"><a-switch v-model="form.is_show"/></Unit_form>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {TypePromotion} from "@/api/promotion_api";
import {Api} from "@/api/api";
import Gvb_upload_image from "@/components/image/gvb_upload_image.vue";
import {BaseLog} from "@/BAIYIN/log";

type Props = {form: TypePromotion.Create; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), { deep: true });

async function ok() {
	if (await formRef.value.validate()) return false
	let result = props.title === "创建广告" ?
		await Api.Create(Api.URL.Promotion, form) :
		await Api.Update(Api.URL.Promotion, form, form.id);
	Object.assign(form, TypePromotion.Create())	// 重新赋值清空
	return BaseLog(result)
}
</script>

<style lang="scss">
</style>