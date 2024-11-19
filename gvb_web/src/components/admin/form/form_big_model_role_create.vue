<template>
	<a-modal :title="props.title" :on-before-ok="ok" width="40%">
		<a-form validate ref="formRef" :model="form" :label-col-props="{span: 6}" :wrapper-col-props="{span:18}">
			<div class="role_add_modal">
				<div class="left">
					<Unit_form label="角色名称" field="name">
						<a-input v-model="form.name"/>
					</Unit_form>
					<Unit_form label="角色图标" field="icon">
						<a-select v-model="form.icon">
							<a-option v-for="item in iconOptions" :value="item.value" :key="item.value">
								<div style="display: flex; align-items: center;">
									<img :src="item.value as string" alt="" height="30px"> {{ item.label }}
								</div>
							</a-option>
						</a-select>
					</Unit_form>
					<Unit_form label="角色标签" field="tag_list">
						<a-select v-model="form.tag_list" :options="tagOptions" multiple allow-create/>
					</Unit_form>
					<Unit_form label="角色简介" field="abstract">
						<a-textarea v-model="form.abstract" :auto-size="{minRows: 1, maxRows: 5}"/>
					</Unit_form>
					<Unit_form label="消耗积分" field="scope">
						<a-input-number v-model="form.scope"/>
					</Unit_form>
					<Unit_form label="开场白" field="prologue">
						<a-textarea v-model="form.prologue" :auto-size="{minRows: 1, maxRows: 5}"/>
					</Unit_form>
					<Unit_form label="角色设定" field="prompt">
						<a-textarea v-model="form.prompt" :auto-size="{minRows: 1, maxRows: 5}"/>
					</Unit_form>
					<Unit_form label="自动回复" field="auto_reply">
						<a-switch v-model="form.auto_reply"/>
					</Unit_form>
					<Unit_form label="是否启用" field="enable">
						<a-switch v-model="form.enable"/>
					</Unit_form>
				</div>

				<div class="right">
					<gvb_big_model_role_preview :data="form"/>
				</div>
			</div>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {BaseLog} from "@/BAIYIN/log";
import {ApiBigModel, TypeBigModel} from "@/api/big_model_api";
import {DefineUtils} from "@/api/api_utils/type";
import Gvb_big_model_role_preview from "@/components/admin/gvb_big_model_role_preview.vue";

type Props = { form: TypeBigModel.RoleCreate; title: string };
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), {deep: true});


const tagOptions = ref<DefineUtils.options[]>([])
const iconOptions = ref<DefineUtils.options[]>([])

async function getIcon() {
	let result = await ApiBigModel.RoleIcon();
	iconOptions.value = result.data.list
}

async function getTagList() {
	let result = await ApiBigModel.TagList();
	tagOptions.value = result.data.list
}

getIcon()
getTagList()

async function createTag() {
	let is = false
	let tag_id: number[] = []
	for (let item of form.tag_list) {
		const s = item as string | number
		if (typeof s === "string") {
			let res = await ApiBigModel.TagCreate({
				title: s,
				color: "#ff0000",
			});
			tag_id.push(res.data);
			is = true;
		} else {
			tag_id.push(item);
		}
	}
	form.tag_list = tag_id
	if (is) await getTagList();
}

async function ok(): Promise<boolean> {
	if (await formRef.value.validate()) return false;
	await createTag()
	let result = form.id === undefined || form.id === 0 ?
		await ApiBigModel.RoleCreate(form) :
		await ApiBigModel.RoleUpdate(form);
	Object.assign(form, TypeBigModel.RoleCreate())	// 重新赋值清空
	return BaseLog(result)
}
</script>

<style scoped lang="scss">
.role_add_modal {
	display: flex;

	.left {
		width: 50%;
	}

	.right {
		width: 50%;
		padding-left: 20px;
	}
}
</style>
