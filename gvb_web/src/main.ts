import {createApp} from 'vue'
import {createPinia} from 'pinia'

import '@/assets/base.css';
import '@/assets/theme.css';

import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';

import App from './App.vue'
import router from './router'

import '@arco-design/web-vue/dist/arco.css';
import 'font-awesome/css/font-awesome.min.css'  // npm i font-awesome --save

import "md-editor-v3/lib/style.css"

// npm install nprogress    // 进度条
// npm install --save-dev @types/nprogress
import 'nprogress/nprogress.css'; // 引入 nprogress 进度条样式

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ArcoVue);
app.use(ArcoVueIcon);

app.mount('#app')

// document.documentElement.clientWidth    // 浏览器宽

