import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './routers/index'
import pinia from './stores/index'
import ElementPlus from 'element-plus';
import "element-plus/dist/index.css";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import 'element-ui/lib/theme-chalk/display.css'


const app = createApp(App);
app.use(router);
app.use(pinia);
app.use(ElementPlus);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}
app.mount('#app'); 