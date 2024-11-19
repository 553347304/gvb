export namespace DefinePromotion {
    export const Columns = [
        {title: '标题', dataIndex: 'title',},
        {title: '图片', slotName: 'images',},
        {title: '链接', slotName: 'href',},
        {title: '是否显示', slotName: 'is_show',},
        {title: '创建时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
}

export namespace TypePromotion {
    export interface info {
        id: number
        title: string
        images: string
        href: string
        is_show: boolean
        created_at: string
    }

    export interface Create {
        id: number

        title: string
        href: string
        images: string
        is_show: boolean
    }

    export const Create = (): Create => ({
        id: 0,

        title: "",
        href: "",
        images: "",
        is_show: false,
    });
}
