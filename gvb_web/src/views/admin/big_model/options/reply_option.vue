<template>
	<div>
		<form_big_model_create :title="title" v-model:visible="createVisible" :form="createForm"/>
		<div class="menu_list_view">
			<Gvb_table
				_title="创建自动回复"
				_search="name"
				:_url="ApiBigModel._url_auto_reply"
				:_columns="DefineBigModel.Columns"
				:_filter="[]"
				@_create="create"
				@_update="update"
			>
			</Gvb_table>
		</div>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {reactive, ref} from "vue";
import {ApiBigModel, DefineBigModel, TypeBigModel} from "@/api/big_model_api";
import Form_big_model_create from "@/components/admin/form/form_big_model_create.vue";

const createVisible = ref(false)
const createForm = reactive(TypeBigModel.AutoReply())
const title = ref("")

function create() {
	title.value = "创建自动回复"
	Object.assign(createForm, TypeBigModel.AutoReply())
	createVisible.value = true
}

function update(record: TypeBigModel.AutoReply) {
	console.log(record)
	title.value = "编辑自动回复"
	Object.assign(createForm, record)
	createVisible.value = true
}


</script>

<style scoped lang="scss">
</style>

