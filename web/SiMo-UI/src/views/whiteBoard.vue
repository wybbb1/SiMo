<template>
    <div class="page-container">
        <!-- 顶部导航栏 (包含返回按钮和时钟) -->
        <div class="top-nav">
            <n-button text @click="router.push('/home')" class="back-button">
                <n-icon size="20px">
                    <ArrowHookUpLeft16Filled />
                </n-icon>
                <span class="back-text">返回</span>
            </n-button>

            <TimeComponent :showAnalog="false" :showSeconds="true" :compact="true" class="header-clock" />

            <!-- 占位元素，保持时钟居中 -->
            <div class="placeholder"></div>
        </div>

        <!-- 白板区域 -->
        <div class="whiteboard-page">
            <WhiteBoardComponent />

            <!-- 音乐组件 - 左下角 -->
            <div class="music-widget">
                <MusicComponent />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import WhiteBoardComponent from '@/components/whiteBoard.tsx'
import TimeComponent from '@/components/time.vue'
import MusicComponent from '@/components/music.vue'
import { useRouter } from "vue-router"
import { NButton, NIcon } from 'naive-ui'
import { ArrowHookUpLeft16Filled } from '@vicons/fluent'

const router = useRouter()
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

/* 占位元素，保持时钟居中 */
.placeholder {
    width: 80px;
    /* 与返回按钮大致相同的宽度 */
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

/* 页面标题 */
.page-title {
    margin: 0;
    font-size: 18px;
    font-weight: 600;
    color: #333;
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
}

/* 白板容器 */
.whiteboard-page {
    flex: 1;
    /* 占据剩余空间 */
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 10px 60px;
    /* 四周留白 */
    box-sizing: border-box;
    position: relative;
    /* 允许子元素绝对定位 */
}

/* 音乐组件定位 */
.music-widget {
    position: absolute;
    bottom: 20px;
    left: 10px;
    z-index: 100;
    /* 确保在白板之上 */
}

/* 白板组件样式 */
.whiteboard-page> :not(.music-widget) {
    width: 100%;
    height: 100%;
    max-width: 1400px;
    /* 最大宽度限制 */
    border-radius: 16px;
    /* 更大的圆角 */
    overflow: hidden;
}

/* 响应式设计 */
@media (max-width: 768px) {
    .top-nav {
        padding: 12px 16px;
    }

    .page-title {
        font-size: 16px;
    }

    .back-text {
        display: none;
        /* 小屏幕隐藏文字 */
    }

    .whiteboard-page {
        padding: 12px;
    }

    .music-widget {
        bottom: 12px;
        left: 12px;
        /* 移动端时调整位置 */
    }

    .whiteboard-page> :not(.music-widget) {
        border-radius: 12px;
    }
}
</style>