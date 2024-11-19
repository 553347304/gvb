<template>
	<div class="gvb_login">
		<Gvb_login_form @ok="okLogin"></Gvb_login_form>
	</div>
</template>

<script setup lang="ts">
import Gvb_login_form from "@/components/common/gvb_login_form.vue";
import router from "@/router";

// 登录成功跳转
function okLogin() {
	let back = window.history.state.back	// 上一次操作历史记录路径
	if (back === null) {
		router.push({name: "index"})
		return
	}
	router.push(back)
}
</script>


<style scoped lang="scss">
.gvb_login {
	background: url("/image/bg/bg.jpg") 50%/cover no-repeat;
	width: 100%;
	height: 100vh;

	.gvb_login_form {
		position: absolute;
		right: 0;
		width: 400px;
		height: 100%;
		background-color: var(--login_bg);
		display: flex;
		justify-content: center;
		align-items: center; // 居中
		padding: 0 40px;
	}
}
</style>