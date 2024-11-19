export namespace DefineMenu {
    export const Columns = [
        {title: '序号', dataIndex: 'sort',},
        {title: '菜单标题', dataIndex: 'title',},
        {title: '路径', dataIndex: 'path',},
        {title: 'slogan', dataIndex: 'slogan',},
        {title: '简介', dataIndex: 'abstract',},
        {title: '简介切换时间', dataIndex: 'abstract_time',},
        {title: 'slogan切换时间', dataIndex: 'banner_time',},
        {title: 'banners', slotName: 'banners',},
        {title: '更新时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
}
export namespace TypeMenu {
    export interface imageList {
        id: number
        path: string
    }

    export interface info {
        title: string
        path: string
        slogan: string
        abstract: string[]
        abstract_time: number
        banner_time: number
        banners: imageList[]
        created_at: string
        sort: number
    }

    export interface Create {
        id: number

        title: string
        path: string
        slogan: string
        abstract: string
        abstract_time: number
        banner: number[]
        banner_time: number
        sort: number
        created_at: string
    }


    export const Create = (): Create => ({
        id: 0,

        title: "",
        path: "",
        slogan: "",
        abstract: "",
        abstract_time: 7,
        banner: [],
        banner_time: 7,
        sort: 1,
        created_at: "",

    })
}










