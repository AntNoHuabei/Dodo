<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useDark } from '@vueuse/core';
import { ElCard, ElSwitch, ElSelect, ElOption, ElButton, ElMessage } from 'element-plus';

const { t, locale, availableLocales } = useI18n();

const isDark = useDark({
  selector: 'html',
  attribute: 'class',
  valueDark: 'dark',
  valueLight: ''
});

const currentLanguage = ref(locale.value);

const languageOptions = [
  { value: 'zh-CN', label: '中文' },
  { value: 'en-US', label: 'English' }
];

const changeLanguage = (lang: string) => {
  locale.value = lang;
  currentLanguage.value = lang;
  ElMessage.success(t('settings.languageChanged'));
};

const toggleTheme = () => {
  isDark.value = !isDark.value;
};

const resetSettings = () => {
  locale.value = 'zh-CN';
  currentLanguage.value = 'zh-CN';
  isDark.value = false;
  ElMessage.success(t('settings.resetSuccess'));
};
</script>

<template>
  <div class="settings">
    <Teleport to="#page-title-target">
      <span>{{ t('settings.title') }}</span>
    </Teleport>


    <div class="settings-content">
      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <span>{{ t('settings.appearance') }}</span>
          </div>
        </template>
        <div class="setting-item">
          <div class="setting-label">
            <span>{{ t('settings.darkMode') }}</span>
            <span class="setting-description">{{ t('settings.darkModeDesc') }}</span>
          </div>
          <el-switch v-model="isDark" @change="toggleTheme" />
        </div>
      </el-card>

      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <span>{{ t('settings.language') }}</span>
          </div>
        </template>
        <div class="setting-item">
          <div class="setting-label">
            <span>{{ t('settings.selectLanguage') }}</span>
            <span class="setting-description">{{ t('settings.languageDesc') }}</span>
          </div>
          <el-select
            v-model="currentLanguage"
            @change="changeLanguage"
          >
            <el-option
              v-for="lang in languageOptions"
              :key="lang.value"
              :label="lang.label"
              :value="lang.value"
            />
          </el-select>
        </div>
      </el-card>

      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <span>{{ t('settings.data') }}</span>
          </div>
        </template>
        <div class="setting-item">
          <div class="setting-label">
            <span>{{ t('settings.resetData') }}</span>
            <span class="setting-description">{{ t('settings.resetDataDesc') }}</span>
          </div>
          <el-button type="danger" @click="resetSettings">
            {{ t('settings.reset') }}
          </el-button>
        </div>
      </el-card>

      <el-card class="settings-card">
        <template #header>
          <div class="card-header">
            <span>{{ t('settings.about') }}</span>
          </div>
        </template>
        <div class="about-info">
          <p><strong>{{ t('settings.version') }}:</strong> 1.0.0</p>
          <p><strong>{{ t('settings.author') }}:</strong> Dodo Team</p>
          <p><strong>{{ t('settings.description') }}:</strong> {{ t('settings.appDesc') }}</p>
        </div>
      </el-card>
    </div>
  </div>
</template>

<style scoped>
.settings {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.settings-header {
  margin-bottom: 20px;
}

.settings-header h2 {
  margin: 0;
  color: #303133;
}

.settings-content {
  flex: 1;
  overflow-y: auto;
}

.settings-card {
  margin-bottom: 16px;
}

.card-header {
  font-weight: 600;
  color: #303133;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 0;
}

.setting-label {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.setting-label span {
  color: #303133;
  font-size: 14px;
}

.setting-description {
  color: #909399 !important;
  font-size: 12px !important;
}

.about-info p {
  margin: 8px 0;
  color: #606266;
  font-size: 14px;
}

.about-info strong {
  color: #303133;
}
</style>