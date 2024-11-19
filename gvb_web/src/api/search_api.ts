import {Base} from "@/BAIYIN/HTTP";
import {Axios} from "@/api/api_index";
import {TypeApi} from "@/api/api";

export namespace DefineSearch {
}
export namespace TypeSearch {
    export interface Search {
        key: string
        title: string
        slug: string
        content: string
    }

    export const Search = (): Search => ({
        key: "",
        title: "",
        slug: "",
        content: "",
    })
}

class _ApiSearch {
    Search = (params?: TypeApi.PageInfo): Promise<Base.ResponseList<TypeSearch.Search>> => {
        return Axios.get("/api/search", {params})
    }
}

export const ApiSearch = new _ApiSearch()











