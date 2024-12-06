<template>
	<div class="gvb_login_form">
		<a-form ref="formRef" :model="form" :label-col-props="{span: 0}" :wrapper-col-props="{span: 24}">
			<div class="title">用户登录</div>
			<a-form-item
				field="user_name" label="用户名"
				:rules="[{required:true,message:'请输入用户名'}]"
				:validate-trigger="['blur']"
			>
				<a-input v-model="form.user_name" placeholder="用户名">
					<template #prefix>
						<IconUser/>
					</template>
				</a-input>
			</a-form-item>
			<a-form-item
				field="password" label="密码"
				:rules="[{required:true,message:'请输入密码'}]"
				:validate-trigger="['blur']"
			>
				<a-input type="password" v-model="form.password" placeholder="密码">
					<template #prefix>
						<IconLock/>
					</template>
				</a-input>
			</a-form-item>
			<a-button @click="loginEmail" type="primary">登录</a-button>

			<div class="other-login">
				<div class="label">第三方登录</div>
				<div class="icons">
					<a href=""><img :src="Files.icon.qq" alt=""></a>
				</div>
			</div>
		</a-form>
	</div>
</template>

<script setup lang="ts">
import "@/assets/font.css"
import {IconLock, IconUser} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from "vue";
import {ApiUser, TypeUser} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";
import {Files} from "@/global/files";

const emits = defineEmits(["ok"]);
const store = useStore()

const form = reactive<TypeUser.LoginEmail>({
	user_name: "",
	password: "",
})
const formRef = ref()

// 组件销毁   清除表单数据
function formReset() {
	formRef.value.resetFields(Object.keys(form));
	formRef.value.clearValidate(Object.keys(form));
}

defineExpose({formReset})	// 抛出给父组件

async function loginEmail() {
	let value = await formRef.value.validate()
	if (value) {
		return
	}
	let result = await ApiUser.LoginEmail(form)
	if (result.code) {
		Message.error(result.message)
	}
	Message.success(result.message)
	await store.setToken(result.data)
	emits("ok");
}

</script>

<style scoped lang="scss">
.gvb_login_form {
	.title {
		font-size: 30px;
		font-weight: 700;
		text-align: center;
		margin-bottom: 10px;
		color: var(--active);
		font-family: huakang_square_round_body_w7, serif;
	}

	.other-login {
		margin-top: 20px;

		.label {
			display: flex;
			align-items: center;
			color: var(--color-text-3);
			justify-content: space-between;

			&::before {
				display: inline-flex;
				width: 110px;
				height: 1px;
				content: "";
				background-color: var(--color-text-3);
			}

			&::after {
				display: inline-flex;
				width: 110px;
				height: 1px;
				content: "";
				background-color: var(--color-text-3);
			}
		}

		.icons {
			display: flex;
			justify-content: center;
			margin-top: 10px;

			// 每个图标间隔
			> a {
				margin-right: 20px;

				// 最后一个位置不变
				&:last-child {
					margin-right: 0;
				}
			}

			img {
				width: 40px;
				height: 40px;
			}
		}

	}
}

</style>