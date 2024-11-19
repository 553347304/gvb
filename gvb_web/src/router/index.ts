import {createRouter, createWebHistory} from 'vue-router'
import {useStore} from "@/stores";
import {Message} from "@arco-design/web-vue";
import NProgress from 'nprogress';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/', name: 'web',
            component: () => import('../views/web/web.vue'),
            children: [
                {path: "", name: "index", component: () => import('../views/web/index.vue'),},
                {path: "chat", name: "chat", component: () => import('../views/web/chat.vue'),},
                {
                    path: "bigModel", name: "bigModel", meta: {title: "AI", Role: "普通用户"},
                    component: () => import('../views/web/big_model/index.vue'),
                    redirect: "/bigModel/square",
                    children: [
                        {path: "square", name: "role_square", component: () => import('../views/web/big_model/role_square.vue'),},
                        {path: "session", name: "role_session", component: () => import('../views/web/big_model/role_session.vue'),},
                    ]
                },
                {path: "search", name: "search", component: () => import('../views/web/search.vue'),},
                {path: "news", name: "news", component: () => import('../views/web/news.vue'),},
                {path: "article/:id", name: "article", component: () => import('../views/web/article.vue'),},
            ]
        },
        {
            path: "/login", name: "login",
            component: () => import('../views/login/index.vue'),
        },
        {
            path: "/admin", name: "admin", meta: {title: "首页"},
            component: () => import("../views/admin/index.vue"),
            children: [
                {path: "", name: "admin_index", meta: {title: "主页", siLogin: true}, redirect: "/admin",},
                {path: "home", name: "home", component: () => import("../views/admin/home/home.vue")},
                {
                    path: "article", name: "article_mgr", meta: {title: "文章管理", Role: "普通用户"},
                    children: [
                        {
                            path: "article_list", name: "article_list", meta: {title: "文章列表"},
                            component: () => import("../views/admin/article/article_list.vue")
                        },
                        {
                            path: "image_list", name: "image_list", meta: {title: "图片列表"},
                            component: () => import("../views/admin/article/image_list.vue")
                        },
                        {
                            path: "comment_list", name: "comment_list", meta: {title: "评论列表"},
                            component: () => import("../views/admin/article/comment_list.vue")
                        },
                    ],
                },
                {
                    path: "chat_group", name: "chat_group", meta: {title: "群聊管理", Role: "普通用户"},
                    children: [
                        {
                            path: "chat_list", name: "chat_list", meta: {title: "群聊列表"},
                            component: () => import("../views/admin/chat_group/chat_list.vue")
                        },
                    ],
                },
                {
                    path: "system", name: "system", meta: {title: "系统管理", Role: "管理员"},
                    children: [
                        {
                            path: "system_list", name: "system_list", meta: {title: "菜单列表"},
                            component: () => import("@/views/admin/system/menu_list.vue")
                        },
                        {
                            path: "promotion_list", name: "promotion_list", meta: {title: "广告列表"},
                            component: () => import("@/views/admin/system/promotion_list.vue")
                        },
                        {
                            path: "config", name: "config", meta: {title: "系统配置"},
                            component: () => import("@/views/admin/system/config/index.vue"),
                            redirect: "/admin/system/config/site",
                            children: [
                                {
                                    path: "site", name: "site_config", meta: {title: "网站配置"},
                                    component: () => import("@/views/admin/system/config/site_config.vue"),
                                },
                                {
                                    path: "email", name: "email_config", meta: {title: "邮箱配置"},
                                    component: () => import("@/views/admin/system/config/email_config.vue"),
                                },
                            ],
                        },
                    ],
                },
                {
                    path: "user_center", name: "user_center", meta: {title: "个人中心"},
                    children: [
                        {
                            path: "user_message", name: "user_message", meta: {title: "我的消息"},
                            component: () => import("../views/admin/user_center/user_message_list.vue")
                        },
                    ],
                },
                {
                    path: "users", name: "users", meta: {title: "用户管理", Role: "管理员"},
                    children: [
                        {
                            path: "user_list", name: "user_list", meta: {title: "用户列表"},
                            component: () => import("@/views/admin/users/user_list.vue")
                        },
                        {
                            path: "message_list", name: "message_list", meta: {title: "消息列表"},
                            component: () => import("@/views/admin/users/message_list.vue")
                        },
                    ],
                },
                {
                    path: "big_model", name: "big_model", meta: {title: "大模型管理", Role: "管理员"},
                    children: [
                        {
                            path: "options", name: "options", meta: {title: "参数配置"},
                            component: () => import("@/views/admin/big_model/options/index.vue"),
                            redirect: "/admin/big_model/options/base",
                            children: [
                                {
                                    path: "base", name: "base_option", meta: {title: "基本配置"},
                                    component: () => import("@/views/admin/big_model/options/base_option.vue")
                                },
                                {
                                    path: "session", name: "session_option", meta: {title: "会话配置"},
                                    component: () => import("@/views/admin/big_model/options/session_option.vue")
                                },
                                {
                                    path: "reply", name: "reply_option", meta: {title: "自动回复"},
                                    component: () => import("@/views/admin/big_model/options/reply_option.vue")
                                },
                            ]
                        },
                        {
                            path: "chat_role", name: "chat_role", meta: {title: "角色配置"},
                            component: () => import("@/views/admin/big_model/chat_role/index.vue"),
                            redirect: "/admin/big_model/chat_role/tag",
                            children: [
                                {
                                    path: "tag", name: "role_tag", meta: {title: "角色标签"},
                                    component: () => import("@/views/admin/big_model/chat_role/role_tag.vue")
                                },
                                {
                                    path: "list", name: "role_list", meta: {title: "角色列表"},
                                    component: () => import("@/views/admin/big_model/chat_role/role_list.vue")
                                },
                            ]
                        },
                        {
                            path: "session", name: "session", meta: {title: "会话管理"},
                            component: () => import("@/views/admin/big_model/session/index.vue")
                        },
                    ],
                },
            ]

        }]
})

export default router

// 前置导航守卫
router.beforeEach((to, from, next) => {
    const store = useStore()
    const isRole = to.meta.Role
    const Role: string = store.LoginRole
    if (isRole !== undefined) {
        if (isRole === "普通用户" && Role === "") {
            Message.warning("未登录账号")
            router.push({name: from.name as string})
            return
        }
        if (isRole === "管理员" && Role !== "管理员") {
            Message.warning("管理员权限不足") // 管理员
            router.push({name: from.name as string})
            return
        }
    }
    NProgress.start(); // 开始进度条
    next()
})

router.afterEach(() => {
    NProgress.done(); // 完成进度条
});

