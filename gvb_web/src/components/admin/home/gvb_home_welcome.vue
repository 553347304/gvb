<template>
	<div class="gvb_home_welcome">
		<div class="title">早安 {{ store.userInfo.nick_name }} 请开始一天的工作吧</div>
		<div class="weather">河北·廊坊市 今日多云 10℃ 天气寒冷, 注意多加点衣服s</div>
		<div class="statistics">
			<span><IconUserGroup/>用户注册：{{ data.user }}</span>
			<span><IconBook/>文章总数：{{ data.article }}</span>
			<span><IconMessage/>群聊消息：{{ data.message }}</span>
			<span><IconBulb/>今日登录：{{ data.login_same_day }}</span>
			<span><IconSwap style="transform: rotate(90deg)"/>今日注册：{{ data.user_same_day }}</span>
		</div>
		<div class="link">
			<div>bilibili：<a :href="link">{{ link }}</a></div>
			<div>bilibili：<a :href="link">{{ link }}</a></div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {IconMessage, IconSwap, IconBulb, IconUserGroup, IconBook} from "@arco-design/web-vue/es/icon";
import {reactive} from "vue";
import {useStore} from "@/stores";
import {ApiDate, TypeDate} from "@/api/date_api";

const link = "https://www.bilibili.com/"
const store = useStore()
const data = reactive(TypeDate.Total())

async function Total() {
	let result = await ApiDate.Total()
	Object.assign(data, result.data)
}

Total()

</script>
<style lang="scss">
.gvb_home_welcome {
	width: 100%;
	background-image: url("/image/bg/bg.jpg");
	background-repeat: no-repeat;
	background-position: right;
	background-size: 500px;
	background-color: var(--color-bg-1);
	padding: 20px;
	border-radius: 5px;
	color: var(--color-text-2);

	.title {
		font-size: 22px;
		margin-top: 10px;
		font-weight: 400;
		color: var(--color-text-1);
	}

	.weather {
		margin: 20px 0 10px 0;
		white-space: pre-wrap;
	}

	.statistics {
		margin: 20px 0 10px 0;
		font-size: 16px;

		> span {
			margin-right: 20px;
		}
	}

	.link {
		margin: 30px 0 20px 0;

		> div {
			margin-bottom: 20px;

			&:last-child {
				margin-bottom: 0;
			}
		}

		a {
			text-decoration: none;
			color: var(--active);
		}
	}
}
</style>