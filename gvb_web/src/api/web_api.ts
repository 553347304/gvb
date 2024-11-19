export namespace DefineWeb {
    export const sortOption = [
        {label: "综合排序", value: ""},
        {label: "最多浏览", value: "look_count desc"},
        {label: "最新发布", value: "date desc"},
        {label: "最多点赞", value: "digg_count desc"},
        {label: "最多评论", value: "comment_count desc"},
        {label: "最多收藏", value: "collects_count desc"},
    ]
    export const categoryOption = [
        {label: "全部分类", value: ""},
        {label: "前端", value: "前端"},
        {label: "后端", value: "后端"},
    ]
    export const tagOption = [
        {label: "全部标签", value: ""},
        {label: "golang", value: "golang"},
        {label: "python", value: "python"},
    ]
    export const bannerIndex = {
        slogan: "Koneko",
        abstract: ['silver hair golden pupils'],
        image: [
            '/image/bg/bg.jpg',
            '/image/bg/bg1.jpg',
            '/image/bg/bg2.jpg',
            '/image/bg/bg3.jpg',
        ],
    }
}
export namespace TypeWeb {
    export interface Search {
        sort: string | number
        category: string | number
        tags: string | number
    }

    export const Search = (): Search => ({
        sort: "",
        category: "",
        tags: "",
    })

    export interface Banner {
        slogan: string
        abstract: string[] | string
        image: string[]
    }
}

class _ApiWeb {
}

export const ApiWeb = new _ApiWeb()











