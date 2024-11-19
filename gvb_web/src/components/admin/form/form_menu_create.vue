<template>
	<a-modal :title="props.title" :on-before-ok="ok">
		<a-form validate ref="formRef" :model="form">
			<Unit_form label="菜单标题" field="title"><a-input v-model="form.title"/></Unit_form>
			<Unit_form label="菜单路径" field="path"><a-input v-model="form.path"/></Unit_form>
			<Unit_form label="slogan" field="slogan"><a-input v-model="form.slogan"/></Unit_form>
			<Unit_form label="简介"><a-textarea v-model="form.abstract" :auto-size="{minRows:2, maxRows:2}"></a-textarea></Unit_form>
			<Unit_form label="简介切换时间"><a-input-number v-model="form.abstract_time"></a-input-number></Unit_form>
			<Unit_form label="切换时间"><a-input-number v-model="form.banner_time"></a-input-number></Unit_form>
			<Unit_form label="序号"><a-input-number v-model="form.sort"></a-input-number></Unit_form>
			<Unit_form label="banner图">
				<a-select v-model="form.banner" multiple>    <!--多选-->
					<a-option v-for="item in imageIdList" :value="item.id" :key="item.id">
						<div class="banner_image_div">
							<img :src="item.path" alt=""> {{ item.name }}
						</div>
					</a-option>
				</a-select>
			</Unit_form>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import {TypeMenu} from "@/api/menu_api";
import Unit_form from "@/components/unit/unit_form.vue";
import {TypeImage} from "@/api/image_api";
import {Api} from "@/api/api";
import {BaseLog} from "@/BAIYIN/log";

// 加载选择图片列表
const imageIdList = ref<TypeImage.Info[]>([])
const imageShowList = async () => imageIdList.value = (await Api.List(Api.URL.Image)).data.list;
imageShowList()

type Props = {form: TypeMenu.Create; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
watch(() => props.form, (data) => Object.assign(form, data), {deep: true});

async function ok(): Promise<boolean> {
	if (await formRef.value.validate()) return false
	let result = props.title === "创建菜单" ?
		await Api.Create(Api.URL.Menu, form) :
		await Api.Update(Api.URL.Menu, form, form.id);
	Object.assign(form, TypeMenu.Create())	// 重新赋值清空
	return BaseLog(result)
}
</script>

<style lang="scss">
@import "@/global/mixin";

.banner_image_div {
	@include image_show_select(40px);
}
</style>
