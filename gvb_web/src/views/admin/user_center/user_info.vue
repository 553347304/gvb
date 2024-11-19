<template></template>
<!--<template>-->
<!--	<div class="user_info_view">-->
<!--		<gvb_cropper ref="clipperRef"-->
<!--					 :type="clipperData.type"-->
<!--					 :allow-type-list="clipperData.allowTypeList"-->
<!--					 :limit-size="clipperData.limitSize"-->
<!--					 :fixed-number="clipperData.fixedNumber"-->
<!--					 :fixed-number-aider="clipperData.fixedNumberAider"-->
<!--					 :preview-width="clipperData.previewWidth"-->
<!--					 @confirm="onConfirm">-->

<!--		</gvb_cropper>-->

<!--	</div>-->
<!--</template>-->

<!--<script setup lang="ts">-->
<!--import {ref} from "vue";-->
<!--import {useStore} from "@/stores";-->

<!--const isLaptops1 = ""//isLaptops-->
<!--const formRef = ref()-->

<!--const store = useStore()-->

<!--interface IClipper {-->
<!--	type: string // 上传类型-->
<!--	allowTypeList: string[] // 接收允许上传的图片类型-->
<!--	limitSize: number // 限制大小-->
<!--	fixedNumber: number[] // 截图框的宽高比例-->
<!--	fixedNumberAider?: number[] // 侧边栏收起截图框的宽高比例-->
<!--	previewWidth: number // 预览宽度-->
<!--	previewWidthAider?: number // 侧边栏收起预览宽度-->
<!--}-->

<!--const clipperData = ref<IClipper>({-->
<!--	type: '',-->
<!--	allowTypeList: [],-->
<!--	limitSize: 1,-->
<!--	fixedNumber: [],-->
<!--	previewWidth: 0-->
<!--})-->

<!--const clipperRef = ref()-->

<!--function showCropper() {-->
<!--	clipperData.value = {-->
<!--		type: 'browserLogo', // 该参数可根据实际要求修改类型-->
<!--		allowTypeList: ['png', 'jpg', 'jpeg',], // 允许上传的图片格式-->
<!--		limitSize: 1, // 限制的大小-->
<!--		fixedNumber: [1, 1],  // 截图比例，可根据实际情况进行修改-->
<!--		previewWidth: 100 // 预览宽度-->
<!--	}-->
<!--	// 打开裁剪组件-->
<!--	clipperRef.value.uploadFile()-->
<!--}-->


<!--function onConfirm(val: string) {-->
<!--	form.avatar = val-->
<!--	userInfoUpdate()-->
<!--}-->


<!--</script>-->
