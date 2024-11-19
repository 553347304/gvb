<template>
	<div :class="{gvb_article_item: true, preview: props.preview}">
		<div class="cover">
			<a-image :src="props.form.banner_url"></a-image>
		</div>
		<div class="info">
			<div class="title">
				<template v-if="props.form.id === ''">{{ props.form.title }}</template>
				<router-link v-else :to="{name: 'article', params: {id: props.form.id}}"
							 v-html="props.form.title"/>
			</div>
			<div class="abstract">
				<a-typography-paragraph :ellipsis="{rows: 2,css:true}">
					{{ props.form.abstract }}
				</a-typography-paragraph>
			</div>
			<div class="data">
				<div class="category">
					<a-tag color="red" v-for="item in props.form.category">{{ item }}</a-tag>
				</div>
				<div class="date">
					<IconClockCircle/>
					{{ Time.Time().NowDay }}
				</div>
				<div class="article_data_col">
					<span><icon-eye/>{{ Random.integer(0, 100) }}</span>
					<span><icon-thumb-up/>{{ Random.integer(0, 100) }}</span>
					<span><icon-message/>{{ Random.integer(0, 100) }}</span>
					<span><icon-star/>{{ Random.integer(0, 100) }}</span>
				</div>
			</div>
		</div>
	</div>
</template>


<script setup lang="ts">
import {TypeArticle} from "@/api/article_api";
import {watch} from "vue";
import {Time} from "@/BAIYIN/time";
import {Random} from "mockjs";
import {Api} from "@/api/api";


type Props = { form: TypeArticle.Create; preview?: boolean };
const props = defineProps<Props>()
watch(() => props.form, (data) => {
	image(data.banner_id)
}, {deep: true});

async function image(id: number) {
	let result = await Api.List(Api.URL.Image, {search: "id", key: id})
	if (result.data.list.length === 0) return ""
	props.form.banner_url = result.data.list[0].path
}
</script>

<style lang="scss">//scoped
.gvb_article_item {
	width: 100%;
	padding: 20px;
	display: flex;
	border-radius: 5px;

	&.preview {
		background-color: var(--color-fill-2);
		transform: scale(0.7); // 整体缩放
		transform-origin: left; // 缩放轴
		width: 605px;
	}

	.cover {
		width: 30%;
		border-radius: 5px;
		overflow: hidden; // 范围溢出容器时 将溢出的内容隐藏

		&:hover {
			img {
				transform: scale(1.05);
			}
		}

		.arco-image {
			width: 100%;
			height: 100%;

			img {
				width: 100%;
				height: 120px;
				object-fit: cover;
				transition: all .3s;
			}
		}
	}

	.info {
		width: 120%;
		padding-left: 20px;
		color: var(--color-text-2);

		.title {
			font-weight: 600;
			font-size: 18px;

			a {
				color: var(--color-text-1);
				text-decoration: none;
			}

			em {
				color: red;
				margin-right: 2px;
			}
		}

		.abstract {
			height: 4rem;
		}

		.data {
			display: flex;
			align-items: center;

			.category {
				margin-right: 10px;

				span {
					margin-right: 10px;
				}
			}

			.date {
				margin-right: 10px;
			}

			.article_data_col {
				> span {
					margin-right: 10px;

					svg {
						margin-right: 2px;
					}
				}
			}
		}
	}
}
</style>