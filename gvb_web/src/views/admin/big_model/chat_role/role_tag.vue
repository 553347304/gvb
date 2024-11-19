<template>
	<div>
		<form_big_model_tag_create :title="title" v-model:visible="createVisible" :form="createForm"/>
		<div class="menu_list_view">
			<Gvb_table
				_title="创建标签"
				_search="title"
				:_url="ApiBigModel._url_tag"
				:_columns="DefineBigModel.ColumnsTag"
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
import Form_big_model_tag_create from "@/components/admin/form/form_big_model_tag_create.vue";

const createVisible = ref(false)
const createForm = reactive(TypeBigModel.TagCreate())
const title = ref("")

function create() {
	title.value = "创建标签"
	Object.assign(createForm, TypeBigModel.TagCreate())
	createVisible.value = true
}

function update(record: TypeBigModel.TagCreate) {
	console.log(record)
	title.value = "更新标签"
	Object.assign(createForm, record)
	createVisible.value = true
}


</script>

<style scoped lang="scss">
.menu_list_view {
	.menu_column_image {
		display: grid;
		grid-template-columns: repeat(3, 1fr); // 图片一行显示三个
		grid-row-gap: 5px; // 行间距
		.arco-image-img {
			margin-right: 5px;
			border-radius: 5px;
		}
	}


}
</style>


