<script setup lang="ts">

import {useRoute} from "vue-router";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {reactive, ref, watch} from "vue";
import {TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import Unit_form from "@/components/unit/unit_form.vue";
import {BaseLog} from "@/BAIYIN/log";
import router from "@/router";
import Gvb_big_model_message from "@/components/admin/gvb_big_model_message.vue";
import {Message} from "@arco-design/web-vue";

const route = useRoute()
const data = reactive(Base.List<TypeBigModel.SessionList>())
const params = reactive(TypeApi.PageInfo())


async function getSessionList() {
	params.search = "role_id"
	params.key = Number(route.query.roleId)
	let result = await ApiBigModel.SessionList(params);
	Object.assign(data, result.data)
}

const formRef = ref()
const form = reactive(TypeBigModel.SessionUpdate())
const visible = ref(false)
const sessionNameIptRef = ref()
const open = () => sessionNameIptRef.value.focus();	// 输入框焦点

async function edit(record: TypeBigModel.SessionList) {
	if (await formRef.value.validate()) return false

	form.session_id = record.id
	form.nick_name = record.session_name
	visible.value = true
}

async function ok() {
	let result = await ApiBigModel.SessionUpdate(form);
	await getSessionList()
	return BaseLog(result)
}

async function SessionRemove(record: TypeBigModel.SessionList) {
	let result = await ApiBigModel.SessionRemove(record.id);
	await getSessionList()
	return BaseLog(result)
}

async function checkSession(record: TypeBigModel.SessionList) {
	await router.push({
		name: route.name,
		query: {
			roleId: route.query.roleId,
			sessionId: record.id,
		}
	})
}

watch(() => route.query.roleId, () => {
	getRoleDetail()
	getSessionList()
}, {immediate: true})

const detail = reactive(TypeBigModel.RoleDetail())

async function getRoleDetail() {
	let result = await ApiBigModel.RoleDetail(Number(route.query.roleId));
	Object.assign(detail, result.data)
}

async function SessionCreate() {
	let result = await ApiBigModel.SessionCreate(Number(route.query.roleId))
	if (result.code) {
		Message.error(result.message)
		return
	}
	await router.push({
		name: route.name as string,
		query: {
			roleId: route.query.roleId,
			sessionId: result.data,
		}
	})
	await getSessionList()
	Message.success(result.message)
}

</script>

<template>
	<div class="session_view">
		<a-modal title="编辑会话" v-model:visible="visible" :on-before-ok="ok" @open="open">
			<a-form ref="formRef" :model="form"></a-form>
			<unit_form label="会话名称" field="nick_name">
				<a-input ref="sessionNameIptRef" v-model="form.nick_name"/>
			</unit_form>
		</a-modal>

		<div class="left">
			<gvb_big_model_message :session_id="Number(route.query.sessionId)" :info="detail"/>
		</div>

		<div class="right">
			<div class="role_info">
				<div>
					<img :src="detail.icon" alt="">
					<div class="info">
						<div class="name">{{ detail.name }}</div>
						<div class="total">{{ detail.chat_total }}</div>
					</div>
				</div>
				<a-typography-text class="abs" :ellipsis="{rows: 2, css: true}">{{ detail.abstract }}
				</a-typography-text>
			</div>
			<div class="session_list">
				<a-button class="add" type="outline" long @click="SessionCreate">创建新会话</a-button>
				<div :class="{item: true, active: item.id === Number(route.query.sessionId)}"
					 v-for="item in data.list"
					 @click="checkSession(item)">
					<IconMessage class="message"/>
					<div class="name">{{ item.session_name }}</div>
					<IconEdit title="编辑" class="edit" @click="edit(item)"/>
					<IconDelete title="删除" class="delete" @click="SessionRemove(item)"/>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped lang="scss">
@mixin icon_size($right_size) {
	position: absolute;
	font-size: 20px;
	cursor: pointer;
	opacity: 0;
	transition: all .3s;
	right: $right_size;
}

.session_view {
	width: 100%;
	display: flex;
	justify-content: space-between;

	.left {
		width: calc(100% - 260px);
		height: calc(100vh - 100px);
		background-color: var(--color-bg-1);
		border-radius: 5px;

		.head {
			text-align: center;
			padding: 20px 0;
			border-bottom: 1px solid var(--bg);
			color: var(--color-text-2);
			font-size: 16px;
		}
	}

	.right {
		width: 240px;
		height: calc(100vh - 100px);
		background-color: var(--color-bg-1);

		.role_info {
			width: 100%;
			padding: 20px;
			border-bottom: 1px solid var(--bg);

			> div {
				display: flex;
				align-items: center;

				img {
					width: 40px;
					height: 40px;
					border-radius: 50%;
				}

				.info {
					margin-left: 10px;

					.name {
						font-size: 16px;
						font-weight: 600;
					}

					.total {
						font-size: 12px;
						margin-top: 5px;
						color: var(--color-text-2);
					}
				}
			}

			.abs {
				margin-top: 10px;
				color: var(--color-text-2);
				margin-bottom: 0;

			}
		}

		.session_list {
			padding: 20px;

			.add {
				border-radius: 5px;
				margin-bottom: 20px;
				height: 40px;
			}

			.item {
				padding: 10px 20px;
				border-radius: 5px;
				background-color: var(--color-fill-2);
				display: flex;
				align-items: center;
				margin-bottom: 20px;
				position: relative;

				&.active {
					background-color: var(--color-fill-4);
				}

				.message {
					font-size: 20px;
					margin-right: 10px;
				}

				.edit {
					@include icon_size(40px);
				}

				.delete {
					@include icon_size(10px);
					color: #d71345;
				}


				&:hover {
					.edit, .delete {
						opacity: 1;
					}
				}

				.name {
					width: 5rem;
					// 文本不换行
					text-overflow: ellipsis;
					white-space: nowrap;
					overflow: hidden; // 超出部分省略...
				}

			}
		}
	}
}
</style>