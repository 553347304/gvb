<template>
	<div>
		<form_big_model_role_create :title="title" v-model:visible="createVisible" :form="createForm"/>
		<div class="role_list_view">
			<Gvb_table
				_title="创建角色"
				_search="name"
				:_url="ApiBigModel._url_role"
				:_columns="DefineBigModel.ColumnsRole"
				:_filter="[]"
				@_create="create"
				@_update="update"
			>
				<template #enable="{record}: {record: TypeBigModel.RoleList}">
					<a-switch :model-value="record.enable"/>
				</template>
				<template #auto_reply="{record}: {record: TypeBigModel.RoleList}">
					<a-switch :model-value="record.auto_reply"/>
				</template>
				<template #icon="{record}: {record: TypeBigModel.RoleList}">
					<a-image :src="record.icon" height="30px"/>
				</template>
				<template #tags="{record}: {record: TypeBigModel.RoleList}">
					<div class="col_tag">
						<a-tag v-for="item in record.tags" :color="item.color">
							{{ item.title }}
						</a-tag>
					</div>
				</template>
				<template #abstract="{record}: {record: TypeBigModel.RoleList}">
					<a-typography-text :ellipsis="{rows: 1, css: true}">{{ record.abstract }}</a-typography-text>
				</template>
				<template #prologue="{record}: {record: TypeBigModel.RoleList}">
					<a-typography-text :ellipsis="{rows: 1, css: true}">{{ record.prologue }}</a-typography-text>
				</template>
				<template #prompt="{record}: {record: TypeBigModel.RoleList}">
					<a-typography-text :ellipsis="{rows: 1, css: true}">{{ record.prompt }}</a-typography-text>
				</template>
			</Gvb_table>
		</div>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {reactive, ref} from "vue";
import {ApiBigModel, DefineBigModel, TypeBigModel} from "@/api/big_model_api";
import Form_big_model_role_create from "@/components/admin/form/form_big_model_role_create.vue";
const createVisible = ref(false)
const createForm = reactive(TypeBigModel.RoleCreate())
const title = ref("")

function create() {
	title.value = "创建角色"
	Object.assign(createForm, TypeBigModel.RoleCreate())
	createVisible.value = true
}

function update(record: TypeBigModel.RoleList) {
	title.value = "更新角色"
	let tagList = []
	for (const tag of record.tags) {
		tagList.push(tag.id)
	}
	Object.assign(createForm, record)
	createForm.tag_list = tagList
	createVisible.value = true
}


</script>

<style scoped lang="scss">
.role_list_view {
	.arco-typography {
		margin-bottom: 0;
	}

	.col_tag {
		display: flex;

		.arco-tag {
			margin-left: 10px;

			&:first-child {
				margin-left: 0;
			}
		}
	}

}
</style>


