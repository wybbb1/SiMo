<template>
    <div class="music-container">
        <!--mini 模式 只有一个音符-->
        <Transition name="mini-mode" mode="out-in">
            <div v-if="isMini" key="mini" class="mini-container">
                <button class="circle-button" @click="toggleMini">
                    <Icon>
                        <MusicNote216Filled />
                    </Icon>
                </button>
            </div>
            <!--full 模式 显示完整内容-->
            <div v-else key="full" class="full-container">
                <div class="music-player-glass">
                    <!-- 按钮组 -->
                    <div class="button-group" style="--delay: 1">
                        <!-- 切换播放列表按钮（右上角） -->
                        <button class="menu-btn" @click="togglePlaylist">
                            <Icon>
                                <AppsListDetail24Regular />
                            </Icon>
                        </button>
                        <!-- 切换迷你模式按钮（左上角） -->
                        <button class="menu-btn" @click="toggleMini">
                            <Icon>
                                <ArrowMoveInward20Filled />
                            </Icon>
                        </button>
                    </div>
                    <div
                        style="display:flex; flex-direction:row; gap: 20px; justify-content: space-between; --delay: 2">
                        <!-- 专辑封面 - 左侧-->
                        <div class="album-cover">
                            <img :src="currentSong.cover || '/src/assets/placeholder-cover.jpg'" alt="Album Cover" />
                        </div>
                        <!-- 歌曲信息 - 右侧-->
                        <div style="display: flex; flex-direction: column; gap: 8px;">
                            <div class="song-info">
                                <h3 class="song-title">{{ currentSong.title || '未知标题' }}</h3>
                                <p class="artist">{{ currentSong.artist || '未知艺术家' }}</p>
                            </div>
                            <!-- 播放控制按钮 -->
                            <div class="controls">
                                <button class="control-btn" @click="previousSong">
                                    <Icon>
                                        <ArrowPrevious24Filled />
                                    </Icon>
                                </button>
                                <button class="play-btn" @click="togglePlay">
                                    <Icon v-if="!isPlaying">
                                        <Play48Filled />
                                    </Icon>
                                    <Icon v-else>
                                        <Pause48Regular />
                                    </Icon>
                                </button>
                                <button class="control-btn" @click="nextSong">
                                    <Icon>
                                        <ArrowNext24Filled />
                                    </Icon>
                                </button>
                            </div>
                        </div>
                    </div>

                    <!-- 进度条 -->
                    <div class="progress-container" style="--delay: 3">
                        <span class="time current-time">{{ formatTime(currentTime) }}</span>
                        <div class="progress-bar" @click="handleProgressClick">
                            <div class="progress-track"></div>
                            <div class="progress-thumb" :style="{ left: progressPercentage + '%' }"></div>
                        </div>
                        <span class="time total-time">{{ formatTime(duration) }}</span>
                    </div>

                    <!-- 播放列表弹窗（简化版） -->
                    <div v-if="showPlaylist" class="playlist-overlay">
                        <div class="playlist-content">
                            <div class="playlist-header">
                                <h4>播放列表</h4>
                                <button @click="togglePlaylist" class="close-btn">
                                    <Icon>
                                        <ShareScreenStop16Filled />
                                    </Icon>
                                </button>
                            </div>
                            <ul class="playlist-items">
                                <li v-for="(song, index) in songs" :key="song.id"
                                    :class="{ active: song.id === currentSong.id }" @click="selectSong(index)">
                                    {{ song.title }} - {{ song.artist }}
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </Transition>

        <!-- 隐藏的 audio 元素 - 放在外面避免被销毁 -->
        <audio ref="audioRef" :src="currentSong.url" @loadedmetadata="onLoadedMetadata" @timeupdate="onTimeUpdate"
            @ended="onEnded" @play="isPlaying = true" @pause="isPlaying = false"></audio>
    </div>

</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue';
import { ArrowPrevious24Filled, ArrowNext24Filled, Play48Filled, Pause48Regular, ShareScreenStop16Filled, MusicNote216Filled, AppsListDetail24Regular, ArrowMoveInward20Filled } from '@vicons/fluent';
import { Icon } from '@vicons/utils'


// 歌曲接口
interface Song {
    id: string;
    title: string;
    artist: string;
    cover?: string;
    url: string;
}

// 示例歌曲数据（可替换为真实API）
const songs: Song[] = [
    {
        id: '1',
        title: '大黑天',
        artist: '杨秉音',
        cover: '/public/cover/songzan.jpg',
        url: '/public/music/大黑天.mp3' // 确保路径正确，或使用在线资源
    },
    {
        id: '2',
        title: '颂赞',
        artist: '吴莫愁',
        cover: '/public/cover/songzan.jpg',
        url: '/public/music/颂赞.mp3'
    },
    {
        id: '3',
        title: '天后',
        artist: '李佳薇',
        cover: '/public/cover/tianhou.jpg',
        url: '/public/music/天后.mp3'
    }
];

// 响应式数据
const currentSong = ref<Song>(songs[0]);
const isPlaying = ref(false);
const currentTime = ref(0);
const duration = ref(0);
const showPlaylist = ref(false);
const isMini = ref(true); // 切换迷你模式

const toggleMini = () => {
    isMini.value = !isMini.value;
};

// 音频元素引用
const audioRef = ref<HTMLAudioElement | null>(null);

// 计算属性
const progressPercentage = computed(() => {
    if (duration.value === 0) return 0;
    return (currentTime.value / duration.value) * 100;
});

// 格式化时间 MM:SS
const formatTime = (time: number): string => {
    const minutes = Math.floor(time / 60);
    const seconds = Math.floor(time % 60);
    return `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
};

// 控制函数
const togglePlay = () => {
    if (!audioRef.value) return;
    if (isPlaying.value) {
        audioRef.value.pause();
    } else {
        audioRef.value.play().catch(err => {
            console.warn('播放失败，可能被浏览器阻止自动播放', err);
        });
    }
};

const nextSong = () => {
    const currentIndex = songs.findIndex(song => song.id === currentSong.value.id);
    const nextIndex = (currentIndex + 1) % songs.length;
    currentSong.value = songs[nextIndex];
    // 重置播放状态
    isPlaying.value = false;
    // 等待 DOM 更新后播放
    nextTick(() => {
        if (audioRef.value) {
            audioRef.value.play().catch(err => {
                console.warn('自动播放被阻止，请用户交互后播放');
            });
        }
    });
};

const previousSong = () => {
    const currentIndex = songs.findIndex(song => song.id === currentSong.value.id);
    const prevIndex = (currentIndex - 1 + songs.length) % songs.length;
    currentSong.value = songs[prevIndex];
    nextTick(() => {
        if (audioRef.value) {
            audioRef.value.play().catch(err => {
                console.warn('自动播放被阻止');
            });
        }
    });
};

const selectSong = (index: number) => {
    currentSong.value = songs[index];
    togglePlaylist(); // 关闭列表
    nextTick(() => {
        if (audioRef.value) {
            audioRef.value.play().catch(err => {
                console.warn('自动播放被阻止');
            });
        }
    });
};

const togglePlaylist = () => {
    showPlaylist.value = !showPlaylist.value;
};

const handleProgressClick = (e: MouseEvent) => {
    if (!audioRef.value) return;
    const progressBar = e.currentTarget as HTMLElement;
    const clickPosition = e.offsetX;
    const progressBarWidth = progressBar.clientWidth;
    const percentage = clickPosition / progressBarWidth;
    audioRef.value.currentTime = percentage * duration.value;
};

// 音频事件处理
const onLoadedMetadata = () => {
    if (audioRef.value) {
        duration.value = audioRef.value.duration;
    }
};

const onTimeUpdate = () => {
    if (audioRef.value) {
        currentTime.value = audioRef.value.currentTime;
    }
};

const onEnded = () => {
    nextSong(); // 自动播放下一首
};

// 在组件挂载后，尝试播放第一首歌（需用户交互触发首次播放）
onMounted(() => {
    // 注意：浏览器会阻止自动播放，除非有用户交互（如点击播放按钮）
    // 所以首次播放必须由用户点击触发
});
</script>

<style scoped>
/* 容器样式 */
.music-container {
    position: relative;
    display: inline-block;
}

.mini-container,
.full-container {
    display: flex;
    align-items: center;
    justify-content: center;
}

/* 过渡动画 */
.mini-mode-enter-active,
.mini-mode-leave-active {
    transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.mini-mode-enter-from {
    opacity: 0;
    transform: scale(0.8) rotateY(90deg);
}

.mini-mode-enter-to {
    opacity: 1;
    transform: scale(1) rotateY(0deg);
}

.mini-mode-leave-from {
    opacity: 1;
    transform: scale(1) rotateY(0deg);
}

.mini-mode-leave-to {
    opacity: 0;
    transform: scale(1.2) rotateY(-90deg);
}

/* Mini模式圆形按钮样式 */
.circle-button {
    width: 60px;
    height: 60px;
    border-radius: 50%;
    background: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(10px);
    border: none;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 1.8rem;
    color: #4CAF50;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
    animation: miniModeEnter 0.6s cubic-bezier(0.34, 1.56, 0.64, 1);
}

/* Mini模式进入动画 */
@keyframes miniModeEnter {
    0% {
        transform: scale(0.3);
        opacity: 0;
        box-shadow: 0 0 0 rgba(76, 175, 80, 0.4);
    }

    50% {
        transform: scale(1.1);
        opacity: 0.8;
        box-shadow: 0 0 20px rgba(76, 175, 80, 0.4);
    }

    70% {
        transform: scale(0.95);
        opacity: 1;
    }

    100% {
        transform: scale(1);
        opacity: 1;
        box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
    }
}

.circle-button::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: linear-gradient(135deg, rgba(76, 175, 80, 0.1), rgba(139, 195, 74, 0.1));
    border-radius: 50%;
    opacity: 0;
    transition: opacity 0.3s ease;
}

.circle-button:hover {
    transform: scale(1.1);
    box-shadow: 0 6px 20px rgba(76, 175, 80, 0.3);
    color: #45a049;
}

.circle-button:hover::before {
    opacity: 1;
}

.circle-button:active {
    transform: scale(0.95);
    animation: ripple 0.6s cubic-bezier(0, 0, 0.2, 1);
}

/* 点击涟漪效果 */
@keyframes ripple {
    0% {
        box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15), 0 0 0 0 rgba(76, 175, 80, 0.7);
    }

    70% {
        box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15), 0 0 0 10px rgba(76, 175, 80, 0);
    }

    100% {
        box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15), 0 0 0 0 rgba(76, 175, 80, 0);
    }
}

.music-player-glass {
    width: 380px;
    padding: 10px;
    background: rgba(255, 255, 255, 0.7);
    backdrop-filter: blur(10px);
    border-radius: 16px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    position: relative;
    overflow: hidden;
    transform-origin: center center;
    animation: fullModeEnter 0.6s cubic-bezier(0.4, 0, 0.2, 1) forwards;
}

/* Full模式进入动画 */
@keyframes fullModeEnter {
    0% {
        opacity: 0;
        transform: scale(0.3) rotateX(20deg);
        filter: blur(8px);
    }

    50% {
        opacity: 0.8;
        transform: scale(0.9) rotateX(10deg);
        filter: blur(2px);
    }

    100% {
        opacity: 1;
        transform: scale(1) rotateX(0deg);
        filter: blur(0px);
    }
}

/* 内部元素渐进显示动画 */
.music-player-glass>* {
    animation: slideInFadeIn 0.8s cubic-bezier(0.4, 0, 0.2, 1) forwards;
    animation-delay: calc(var(--delay, 0) * 0.1s);
}

@keyframes slideInFadeIn {
    0% {
        opacity: 0;
        transform: translateY(20px);
    }

    100% {
        opacity: 1;
        transform: translateY(0);
    }
}

.button-group {
    position: absolute;
    top: 10px;
    right: 10px;
    display: flex;
    flex-direction: column;
    z-index: 10;
}

.menu-btn {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1.2rem;
    color: #666;
    padding: 5px;
    border-radius: 50%;
    transition: transform 0.2s ease, background-color 0.2s;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 36px;
    height: 36px;
}

.menu-btn:hover {
    transform: scale(1.1);
    color: #4CAF50;
    background-color: rgba(76, 175, 80, 0.1);
}

.album-cover {
    width: 100px;
    height: 100px;
    border-radius: 8px;
    overflow: hidden;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.album-cover img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.song-info {
    text-align: center;
    margin: 0;
}

.song-title {
    font-size: 1.1rem;
    font-weight: 600;
    color: #222;
    margin: 0 0 4px 0;
}

.artist {
    font-size: 0.9rem;
    color: #666;
    margin: 0;
}

.controls {
    display: flex;
    gap: 12px;
    align-items: center;
}

.control-btn,
.play-btn {
    background: none;
    border: none;
    cursor: pointer;
    padding: 10px;
    border-radius: 50%;
    transition: transform 0.2s ease, background-color 0.2s;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1.2rem;
    color: #555;
}

.control-btn:hover {
    background-color: rgba(0, 0, 0, 0.05);
    transform: scale(1.1);
}

.play-btn {
    background-color: #4CAF50;
    color: white;
    width: 50px;
    height: 50px;
    font-size: 1.4rem;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.play-btn:hover {
    background-color: #45a049;
    transform: scale(1.05);
}

.progress-container {
    display: flex;
    align-items: center;
    width: 100%;
    gap: 8px;
}

.time {
    font-size: 0.75rem;
    color: #888;
    min-width: 36px;
    text-align: center;
}

.progress-bar {
    flex-grow: 1;
    height: 4px;
    background-color: #ddd;
    border-radius: 2px;
    position: relative;
    cursor: pointer;
}

.progress-track {
    width: 100%;
    height: 100%;
    background: linear-gradient(to right, #4CAF50, #8BC34A);
    border-radius: 2px;
    opacity: 0.3;
}

.progress-thumb {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    width: 12px;
    height: 12px;
    background-color: #4CAF50;
    border-radius: 50%;
    box-shadow: 0 0 6px rgba(0, 0, 0, 0.3);
    z-index: 2;
    transition: left 0.1s ease;
}

/* 播放列表弹窗样式 */
.playlist-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 100;
}

.playlist-content {
    background: white;
    padding: 0;
    border-radius: 12px;
    max-width: 300px;
    width: 80%;
    max-height: 90%;
    display: flex;
    flex-direction: column;
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}

.playlist-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 5px 20px;
    border-bottom: 1px solid #eee;
    background-color: #f9f9f9;
    border-radius: 12px 12px 0 0;
    position: sticky;
    top: 0;
    z-index: 10;
}

.playlist-header h4 {
    margin: 0;
    color: #333;
    font-size: 1rem;
    font-weight: 600;
}

.close-btn {
    background: none;
    border: none;
    cursor: pointer;
    padding: 5px;
    border-radius: 4px;
    transition: background-color 0.2s;
    display: flex;
    align-items: center;
    justify-content: center;
}

.close-btn:hover {
    background-color: #e0e0e0;
}

.playlist-items {
    list-style: none;
    padding: 0;
    margin: 0;
    overflow-y: auto;
    /* 隐藏滚动条 */
    scrollbar-width: none;
    flex-grow: 1;
    padding: 10px 0;
}

.playlist-items li {
    padding: 10px 20px;
    border-radius: 0;
    cursor: pointer;
    transition: background-color 0.2s;
    margin-bottom: 0;
    font-size: 0.9rem;
    border-bottom: 1px solid #f5f5f5;
    border-radius: 15px;
}

.playlist-items li:last-child {
    border-bottom: none;
}

.playlist-items li.active {
    background-color: #e8f5e9;
    color: #4CAF50;
    font-weight: 500;
}

.playlist-items li:hover {
    background-color: #e9fed3;
}

/* 隐藏 audio 元素 */
audio {
    display: none;
}

/* 适配移动端 */
@media (max-width: 400px) {
    .music-player-glass {
        width: 95%;
        padding: 15px;
    }

    .album-cover {
        width: 80px;
        height: 80px;
    }

    .play-btn {
        width: 44px;
        height: 44px;
    }

    .circle-button {
        width: 50px;
        height: 50px;
        font-size: 1.5rem;
        box-shadow: 0 3px 12px rgba(0, 0, 0, 0.12);
    }
}
</style>