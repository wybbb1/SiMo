import { createApp } from 'vue'
import { createPinia } from 'pinia'
import '@/style.css'

import VueMarkdownEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import vuepressTheme from '@kangc/v-md-editor/lib/theme/vuepress.js';
import '@kangc/v-md-editor/lib/theme/style/vuepress.css';
import Prism from 'prismjs';

import App from './App.vue'
import router from './router'

const app = createApp(App)

VueMarkdownEditor.use(vuepressTheme, {
    Prism,
});

app.use(createPinia())
app.use(router)
app.use(VueMarkdownEditor);

app.mount('#app')
