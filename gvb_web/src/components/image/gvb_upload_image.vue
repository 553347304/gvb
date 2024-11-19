<template>
	<div class="gvb_upload_image">
		<div class="line">
			<a-input v-model="src"/>
			<a-upload action="/api/image" name="image"
					  :headers="headers" @success="ok"
					  :show-file-list="false"/>
		</div>
		<a-image :src="src" v-if="src"/>
	</div>
</template>
<script setup lang="ts">
import {useStore} from "@/stores";
import {type FileItem, Message} from "@arco-design/web-vue";
import {ref, watch} from "vue";
import {Base} from "@/BAIYIN/HTTP";

const store = useStore()
const headers = {token: store.userInfo.token}
const props = defineProps({modelValue: {type: String}})
const emits = defineEmits(["update:modelValue"])
const src = ref("")

function ok(file: FileItem) {
	const response = file.response as Base.Response<string>
	if (response.code !== 0) {
		Message.error(response.message)
		return
	}
	Message.success(response.message)
	src.value = response.data[0]
}

watch(() => props.modelValue, () => {
	src.value = props.modelValue as string
})
watch(src, () => {
	emits("update:modelValue", src.value)
})
</script>

<style lang="scss">
.gvb_upload_image {
	width: 100%;

	.line {
		display: flex;

		.arco-input-wrapper {
			margin-right: 10px;
		}
	}

	.arco-image {
		margin-top: 10px;

		.arco-image-img {
			border-radius: 5px;
			// 固定图片大小
			height: 80px;
			object-fit: cover;
			max-width: 100%;
		}
	}
}
</style>