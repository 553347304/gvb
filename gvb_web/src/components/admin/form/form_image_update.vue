<template>
	<a-modal :title="props.title">
		<a-upload
			list-type="picture-card"
			action="/api/image"
			name="image"
			:headers="global.headers"
			v-model:file-list="fileList"
			image-preview
			multiple
			class="gvb_image_upload"
			@success="ok"
		/>
	</a-modal>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import {TypeImage} from "@/api/image_api";
import {type FileItem} from "@arco-design/web-vue";
import {global} from "@/global/global";
import {Base} from "@/BAIYIN/HTTP";
import {BaseLog} from "@/BAIYIN/log";

type Props = {form: TypeImage.Create; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
const fileList = ref([])
watch(() => props.form, (data) => {
	Object.assign(form, data)
}, {deep: true});

function ok(file: FileItem) {
	const result = file.response as Base.ResponseList<TypeImage.Create>
	if (!BaseLog(result)) return
	const data = result.data.list
	data.forEach((item) => {
		let code = item.is_success ? 0 : 1000
		if (!BaseLog({code: code, message: item.message, data: {}})) return
	})
}

</script>

<style lang="scss">
.gvb_image_upload {
	// 保证图片不变形
	//.arco-upload-list-picture{
	//	width: inherit;
	//}
}
</style>