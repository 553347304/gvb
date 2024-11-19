<template>
	<div class="web_promotion">
		<a-carousel>
			<a-carousel-item v-for="item in List.list">
				<a class="link" :href="item.href" target="_blank">
					<img :src="item.images" :alt="item.title"/>
				</a>
			</a-carousel-item>
		</a-carousel>
	</div>
</template>

<script setup lang="ts">
import {reactive} from 'vue'
import { Api } from '@/api/api.js'
import {TypePromotion} from "@/api/promotion_api";
import {Base} from "@/BAIYIN/HTTP";

const List = reactive(Base.List<TypePromotion.Create>())
async function data(){
	let result = await Api.List(Api.URL.Promotion)
	List.list = result.data.list
}
data()
</script>

<style scoped lang="scss">
.web_promotion {
	height: 200px;

	.arco-carousel {
		height: 100%;

		.link {
			width: 100%;
			height: 100%;
			display: block; // 继承父级

			img {
				width: 100%;
				height: 100%;
				object-fit: cover; // 自适应图片尺寸
			}
		}
	}
}

</style>
