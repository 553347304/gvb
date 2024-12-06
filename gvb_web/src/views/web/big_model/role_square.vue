<script setup lang="ts">
import "@/assets/font.css"
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {h, ref, type VNode} from "vue";
import {useStore} from "@/stores";
import {Notification, Button} from '@arco-design/web-vue';
import {BaseLog} from "@/BAIYIN/log";

const store = useStore()
const tagList = ref<TypeBigModel.Square[]>([])
const useRoleList = ref<TypeBigModel.RoleSquare[]>([])
const activeTag = ref(0)

async function getData() {
	let result = await ApiBigModel.Square();
	tagList.value = result.data.list
	if (tagList.value.length > 0) checkTag(tagList.value[0])
}

getData()

function checkTag(record: TypeBigModel.Square) {
	useRoleList.value = record.role_list
	activeTag.value = record.id
}

// 领取积分
async function UserScope() {
	let result = await ApiBigModel.UserScope(true)
	BaseLog(result)
	Notification.clear()	// 清空提示框
}

async function UserIsScope() {
	let result = await ApiBigModel.UserIsScope()
	if (result.data) {
		const id = `${Date.now()}`;
		Notification.info({
			id,
			title: '积分领取通知',
			content: result.message,
			duration: 0,
			footer: (): VNode[] => [
				h(Button, {style: {marginRight: "10px"}, onClick: () => Notification.remove(id),}, () => '取消'),
				h(Button, {type: 'primary', onClick: () => UserScope(),}, () => '领取'),
			],
			closable: true,
		})
	}
}

UserIsScope()

</script>

<template>
	<div class="role_square_view">
		<div class="banner">
			<div class="head">{{ store.bigModelInfo.setting.title }}</div>
			<div class="slogan">{{ store.bigModelInfo.setting.slogan }}</div>
		</div>
		<div class="role_tags">
			<span :class="{item: true, active: item.id === activeTag}" v-for="item in tagList" @click="checkTag(item)">
				{{ item.title }}
			</span>
		</div>

		<div class="role_list">
			<div class="item" v-for="item in useRoleList" @click="ApiBigModel.CheckRole(item)">
				<img :src="item.icon" alt="">
				<a-typography-text class="title" :ellipsis="{rows: 1, css: true}">{{ item.name }}</a-typography-text>
				<a-typography-text class="abs" :ellipsis="{rows: 2, css: true}">{{ item.abstract }}</a-typography-text>
			</div>
		</div>
	</div>
</template>

<style scoped lang="scss">
.role_square_view {
	width: 100%;
	height: 100%;
	padding: 20px;
	border-radius: 5px;
	background-color: var(--color-bg-1);
	overflow-y: auto;

	.banner {
		text-align: center;

		.head {
			color: var(--active);
			font-size: 40px;
			margin-top: 5px;
			font-family: huakang_square_round_body_w7, serif;
		}

		.slogan {
			color: var(--color-text-2);
			margin-top: 20px;
			font-size: 20px;

		}
	}

	.role_tags {
		margin-top: 40px;

		.item {
			cursor: pointer;
			margin-right: 20px;
			background-color: var(--color-fill-2);
			padding: 10px 30px;
			border-radius: 20px;
			color: var(--color-text-2);

			// 鼠标
			&:hover, &.active {
				background-color: var(--active);
				color: white;
			}
		}
	}

	.role_list {
		// 图片排列
		display: grid;
		grid-template-columns: repeat(6, 1fr); // 图片一行显示n个
		grid-row-gap: 40px; // 行间距
		grid-column-gap: 40px; // 列间距
		width: 100%;
		margin-top: 60px;

		.item {
			width: 200px;
			height: 240px;
			border-radius: 5px;
			background-color: var(--color-fill-2);
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			padding: 10px 20px;
			cursor: pointer;
			transition: all .3s;

			&:hover {
				transform: translateY(-10px);
				background-color: var(--color-fill-3);
				box-shadow: 0 5px 5px 5px rgba(0, 0, 0, 0.05); // 阴影
			}

			> .title {
				font-size: 18px;
				margin-top: 10px;
			}

			> .abs {
				font-size: 14px;
				color: var(--color-text-2);
				margin-top: 5px;
			}
		}

		img {
			width: 40px;
			height: 40px;
			border-radius: 50%;
		}
	}
}
</style>