export namespace DefineComment {
    export const Columns = [
        {title: '标题', slotName: 'title',},
        {title: '分类', dataIndex: 'category',},
        {title: '封面', slotName: 'banner_url',},
        {title: '作者', dataIndex: 'user_name',},
        {title: '标签', slotName: 'tags',},
        {title: '文章数据', slotName: 'data',},
        {title: '创建时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
    const category = [
        {label: "前端", value: "前端"},
        {label: "后端", value: "后端"},
    ];
    const tags = [
        {label: "golang", value: "golang"},
        {label: "python", value: "python"},
    ];
    export const Filter = [
        {label: "文章分类", value: 0, search: "category", option: category},
        {label: "文章标签", value: 0, search: "tags", option: tags},
    ];
}
export namespace TypeComment {
    export interface List {
        id: number
        parent_id: number
        article_id: string
        user: number
        content: string
        created_at: string
    }


    export interface Create {
        parent_id: number
        article_id: string
        content: string
    }

    export const Create = (): Create => ({
        parent_id: 0,
        article_id: "",
        content: "",
    })


}
