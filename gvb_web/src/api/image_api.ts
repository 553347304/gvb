import {Tag} from "ant-design-vue";
import {h} from "vue";
import {Base} from "@/BAIYIN/HTTP";
import {Axios} from "@/api/api_index";
import {Api} from "@/api/api";

export namespace DefineImage {
    export const Columns = [
        {title: '文件名', dataIndex: 'name'},
        {title: '图片', slotName: 'path',},
        {
            title: '类型', dataIndex: 'image_type', render: (data: any) => {
                const record = data.record as TypeImage.Info
                let color: string = "red"
                if (record.image_type === "本地") color = "blue"
                return h(Tag, {color: color}, {default: () => record.image_type})
            }
        },
        {title: '上传时间', slotName: 'created_at',},
        {title: '操作', slotName: 'action',},
    ];
}
export namespace TypeImage {
    export interface Info {
        id: number
        created_at: string
        path: string
        hash: string
        name: string
        image_type: "本地" | "七牛云"
    }


    export const Info = (): Info => ({
        id: 0,
        created_at: "",
        path: "",
        hash: "",
        name: "",
        image_type: "本地",
    })

    export interface Create {
        name: string
        path: string
        size: string
        is_success: boolean
        message: string
    }

    export const Create = (): Create => ({
        name: "",
        path: "",
        size: "",
        is_success: false,
        message: "",
    })
}

class _ApiImage {

    Upload(file: File): Promise<Base.ResponseList<TypeImage.Create>> {
        return new Promise((resolve, reject) => {
            const form = new FormData()
            form.set("image", file)
            Axios.post(Api.URL.Image + "_upload", form, {
                headers: {"Content-Type": "multipart/form-data"}
            }).then((result: any) => resolve(result)).catch(error => reject(error))
        })
    }

    OnUploadImage = async (files: Array<File>, callback: (urls: Array<string>) => void): Promise<void> => {
        const resultList = await Promise.all(
            files.map((file) => {
                return this.Upload(file)
            })
        );
        const pathList: string[] = []
        resultList.forEach(r => {
            pathList.push(r.data.list[0].path)
        })
        callback(pathList)
    };
}

export const ApiImage = new _ApiImage()

