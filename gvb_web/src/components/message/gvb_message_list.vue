<template>
	<div class="gvb_list_view">
		<div :class="{item: true, active: active === item.user_id}" v-for="item in props._data" @click="check(item)">
			<img class="avatar" :src="item.avatar" alt="">
			<div class="nick_name">
				<!--超出部分省略...-->
				<a-typography-paragraph :ellipsis="{rows:1,showTooltip:true,css:true}">
					{{ item.nike_name }}
				</a-typography-paragraph>
			</div>
			<div class="count">{{ item.total }}</div>
		</div>
	</div>
</template>

<script setup lang="ts">

import type {TypeMessage} from "@/api/message_api";
import {ref} from "vue";

interface Props {
	_data: TypeMessage.Info[]
}


const props = defineProps<Props>()
const active = ref(0)
const emits = defineEmits(["_check"])


function check(record: TypeMessage.Info) {
	active.value = record.user_id
	emits("_check", record)
}
</script>
<style scoped lang="scss">
.gvb_list_view {
	.item {
		padding: 10px 20px;
		width: 100%;
		height: 60px;
		display: flex;
		justify-content: space-between;
		align-items: center; // 文本居中
		cursor: pointer;

		.nick_name {
			width: 8rem;

			> .arco-typography {
				margin-bottom: 0;
			}
		}

		img {
			width: 40px;
			height: 40px;
			border-radius: 50%;
		}

		&.active {
			background-color: var(--color-fill-2);
			border-radius: 5px;
		}
	}
}
</style>
