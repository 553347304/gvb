<script setup lang="ts">

</script>

<template>
	<div class="big_model_option_view">
		<div class="left">
			<router-link :to="{name: 'base_option'}">基本配置</router-link>
			<router-link :to="{name: 'session_option'}">会话配置</router-link>
			<router-link :to="{name: 'reply_option'}">自动回复</router-link>
		</div>
		<div class="right">
			<router-view/>
		</div>
	</div>
</template>

<style scoped lang="scss">
.big_model_option_view {
	display: flex;

	.left {
		width: 160px;
		height: calc(100vh - 130px);
		background-color: var(--color-bg-1);
		border-radius: 5px;
		display: flex;
		flex-direction: column;
		padding: 20px;

		a {
			color: var(--color-text-2);
			text-decoration: none;
			display: flex;
			border-radius: 5px;
			background-color: var(--color-fill-2);
			height: 50px;
			justify-content: center;
			align-items: center;
			margin-top: 10px;
		}

		a.router-link-exact-active {
			background-color: var(--active);
			color: white;
		}
	}

	.right {
		width: 100%;
		margin-left: 20px;
		background-color: var(--color-bg-1);
		border-radius: 5px;
	}

}
</style>