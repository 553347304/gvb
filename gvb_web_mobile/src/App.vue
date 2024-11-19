<template>
	<RouterView/>
</template>

<script setup lang="ts">
import {RouterView} from "vue-router";
import {useStore} from "@/stores";

const store = useStore()
store.loadThen()
store.loadToken()
</script>
