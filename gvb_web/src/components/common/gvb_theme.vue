<template>
	<div class="action.icon" @click="store.setTheme">
		<IconSun v-if="store.theme"></IconSun>
		<IconMoon v-else></IconMoon>
	</div>

</template>

<script setup lang="ts">

import {
	IconSun,
	IconMoon,
} from "@arco-design/web-vue/es/icon";
import {useStore} from "@/stores";
const store = useStore()
</script>


<style scoped lang="scss">
</style>