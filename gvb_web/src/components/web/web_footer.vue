<template>
	<div class="web_footer_view">
		<div class="footer_container">
			<div class="left">
				<div class="version">
					<span>version</span>
					<span>{{ store.siteInfo.version }}</span>
				</div>
				<div class="date">
					<span class="create_site_date">建站日期: {{ store.siteInfo.created_at }}</span>
					<span class="site_running_date">网站已运行: {{ countDown }}</span>
				</div>
				<div class="filings">
					<a href="https://beian.miit.gov.cn/#/Integrated/index" target="_blank">
						<img src="/image/icon/filings.png" alt="">
						{{ store.siteInfo.bei_an }}
					</a>
				</div>
			</div>

			<div class="right">
				<a :href="store.siteInfo.bilibili_url" target="_blank"><img src="/image/icon/bilibili.png" alt=""></a>
				<a :href="store.siteInfo.gitee_url" target="_blank"><img src="/image/icon/gitee.png" alt=""></a>
				<a :href="store.siteInfo.github_url" target="_blank"><img src="/image/icon/githup.png" alt=""></a>
				<a :href="store.siteInfo.qq_image" target="_blank"><img src="/image/icon/qq.png" alt=""></a>
				<a :href="store.siteInfo.wechat_image" target="_blank"><img src="/image/icon/wx.png" alt=""></a>
			</div>
		</div>


	</div>
</template>

<script setup lang="ts">
import {useStore} from "@/stores";
import {Time} from "@/BAIYIN/time";
import {onUnmounted, ref} from "vue";

const store = useStore()
const countDown = ref("")
let timer: number | undefined = undefined

// 定时获取时间
function countDownFunc() {
	timer = setInterval(() => {
		let date = store.siteInfo.created_at
		const oldTime = Time.Time(date).ms
		const newTime = Time.Time().ms
		countDown.value = Time.Convert(newTime - oldTime).Now
	}, 1000)
}

countDownFunc();
// 组件销毁
onUnmounted(() => {
	clearInterval(timer)
})
</script>

<style scoped lang="scss">
.web_footer_view {
	width: 100%;
	display: flex;
	justify-content: center;
	background-color: var(--color-bg-1);
	padding: 40px 0;
	color: var(--color-text-2);

	.footer_container {
		display: flex;
		width: var(--container_width);

		.left {
			width: 60%;
			display: flex;
			flex-direction: column;
			align-items: center;

			.version {
				margin-bottom: 25px;

				span {
					padding: 8px 15px;
					background-color: var(--color-fill-2);

				}

				span:nth-child(1) {
					background-color: var(--active);
					color: white;
					border-radius: 5px 0 0 5px;
				}

				span:nth-child(2) {
					border-radius: 5px 5px 0 0;
				}
			}

			.date {
				margin-bottom: 10px;

				.create_site_date {
					margin-right: 10px;
				}
			}

			.filings {
				a {
					display: flex;
					align-items: center;
					color: var(--color-text-2);
					text-decoration: none;

					img {
						margin-right: 5px;
					}
				}
			}
		}

		.right {
			width: 40%;
			display: flex;
			align-items: center;
			justify-content: center;

			a {
				margin-right: 40px;

				&:last-child {
					margin-right: 0;
				}

				img {
					width: 40px;
					height: 40px;
				}
			}
		}
	}
}
</style>