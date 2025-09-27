<template>
    <div class="page-container">
        <!-- 顶部导航栏 (包含返回按钮、时钟和功能按钮) -->
        <div class="top-nav">
            <n-button text @click="router.push('/home')" class="back-button">
                <n-icon size="20px">
                    <ArrowHookUpLeft16Filled />
                </n-icon>
                <span class="back-text">返回</span>
            </n-button>

            <TimeComponent :showAnalog="false" :showSeconds="true" :compact="true" class="header-clock" />

            <n-button-group class="mode-toggle">
                <n-button quaternary :color="isRichText ? '#000000' : '#686868'" @click="() => { isRichText = true }">
                    富文本
                </n-button>
                <n-button quaternary :color="isRichText ? '#686868' : '#000000'" @click="() => { isRichText = false }">
                    Markdown
                </n-button>
            </n-button-group>
        </div>

        <!-- 写作区域 -->
        <div class="write-page">
            <div class="editor-container">
                <div v-if="isRichText">
                    <Toolbar :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode"
                        style="border-bottom: 1px solid #ccc" />
                    <Editor :defaultConfig="editorConfig" :mode="mode" v-model="valueHtml"
                        style="height: 500px; overflow-y: hidden;" @onCreated="handleCreated"
                        @onDestroyed="handleDestroyed" />
                </div>
                <div v-else style="width: 100%;">
                    <v-md-editor v-model="text" height="550px"></v-md-editor>
                </div>
            </div>

            <!-- 音乐组件 - 左下角 -->
            <div class="music-widget">
                <MusicComponent />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, ref, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { NButton, NInput, NButtonGroup, NIcon } from 'naive-ui'
import { useRouter } from "vue-router"
import { ArrowHookUpLeft16Filled } from '@vicons/fluent'
import TimeComponent from '@/components/time.vue'
import MusicComponent from '@/components/music.vue'

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

const handleCreated = (editor: any) => {
    editorRef.value = editor
}

const handleDestroyed = (editor: any) => {
    editor.distory()
    editorRef.value = null
}


</script>

<style scoped>
/* 整体页面容器 */
.page-container {
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-direction: column;
}

/* 顶部导航栏 */
.top-nav {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 12px 24px;
    z-index: 10;
}

/* 返回按钮样式 */
.back-button {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 8px 12px !important;
    border-radius: 8px;
    color: #666 !important;
    transition: all 0.2s ease;
}

.back-button:hover {
    color: #333 !important;
}

.back-text {
    font-size: 14px;
    font-weight: 500;
}

/* 模式切换按钮组 */
.mode-toggle {
    display: flex;
    gap: 0;
}

/* 写作区域容器 */
.write-page {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: flex-start;
    padding: 20px 60px;
    box-sizing: border-box;
    position: relative;
    overflow: hidden;
    z-index: 100;
}

/* 编辑器容器 */
.editor-container {
    width: 100%;
    max-width: 1400px;
    height: 100%;
    display: flex;
    flex-direction: column;
    border-radius: 16px;
    overflow: hidden;
    box-shadow:
        0 8px 32px rgba(0, 0, 0, 0.08),
        0 2px 8px rgba(0, 0, 0, 0.04);
    border: 1px solid rgba(0, 0, 0, 0.05);
    background: white;
}

/* 音乐组件定位 */
.music-widget {
    position: absolute;
    bottom: 20px;
    left: 10px;
    z-index: 100;
}

/* 富文本编辑器样式调整 */
.editor-container :deep(.w-e-toolbar) {
    border-top-left-radius: 16px;
    border-top-right-radius: 16px;
    border-bottom: 1px solid #e8e8e8;
}

.editor-container :deep(.w-e-text-container) {
    border-bottom-left-radius: 16px;
    border-bottom-right-radius: 16px;
}

/* Markdown 编辑器样式调整 */
.editor-container :deep(.v-md-editor) {
    border-radius: 16px;
    border: none;
    box-shadow: none;
}

/* 响应式设计 */
@media (max-width: 768px) {
    .top-nav {
        padding: 12px 16px;
    }

    .back-text {
        display: none;
    }

    .write-page {
        padding: 12px;
    }

    .music-widget {
        bottom: 12px;
        left: 12px;
    }

    .editor-container {
        border-radius: 12px;
    }

    .mode-toggle {
        transform: scale(0.9);
    }
}

/* 模式切换按钮的特殊样式 */
.mode-toggle .n-button {
    font-size: 12px;
    padding: 6px 16px;
}
</style>
