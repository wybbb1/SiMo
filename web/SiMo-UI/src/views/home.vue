<template>

    <h1>Home Page</h1>
    <p>
        Welcome to the home page!
    </p>
    <n-button-group>
        <n-button @click="router.push('/write')">
            <n-icon size="20px">
                <CalligraphyPen20Regular />
            </n-icon>
        </n-button>
        <n-button @click="router.push('/whiteBoard')">
            <n-icon size="20px">
                <Whiteboard20Regular />
            </n-icon>
        </n-button>
    </n-button-group>

    <music />
    
</template>

<script setup lang="ts">
import music from '@/components/music.vue';
import { useRouter } from "vue-router"
import { NButton, NIcon, NButtonGroup } from 'naive-ui'
import { CalligraphyPen20Regular, Whiteboard20Regular } from '@vicons/fluent'

const router = useRouter()

</script>