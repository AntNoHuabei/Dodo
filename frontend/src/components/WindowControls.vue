<script setup lang="ts">
import {Window} from '@wailsio/runtime';
import {onMounted, onUnmounted, ref} from "vue";
import {useI18n} from 'vue-i18n';

// 国际化
const {t} = useI18n();

const maximized = ref(false);

async function refreshMaximized() {
  try {
    maximized.value = await Window.IsMaximised();
  } catch {
    maximized.value = false;
  }
}

function onMinimize() {
  void Window.Minimise();
}

async function onToggleMaximize() {
  try {
    await Window.ToggleMaximise();
    await refreshMaximized();
  } catch {
    /* noop: dev browser / no runtime */
  }
}

function onClose() {
  void Window.Close();
}

function onResize() {
  void refreshMaximized();
}

onMounted(() => {
  void refreshMaximized();
  window.addEventListener('resize', onResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', onResize);
});
</script>

<template>
  <div class="custom-titlebar">
    <div class="titlebar-left window-titlebar-drag ">
      <slot name="title">Dodo 应用</slot>
    </div>
    <div class="custom-titlebar-page-title" id="page-title-target"></div>
    <div class="titlebar-right">
      <button
          type="button"
          class="titlebar-btn"
          :title="t('titlebar.minimize')"
          @click="onMinimize"
      >
        <svg viewBox="0 0 12 12" width="12" height="12" aria-hidden="true">
          <rect x="0" y="5.25" width="12" height="1.5" fill="currentColor" rx="0.25"/>
        </svg>
      </button>
      <button
          type="button"
          class="titlebar-btn"
          :title="maximized ? t('titlebar.restore') : t('titlebar.maximize')"
          @click="onToggleMaximize"
      >
        <svg v-if="!maximized" viewBox="0 0 12 12" width="12" height="12" aria-hidden="true">
          <rect x="1" y="1" width="10" height="10" fill="none" stroke="currentColor" stroke-width="1.25" rx="0.5"/>
        </svg>
        <svg v-else viewBox="0 0 12 12" width="12" height="12" aria-hidden="true">
          <path
              fill="none"
              stroke="currentColor"
              stroke-width="1.25"
              d="M3 4.5h6.5v6.5H3zM4.5 4.5V3h4.75v4.75H8"
          />
        </svg>
      </button>
      <button
          type="button"
          class="titlebar-btn titlebar-btn-close"
          :title="t('titlebar.close')"
          @click="onClose"
      >
        <svg viewBox="0 0 12 12" width="12" height="12" aria-hidden="true">
          <path
              fill="none"
              stroke="currentColor"
              stroke-width="1.25"
              stroke-linecap="round"
              d="M2.5 2.5l7 7M9.5 2.5l-7 7"
          />
        </svg>
      </button>
    </div>
  </div>
</template>

<style scoped>
/* 自定义标题栏 */
.custom-titlebar {
  height: 40px;
  background-color: var(--el-menu-bg-color);
  color: white;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 16px;
  user-select: none;
  cursor: default;
  -webkit-app-region: drag;
  position: relative;
}

.custom-titlebar-page-title {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  color: var(--el-text-color-primary);
  font-size: 14px;
  font-weight: 500;
}
.window-titlebar-drag {
  flex: 1;
  min-width: 0;
  display: flex;
  align-items: center;
  min-height: 0;
  /* Wails v3 reads --wails-draggable; --wails-window-drag kept for requested naming */
  --wails-draggable: drag;
  --wails-window-drag: drag;
}

.titlebar-left {
  font-size: 14px;
  font-weight: 500;
}

.titlebar-right {
  display: flex;
  -webkit-app-region: no-drag;

}

.titlebar-btn {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: transparent;
  border: none;
  color : var(--el-text-color-primary);
  cursor: pointer;
  font-size: 16px;
  margin-left: 8px;
  border-radius: 4px;
  transition: all 0.3s;
}

.titlebar-btn:hover {
  background-color: color-mix(in srgb, var(--el-text-color-primary), transparent 80%);
}

.titlebar-btn-close:hover {
  background-color: #ff4d4f;
}


</style>