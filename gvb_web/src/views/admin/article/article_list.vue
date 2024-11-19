<template>
	<div>
		<Drawer_article_create title="发布文章" v-model:visible="createVisible"/>
		<Drawer_article_edit title="编辑正文" v-model:visible="editVisible"  :form="updateForm"/>
		<Form_article_update title="编辑文章" v-model:visible="updateVisible" :form="updateForm"/>
		<Gvb_table
			_title="发布文章"
			_search="title"
			:_url="Api.URL.Article"
			:_columns="DefineArticle.Columns"
			:_filter="DefineArticle.Filter"
			@_create="create"
			@_update="update"
		>
			<template #banner_url="{record}:{record:TypeArticle.List}">
				<a-image :src="record.banner_url" height="50px" style="border-radius: 5px"/>
			</template>

			<template #data="{record}:{record:TypeArticle.List}">
				<div class="article_data_col">
					<span><icon-eye/>1</span>
					<span><icon-thumb-up/>2</span>
					<span><icon-message/>3</span>
					<span><icon-star/>4</span>
				</div>
			</template>

			<template #tags="{record}:{record:TypeArticle.List}">
				<div class="article_tags_col">
					<a-tag v-for="(item, index) in record.tags" :color="Color.List[index]">{{ item }}</a-tag>
				</div>
			</template>

			<template #title="{record}:{record:TypeArticle.List}">
				<div class="article_title_highlight" v-html="record.title"></div>
			</template>

			<template #action_left="{record}:{record:TypeArticle.List}">
				<a-button type="primary" @click="edit(record)">编辑正文</a-button>
			</template>

		</Gvb_table>
	</div>
</template>

<script setup lang="ts">
import Gvb_table from "@/components/admin/gvb_table.vue";
import Form_article_update from "@/components/admin/form/form_article_update.vue";
import Drawer_article_create from "@/components/admin/drawer/drawer_article_create.vue";
import Drawer_article_edit from "@/components/admin/drawer/drawer_article_edit.vue";
import {reactive, ref} from "vue";
import {Color} from "@/global/color";
import {DefineArticle, TypeArticle} from "@/api/article_api";
import {Api} from "@/api/api";

const createVisible = ref(false)
const updateVisible = ref(false)
const editVisible = ref(false)
const updateForm = reactive(TypeArticle.Create())

function edit(record: TypeArticle.List) {
	Object.assign(updateForm, record)
	updateForm.id = record.id
	editVisible.value = true
}

function create() {
	createVisible.value = true
}

function update(record: TypeArticle.Create) {
	Object.assign(updateForm, record)
	updateForm.id = record.id
	updateVisible.value = true
}


</script>

<style  lang="scss">
@import '@/global/mixin.scss';

.article_data_col {

	> span {
		color: var(--color-text-2);
		@include margin-right(10px);
	}
}

.article_tags_col {

	.arco-tag {
		@include margin-right(10px);
	}
}

.article_title_highlight {
	em {
		color: rgb(var(--red-6));
	}
}
</style>