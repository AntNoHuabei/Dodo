<script setup lang="ts">
import {ref, computed, onMounted} from 'vue';
import { useI18n } from 'vue-i18n';
import { ElConfigProvider } from 'element-plus';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';
import WindowControls from './components/WindowControls.vue';
import Sidebar from './components/Sidebar.vue';
import TaskList from './components/TaskList.vue';
import Weekly from './components/Weekly.vue';
import Settings from './components/Settings.vue';

// 国际化
const { t, locale } = useI18n();

// 当前选中菜单
const activeMenu = ref('0');

onMounted(()=>{
  activeMenu.value = '1'
})

const handleMenuChange = (index: string) => {
  activeMenu.value = index;
};

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
        <Sidebar @menu-change="handleMenuChange" />

        <!-- 右侧内容区 -->
        <div class="content-area">
          <TaskList v-if="activeMenu === '1'" />
          <Weekly v-else-if="activeMenu === '2'" />
          <Settings v-else-if="activeMenu === '3'" />
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
  background-color: var(--el-bg-color);
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

  transition: background-color 0.3s;
  border-left: solid 1px var(--el-menu-border-color);
  border-top: solid 1px var(--el-menu-border-color);
  border-top-left-radius: 10px;
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