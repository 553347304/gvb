<template>
	<div>
		<drawer_big_model_chat v-model:visible="createVisible" :title="title" :id="id"/>
		<Gvb_table
			_title="创建会话"
			_search="name"
			:_url="ApiBigModel._url_session"
			:_columns="DefineBigModel.ColumnsSession"
			:_filter="[]"
			@_create="create"
			@_update="update"
		>
		</Gvb_table>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {reactive, ref} from "vue";
import {ApiBigModel, DefineBigModel, TypeBigModel} from "@/api/big_model_api";
import Drawer_big_model_chat from "@/components/admin/drawer/drawer_big_model_chat.vue";

const createVisible = ref(false)
const createForm = reactive(TypeBigModel.ChatList())
const id = ref(0)
const title = ref("")

function create() {
	Object.assign(createForm, TypeBigModel.ChatList())
	createVisible.value = true
}

async function update(record: TypeBigModel.SessionList) {
	title.value = `${record.nick_name}-${record.role_name} ${record.session_name}`
	id.value = record.id
	createVisible.value = true
}


</script>

<style scoped lang="scss">

</style>


