<template>
    <div class="login-container">
        <div class="login-card">
            <div class="login-header">
                <h1>LOG IN</h1>
            </div>

            <div class="form-section">
                <div class="input-group">
                    <label>username</label>
                    <n-input v-model:value="username" round size="large" />
                </div>

                <div class="input-group">
                    <div class="password-label-row">
                        <label>password</label>
                    </div>
                    <n-input v-model:value="password" type="password" show-password-on="mousedown" round size="large"/>
                    <div class="forgot-password">
                        <a href="#" class="forgot-link">forget password</a>
                    </div>
                </div>
            </div>

            <div class="button-section">
                <n-button class="login-btn" @click="handleLogin" color="#4ba0d1" round>登录</n-button>
                <n-button class="guest-btn" color="#3dad8d" round>以游客身份登录</n-button>
                <n-button class="home-btn" :loading="isWaitingLogin" @click="goHome()" round>Go to Home(Test)</n-button>
            </div>
        </div>

        <footer class="login-footer">
            <p>contact admin to get the auth</p>
        </footer>
    </div>
</template>

<script setup lang="ts">
import { NButton, NInput } from "naive-ui"
import { useRouter } from "vue-router"
import { ref } from "vue"

const router = useRouter()
const username = ref('')
const password = ref('')
const isWaitingLogin = ref(false)

const handleLogin = () => {
    // 登录逻辑
    
}

const goHome = () => {
    isWaitingLogin.value = true
    setTimeout(() => {
        isWaitingLogin.value = false
        router.push('/home')
    }, 2000)
}
</script>

<style scoped>
.login-container {
    width: 100%;
    height: 100vh;
    /* background: linear-gradient(135deg, #9cc0ea 0%, #205898 100%); */
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    position: relative;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

.login-card {
    background: #E8E8E8;
    border-radius: 12px;
    padding: 40px;
    width: 380px;
    box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    gap: 24px;
}

.login-header h1 {
    text-align: center;
    font-size: 24px;
    font-weight: 400;
    color: #333;
    margin: 0;
    letter-spacing: 2px;
}

.form-section {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.input-group {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.input-group label {
    font-size: 14px;
    color: #333;
    font-weight: normal;
}

.password-label-row {
    display: flex;
    justify-content: flex-start;
    align-items: center;
}

.forgot-password {
    display: flex;
    justify-content: flex-end;
    margin-top: 4px;
}

.forgot-link {
    color: #4A90E2;
    font-size: 12px;
    text-decoration: none;
}

.forgot-link:hover {
    text-decoration: underline;
}

.button-section {
    display: flex;
    flex-direction: column;
    gap: 12px;
    margin-top: 8px;
}

.login-btn {
    height: 36px;
    font-size: 14px;
    font-weight: normal;
}

.guest-btn,
.home-btn {
    height: 36px;
    font-size: 14px;
}

.login-footer {
    position: absolute;
    bottom: 40px;
    text-align: center;
}

.login-footer p {
    color: #eb7456;
    font-size: 16px;
    margin: 0;
    font-weight: normal;
}

</style>