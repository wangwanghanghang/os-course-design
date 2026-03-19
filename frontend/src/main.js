import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus' // 引入 Element Plus
import 'element-plus/dist/index.css'   // 引入 Element Plus 样式
import { createPinia } from 'pinia'     // 引入 Pinia

const app = createApp(App)
const pinia = createPinia()

app.use(ElementPlus) // 使用 Element Plus
app.use(pinia)       // 使用 Pinia

app.mount('#app')