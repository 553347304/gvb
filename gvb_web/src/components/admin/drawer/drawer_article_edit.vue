<template>
	<a-drawer :title="props.title" width="85%" unmount-on-close>
		<MdEditor v-model="form.content" :on-upload-img="ApiImage.OnUploadImage" :theme="store.themeString"/>
		<template #footer>
			<a-button type="primary">取消</a-button>
			<a-button type="primary" status="success" @click="ok">发布</a-button>
		</template>
	</a-drawer>
</template>

<script setup lang="ts">
import "md-editor-v3/lib/style.css"
import {MdEditor} from "md-editor-v3";
import {reactive, watch} from "vue";
import {useStore} from "@/stores";
import {TypeArticle} from "@/api/article_api";
import {Api} from "@/api/api";
import {ApiImage} from "@/api/image_api";
import {BaseLog} from "@/BAIYIN/log";

const store = useStore()

type Props = {form: TypeArticle.Create; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
watch(() => props.form, (data) => {
	Object.assign(form, data);
}, {deep: true});

async function ok() {
	let result = await Api.Update(Api.URL.Article, form, form.id);
	return BaseLog(result)
}
</script>

<style scoped  lang="scss">
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
