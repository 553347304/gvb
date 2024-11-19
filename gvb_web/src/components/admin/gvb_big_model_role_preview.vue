<script setup lang="ts">
import {MdPreview} from "md-editor-v3";
import {useStore} from "@/stores";
import {Time} from "@/BAIYIN/time";
import {TypeBigModel} from "@/api/big_model_api";

const store = useStore();
type Props = { data: TypeBigModel.RoleCreate }
const props = defineProps<Props>()
</script>

<template>
	<div class="gvb_big_model_role_preview_view">
		<div class="head">
			<div>预览角色</div>
			<div>{{ props.data.name }}</div>
		</div>
		<div class="body">
			<div class="item">
				<div class="avatar">
					<img :src="props.data.icon" alt="">
				</div>
				<div class="main">
					<div class="date">{{ Time.Time().Now }}</div>
					<md-preview :model-value="props.data.prologue" class="content"/>
				</div>
			</div>
		</div>
		<div class="menu">
			<div class="btn">
				新话题
			</div>
			<div class="ipt">
				<IconSend/>
				<a-input placeholder="来说点什么吧"></a-input>
			</div>
		</div>

	</div>
</template>

<style scoped lang="scss">
.gvb_big_model_role_preview_view {
	border: 1px solid var(--bg);
	border-radius: 5px;

	.head {
		border-bottom: 1px solid var(--bg);
		padding: 5px 0;

		> div {
			display: flex;
			justify-content: center;
			align-items: center;

			&:last-child {
				font-size: 12px;
			}
		}
	}

	.body {
		padding: 20px;
		min-height: 400px;

		.item {
			display: flex;

			.avatar {
				width: 30px;
				height: 30px;

				img {
					width: 100%;
					height: 100%;
					object-fit: cover;
					border-radius: 5px;
				}
			}

			.main {
				margin-left: 20px;

				.date {
					font-size: 12px;
				}

				.content {
					background-color: var(--bg);

					padding: 5px;
					border-radius: 5px;
					margin-top: 5px;

					article > p:first-child {
						margin-top: 5px;
					}
				}
			}


		}
	}

	.menu {
		padding: 20px;
		display: flex;
		align-items: center;
		width: 100%;

		.btn {
			display: flex;
			width: 80px;
			height: 30px;
			border-radius: 20px;
			justify-content: center;
			align-items: center;
			background-color: var(--active);
			color: white;
		}

		.ipt {
			width: 100%;
			margin-left: 10px;
			position: relative;

			svg {
				position: absolute;
				right: 10px;
				top: 50%;
				transform: translateY(-50%);
				color: var(--color-text-2);
			}

			.arco-input-wrapper {
				border-radius: 10px;
				pointer-events: none; // 输入框禁止输入
			}
		}
	}
}
</style>