<script setup lang="ts">
import Gvb_tabbar from "@/components/gvb_tabbar.vue";
import {Button, Popup} from 'vant';
import {Form, Field, CellGroup} from 'vant';
import {useStore} from "@/stores";
import router from "@/router";
import {reactive, ref} from "vue";
import {ApiUser, TypeUser} from "@/api/user_api";
import {BaseLog} from "@/BAIYIN/log";
import {Bind} from "@/components/bind/bind";

const store = useStore()
const showLogin = ref(false)
const form = reactive(TypeUser.LoginEmail())

function goto(name: string) {
	router.push({name: name})
}

async function onSubmit() {
	let result = await ApiUser.LoginEmail(form)
	let is = BaseLog(result)
	if (is) {
		await store.setToken(result.data)
	}
	showLogin.value = false
}

const userList = [
	{i: "fa fa-star", title: "我的收藏", click: () => goto('coll'),},
	{i: "fa fa-lock", title: "修改密码"},
	{i: "fa fa-envelope-o", title: "绑定邮箱"},
	{i: "fa fa-mail-forward", title: "注销退出", click: () => store.logout(),},
]
</script>

<template>
	<div class="my_view">
		<div class="head">
			<div class="info no_login" v-if="!store.isLogin">
				<div class="title">欢迎来到Koneko</div>
				<div class="abstract">登录享更多精彩</div>
			</div>
			<div class="info login_data" v-else>
				<img :src="store.userInfo.avatar" alt="" class="avatar">
				<div class="title">{{ store.userInfo.user_name }}</div>
				<div class="edit_info">
					<button>编辑资料</button>
				</div>
				<div class="abstract">{{ store.userInfo.nick_name }}</div>
			</div>
		</div>
		<div class="body">
			<div class="login_button" v-if="!store.isLogin">
				<Button type="primary" :block="true" @click="showLogin = true">登录/注册</Button>
			</div>
			<div class="action" v-else>
				<template v-for="item in userList">
					<div class="item" @click="item.click">
						<i :class="item.i"></i>
						<span>{{ item.title }}</span>
					</div>
				</template>
			</div>
		</div>

		<Popup v-model:show="showLogin" position="bottom" style="padding: 10px 0">
			<Form @submit="onSubmit">
				<CellGroup inset>
					<Field v-model="form.user_name" v-bind="Bind.form('用户名')"/>
					<Field v-model="form.password" v-bind="Bind.form('密码')" type="password"/>
				</CellGroup>
				<div style="margin: 16px;">
					<Button round block type="primary" native-type="submit">登录</Button>
				</div>
			</Form>
		</Popup>

		<gvb_tabbar/>
	</div>
</template>

<style scoped lang="scss">
.my_view {
	min-height: 100vh;
	background-color: var(--bg);

	.head {
		background-image: url("/image/bg/bg.jpg");
		background-repeat: no-repeat;
		background-size: cover;
		width: 100vw;
		height: 200px;
		display: flex;
		justify-content: center;
		position: relative;

		.info {
			width: 90vw;
			height: 100px;
			background-color: white;
			border-radius: 5px;
			position: absolute;
			bottom: 0;
			transform: translateY(50%);
			padding: 10px;
			display: flex;
			justify-content: center;
			align-items: center;

			.avatar {
				background-color: white;
				width: 50px;
				height: 50px;
				border-radius: 50%;
				position: absolute;
				left: 20px;
				top: 0;
				transform: translateY(-50%);
			}

			.title {
				color: #444;
				position: absolute;
				left: 80px;
				top: 13px;
				transform: translateY(-50%);
				font-weight: 600;
				font-size: 14px;

			}

			.abstract {
				color: #444;
			}

			.edit_info {
				position: absolute;
				right: 10px;
				top: 10px;
				height: 30px;

				button {
					background-color: #1989fa;
					border: none;
					color: white;
					font-size: 14px;
					padding: 4px 10px;
					border-radius: 10px;
				}
			}

		}
	}

	.body {
		margin-top: 70px;
		padding: 20px;

		.action {
			background-color: white;
			border-radius: 5px;
			display: flex;


			.item {
				width: 25%;
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: center;
				padding: 20px 0;
				color: #333;
				font-size: 14px;

				i {
					font-size: 20px;
					margin-bottom: 5px;
				}
			}
		}
	}
}
</style>