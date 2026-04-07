<script setup lang="ts">
import { ref, computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { useDark } from '@vueuse/core';
import { ElMenu, ElMenuItem, ElTooltip, ElIcon } from 'element-plus';

const emit = defineEmits<{
  (e: 'menu-change', index: string): void;
}>();

// 国际化
const { t, locale } = useI18n();

// 主题状态
const isDark = useDark({
  selector: 'html',
  attribute: 'class',
  valueDark: 'dark',
  valueLight: 'light'
});

// 切换语言
const switchLanguage = () => {
  locale.value = locale.value === 'zh-CN' ? 'en-US' : 'zh-CN';
};

// 菜单数据
const menuItems = computed(() => [
  { id: 1, icon: 'List', title: t('app.menu.task') },
  { id: 2, icon: 'Calendar', title: t('app.menu.weekly') },
  { id: 3, icon: 'Setting', title: t('app.menu.settings') }

]);

// 当前选中的菜单项
const activeIndex = ref('1');

const handleMenuSelect = (index: string) => {
  activeIndex.value = index;
  emit('menu-change', index);
};
</script>

<template>
  <div class="sidebar">
    <!-- 上部菜单区域 -->
    <div class="sidebar-menu-area">
      <el-menu
        :default-active="activeIndex"
        class="sidebar-menu"
        mode="vertical"
        @select="handleMenuSelect"
      >
        <el-menu-item
          v-for="item in menuItems"
          :key="item.id"
          :index="item.id.toString()"
        >
          <el-tooltip
            :content="item.title"
            placement="right"
            effect="dark"
          >
            <el-icon class="menu-icon">
              <component :is="item.icon" />
            </el-icon>
          </el-tooltip>
        </el-menu-item>
      </el-menu>
    </div>

    <!-- 底部功能区域 -->
    <div class="sidebar-footer">
      <!-- 主题切换按钮 -->
      <el-tooltip
        :content="isDark ? '切换到亮色模式' : '切换到暗色模式'"
        placement="right"
        effect="dark"
      >
        <button class="sidebar-footer-btn el-menu-item" @click="isDark = !isDark">
          <el-icon class="footer-icon">
            <component :is="isDark ? 'Sunny' : 'Moon'" />
          </el-icon>
        </button>
      </el-tooltip>

      <!-- 语言切换按钮 -->
      <el-tooltip
        :content="locale === 'zh-CN' ? 'Switch to English' : '切换到中文'"
        placement="right"
        effect="dark"
      >
        <button class="sidebar-footer-btn language-btn el-menu-item" @click="switchLanguage">
          {{ locale === 'zh-CN' ? 'CN' : 'EN' }}
        </button>
      </el-tooltip>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  width: 64px;
  background-color: var(--el-menu-bg-color);
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  transition: background-color 0.3s;
}

.sidebar-menu-area {
  flex: 1;
  overflow-y: auto;
}

.sidebar-menu {
  width: 100%;
  height: 100%;
  border-right: none;
}

.menu-icon {
  font-size: 20px;
}

.sidebar-footer {
  padding: 8px 0;

  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.sidebar-footer-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border: none;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.3s;
  margin: 0;
}



.language-btn {
  font-size: 14px;
  font-weight: 500;
}

.footer-icon {
  margin: 0 !important;
}

</style>