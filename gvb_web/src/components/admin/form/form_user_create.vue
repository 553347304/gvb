<template>
	<a-modal :title="props.title" :on-before-ok="ok">
		<a-form validate ref="formRef" :model="form">	<!--:label-col-props="{span: 0}" :wrapper-col-props="{span: 24}"-->
			<Unit_form label="昵称" field="nick_name" ><a-input v-model="form.nick_name"/></Unit_form>
			<Unit_form label="用户名" field="user_name" ><a-input v-model="form.user_name"/></Unit_form>
			<Unit_form label="密码" field="password" ><a-input v-model="form.password"/></Unit_form>
			<Unit_form label="确认密码" field="pass_word" ><a-input v-model="form.pass_word"/></Unit_form>
			<Unit_form label="邮箱" field="email" ><a-input v-model="form.email"/></Unit_form>
			<Unit_form label="手机号" field="tel" ><a-input v-model="form.tel"/></Unit_form>
			<Unit_form label="权限" field="role" ><a-select v-model="form.role" :options="DefineUtils.Role"></a-select></Unit_form>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {Message} from "@arco-design/web-vue";
import {ApiUser, TypeUser} from "@/api/user_api";
import {reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {DefineUtils} from "@/api/api_utils/type";
import {BaseLog} from "@/BAIYIN/log";

type Props = {form: TypeUser.Create; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), { deep: true });

async function ok(): Promise<boolean> {
	if (await formRef.value.validate()) return false
	if (form.password !== form.pass_word) {
		Message.warning("两次密码不一致")
		return false
	}
	let result = await ApiUser.Create(form)
	Object.assign(form, TypeUser.Create())	// 重新赋值清空
	return BaseLog(result)
}
</script>

<style scoped lang="scss">

</style>
