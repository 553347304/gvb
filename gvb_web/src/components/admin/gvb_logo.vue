<template>
	<div class="gvb_logo">
		<img :src="store.userInfo.avatar" alt="">
		<div class="logo_head">
			<div>{{store.userInfo.nick_name}}</div>
			<div>{{store.userInfo.user_name}}</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {useStore} from "@/stores";

const store = useStore()
</script>

<style scoped lang="scss">
.gvb_logo {
	height: 90px;
	display: flex;
	padding: 20px;
	align-items: center;
	border-bottom: 1px solid var(--bg);
	transition: all .3s;

	img {
		width: 60px; // 头像大小
		height: 60px;
		border-radius: 50%; // 圆形
		transition: all .3s;
	}

	.logo_head {
		margin-left: 20px;
		transition: all .3s;
		transform: scale(1);	// 基于自身缩放
		transform-origin: left;	// 中心点位置

		> div:nth-child(1) {
			font-size: 22px;
			font-weight: 600;
			margin-bottom: 2px;
		}

		> div:nth-child(2) {
			font-size: 12px;
		}

	}
}

.collapsed {
	.gvb_logo {
		padding: 2px;

		img {
			width: 40px;
			height: 40px;
			border-radius: 50%; // 圆形
		}

		.logo_head {
			transform: scale(0);	// 基于自身缩放
		}
	}
}
</style>