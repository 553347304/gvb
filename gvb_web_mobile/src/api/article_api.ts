import {Base} from "@/BAIYIN/HTTP";
import {Axios} from "@/api/index";
import {TypeApi} from "@/api/api";

export namespace DefineArticle {
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
export namespace TypeArticle {
    export interface List {
        id: string
        user_id: number
        user_name: string
        user_head: string
        look_count: number
        comment_count: number
        digg_count: number
        collects_count: number
        banner_id: number
        banner_url: string
        title: string
        category: string
        keyword: string
        abstract: string
        content: string
        tags: string[]
        source: string
        link: string
        created_at: string
        updated_at: string
    }
    export const List = (): List => ({
        id:"",
        user_id:0,
        user_name:"",
        user_head:"",
        look_count:0,
        comment_count:0,
        digg_count:0,
        collects_count:0,
        banner_id:0,
        banner_url:"",
        title:"",
        category:"",
        keyword:"",
        abstract:"",
        content:"",
        tags:[],
        source:"",
        link:"",
        created_at:"",
        updated_at:"",
    })

    export interface Create {
        id: string
        banner_id: number
        banner_url: string
        title?: string
        abstract: string
        category?: string[]
        tags?: string[]
        content?: string
        source?: string
        link?: string
    }

    export const Create = (): Create => ({
        id: "",
        banner_id: 1,
        banner_url: "",
        title: "",
        abstract: "",
        category: [],
        tags: [],
        content: "",
        source: "",
        link: "",
    })

    export interface Calendar {
        date: string[]
        total: [string, number][]
    }
    export const Calendar = (): Calendar => ({
        date: [],
        total: [],
    })
}

class _ApiArticle {
    Calendar = (): Promise<Base.ResponseList<[string, number]>> => {
        return Axios.get( "/api/articles/calendar")
    }

}

export const ApiArticle = new _ApiArticle()
