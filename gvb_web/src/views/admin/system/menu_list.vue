<template>
	<div>
		<Form_menu_create :title="title" v-model:visible="updateVisible" :form="updateForm"></Form_menu_create>
		<div class="menu_list_view">
			<Gvb_table
				_title="创建菜单"
				_search="title"
				:_url="Api.URL.Menu"
				:_columns="DefineMenu.Columns"
				:_filter="[]"
				@_create="create"
				@_update="update"
			>

				<template #banners="{record}:{record:TypeMenu.info}">
					<div class="menu_column_image">
						<a-image v-for="item in record.banners" :key="item.id"
								 :src="item.path"
								 height="40px"
						></a-image>
					</div>
				</template>
			</Gvb_table>
		</div>
	</div>
</template>

<script setup lang="ts">

import Gvb_table from "@/components/admin/gvb_table.vue";
import Form_menu_create from "@/components/admin/form/form_menu_create.vue";
import {reactive, ref} from "vue";
import {DefineMenu, TypeMenu} from "@/api/menu_api";
import {Api} from "@/api/api";

const updateVisible = ref(false)
const updateForm = reactive(TypeMenu.Create())
const title = ref("")

function create() {
	title.value = "创建菜单"
	Object.assign(updateForm, TypeMenu.Create())
	updateVisible.value = true
}

function update(record: TypeMenu.info) {
	console.log(record)
	title.value = "编辑菜单"
	Object.assign(updateForm, record)
	updateForm.banner = []
	record.banners.forEach((item) => {
		updateForm.banner.push(item.id)
	})
	updateVisible.value = true
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

