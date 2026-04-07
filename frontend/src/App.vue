<script setup lang="ts">
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { ElConfigProvider } from 'element-plus';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';
import WindowControls from './components/WindowControls.vue';
import Sidebar from './components/Sidebar.vue';

// 国际化
const { t, locale } = useI18n();

// Element Plus 语言包映射
const elementLocale = computed(() => {
  return locale.value === 'zh-CN' ? zhCn : en;
});
</script>

<template>
  <el-config-provider :locale="elementLocale">
    <div class="app-container">
      <!-- 自定义标题栏 -->
      <WindowControls>
        <template #title>
          <span class="app-title">{{ t('app.title') }}</span>
        </template>
      </WindowControls>

      <!-- 主内容区 -->
      <div class="main-content">
        <!-- 左侧边栏 -->
        <Sidebar />

        <!-- 右侧内容区 -->
        <div class="content-area">
          <h1>{{ t('app.content.welcome') }}</h1>
          <p>{{ t('app.content.description') }}</p>
        </div>
      </div>
    </div>
  </el-config-provider>
</template>

<style scoped>
.app-container {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f5f7fa;
  overflow: hidden;
  transition: background-color 0.3s;
}

/* 主内容区 */
.main-content {
  flex: 1;
  display: flex;
  overflow: hidden;
}

/* 右侧内容区 */
.content-area {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
  background-color: #f5f7fa;
  transition: background-color 0.3s;
}

.content-area h1 {
  margin-bottom: 16px;
  color: #303133;
  transition: color 0.3s;
}

.content-area p {
  color: #606266;
  font-size: 14px;
  transition: color 0.3s;
}
.app-title{
  color: var(--el-text-color-primary);
}

</style>