<template>
	<div v-if="typeof props.strings === 'string'"></div>
	<Typed :options="options" id="libTypewriter" v-else>
		<div class="typing"></div>
	</Typed>
</template>

<script setup lang="ts">
import {Typed} from "@duskmoon/vue3-typed-js";
import type {TypedOptions} from "@duskmoon/vue3-typed-js";
import {reactive} from "vue";

const options = reactive<TypedOptions>({
	strings: [""],
	loop: false,
	typeSpeed: 50,	// 打字速度	每毫秒出现几个字符
	backSpeed: 20,	// 回退速度
})
type Props = { strings: string[] | string, loop?: boolean, typeSpeed?: number | undefined, backSpeed?: number }
const props = defineProps<Props>()
options.loop = props.loop
if (props.typeSpeed) options.typeSpeed = props.typeSpeed;
if (props.backSpeed) options.backSpeed = props.backSpeed;
if (typeof props.strings === "object") options.strings = props.strings;
</script>
<style lang="scss" scoped>
#libTypewriter {
	display: flex;
	justify-content: center;
}
</style>