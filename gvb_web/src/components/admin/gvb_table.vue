<template>
	<div class="gvb_table">
		<div class="gvb_table_head">

			<div class="action_create">
				<a-button @click="emits('_create')" type="primary">{{ props._title }}</a-button>
			</div>

			<div class="action_group">
				<a-select placeholder="操作" :options="DefineUser.Operation" v-model="actionValue"
						  allow-clear style="width: 120px"></a-select>
				<a-button type="primary" status="danger" @click="t.Batch">执行</a-button>
			</div>

			<div class="action_search">
				<a-input-search :placeholder="'搜索'+t._search" v-model="PageInfo.key"
								@keydown.enter="t.Search"
								@search="t.Search"></a-input-search>
			</div>

			<div class="action_filter" v-if="props._filter.length">
				<a-select :placeholder="item.label" :options="item.option"
						  @change="t.Filter($event,item)" style="width: 120px"
						  v-for="item in props._filter"></a-select>
			</div>

			<div class="action_flush">
				<a-button @click="t.Refresh">
					<IconRefresh></IconRefresh>
				</a-button>
			</div>
		</div>
		<a-spin class="gvb_table_data" :loading="isLoading" tip="加载中">
			<div class="gvb_table_source">

				<a-table row-key="id" :columns="props._columns" :data="data.list"
						 :row-selection="{type: 'checkbox',showCheckedAll: true,onlyCurrent: false}"
						 v-model:selectedKeys="selectedKeys" :pagination="false"
						 :stripe="true" :scroll="{ x: 2000,y: 1000}"
						 column-resizable :bordered="{cell:true}"
				>
					<template #columns>
						<template v-for="item in props._columns">

							<a-table-column v-if="item.render" :title="item.title as string">
								<!--自定义最先渲染-->
								<template #cell="data">
									<component :is="item.render(data)"/>
								</template>
							</a-table-column>

							<!--自定义-->
							<a-table-column v-else-if="!item.slotName" :title="item.title as string"
											:data-index="item.dataIndex"/>
							<a-table-column v-else :title="item.title as string">

								<template #cell="{record}" v-if="item.slotName === 'action'">
									<div class="gvb_cell_actions">

										<slot name="action_left" :record="record"></slot>

										<a-button type="primary" @click="emits('_update', record)">编辑</a-button>

										<a-popconfirm content="是否确认删除" @ok="t.Remove(record)">
											<a-button type="primary" status="danger">删除</a-button>
										</a-popconfirm>

									</div>
								</template>

								<template #cell="{record}" v-else-if="item.slotName === 'created_at'">
									<span>{{ Time.Time(record.created_at).Now }}</span>
								</template>

								<template #cell="{record}" v-else-if="item.slotName === 'avatar'">
									<a-avatar :image-url="record.avatar"/>
								</template>

								<template #cell="{record}" v-else>
									<slot :name="item.slotName" :record="record"/>
								</template>

							</a-table-column>
						</template>
					</template>
				</a-table>
			</div>
			<div class="gvb_table_page">
				<a-pagination :total="data.total"
							  @change="t.List"
							  v-model:current="PageInfo.page"
							  :default-page-size="PageInfo.limit"
							  :show-total="true"
							  show-jumper/>
			</div>
		</a-spin>
	</div>
</template>

<script setup lang="ts">
import {IconRefresh} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from 'vue';
import {DefineUser} from "@/api/user_api";
import {Message, type TableColumnData} from "@arco-design/web-vue"
import {Time} from "@/BAIYIN/time";
import {Api, TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import {DefineUtils} from "@/api/api_utils/type";
import {BaseLog} from "@/BAIYIN/log";

const PageInfo = reactive(TypeApi.PageInfo())
const isLoading = ref(false)	// 加载中
const emits = defineEmits(["_create", "_update"])


interface props {
	_columns: TableColumnData[]
	_title: string
	_url: string
	_search: string
	_filter: DefineUtils.optionsList[]
}

const props = defineProps<props>();

// 批量操作
const selectedKeys = ref([""]);
const actionValue = ref<number | string | undefined>(undefined)
const data = reactive(Base.List<any>())	// 按钮

class table {
	private _url = props._url
	_search = props._search

	private Reset = async (s: string) => {
		Message.success(s)
		await this.List()
	}

	List = async () => {
		isLoading.value = true
		PageInfo.source = "like"
		let result = await Api.List(this._url, PageInfo)
		data.list = result.data.list
		data.total = result.data.total
		isLoading.value = false
	}

	Refresh = async () => {
		await this.Reset("刷新成功")
	}

	Remove = async (record: DefineUtils.uid) => {
		let R = await Api.Remove(this._url, [record.id])
		await this.Reset(R.message)
	}

	Search = async () => {
		PageInfo.search = this._search
		await this.Reset("查询成功")
	}

	// 过滤器
	Filter = async (event: any, item: DefineUtils.optionsList) => {
		PageInfo.search = item.search
		PageInfo.key = event
		await this.Reset("筛选成功")
	}

	// 批量操作
	Batch = async () => {
		let value = actionValue.value
		let data = selectedKeys.value
		if (value === undefined || value === "" || data.length === 0) {
			Message.warning("未选择")
			return
		}
		switch (value) {
			case 0:
				let result = await Api.Remove(this._url, data);
				if (!BaseLog(result)) return
				break;
		}
		await this.List();
		selectedKeys.value = []
	}
}

const t = new table()
t.List()
</script>

<style lang="scss">
.gvb_table {
	padding: 10px 20px 20px 20px;
	background-color: var(--color-bg-1);
	border-radius: 5px;

	.gvb_table_head {
		padding: 20px 20px 10px 20px;
		border-bottom: 1px solid var(--bg);
		display: flex; // 横向布局
		align-items: center; // 居中
		position: relative;

		> div {
			margin-right: 10px; // 给每一个子元素  内间距
		}

		.action_group {
			margin-left: 10px;
			display: flex;
			align-items: center; // 居中
			button {
				margin-left: 10px;
			}
		}

		.action_filter {
			display: flex;

			> .arco-select {
				margin-right: 10px;
			}
		}

		.action_flush {
			position: absolute;
			margin-right: 0;
			right: 30px;

			button {
				padding: 0 10px;
			}
		}
	}

	.gvb_table_data {
		width: 100%;
		padding: 10px 20px 20px 20px;

		.gvb_cell_actions {
			> button {
				margin-right: 10px;
			}
		}

		.gvb_table_page {
			display: flex; // 容器
			justify-content: center; // 水平对齐
			margin-top: 20px;
		}
	}
}
</style>