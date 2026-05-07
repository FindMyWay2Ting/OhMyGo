<template>
  <div class="login-container">
    <el-card class="login-card">
      <template #header>
        <div class="card-header">
          <h2>登录</h2>
        </div>
      </template>
      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="loginRules"
        label-width="80px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
          />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="loginForm.password"
            placeholder="请输入密码"
            type="password"
            show-password
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            @click="handleLogin"
            style="width: 100%"
          >
            登录
          </el-button>
        </el-form-item>
        <el-form-item>
          <el-button
            type="text"
            @click="$router.push('/register')"
            style="width: 100%"
          >
            还没有账号？去注册
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import api from '../utils/api'

export default {
  name: 'LoginView',
  setup() {
    const router = useRouter()
    const loginFormRef = ref()
    const loading = ref(false)
    const loginForm = ref({
      username: '',
      password: ''
    })

    const loginRules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, message: '密码长度不能少于6位', trigger: 'blur' }
      ]
    }

    const handleLogin = async () => {
      try {
        await loginFormRef.value.validate()
        loading.value = true
        const response = await api.post('/user/login', {
          username: loginForm.value.username,
          password: loginForm.value.password
        })
        if (response.data.status_code === 1000) {
          localStorage.setItem('token', response.data.token)
          ElMessage.success('登录成功')
          router.push('/menu')
        } else {
          ElMessage.error(response.data.status_msg || '登录失败')
        }
      } catch (error) {
        console.error('Login error:', error)
        ElMessage.error('登录失败，请重试')
      } finally {
        loading.value = false
      }
    }

    return {
      loginFormRef,
      loading,
      loginForm,
      loginRules,
      handleLogin
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: radial-gradient(1200px 700px at 20% 10%, rgba(102,126,234,0.85), rgba(102,126,234,0.10) 55%, transparent 70%),
              radial-gradient(1000px 620px at 80% 18%, rgba(240,147,251,0.72), rgba(118,75,162,0.10) 55%, transparent 72%),
              radial-gradient(900px 600px at 55% 85%, rgba(245,87,108,0.55), rgba(245,87,108,0.06) 55%, transparent 70%),
              linear-gradient(135deg, #0b1020 0%, #0b1633 50%, #0b1020 100%);
  position: relative;
  overflow: hidden;
}

.login-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background:
    radial-gradient(800px 500px at 20% 15%, rgba(255,255,255,0.16), transparent 60%),
    radial-gradient(700px 420px at 78% 22%, rgba(255,255,255,0.12), transparent 62%),
    radial-gradient(600px 380px at 55% 88%, rgba(255,255,255,0.10), transparent 64%);
  filter: blur(0px);
  opacity: 0.9;
  animation: aurora 14s cubic-bezier(0.22, 1, 0.36, 1) infinite;
}

.login-container::after {
  content: '';
  position: absolute;
  inset: 0;
  background-image: url('data:image/svg+xml,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 120 120"><filter id="n"><feTurbulence type="fractalNoise" baseFrequency="0.9" numOctaves="3" stitchTiles="stitch"/></filter><rect width="120" height="120" filter="url(%23n)" opacity="0.28"/></svg>');
  opacity: 0.08;
  mix-blend-mode: overlay;
  pointer-events: none;
}

@keyframes aurora {
  0% { transform: translate3d(0, 0, 0) scale(1); }
  33% { transform: translate3d(-18px, -14px, 0) scale(1.02); }
  66% { transform: translate3d(16px, 10px, 0) scale(1.03); }
  100% { transform: translate3d(0, 0, 0) scale(1); }
}

.login-card {
  width: 420px;
  max-width: calc(100vw - 32px);
  margin: 16px;
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(18px);
  border-radius: 20px;
  box-shadow: 0 24px 70px rgba(0, 0, 0, 0.28);
  border: 1px solid rgba(255, 255, 255, 0.35);
  animation: slideIn 0.8s ease-out;
  position: relative;
  z-index: 1;
}

.login-card::before {
  content: '';
  position: absolute;
  inset: -1px;
  border-radius: 20px;
  padding: 1px;
  background: linear-gradient(135deg, rgba(102,126,234,0.55), rgba(240,147,251,0.42), rgba(245,87,108,0.35));
  -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
  -webkit-mask-composite: xor;
  mask-composite: exclude;
  pointer-events: none;
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(24px) scale(0.98);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

.card-header {
  text-align: center;
  padding: 30px 0 20px 0;
}

.card-header h2 {
  margin: 0;
  background: linear-gradient(135deg, #1f2937 0%, rgba(31,41,55,0.75) 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  font-size: 28px;
  font-weight: 600;
}

.el-form-item {
  margin-bottom: 24px;
}

:deep(.el-form-item__label) {
  color: rgba(31, 41, 55, 0.72);
  font-weight: 600;
}

.el-input {
  transition: all 0.3s ease;
}

.el-input:focus-within {
  transform: scale(1.02);
}

.el-button {
  height: 48px;
  border-radius: 12px;
  font-weight: 600;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.el-button::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s;
}

.el-button:hover::before {
  left: 100%;
}

.el-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 14px 40px rgba(15, 23, 42, 0.18);
}

:deep(.el-button.is-text) {
  height: auto;
  padding: 6px 0;
  font-weight: 600;
}

.login-link {
  text-align: center;
  margin-top: 20px;
  animation: fadeIn 1s ease-out 0.5s both;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
