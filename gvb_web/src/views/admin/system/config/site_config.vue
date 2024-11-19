<template>
	<div class="site_config">
		<div class="left">
			<div class="site_info">
				<Gvb_title title="站点配置"></Gvb_title>
				<a-form ref="formRef" :model="form" :label-col-props="{span: 5}" :wrapper-col-props="{span: 18}">
					<Unit_form label="网站标题" field="title"><a-input v-model="form.title"/></Unit_form>
					<Unit_form label="slogan" field="slogan"><a-input v-model="form.slogan"/></Unit_form>
					<Unit_form label="英文slogan" field="slogan_en"><a-input v-model="form.slogan_en"/></Unit_form>
					<Unit_form label="备案号" field="title"><a-input v-model="form.bei_an"/></Unit_form>
					<Unit_form label="版本号" field="version"><a-input v-model="form.version"/></Unit_form>
					<Unit_form label="建站日期" field="created_at"><a-input v-model="form.created_at"/></Unit_form>
					<Unit_form label="qq二维码"><Gvb_upload_image v-model="form.qq_image"/></Unit_form>
					<Unit_form label="微信二维码"><Gvb_upload_image v-model="form.wechat_image"/></Unit_form>
				</a-form>
				<div class="site_config_update">
					<a-button type="primary" @click="setInfo">更新</a-button>
				</div>
			</div>

		</div>


		<div class="right">
			<div class="user_info">
				<Gvb_title title="用户信息"></Gvb_title>
				<a-form :model="form" :label-col-props="{span: 3}" :wrapper-col-props="{span: 20}">
					<Unit_form label="昵称"><a-input v-model="form.name"/></Unit_form>
					<Unit_form label="工作"><a-input v-model="form.job"/></Unit_form>
					<Unit_form label="地址"><a-input v-model="form.addr"/></Unit_form>
					<Unit_form label="邮箱"><a-input v-model="form.email"/></Unit_form>
				</a-form>
			</div>
			<div class="link_info">
				<Gvb_title title="链接信息"></Gvb_title>
				<a-form :model="form" :label-col-props="{span: 3}" :wrapper-col-props="{span: 20}">
					<Unit_form label="bilibili"><a-input v-model="form.bilibili_url"/></Unit_form>
					<Unit_form label="gitee"><a-input v-model="form.gitee_url"/></Unit_form>
					<Unit_form label="githup"><a-input v-model="form.github_url"/></Unit_form>
				</a-form>
			</div>


		</div>
	</div>
</template>
<script setup lang="ts">

import Gvb_title from "@/components/system/gvb_title.vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {reactive, ref} from "vue";
import {ApiSetting, TypeSetting} from "@/api/setting_api";
import Gvb_upload_image from "@/components/image/gvb_upload_image.vue";
import {BaseLog} from "@/BAIYIN/log";

const url = "site"
const formRef = ref()
const form = reactive(TypeSetting.info())
async function getDate(){
	let res = await ApiSetting.List(url)
	Object.assign(form,res.data)
}
getDate()
async function setInfo(){
	if (await formRef.value.validate()) return false
	let result = await ApiSetting.Update(url, form)
	if (!BaseLog(result)) return
}


</script>
<style scoped lang="scss">
.site_config {
	display: flex;

	.arco-form{
		margin-top: 10px;
	}

	.left {
		width: 50%;

		.site_config_update{
			margin-top: 10px;
		}
	}

	.right {
		width: 50%;
	}
}
</style>