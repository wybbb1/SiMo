<template>
    <n-button text style="font-size: 24px;margin:5px" @click="router.push('/home')">
        <n-icon>
            <ArrowHookUpLeft16Filled/>
        </n-icon>
    </n-button>
    <div class="write-container">
        <n-button-group class="button-group">
            <n-button :type="isRichText ? 'primary' : 'default'" @click="() => {isRichText = true}">
                RichText
            </n-button>
            <n-button :type="!isRichText ? 'primary' : 'default'" @click="() => {isRichText = false}">
                Markdown
            </n-button>
        </n-button-group>
        <div class="editor-container">
            <div v-if="isRichText">
                <Toolbar :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode"
                    style="border-bottom: 1px solid #ccc" />
                <Editor :defaultConfig="editorConfig" :mode="mode" v-model="valueHtml"
                    style="height: 500px; overflow-y: hidden;" @onCreated="handleCreated"
                    @onDestroyed="handleDestroyed" />
            </div>
            <div v-else style="width: 90vw;">
                <v-md-editor v-model="text" height="550px"></v-md-editor>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, ref, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { NButton, NInput, NButtonGroup, NIcon} from 'naive-ui'
import { useRouter } from "vue-router"
import { ArrowHookUpLeft16Filled } from '@vicons/fluent'

const router = useRouter()
const isRichText = ref(true)

// richtext
const mode = 'default'
const editorRef = shallowRef()
const valueHtml = ref('')
const toolbarConfig = {}
const editorConfig = {
    placeholder: '请输入内容...'
}

// markdown
const text = ref('')

const handleCreated = (editor:any) => {
    editorRef.value = editor
}

const handleDestroyed = (editor: any) => {
    editor.distory()
    editorRef.value = null
}


</script>

<style scoped>
.write-container {
    display: flex;
    flex-direction: column;
    height: 90vh;
    align-items: center;
    justify-content: flex-start; /* 改为从顶部开始 */
    background-color: #f5f5f5;
    margin-left: 2%;
    margin-right: 2%;
    border-radius: 15px;
    padding-top: 20px; /* 添加顶部内边距 */
}

.button-group {
    margin: 10px 0;
    align-self: center; /* 按钮组居中对齐 */
    position: sticky; /* 固定位置 */
    top: 0; /* 固定在顶部 */
    background-color: #fbf4f4;
    padding: 10px;
    border-radius: 8px;
    z-index: 10; /* 确保在最上层 */
}

/* 编辑器容器占满剩余空间 */
.editor-container {
    flex: 1;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 20px;
    box-sizing: border-box;
}
</style>
