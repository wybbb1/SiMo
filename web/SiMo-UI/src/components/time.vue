<template>
    <div class="clock-container" :class="{ compact: compact }">
        <!-- 数字时钟 -->
        <div class="digital-clock">
            <div class="time-display">
                {{ currentTime }}
            </div>
            <div class="date-display" v-if="!compact">
                {{ currentDate }}
            </div>
        </div>

        <!-- 模拟时钟 -->
        <div class="analog-clock" v-if="showAnalog">
            <div class="clock-face">
                <!-- 时刻度 -->
                <div v-for="hour in 12" :key="`hour-${hour}`" class="hour-mark"
                    :style="{ transform: `rotate(${hour * 30}deg)` }">
                    <div class="hour-mark-line"></div>
                </div>

                <!-- 分钟刻度 -->
                <div v-for="minute in 60" :key="`minute-${minute}`" class="minute-mark"
                    :style="{ transform: `rotate(${minute * 6}deg)` }" v-show="minute % 5 !== 0">
                    <div class="minute-mark-line"></div>
                </div>

                <!-- 指针 -->
                <div class="hand hour-hand" :style="{ transform: `rotate(${hourAngle}deg)` }"></div>
                <div class="hand minute-hand" :style="{ transform: `rotate(${minuteAngle}deg)` }"></div>
                <div class="hand second-hand" :style="{ transform: `rotate(${secondAngle}deg)` }"></div>

                <!-- 中心点 -->
                <div class="center-dot"></div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'

// 组件属性
interface Props {
    showAnalog?: boolean    // 是否显示模拟时钟
    format24?: boolean      // 是否使用24小时制
    showSeconds?: boolean   // 是否显示秒数
    compact?: boolean       // 是否紧凑显示
}

const props = withDefaults(defineProps<Props>(), {
    showAnalog: true,
    format24: true,
    showSeconds: true,
    compact: false
})

// 响应式数据
const currentTime = ref('')
const currentDate = ref('')
const now = ref(new Date())

let timer: number | null = null

// 计算指针角度
const hourAngle = computed(() => {
    const hours = now.value.getHours() % 12
    const minutes = now.value.getMinutes()
    return (hours * 30) + (minutes * 0.5) - 90 // -90度是为了让12点指向顶部
})

const minuteAngle = computed(() => {
    const minutes = now.value.getMinutes()
    const seconds = now.value.getSeconds()
    return (minutes * 6) + (seconds * 0.1) - 90
})

const secondAngle = computed(() => {
    const seconds = now.value.getSeconds()
    return (seconds * 6) - 90
})

// 更新时间
const updateTime = () => {
    now.value = new Date()

    // 格式化时间显示
    const hours = now.value.getHours()
    const minutes = now.value.getMinutes().toString().padStart(2, '0')
    const seconds = now.value.getSeconds().toString().padStart(2, '0')

    let timeStr = ''
    if (props.format24) {
        timeStr = `${hours.toString().padStart(2, '0')}:${minutes}`
    } else {
        const hour12 = hours % 12 || 12
        const ampm = hours >= 12 ? 'PM' : 'AM'
        timeStr = `${hour12}:${minutes} ${ampm}`
    }

    if (props.showSeconds) {
        timeStr += `:${seconds}`
    }

    currentTime.value = timeStr

    // 格式化日期显示
    const year = now.value.getFullYear()
    const month = (now.value.getMonth() + 1).toString().padStart(2, '0')
    const date = now.value.getDate().toString().padStart(2, '0')
    const weekdays = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']
    const weekday = weekdays[now.value.getDay()]

    currentDate.value = `${year}年${month}月${date}日 ${weekday}`
}

// 生命周期
onMounted(() => {
    updateTime()
    timer = setInterval(updateTime, 1000)
})

onUnmounted(() => {
    if (timer) {
        clearInterval(timer)
    }
})
</script>

<style scoped>
.clock-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;
    padding: 20px;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* 紧凑模式样式 */
.clock-container.compact {
    gap: 8px;
    padding: 8px 16px;
}

.clock-container.compact .time-display {
    font-size: 1.5rem;
    margin-bottom: 0;
    letter-spacing: 1px;
}

/* 数字时钟样式 */
.digital-clock {
    text-align: center;
    user-select: none;
}

.time-display {
    font-size: 2.5rem;
    font-weight: 700;
    color: #2c3e50;
    margin-bottom: 8px;
    letter-spacing: 2px;
    text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.date-display {
    font-size: 1rem;
    color: #7f8c8d;
    font-weight: 500;
}

/* 模拟时钟样式 */
.analog-clock {
    position: relative;
}

.clock-face {
    width: 200px;
    height: 200px;
    border: 4px solid #34495e;
    border-radius: 50%;
    position: relative;
    background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
    box-shadow:
        0 8px 32px rgba(0, 0, 0, 0.1),
        inset 0 2px 8px rgba(0, 0, 0, 0.05);
}

/* 时刻度 */
.hour-mark {
    position: absolute;
    width: 2px;
    height: 100px;
    top: 0;
    left: 50%;
    transform-origin: bottom center;
}

.hour-mark-line {
    width: 4px;
    height: 20px;
    background: #2c3e50;
    margin: 0 auto;
    border-radius: 2px;
}

/* 分钟刻度 */
.minute-mark {
    position: absolute;
    width: 1px;
    height: 100px;
    top: 0;
    left: 50%;
    transform-origin: bottom center;
}

.minute-mark-line {
    width: 1px;
    height: 10px;
    background: #7f8c8d;
    margin: 0 auto;
}

/* 指针样式 */
.hand {
    position: absolute;
    transform-origin: bottom center;
    border-radius: 10px;
    transition: transform 0.1s ease-out;
}

.hour-hand {
    width: 6px;
    height: 60px;
    background: linear-gradient(to top, #2c3e50, #34495e);
    bottom: 50%;
    left: calc(50% - 3px);
    z-index: 3;
}

.minute-hand {
    width: 4px;
    height: 80px;
    background: linear-gradient(to top, #3498db, #5dade2);
    bottom: 50%;
    left: calc(50% - 2px);
    z-index: 2;
}

.second-hand {
    width: 2px;
    height: 90px;
    background: linear-gradient(to top, #e74c3c, #ec7063);
    bottom: 50%;
    left: calc(50% - 1px);
    z-index: 4;
}

/* 中心点 */
.center-dot {
    width: 12px;
    height: 12px;
    background: #2c3e50;
    border-radius: 50%;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    z-index: 5;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

/* 响应式设计 */
@media (max-width: 768px) {
    .clock-container {
        gap: 15px;
        padding: 15px;
    }

    .time-display {
        font-size: 2rem;
    }

    .date-display {
        font-size: 0.9rem;
    }

    .clock-face {
        width: 150px;
        height: 150px;
    }

    .hour-hand {
        height: 45px;
    }

    .minute-hand {
        height: 60px;
    }

    .second-hand {
        height: 65px;
    }
}
</style>