import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/', name: 'home', component: () => import('../views/index.vue')},
        {path: '/article/:id', name: 'article', component: () => import('../views/article.vue')},
        {
            path: '/my', name: 'my', component: () => import('../views/my/index.vue'), children: [
                {path: '', name: 'my_index', component: () => import('../views/my/my.vue'),},
                {path: 'coll', name: 'coll', component: () => import('../views/my/coll.vue'),}
            ]
        },
    ]
})

export default router
