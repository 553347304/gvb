<script setup lang="ts">
import {Tabbar, TabbarItem} from "vant";
import {ref} from "vue";
const tab = ref()
</script>

<template>
	<Tabbar v-model="tab" route>
		<TabbarItem icon="wap-home" to="/">首页</TabbarItem>
		<TabbarItem icon="manager" to="/my">我的</TabbarItem>
	</Tabbar>
</template>

<style scoped lang="scss">

</style>