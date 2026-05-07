<template>
  <div class="menu-container">
    <el-header class="header">
      <h1>AI应用平台</h1>
      <el-button type="danger" @click="handleLogout">退出登录</el-button>
    </el-header>
    <el-main class="main">
      <div class="menu-grid">
        <el-card class="menu-item" @click="$router.push('/ai-chat')">
          <div class="card-content">
            <el-icon size="48" color="#409eff"><ChatDotRound /></el-icon>
            <h3>AI聊天</h3>
            <p>与AI进行智能对话</p>
          </div>
        </el-card>
        <el-card class="menu-item" @click="$router.push('/image-recognition')">
          <div class="card-content">
            <el-icon size="48" color="#67c23a"><Camera /></el-icon>
            <h3>图像识别</h3>
            <p>上传图片进行AI识别</p>
          </div>
        </el-card>
      </div>
    </el-main>
  </div>
</template>

<script>
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ChatDotRound, Camera } from '@element-plus/icons-vue'

export default {
  name: 'MenuView',
  components: {
    ChatDotRound,
    Camera
  },
  setup() {
    const router = useRouter()

    const handleLogout = async () => {
      try {
        await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        localStorage.removeItem('token')
        ElMessage.success('退出登录成功')
        router.push('/login')
      } catch {
        // 用户取消操作
      }
    }

    return {
      handleLogout
    }
  }
}
</script>

<style scoped>
.menu-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: transparent;
  position: relative;
  overflow: hidden;
}

.menu-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background:
    radial-gradient(900px 420px at 18% 8%, rgba(102, 126, 234, 0.10), transparent 60%),
    radial-gradient(700px 360px at 88% 18%, rgba(118, 75, 162, 0.08), transparent 62%);
}

.header {
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(14px);
  color: #111827;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
  padding: 0 30px;
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.06);
  border-bottom: 1px solid rgba(15, 23, 42, 0.06);
  position: relative;
  z-index: 2;
}

.header h1 {
  margin: 0;
  font-size: 20px;
  font-weight: 700;
  color: rgba(17, 24, 39, 0.92);
}

.el-button {
  background: rgba(245, 108, 108, 0.10);
  border: 1px solid rgba(245, 108, 108, 0.18);
  color: rgba(153, 27, 27, 0.92);
  transition: all var(--dur-1) var(--ease-out);
}

.el-button:hover {
  background: rgba(245, 108, 108, 0.14);
  transform: translateY(-2px);
  box-shadow: 0 14px 40px rgba(15, 23, 42, 0.10);
}

.main {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  z-index: 1;
}

.menu-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 40px;
  max-width: 900px;
  width: 100%;
  padding: 40px;
  animation: gridFadeIn 700ms var(--ease-out);
}

@keyframes gridFadeIn {
  from {
    opacity: 0;
    transform: translateY(18px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.menu-item {
  cursor: pointer;
  background: rgba(255, 255, 255, 0.90);
  backdrop-filter: blur(12px);
  border-radius: 20px;
  box-shadow: 0 12px 40px rgba(15, 23, 42, 0.10);
  border: 1px solid rgba(15, 23, 42, 0.06);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  position: relative;
  overflow: hidden;
  animation: cardSlideIn 0.8s ease-out both;
}

.menu-item:nth-child(1) { animation-delay: 0.1s; }
.menu-item:nth-child(2) { animation-delay: 0.2s; }

@keyframes cardSlideIn {
  from {
    opacity: 0;
    transform: translateY(36px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.menu-item::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(102,126,234,0.10), transparent);
  transition: left 0.6s;
}

.menu-item:hover::before {
  left: 100%;
}

.menu-item:hover {
  transform: translateY(-6px) scale(1.03);
  box-shadow: 0 18px 60px rgba(15, 23, 42, 0.16);
}

.card-content {
  text-align: center;
  padding: 50px 30px;
  position: relative;
  z-index: 1;
}

.el-icon {
  display: block;
  margin: 0 auto 20px;
  transition: all 0.3s ease;
}

.menu-item:hover .el-icon {
  transform: scale(1.1);
}

.card-content h3 {
  margin: 0 0 15px 0;
  color: #2c3e50;
  font-size: 24px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.menu-item:hover h3 {
  color: rgba(17, 24, 39, 0.92);
  transform: translateY(-2px);
}

.card-content p {
  margin: 0;
  color: #7f8c8d;
  font-size: 16px;
  line-height: 1.6;
  transition: all 0.3s ease;
}

.menu-item:hover p {
  color: #34495e;
  transform: translateY(-1px);
}
</style>
