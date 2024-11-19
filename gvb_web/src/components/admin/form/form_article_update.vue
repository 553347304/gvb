<template>
	<a-modal :title="props.title" :on-before-ok="ok" width="40%" modal-class="gvb_article_modal_body">
		<a-form ref="formRef" :model="form" :label-col-props="{span: 4}">
			<Unit_form label="文章标题" field="title"><a-input v-model="form.title"/></Unit_form>
			<Unit_form label="文章简介"><a-textarea v-model="form.abstract" :auto-size="{minRows: 4, maxRows: 5}"/></Unit_form>
			<Unit_form label="文章内容"><a-textarea v-model="form.content" :auto-size="{minRows: 4, maxRows: 5}"/></Unit_form>
			<Unit_form label="文章分类" field="category"><a-select v-model="form.category" :options="optionsCategory" multiple allow-create/></Unit_form>
			<Unit_form label="文章标签" field="tags"><a-select v-model="form.tags" :options="optionsTag" multiple allow-create/></Unit_form>
			<Unit_form label="文章封面" field="banner_id">
				<a-select v-model="form.banner_id">
					<a-option v-for="item in imageIdList" :value="item.id" :key="item.id">
						<div class="banner_image_div">
							<img :src="item.path" alt=""> {{ item.name }}
						</div>
					</a-option>
					<template #label="{ data }">
						<img :src="coverSrc(data.value).value" alt=""
							 style="height: 30px; border-radius: 5px" >
						<span style="margin-left: 10px">{{data.label}}</span>
					</template>
				</a-select>
			</Unit_form>
			<Unit_form label="文章来源"><a-input v-model="form.source"/></Unit_form>
			<Unit_form label="原文地址"><a-input v-model="form.link"/></Unit_form>
			<Unit_form label="预览"><form_article_item :form="form" preview/></Unit_form>
		</a-form>
	</a-modal>
</template>

<script setup lang="ts">
import {computed, reactive, ref, watch} from "vue";
import Unit_form from "@/components/unit/unit_form.vue";
import {Api, TypeApi} from "@/api/api";
import {TypeArticle} from "@/api/article_api";
import form_article_item from "@/components/admin/form/form_article_item.vue";
import {TypeImage} from "@/api/image_api";
import {BaseLog} from "@/BAIYIN/log";
import {DefineUtils} from "@/api/api_utils/type";

// 加载选择图片列表
const imageIdList = ref<TypeImage.Info[]>([])
const imageShowList = async () => imageIdList.value = (await Api.List(Api.URL.Image)).data.list;
imageShowList()

type Props = {form: TypeArticle.Create; title: string};
const props = defineProps<Props>()
const form = reactive({...props.form})
const formRef = ref()
const previewData = reactive(TypeArticle.Create())
watch(() => props.form, (data) =>  {
	Object.assign(form, data);
	Object.assign(previewData, data);
}, { deep: true });

const coverSrc = (value: number)=>{
	return computed(() :string => {
		let result = imageIdList.value.find((item)=> item.id === value)
		if (result === undefined) return ""
		return result.path
	})
}

const PageInfo = reactive(TypeApi.PageInfo())
const optionsTag = ref<DefineUtils.options[]>([])
const optionsCategory = ref<DefineUtils.options[]>([])

class Options {
	private Search = async (search: string) => {
		PageInfo.search = search
		let res = await Api.Search(Api.URL.Article, PageInfo)
		return res.data.list
	}
	Tag = async () => {
		optionsTag.value = await this.Search("tags")
	}
	Category = async () => {
		optionsCategory.value = await this.Search("category")
	}
}

const o = new Options()
o.Tag()
o.Category()

async function ok() {
	if (await formRef.value.validate()) return false
	let result = props.title === "发布文章" ?
		await Api.Create(Api.URL.Article, form) :
		await Api.Update(Api.URL.Article, form, form.id);
	return BaseLog(result)
}
</script>

<style lang="scss">
@import "@/global/mixin";

.banner_image_div {
	@include image_show_select(40px);
}

.promotion_list_view{
	max-width: inherit;
}

.gvb_article_modal_body .arco-modal-body {
	overflow-x: hidden;
 }

.arco-form-item-content-flex{
	display: block;
}
</style>