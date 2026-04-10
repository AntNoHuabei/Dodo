<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { Plus } from '@element-plus/icons-vue';
import { ElCard, ElCheckbox, ElButton, ElInput, ElTag, ElEmpty, ElDialog, ElSelect, ElOption, ElMessageBox } from 'element-plus';
import { Open, Execute, Query } from '../../bindings/github.com/wailsapp/wails/v3/pkg/services/sqlite/sqliteservice';

const { t } = useI18n();

interface Task {
  id: number;
  number: number;
  title: string;
  completed: boolean;
  priority: 'low' | 'medium' | 'high';
  type: 'bug' | 'feature' | 'refactor' | 'performance' | 'security' | 'chore' | 'migration' | 'ci' | 'test' | 'review' | 'doc' | 'subtask';
}

const dbReady = ref(false);

const initDatabase = async () => {
  try {
    await Open();
    // 创建表（如果不存在）
    await Execute(`
      CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        number INTEGER DEFAULT 0,
        title TEXT NOT NULL,
        completed INTEGER DEFAULT 0,
        priority TEXT DEFAULT 'medium',
        type TEXT DEFAULT 'feature'
      )
    `);
    dbReady.value = true;
    await loadTasks();
  } catch (error) {
    console.error('数据库初始化失败:', error);
  }
};

const loadTasks = async () => {
  if (!dbReady.value) return;
  try {
    // 尝试查询包含 number 和 type 字段的任务
    let rows;
    try {
      rows = await Query('SELECT id, number, title, completed, priority, type FROM tasks ORDER BY id DESC') as any[];
    } catch (e) {
      
    }
    tasks.value = rows.map(row => ({
      id: row.id,
      number: row.number || 0,
      title: row.title,
      completed: Boolean(row.completed),
      priority: row.priority as 'low' | 'medium' | 'high',
      type: (row.type || 'feature') as 'bug' | 'feature' | 'refactor' | 'performance' | 'security' | 'chore' | 'migration' | 'ci' | 'test' | 'review' | 'doc' | 'subtask'
    }));
  } catch (error) {
    console.error('加载任务失败:', error);
  }
};

const tasks = ref<Task[]>([]);

onMounted(() => {
  initDatabase();
});

const newTaskTitle = ref('');
const dialogVisible = ref(false);
const newTaskPriority = ref<'low' | 'medium' | 'high'>('medium');
const newTaskType = ref<'bug' | 'feature' | 'refactor' | 'performance' | 'security' | 'chore' | 'migration' | 'ci' | 'test' | 'review' | 'doc' | 'subtask'>('feature');
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
    newTaskType.value = task.type;
  } else {
    editingTask.value = null;
    newTaskTitle.value = '';
    newTaskPriority.value = 'medium';
    newTaskType.value = 'feature';
  }
  dialogVisible.value = true;
};

const saveTask = async () => {
  if (newTaskTitle.value.trim()) {
    if (editingTask.value) {
      editingTask.value.title = newTaskTitle.value.trim();
      editingTask.value.priority = newTaskPriority.value;
      editingTask.value.type = newTaskType.value;
      if (dbReady.value) {
        await Execute('UPDATE tasks SET title = ?, priority = ?, type = ? WHERE id = ?', 
          newTaskTitle.value.trim(), newTaskPriority.value, newTaskType.value, editingTask.value.id);
      }
    } else {
      // 生成新的编号
      let newNumber = 1;
      if (dbReady.value) {
        const result = await Query('SELECT MAX(number) as maxNumber FROM tasks') as any[];
        if (result.length > 0 && result[0].maxNumber) {
          newNumber = result[0].maxNumber + 1;
        }
      }
      const newId = Date.now();
      tasks.value.push({
        id: newId,
        number: newNumber,
        title: newTaskTitle.value.trim(),
        completed: false,
        priority: newTaskPriority.value,
        type: newTaskType.value
      });
      if (dbReady.value) {
        await Execute('INSERT INTO tasks (id, number, title, completed, priority, type) VALUES (?, ?, ?, 0, ?, ?)',
          newId, newNumber, newTaskTitle.value.trim(), newTaskPriority.value, newTaskType.value);
      }
    }
    dialogVisible.value = false;
  }
};

const toggleTask = async (task: Task) => {
  task.completed = !task.completed;
  if (dbReady.value) {
    await Execute('UPDATE tasks SET completed = ? WHERE id = ?', task.completed ? 1 : 0, task.id);
  }
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
    if (dbReady.value) {
      await Execute('DELETE FROM tasks WHERE id = ?', id);
    }
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

const getTypeColor = (type: string): { backgroundColor: string; textColor: string } => {
  const colors: Record<string, { backgroundColor: string; textColor: string }> = {
    bug: { backgroundColor: '#F56C6C', textColor: '#FFFFFF' }, // 红色背景，白色文本
    feature: { backgroundColor: '#409EFF', textColor: '#FFFFFF' }, // 蓝色背景，白色文本
    refactor: { backgroundColor: '#E6A23C', textColor: '#FFFFFF' }, // 黄色背景，白色文本
    performance: { backgroundColor: '#67C23A', textColor: '#FFFFFF' }, // 绿色背景，白色文本
    security: { backgroundColor: '#909399', textColor: '#FFFFFF' }, // 灰色背景，白色文本
    chore: { backgroundColor: '#909399', textColor: '#FFFFFF' }, // 灰色背景，白色文本
    migration: { backgroundColor: '#E6A23C', textColor: '#FFFFFF' }, // 黄色背景，白色文本
    ci: { backgroundColor: '#409EFF', textColor: '#FFFFFF' }, // 蓝色背景，白色文本
    test: { backgroundColor: '#67C23A', textColor: '#FFFFFF' }, // 绿色背景，白色文本
    review: { backgroundColor: '#909399', textColor: '#FFFFFF' }, // 灰色背景，白色文本
    doc: { backgroundColor: '#909399', textColor: '#FFFFFF' }, // 灰色背景，白色文本
    subtask: { backgroundColor: '#E6A23C', textColor: '#FFFFFF' } // 黄色背景，白色文本
  };
  return colors[type] || { backgroundColor: '#909399', textColor: '#FFFFFF' };
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
          <span class="task-number">{{ "#"+task.number }}</span>
          <el-checkbox
            :model-value="task.completed"
            @change="toggleTask(task)"
          />
          <span 
            class="task-title" 
            @dblclick="openDialog(task)"
            :class="`priority-${task.priority}`"
          >{{ task.title }}</span>
          <el-tag 
            :color="getTypeColor(task.type).backgroundColor" 
            :style="{ color: getTypeColor(task.type).textColor }" 
            size="small"
          >{{ t(`task.type.${task.type}`) }}</el-tag>
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
        <div class="form-group">
          <label>{{ t('task.type.label') }}</label>
          <el-select v-model="newTaskType" class="w-full">
            <el-option :label="t('task.type.bug')" value="bug" />
            <el-option :label="t('task.type.feature')" value="feature" />
            <el-option :label="t('task.type.refactor')" value="refactor" />
            <el-option :label="t('task.type.performance')" value="performance" />
            <el-option :label="t('task.type.security')" value="security" />
            <el-option :label="t('task.type.chore')" value="chore" />
            <el-option :label="t('task.type.migration')" value="migration" />
            <el-option :label="t('task.type.ci')" value="ci" />
            <el-option :label="t('task.type.test')" value="test" />
            <el-option :label="t('task.type.review')" value="review" />
            <el-option :label="t('task.type.doc')" value="doc" />
            <el-option :label="t('task.type.subtask')" value="subtask" />
          </el-select>
        </div>
        <div class="form-group">
          <label>{{ t('task.priority.label') }}</label>
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

.task-content .el-button {
  opacity: 0;
  transition: opacity 0.3s;
}

.task-item:hover .task-content .el-button {
  opacity: 1;
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

.task-title.priority-high {
  color: var(--el-color-danger);
}

.task-title.priority-medium {
  color: var(--el-color-warning);
}

.task-title.priority-low {
  color: var(--el-color-success);
}



.task-id {
  font-size: 12px;
  color: var(--el-text-color-disabled);
  margin-right: 8px;
}

.task-number {
  width: 24px;
  text-align: center;
  font-size: 12px;
  color: var(--el-text-color-primary);
  margin-right: 8px;
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

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  font-size: 14px;
  color: var(--el-text-color-secondary);
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