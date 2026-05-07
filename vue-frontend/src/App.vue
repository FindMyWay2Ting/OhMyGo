<template>
  <div id="app">
    <router-view v-slot="{ Component }">
      <transition name="page" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
:root {
  --brand-500: #667eea;
  --brand-600: #5a6fe6;
  --brand-700: #4b5fe1;
  --accent-500: #764ba2;
  --success-500: #67c23a;
  --danger-500: #f56c6c;

  --bg-page: #f6f7fb;
  --bg-elevated: rgba(255, 255, 255, 0.96);
  --bg-glass: rgba(255, 255, 255, 0.92);

  --text-900: #1f2937;
  --text-700: rgba(31, 41, 55, 0.78);
  --text-500: rgba(31, 41, 55, 0.56);

  --border-100: rgba(15, 23, 42, 0.08);
  --border-200: rgba(15, 23, 42, 0.12);

  --radius-sm: 10px;
  --radius-md: 12px;
  --radius-lg: 16px;
  --radius-xl: 20px;

  --shadow-card: 0 12px 40px rgba(0, 0, 0, 0.10);
  --shadow-hover: 0 18px 60px rgba(0, 0, 0, 0.16);
  --shadow-soft: 0 6px 20px rgba(0, 0, 0, 0.08);

  --ease-out: cubic-bezier(0.22, 1, 0.36, 1);
  --dur-1: 180ms;
  --dur-2: 240ms;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  height: 100%;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', '微软雅黑', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background: var(--bg-page);
  color: var(--text-900);
}

body {
  background-image:
    radial-gradient(900px 420px at 18% 8%, rgba(102, 126, 234, 0.10), transparent 60%),
    radial-gradient(700px 360px at 88% 18%, rgba(118, 75, 162, 0.08), transparent 62%);
  background-attachment: fixed;
}

#app {
  height: 100%;
}

/* 页面切换动画 */
.page-enter-active,
.page-leave-active {
  transition: opacity var(--dur-2) var(--ease-out), transform var(--dur-2) var(--ease-out);
}

.page-enter-from {
  opacity: 0;
  transform: translateY(10px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* 全局滚动条样式 */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: rgba(102, 126, 234, 0.28);
  border-radius: 4px;
  transition: background 0.3s ease;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(102, 126, 234, 0.5);
}

/* Element Plus 组件样式覆盖 */
.el-button {
  font-weight: 500;
  border-radius: var(--radius-md);
}

.el-input {
  border-radius: var(--radius-md);
}

.el-card {
  border-radius: var(--radius-xl);
}

.el-message {
  border-radius: var(--radius-md);
}

.el-input__wrapper {
  border-radius: var(--radius-md);
}

.el-button.is-text {
  color: rgba(31, 41, 55, 0.62);
}

.el-button.is-text:hover {
  color: var(--brand-600);
  background: transparent;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .page-enter-from,
  .page-leave-to {
    transform: translateX(0);
    opacity: 0;
  }

  .page-enter-active,
  .page-leave-active {
    transition: opacity 0.3s ease;
  }
}
</style>
