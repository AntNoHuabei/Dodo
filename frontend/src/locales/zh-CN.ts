// 中文语言包
export default {
  app: {
    title: 'Dodo',
    menu: {
      task: '任务',
      weekly: '周报',
      settings: '设置',
      profile: '个人中心'
    },
    content: {
      welcome: '欢迎使用 Dodo 应用',
      description: '这是一个基于 Wails 3 和 Element Plus 的 PC 端应用'
    }
  },
  titlebar: {
    minimize: '最小化',
    maximize: '最大化',
    restore: '还原',
    close: '关闭',
    pin: '置顶',
    unpin: '取消置顶'
  },
  common: {
    cancel: '取消',
    confirm: '确认',
    save: '保存'
  },
  task: {
    title: '任务列表',
    inputPlaceholder: '请输入新任务...',
    add: '添加',
    edit: '编辑',
    save: '保存',
    delete: '删除',
    deleteConfirm: '确定要删除这个任务吗？',
    deleteConfirmTitle: '确认删除',
    empty: '暂无任务',
    expandCompleted: '展开已完成',
    collapseCompleted: '折叠已完成',
    priority: {
      label: '优先级',
      high: '高',
      medium: '中',
      low: '低',
      select: '选择优先级'
    },
    type: {
      label: '任务类型',
      bug: 'Bug',
      feature: '功能',
      refactor: '重构',
      performance: '性能',
      security: '安全',
      chore: '杂务',
      migration: '迁移',
      ci: 'CI',
      test: '测试',
      review: '审查',
      doc: '文档',
      subtask: '子任务'
    }
  },
  weekly: {
    title: '周视图',
    previous: '上一周',
    next: '下一周',
    today: '今天',
    month: '月',
    noSchedule: '暂无日程',
    mon: '周一',
    tue: '周二',
    wed: '周三',
    thu: '周四',
    fri: '周五',
    sat: '周六',
    sun: '周日'
  },
  settings: {
    title: '设置',
    appearance: '外观',
    darkMode: '深色模式',
    darkModeDesc: '切换深色/浅色主题',
    language: '语言',
    selectLanguage: '选择语言',
    languageDesc: '选择界面显示语言',
    languageChanged: '语言已切换',
    data: '数据管理',
    resetData: '重置设置',
    resetDataDesc: '恢复所有设置为默认值',
    reset: '重置',
    resetSuccess: '设置已重置',
    about: '关于',
    version: '版本',
    author: '作者',
    description: '描述',
    appDesc: '一个基于 Wails 3 和 Vue 3 的桌面应用'
  }
};