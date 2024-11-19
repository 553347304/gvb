// npm i sass -D
// npm i axios
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import "@/assets/base.css"

import App from './App.vue'
import router from './router'

// npm install vant
// npm i babel-plugin-import -D
// import Vant from 'vant';
import 'vant/lib/index.css';

import 'font-awesome/css/font-awesome.min.css'  // npm i font-awesome --save

const app = createApp(App)

app.use(createPinia())
app.use(router)
// app.use(Vant);


app.mount('#app')



// npm i md-editor-v3