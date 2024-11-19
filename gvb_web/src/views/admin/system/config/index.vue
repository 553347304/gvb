<template>
	<div class="system_view">
		<div class="left">
			<router-link :to="{name: 'site_config'}">网站</router-link>
			<router-link :to="{name: 'email_config'}">邮箱</router-link>
		</div>

		<div class="right">
			<router-view/>
		</div>
	</div>
</template>

<script setup lang="ts">

</script>

<style scoped lang="scss">
.system_view {
	display: flex;

	.left {
		width: 160px;
		height: calc(100vh - 130px);
		border-radius: 5px;
		background-color: var(--color-bg-1);
		flex-direction: column;
		padding: 20px;

		a {
			color: var(--color-text-2);
			text-decoration: none;
			display: flex;
			border-radius: 5px;
			background-color: var(--color-fill-2);
			height: 50px;
			justify-content: center; // 文本水平居中
			align-items: center;
			margin-top: 10px;

			&.router-link-exact-active {
				background-color: var(--active);
				color: white;
			}
		}
	}

	.right {
		width: calc(100% - 180px);
		min-height: calc(100vh - 150px);
		border-radius: 5px;
		margin-left: 20px;
		background-color: var(--color-bg-1);
		padding: 20px;

	}
}
</style>