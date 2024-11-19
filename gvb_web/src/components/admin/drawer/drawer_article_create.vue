<template>
	<a-drawer  class="gvb_article_drawer" width="85%" unmount-on-close>
		<Form_article_update title="发布文章" v-model:visible="createVisible" :form="createForm"/>

		<MdEditor v-model="text" :on-upload-img="ApiImage.OnUploadImage" :theme="store.themeString"/>
		<template #footer>
			<a-button type="primary">取消</a-button>
			<a-button type="primary" status="success" @click="create">发布</a-button>
		</template>
	</a-drawer>
</template>

<script setup lang="ts">
import "md-editor-v3/lib/style.css"
import {MdEditor} from "md-editor-v3";
import {reactive, ref} from "vue";
import {useStore} from "@/stores";
import Form_article_update from "@/components/admin/form/form_article_update.vue";
import {TypeArticle} from "@/api/article_api";
import {ApiImage} from "@/api/image_api";

const text = ref("")
const store = useStore()

const createVisible = ref(false)
const createForm = reactive(TypeArticle.Create())

function create() {
	Object.assign(createForm, TypeArticle.Create())
	createForm.content = text.value
	createVisible.value = true
}

</script>

<style scoped lang="scss">
.md-editor-dark {
	--md-color: #999;
	--md-hover-color: #bbb;
	--md-bk-color: #141414;  // 黑夜颜色
	--md-bk-color-outstand: #111;
	--md-bk-hover-color: #1b1a1a;
	--md-border-color: #2d2d2d;
	--md-border-hover-color: #636262;
	--md-border-active-color: #777;
	--md-modal-mask: #00000073;
	--md-scrollbar-bg-color: #0f0f0f;
	--md-scrollbar-thumb-color: #2d2d2d;
	--md-scrollbar-thumb-hover-color: #3a3a3a;
	--md-scrollbar-thumb-active-color: #3a3a3a;
}
.md-editor {
	height: calc(100vh - 130px);
}
</style>