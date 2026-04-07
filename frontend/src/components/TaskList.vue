<script setup lang="ts">
import { ref, computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { Plus } from '@element-plus/icons-vue';
import { ElCard, ElCheckbox, ElButton, ElInput, ElTag, ElEmpty, ElDialog, ElSelect, ElOption, ElMessageBox } from 'element-plus';

const { t } = useI18n();

interface Task {
  id: number;
  title: string;
  completed: boolean;
  priority: 'low' | 'medium' | 'high';
}

const tasks = ref<Task[]>([
  { id: 1, title: '完成项目需求分析', completed: true, priority: 'high' },
  { id: 2, title: '设计数据库结构', completed: false, priority: 'high' },
  { id: 3, title: '编写前端组件', completed: false, priority: 'medium' },
  { id: 4, title: '编写后端API', completed: false, priority: 'medium' },
  { id: 5, title: '编写测试用例', completed: false, priority: 'low' }
]);

const newTaskTitle = ref('');
const dialogVisible = ref(false);
const newTaskPriority = ref<'low' | 'medium' | 'high'>('medium');
const showCompleted = ref(false);
const editingTask = ref<Task | null>(null);

const priorityOrder = (p: string) => {
  const order: Record<string, number> = { high: 1, medium: 2, low: 3 };
  return order[p] || 4;
};

const activeTasks = computed(() => 
  tasks.value
    .filter(t => !t.completed)
    .sort((a, b) => priorityOrder(a.priority) - priorityOrder(b.priority))
);
const completedTasks = computed(() => 
  tasks.value
    .filter(t => t.completed)
    .sort((a, b) => priorityOrder(a.priority) - priorityOrder(b.priority))
);
const displayTasks = computed(() => showCompleted.value ? [...activeTasks.value, ...completedTasks.value] : activeTasks.value);

const isEditing = computed(() => editingTask.value !== null);

const openDialog = (task?: Task) => {
  if (task) {
    editingTask.value = task;
    newTaskTitle.value = task.title;
    newTaskPriority.value = task.priority;
  } else {
    editingTask.value = null;
    newTaskTitle.value = '';
    newTaskPriority.value = 'medium';
  }
  dialogVisible.value = true;
};

const saveTask = () => {
  if (newTaskTitle.value.trim()) {
    if (editingTask.value) {
      editingTask.value.title = newTaskTitle.value.trim();
      editingTask.value.priority = newTaskPriority.value;
    } else {
      tasks.value.push({
        id: Date.now(),
        title: newTaskTitle.value.trim(),
        completed: false,
        priority: newTaskPriority.value
      });
    }
    dialogVisible.value = false;
  }
};

const toggleTask = (task: Task) => {
  task.completed = !task.completed;
};

const deleteTask = async (id: number) => {
  try {
    await ElMessageBox.confirm(
      t('task.deleteConfirm'),
      t('task.deleteConfirmTitle'),
      {
        confirmButtonText: t('common.confirm'),
        cancelButtonText: t('common.cancel'),
        type: 'warning'
      }
    );
    tasks.value = tasks.value.filter(t => t.id !== id);
  } catch {
    // 用户取消删除
  }
};

const getPriorityColor = (priority: string) => {
  const colors: Record<string, 'danger' | 'warning' | 'success' | 'info'> = {
    high: 'danger',
    medium: 'warning',
    low: 'success'
  };
  return colors[priority] || 'info';
};

const getPriorityLabel = (priority: string) => {
  const labels: Record<string, string> = {
    high: t('task.priority.high'),
    medium: t('task.priority.medium'),
    low: t('task.priority.low')
  };
  return labels[priority] || priority;
};
</script>

<template>
  <div class="task-list">
    <Teleport to="#page-title-target">
      <span>{{ t('task.title') }}</span>
    </Teleport>

    <div class="task-items">
      <el-empty v-if="tasks.length === 0" :description="t('task.empty')" />
      <el-card
        v-for="task in displayTasks"
        :key="task.id"
        class="task-item"
        :class="{ completed: task.completed }"
      >
        <div class="task-content">
          <el-checkbox
            :model-value="task.completed"
            @change="toggleTask(task)"
          />
          <span class="task-title" @dblclick="openDialog(task)">{{ task.title }}</span>
          <el-tag :type="getPriorityColor(task.priority)" size="small">
            {{ getPriorityLabel(task.priority) }}
          </el-tag>
          <el-button
            type="danger"
            size="small"
            text
            @click="deleteTask(task.id)"
          >
            {{ t('task.delete') }}
          </el-button>
        </div>
      </el-card>

      <div v-if="completedTasks.length > 0" class="completed-toggle">
        <el-button link type="primary" @click="showCompleted = !showCompleted">
          {{ showCompleted ? t('task.collapseCompleted') : t('task.expandCompleted') }}
          ({{ completedTasks.length }})
        </el-button>
      </div>
    </div>

    <el-button
      class="add-float-btn"
      type="primary"
      :icon="Plus"
      circle
      size="large"
      @click="openDialog()"
    />

    <el-dialog
      v-model="dialogVisible"
      :title="isEditing ? t('task.edit') : t('task.add')"
      width="400px"
      :close-on-click-modal="false"
    >
      <div class="dialog-form">
        <el-input
          v-model="newTaskTitle"
          :placeholder="t('task.inputPlaceholder')"
          @keyup.enter="saveTask"
        />
        <div class="priority-tags">
          <el-tag
            :type="newTaskPriority === 'low' ? 'success' : 'info'"
            :effect="newTaskPriority === 'low' ? 'dark' : 'plain'"
            class="priority-tag"
            @click="newTaskPriority = 'low'"
          >
            {{ t('task.priority.low') }}
          </el-tag>
          <el-tag
            :type="newTaskPriority === 'medium' ? 'warning' : 'info'"
            :effect="newTaskPriority === 'medium' ? 'dark' : 'plain'"
            class="priority-tag"
            @click="newTaskPriority = 'medium'"
          >
            {{ t('task.priority.medium') }}
          </el-tag>
          <el-tag
            :type="newTaskPriority === 'high' ? 'danger' : 'info'"
            :effect="newTaskPriority === 'high' ? 'dark' : 'plain'"
            class="priority-tag"
            @click="newTaskPriority = 'high'"
          >
            {{ t('task.priority.high') }}
          </el-tag>
        </div>
      </div>
      <template #footer>
        <el-button @click="dialogVisible = false">{{ t('common.cancel') }}</el-button>
        <el-button type="primary" @click="saveTask">{{ isEditing ? t('task.save') : t('task.add') }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped>
.task-list {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.task-header {
  margin-bottom: 20px;
      text-align: left;
}

.task-header h2 {
  margin: 0;
  color: var(--el-text-color-primary);
}

.task-items {
  flex: 1;
  overflow-y: auto;
}

.task-item {
  margin-bottom: 12px;
}

.task-item.completed .task-title {
  text-decoration: line-through;
  color: var(--el-text-color-disabled);
}

.task-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.task-title {
  flex: 1;
  font-size: 14px;
  color: var(--el-text-color-secondary);
  text-align: left;
  cursor: pointer;
}

.task-title:hover {
  color: var(--el-color-primary);
}

.add-float-btn {
  position: fixed;
  right: 40px;
  bottom: 40px;
  z-index: 100;
}

.dialog-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.priority-tags {
  display: flex;
  gap: 12px;
}

.priority-tag {
  cursor: pointer;
  user-select: none;
}

.completed-toggle {
  text-align: center;
  padding: 12px 0;
}
</style>