<template>
	<div class="gvb_user_info_menu">
		<a-dropdown @select="select">
			<div class="gvb_user_info_dropdown">
				<img :src="store.userInfo.avatar" alt="">
				<span class="gvb_user_info_dropdown_span">{{ store.userInfo.nick_name }}</span>
				<IconDown></IconDown>
			</div>
			<template #content>
				<a-doption
					class="gvb_bg_line"
					v-for="item in dopList"
					:value="item.name"
				>{{ item.title }}
				</a-doption>
			</template>
		</a-dropdown>
	</div>
</template>

<script setup lang="ts">
import {IconDown} from "@arco-design/web-vue/es/icon";
import {useStore} from "@/stores";
import router from "@/router";

const store = useStore()

import type {tabType} from "@/types";

const dopList: tabType[] = [
	{name: "home", title: "后台首页",},
	{name: "article_list", title: "文章列表",},
	{name: "user_list", title: "用户列表",},
	//{name: "chat_list", title: "群聊管理",},
	// {name: "log_list", title: "系统日志",},
	{name: "logout", title: "退出登录",},
]

function select(value: any) {
	// 注销用户
	if (value === "logout") {
		store.logout()
		router.push({name: "index"})	// 回到首页
		return
	}
	router.push({
		name: value
	})
}

</script>

<style scoped lang="scss">
.gvb_user_info_menu {
	img {
		width: 30px;
		height: 30px;
		border-radius: 50%;
	}

	.gvb_user_info_dropdown {
		display: flex;
		cursor: pointer;
		align-items: center;

		.gvb_user_info_dropdown_span {
			margin: 0 5px;
		}
	}

}

.gvb_bg_line {
	border-bottom: 1px solid var(--bg);
}

</style>