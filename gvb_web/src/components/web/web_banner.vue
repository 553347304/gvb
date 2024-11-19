<template>
	<div class="web_banner_view">
		<div class="head">
			<div class="slogan">{{ props.data.slogan }}</div>
			<div class="abstract">
				<typewriter :strings="props.data.abstract" loop/>
			</div>
		</div>
		<a-carousel :auto-play="{interval: 5 * 1000}" indicator-type="dot" show-arrow="hover">
			<a-carousel-item v-for="image in props.data.image"><img :src="image" alt=""/></a-carousel-item>
		</a-carousel>
	</div>
</template>

<script setup lang="ts">
import typewriter from "@/lib/typewriter.vue";
import {TypeWeb} from "@/api/web_api";

type Props = { data: TypeWeb.Banner }
const props = defineProps<Props>()
</script>

<style scoped lang="scss">
.web_banner_view {
	width: 100%;
	height: 700px;
	position: relative;

	.head {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
		z-index: 2;
		color: white;

		.slogan {
			text-align: center;
			font-size: 44px;
			color: var(--nav_text_color);
		}

		.abstract {
			font-size: 24px;
			text-align: center;
			margin-top: 20px;
			color: var(--nav_text_color);
		}
	}

	.arco-carousel {
		height: 100%;

		img {
			width: 100%;
			height: 100%;
			object-fit: cover;
		}
	}
}
</style>