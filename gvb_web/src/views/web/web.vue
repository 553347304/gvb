<template>
	<router-view class="web_view"/>
</template>

<script setup lang="ts">
import {useStore} from "@/stores";

const store = useStore()
store.loadSiteInfo()

</script>

<style scoped lang="scss">
// 媒体查询
@media screen and (min-width: 1400px) and (max-width: 1800px) {
	body {
		--container_width: 1200px;
	}
}

@media screen and (min-width: 1200px) and (max-width: 1400px) {
	body {
		--container_width: 1000px;
	}
}

@media screen and (min-width: 1000px) and (max-width: 1200px) {
	body {
		--container_width: 800px;
	}
}

.web_view {
	background-color: var(--bg);
}
</style>