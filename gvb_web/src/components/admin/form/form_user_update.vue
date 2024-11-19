<template>
	<a-modal :title="props.title" :on-before-ok="ok">
		<a-form ref="formRef" :model="form">
			<Unit_form label="昵称" field="nick_name">
				<a-input v-model="form.nick_name"/>
			</Unit_form>
			<Unit_form label="权限" field="role">
				<a-select v-model="form.role" :options="DefineUtils.Role"></a-select>
			</Unit_form>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {ApiUser, TypeUser} from "@/api/user_api";
import {reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {DefineUtils} from "@/api/api_utils/type";
import {BaseLog} from "@/BAIYIN/log";


type Props = {form: TypeUser.Update; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), {deep: true});

async function ok(): Promise<boolean> {
	if (await formRef.value.validate()) return false
	let result = await ApiUser.Update(form)
	return BaseLog(result)
}
</script>

<style scoped lang="scss">

</style>
