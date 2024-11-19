<template>
	<div>
		<Form_user_create title="创建用户" v-model:visible="createVisible" :form="createForm"/>
		<Form_user_update title="编辑用户" v-model:visible="updateVisible" :form="updateForm"/>
		<Gvb_table
			_title="创建用户"
			_search="user_name"
			:_url="Api.URL.User"
			:_columns="DefineUser.Columns"
			:_filter="DefineUser.Filter"
			@_create="create"
			@_update="update"
		></Gvb_table>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import Form_user_create from "@/components/admin/form/form_user_create.vue";
import Form_user_update from "@/components/admin/form/form_user_update.vue";
import {DefineUser, TypeUser} from "@/api/user_api";
import {reactive, ref} from "vue";
import {Api} from "@/api/api";


const createVisible = ref(false)
const updateVisible = ref(false)
const createForm = reactive(TypeUser.Create())
const updateForm = reactive(TypeUser.Update())

function create() {
	Object.assign(createForm, TypeUser.Create())
	createVisible.value = true
}

function update(record: TypeUser.Info) {
	updateForm.role = record.role
	updateForm.nick_name = record.nick_name
	updateForm.user_id = record.id
	updateVisible.value = true
}


</script>

<style scoped lang="scss">

</style>

