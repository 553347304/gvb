<template>
	<a-modal :title="props.title" :on-before-ok="ok">
		<a-form validate ref="formRef" :model="form">
			<Unit_form label="标签名称" field="title"><a-input v-model="form.title"/></Unit_form>
			<Unit_form label="标签颜色" field="color">
				<a-input v-model="form.color"/>
				<input type="color" v-model="form.color" style="margin-left: 10px">
			</Unit_form>

		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {BaseLog} from "@/BAIYIN/log";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";

type Props = { form: TypeBigModel.TagCreate; title: string };
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), {deep: true});

async function ok(): Promise<boolean> {
	if (await formRef.value.validate()) return false;
	let result = props.title === "创建标签" ?
		await ApiBigModel.TagCreate(form) :
		await ApiBigModel.TagUpdate(form);
	Object.assign(form, TypeBigModel.TagCreate())	// 重新赋值清空
	return BaseLog(result)
}
</script>

<style scoped lang="scss">

</style>
