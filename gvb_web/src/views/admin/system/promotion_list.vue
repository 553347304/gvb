<template>
	<div>
		<Form_promotion_create :title="title" v-model:visible="updateVisible" :form="updateForm"></Form_promotion_create>
		<div class="promotion_list_view">
			<Gvb_table
				_title="创建广告"
				_search="title"
				:_url="Api.URL.Promotion"
				:_columns="DefinePromotion.Columns"
				:_filter="[]"
				@_create="create"
				@_update="update"
			>
				<template #images="{record}:{record:TypePromotion.info}">
					<a-image :src="record.images" height="80px" style="border-radius: 5px"></a-image>
				</template>
				<template #href="{record}:{record:TypePromotion.info}">
					<a-link :href="record.href" target="_blank">{{ record.href }}</a-link>
				</template>
				<template #is_show="{record}:{record:TypePromotion.info}">
					<a-tag v-if="record.is_show" color="#0fc6c2">显示</a-tag>
					<a-tag v-else color="#ff5722">不显示</a-tag>
				</template>
			</Gvb_table>
		</div>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import Form_promotion_create from "@/components/admin/form/form_promotion_create.vue";
import {reactive, ref} from "vue";
import {DefinePromotion, TypePromotion} from "@/api/promotion_api";
import {Api} from "@/api/api";

const updateVisible = ref(false)
const updateForm = reactive(TypePromotion.Create())
const title = ref("")

function create() {
	title.value = "创建广告"
	Object.assign(updateForm, TypePromotion.Create())
	updateVisible.value = true
}

function update(record: TypePromotion.info) {
	title.value = "编辑广告"
	Object.assign(updateForm, record)
	updateVisible.value = true
}


</script>

<style scoped lang="scss">
</style>

