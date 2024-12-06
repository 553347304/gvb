<script setup lang="ts">
import "@/assets/font.css"
import Gvb_theme from "@/components/common/gvb_theme.vue";
import {useStore} from "@/stores";
import Gvb_user_info_menu from "@/components/common/gvb_user_info_menu.vue";
import {onUnmounted, ref} from "vue";
import Web_search from "@/components/web/web_search.vue";

const props = defineProps<{ noScroll?: boolean }>()
const {noScroll = false} = props

const store = useStore()
console.log(store.userInfo)
const navList = ref([
	{title: "首页", path: "index"},
	{title: "新闻", path: "https://www.qq.com/"},
	{title: "搜索", path: "search"},
	{title: "聊天室", path: "chat"},
	{title: "AI", path: "bigModel"},
	{title: "官方文档", path: "https://www.bilibili.com"},
	{title: "管理后台", path: "home"},
])

const isShow = ref(true)
const scroll = () => isShow.value = document.documentElement.scrollTop >= 110
if (!noScroll) {
	isShow.value = false
	window.addEventListener("scroll", scroll)
}


onUnmounted(() => {
	if (!noScroll) {
		window.removeEventListener("scroll", scroll)
	}
})


</script>

<template>
	<div :class="{web_nav_view: true, isShow: isShow}">
		<div class="nav_container">
			<div class="left">
				<div class="logo">
					<div class="slogan">{{ store.siteInfo.slogan }}</div>
					<div class="en_slogan">{{ store.siteInfo.slogan_en }}</div>
				</div>
				<div class="nav">
					<template v-for="item in navList">
						<a v-if="item.path.startsWith('http')" :title="item.title"
						   :href="item.path" target="_blank">{{ item.title }}</a>
						<router-link v-else :to="{name: item.path}">{{ item.title }}</router-link>
					</template>
				</div>
				<web_search/>
			</div>

			<div class="right">
				<div class="login" v-if="!store.isLogin">
					<router-link :to="{name: 'login'}">登录</router-link>
				</div>
				<gvb_user_info_menu class="user_info" v-else/>
				<Gvb_theme class="theme"/>
			</div>
		</div>
	</div>
</template>

<style lang="scss">
.web_nav_view {
	top: 0;
	width: 100%;
	display: flex;
	justify-content: center;
	position: fixed;
	transition: all .3s;
	color: var(--nav_text_color);
	z-index: 100;

	&.isShow {
		background-color: var(--color-bg-1);
		color: var(--color-text-1);

		a {
			color: var(--color-text-2);
		}
	}

	a {
		text-decoration: none; // 去掉a标签下划线
		color: var(--nav_text_color);
		// 选中效果
		&.router-link-exact-active {
			color: var(--active);
		}
	}

	.nav_container {
		width: var(--container_width);
		height: 60px;
		display: flex;
		align-items: center;

		.left {
			width: 70%;
			display: flex;
			align-items: center;

			.logo {
				margin-right: 40px;

				.slogan {
					font-size: 21px;
					font-family: huakang_square_round_body_w7, serif;
				}

				.en_slogan {
					font-size: 12px;
					margin-top: 2px;
				}
			}

			.nav {
				> a {
					font-size: 16px;
					margin-right: 30px;
				}
			}


			.search {
				svg {
					cursor: pointer;
				}
			}
		}

		.right {
			width: 30%;
			display: flex;
			align-items: center;
			justify-content: end;

			.theme {
				margin-left: 20px;

				svg {
					cursor: pointer;
				}
			}
		}
	}
}
</style>