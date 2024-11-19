<template>
	<div class="email_config">
		<div class="left">
			<div class="site_info">
				<Gvb_title title="站点配置"></Gvb_title>
				<a-alert style="margin-top: 20px">配置邮箱之后, 系统将会通知你的邮箱</a-alert>
				<a-form ref="formRef" :model="form" :label-col-props="{span: 7}" :wrapper-col-props="{span: 17}">
					<Unit_form label="邮箱域名" field="host"><a-input v-model="form.host"/></Unit_form>
					<Unit_form label="邮箱端口" field="port"><a-input-number v-model="form.port"/></Unit_form>
					<Unit_form label="发信邮箱" field="user"><a-input v-model="form.user"/></Unit_form>
					<Unit_form label="密码" field="password"><a-input v-model="form.password"/></Unit_form>
					<Unit_form label="发信名称" field="default_from_email"><a-input v-model="form.default_from_email"/></Unit_form>
					<Unit_form label="ssl" field="use_ssl"><a-switch v-model="form.use_ssl"/></Unit_form>
					<Unit_form label="tls" field="user_tls"><a-switch v-model="form.user_tls"/></Unit_form>
				</a-form>
				<div class="config_update">
					<a-button type="primary" @click="setInfo">更新</a-button>
					<a-button type="primary" status="success" @click="send">测试连接</a-button>
				</div>
			</div>
		</div>

		<div class="right">
			<Gvb_title title="帮助"></Gvb_title>
			<div class="content">
				<div class="title">参考链接:</div>
				<a :href="help" target="_blank">{{ help }}</a>
			</div>
		</div>
	</div>
</template>
<script setup lang="ts">

import Gvb_title from "@/components/system/gvb_title.vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {reactive, ref} from "vue";
import {ApiSetting, TypeSetting} from "@/api/setting_api";
import {BaseLog} from "@/BAIYIN/log";

const url = "email"
const help = "https://wx.mail.qq.com/list/readtemplate?name=app_intro.html#/agreement/authorizationCode"
const formRef = ref()
const form = reactive(TypeSetting.Email())

async function getDate() {
	let result = await ApiSetting.List(url)
	Object.assign(form, result.data)
}

getDate()

async function setInfo() {
	if (await formRef.value.validate()) return false
	let result = await ApiSetting.Update(url, form)
	if (!BaseLog(result)) return
}

async function send() {
	if (await formRef.value.validate()) return false
	let result = await ApiSetting.Send({receive: form.user})
	if (!BaseLog(result)) return
}

</script>
<style scoped lang="scss">
.email_config {
	display: flex;

	.arco-form {
		margin-top: 10px;
	}

	.left {
		width: 30%;

		.config_update {
			margin-top: 10px;

			button {
				margin-right: 10px;
			}
		}
	}

	.right {
		width: 70%;
		margin-left: 20px;

		.content {
			display: flex;
			margin-top: 10px;

			.title {
				margin-right: 10px;
			}
		}
	}
}
</style>