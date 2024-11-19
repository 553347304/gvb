<template>
	<div>
		<Form_image_update title="上传图片" v-model:visible="createVisible" :form="createForm"></Form_image_update>
		<Gvb_table
			_title="上传图片"
			_search="name"
			:_url="Api.URL.Image"
			:_columns="DefineImage.Columns"
			:_filter="[]"
			@_create="create"
		>
			<template #path="{record}: {record: TypeImage.Info}">
				<a-image :src="record.path" height="50px" style="border-radius: 5px"></a-image>
			</template>

			<!--下载文件-->
			<template #action_left="{record}: {record: TypeImage.Info}">
				<a :href="record.path" :download="record.name" style="margin-right: 10px">
					<a-button type="primary" status="success">
						<template #icon><IconDownload/></template>
						下载</a-button>
				</a>
			</template>
		</Gvb_table>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import {DefineImage, TypeImage} from "@/api/image_api";
import {reactive, ref} from "vue";
import {IconDownload} from "@arco-design/web-vue/es/icon";
import Form_image_update from "@/components/admin/form/form_image_update.vue";
import {Api} from "@/api/api";

const createVisible = ref(false)
const createForm = reactive(TypeImage.Create())

function create() {
	Object.assign(createForm, TypeImage.Create())
	createVisible.value = true
}
</script>

<style scoped lang="scss">

</style>
