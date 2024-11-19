import {defineStore} from 'pinia'
import {Message} from "@/BAIYIN/message/message";
import {HTTP} from "@/BAIYIN/HTTP";
import {JWT} from "@/BAIYIN/jwt";
import {TypeUser} from "@/api/user_api";
import router from "@/router";
import type {Themes} from "md-editor-v3";
import {ApiSetting, TypeSetting} from "@/api/setting_api";
import {nextTick} from "vue";

export interface userStoreInfoType {
    user_name: string
    nick_name: string
    role: string
    user_id: number
    avatar: string
    token: string
    exp: number     // 过期时间  需要x1000
}

const userInfo: userStoreInfoType = {
    user_name: "",
    nick_name: "",
    role: "",
    user_id: 0,
    avatar: "",
    token: "",
    exp: 0,
}

const theme: boolean = true  // 主题明暗颜色
const collapsed: boolean = false  // 主题明暗颜色
const siteInfo = TypeSetting.info()
export const useStore = defineStore('counter', {
    state() {
        return {
            theme: theme,
            collapsed: collapsed,   // 后台侧边栏收缩状态  默认展开
            userInfo: userInfo,
            siteInfo: siteInfo,
        }
    },
    actions: {
        // 黑白主题持久化
        setTheme() {
            this.theme = !this.theme
            document.documentElement.style.colorScheme = this.themeString
            document.body.setAttribute("arco-theme", this.themeString)
            localStorage.setItem("theme", JSON.stringify(this.theme))
        },
        loadThen() {
            let value = localStorage.getItem("theme")
            if (value !== null) {
                try {
                    this.theme = !JSON.parse(value)
                    this.setTheme()
                } catch (e) {
                    return;
                }
            }
        },
        setCollapsed(collapsed: boolean) {
            this.collapsed = collapsed
        },
        async setToken(token: string) {
            this.userInfo.token = token
            let info = JWT.Parse(token)
            Object.assign(this.userInfo, info)

            let res = await HTTP.Get<TypeUser.Info>("/api/user_info")
            let data = res.data
            this.userInfo = {
                user_name: data.user_name,
                nick_name: data.nick_name,
                role: data.role,
                user_id: data.id,
                avatar: data.avatar,
                token: token,
                exp: info.exp,
            }
            localStorage.setItem("userInfo", JSON.stringify(this.userInfo))

        },
        loadToken() {
            let value = localStorage.getItem("userInfo")
            if (value === null) {
                return
            }
            try {
                this.userInfo = JSON.parse(value)
            } catch (e) {
                console.log(e)
                return;
            }
            // 判断token是否过期
            let exp = Number(this.userInfo.exp) * 1000
            let nowTime = new Date().getTime()
            if (exp - nowTime <= 0) {
                Message.warning("登录已过期")
                router.push({name: "login"})
                localStorage.removeItem("userInfo")
                return;
            }
        },
        async logout() {
            Message.success("退出登录成功")
            await HTTP.Post("/api/user_logout")
            localStorage.removeItem("userInfo")
            // await router.push({name: "login"})
        },
        async loadSiteInfo() {
            const value = sessionStorage.getItem("siteInfo")
            if (value !== null) {
                try {
                    this.siteInfo = JSON.parse(value)
                    return
                } catch (e) {
                }
            }

            let result = await ApiSetting.List("site")
            this.siteInfo = result.data as TypeSetting.info

            sessionStorage.setItem("siteInfo", JSON.stringify(this.siteInfo))

        }
    },
    getters: {
        isLogin(): boolean {
            return this.userInfo.role !== ""
        },
        themeString(): Themes {
            return this.theme ? "light" : "dark"
        },
        LoginRole(): string {
            return this.userInfo.role
        },
    },
})
