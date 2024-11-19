<template>
	<div class="web_user_card">
		<div class="title">
			<i class="fa fa-quote-left"></i>
			<i class="fa fa-quote-right"></i>
			白音博客
		</div>
		<div class="info">
			<div class="item">
				<span>Name</span>
				<span>{{ store.siteInfo.name }}</span>
			</div>
			<div class="item">
				<span>JOB</span>
				<span>{{ store.siteInfo.job }}</span>
			</div>
			<div class="item">
				<span>ADDR</span>
				<span>{{ store.siteInfo.addr }}</span>
			</div>
			<div class="item">
				<span>WEB</span>
				<a :href="store.siteInfo.web">{{ store.siteInfo.web }}</a>
			</div>
		</div>
		<div class="image">
			<div class="qq">
				<img :src="store.siteInfo.qq_image" alt="">
				<span>QQ</span>
			</div>
			<div class="wx">
				<img :src="store.siteInfo.wechat_image" alt="">
				<span>微信</span>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import {useStore} from "@/stores";

const store = useStore()
</script>

<style scoped lang="scss">
.web_user_card {
	.title {
		background-color: var(--color-fill-2);
		border-radius: 5px;
		margin-bottom: 20px;
		color: var(--color-text-1);
		text-align: center;
		width: 100%;
		padding: 40px 0;
		position: relative;

		i {
			color: var(--color-text-4);
		}

		i:nth-child(1) {
			position: absolute;
			left: 10px;
			top: 10px;
		}

		i:nth-child(2) {
			position: absolute;
			right: 10px;
			bottom: 10px;
		}
	}

	.info {
		.item {
			margin-bottom: 10px;

			> span:nth-child(1) {
				margin-right: 5px;
				color: var(--color-text-1);
				font-weight: 600;
			}

			a {
				text-decoration: none;
				color: var(--active);
			}
		}
	}

	.image {
		display: flex;
		justify-content: space-around;
		margin-top: 20px;

		> div {
			display: flex;
			flex-direction: column;
			align-items: center;

			img {
				width: 80px;
				height: 80px;
			}

			span {
				margin-top: 10px;
			}
		}

	}
}
</style>