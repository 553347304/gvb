<template>
	<div class="chat_view">
		<web_nav no-scroll/>
		<main>
			<div class="container"><chat_list/></div>
		</main>
	</div>
</template>

<script setup lang="ts">
import Web_nav from "@/components/web/web_nav.vue";
import Chat_list from "@/views/admin/chat_group/chat_list.vue";


</script>

<style scoped lang="scss">
.chat_view {
	height: 100vh;

	main {
		width: 100%;
		display: flex;
		justify-content: center;
		padding-top: 80px;

		.container {
			width: var(--container_width);
		}
	}
}
</style>