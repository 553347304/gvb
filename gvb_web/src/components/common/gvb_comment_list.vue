<template>
	<div class="gvb_comment_list_components_view">

		<div class="create">
			<a-textarea placeholder="评论内容"
						v-model="create.content"
						:auto-size="{minRows: 4, maxRows: 4}"
			></a-textarea>
			<a-button @click="c.Create" type="primary">发布评论</a-button>
		</div>

		<a-comment
			:author="store.userInfo.nick_name"
			:content="item.content"
			:datetime="Time.Time(item.created_at).Now"
			v-for="item in data.list"
			:avatar="store.userInfo.avatar"
		>
			<template #actions>
				<span class="action" @click="c.Remove(item)"><IconDelete/>删除</span>
			</template>
		</a-comment>
	</div>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import {Api, TypeApi} from "@/api/api";
import {Base} from "@/BAIYIN/HTTP";
import {TypeComment} from "@/api/comment_api";
import {Time} from "@/BAIYIN/time";
import {useStore} from "@/stores";
import {BaseLog} from "@/BAIYIN/log";
import {Message} from "@arco-design/web-vue";

const store = useStore()
type Props = { articleId: string; };
const props = defineProps<Props>()
const PageInfo = reactive(TypeApi.PageInfo())
const data = reactive(Base.List<TypeComment.List>())
const create = reactive(TypeComment.Create())


class Comment {
	List = async () => {
		PageInfo.search = "article_id"
		PageInfo.key = props.articleId
		let result = await Api.List(Api.URL.Comment, PageInfo)
		data.list = result.data.list
		data.total = result.data.total
	}
	Remove = async (record: TypeComment.List) => {
		let result = await Api.Remove(Api.URL.Comment, [record.id])
		BaseLog(result)
		await this.List()
	}
	Create = async () => {
		if (create.content.trim() === "") {
			Message.warning("内容不能为空")
			return
		}
		create.article_id = props.articleId
		let result = await Api.Create(Api.URL.Comment, create)
		if (!BaseLog(result)) return
		await this.List()
	}

}

const c = new Comment()

watch(() => props.articleId, () => {
	if (props.articleId !== "") c.List()
})


</script>

<style lang="scss">
.gvb_comment_list_components_view {
	background-color: var(--color-bg-1);
	border-radius: 5px;
	width: 100%;
	padding: 20px;

	.create {
		position: relative;

		button {
			position: absolute;
			right: 20px;
			bottom: 15px;
			z-index: 1;
		}
	}

	.action {
		> span {
			cursor: pointer;

		}
	}
}
</style>