<script setup lang="ts">
import { ref, computed } from 'vue';
import { useI18n } from 'vue-i18n';
import { ElCard, ElButton, ElButtonGroup } from 'element-plus';

const { t } = useI18n();

const currentDate = ref(new Date());

const weekDays = computed(() => {
  const days: { date: Date; day: number; isToday: boolean }[] = [];
  const startOfWeek = new Date(currentDate.value);
  const day = startOfWeek.getDay();
  const diff = startOfWeek.getDate() - day + (day === 0 ? -6 : 1);
  startOfWeek.setDate(diff);

  const today = new Date();
  today.setHours(0, 0, 0, 0);

  for (let i = 0; i < 7; i++) {
    const date = new Date(startOfWeek);
    date.setDate(startOfWeek.getDate() + i);
    const dateOnly = new Date(date);
    dateOnly.setHours(0, 0, 0, 0);
    days.push({
      date,
      day: date.getDate(),
      isToday: dateOnly.getTime() === today.getTime()
    });
  }
  return days;
});

const weekRange = computed(() => {
  const start = weekDays.value[0].date;
  const end = weekDays.value[6].date;
  const startMonth = start.getMonth() + 1;
  const endMonth = end.getMonth() + 1;
  const startDate = start.getDate();
  const endDate = end.getDate();

  if (startMonth === endMonth) {
    return `${startMonth}${t('weekly.month')} ${startDate} - ${endDate}`;
  } else {
    return `${startMonth}${t('weekly.month')}${startDate} - ${endMonth}${t('weekly.month')}${endDate}`;
  }
});

interface ScheduleItem {
  id: number;
  title: string;
  time: string;
}

const schedules = ref<Record<number, ScheduleItem[]>>({
  0: [
    { id: 1, title: '团队周会', time: '09:00 - 10:00' },
    { id: 2, title: '项目评审', time: '14:00 - 15:00' }
  ],
  1: [
    { id: 3, title: '代码审查', time: '10:00 - 11:00' }
  ],
  2: [
    { id: 4, title: '技术分享', time: '15:00 - 16:00' }
  ],
  3: [
    { id: 5, title: '一对一会议', time: '11:00 - 11:30' }
  ],
  4: [
    { id: 6, title: '周报总结', time: '16:00 - 17:00' }
  ],
  5: [],
  6: []
});

const getDayName = (index: number) => {
  const days = [
    t('weekly.mon'),
    t('weekly.tue'),
    t('weekly.wed'),
    t('weekly.thu'),
    t('weekly.fri'),
    t('weekly.sat'),
    t('weekly.sun')
  ];
  return days[index];
};

const previousWeek = () => {
  const newDate = new Date(currentDate.value);
  newDate.setDate(newDate.getDate() - 7);
  currentDate.value = newDate;
};

const nextWeek = () => {
  const newDate = new Date(currentDate.value);
  newDate.setDate(newDate.getDate() + 7);
  currentDate.value = newDate;
};

const goToToday = () => {
  currentDate.value = new Date();
};
</script>

<template>
  <div class="weekly-view">
    <Teleport to="#page-title-target">
      <span>{{ t('weekly.title') }}</span>
    </Teleport>
    <div class="weekly-header">
     
      <div class="week-navigation">
        <el-button-group>
          <el-button @click="previousWeek">{{ t('weekly.previous') }}</el-button>
          <el-button @click="goToToday">{{ t('weekly.today') }}</el-button>
          <el-button @click="nextWeek">{{ t('weekly.next') }}</el-button>
        </el-button-group>
        <span class="week-range">{{ weekRange }}</span>
      </div>
    </div>

    <div class="week-grid">
      <div
        v-for="(day, index) in weekDays"
        :key="index"
        class="day-column"
        :class="{ today: day.isToday }"
      >
        <div class="day-header">
          <span class="day-name">{{ getDayName(index) }}</span>
          <span class="day-date" :class="{ today: day.isToday }">{{ day.day }}</span>
        </div>
        <div class="day-schedules">
          <el-card
            v-for="item in schedules[index] || []"
            :key="item.id"
            class="schedule-item"
            shadow="hover"
          >
            <div class="schedule-time">{{ item.time }}</div>
            <div class="schedule-title">{{ item.title }}</div>
          </el-card>
          <div v-if="!schedules[index]?.length" class="no-schedule">
            {{ t('weekly.noSchedule') }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.weekly-view {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.weekly-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.weekly-header h2 {
  margin: 0;
  color: #303133;
}

.week-navigation {
  display: flex;
  align-items: center;
  gap: 16px;
}

.week-range {
  font-size: 14px;
  color: #606266;
}

.week-grid {
  flex: 1;
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 8px;
  overflow-x: auto;
}

.day-column {
  background: #fff;
  border-radius: 8px;
  min-width: 120px;
  display: flex;
  flex-direction: column;
}

.day-column.today {
  background: #ecf5ff;
  border: 2px solid #409eff;
}

.day-header {
  padding: 12px 8px;
  text-align: center;
  border-bottom: 1px solid #ebeef5;
}

.day-name {
  display: block;
  font-size: 12px;
  color: #909399;
  margin-bottom: 4px;
}

.day-date {
  display: inline-block;
  width: 28px;
  height: 28px;
  line-height: 28px;
  border-radius: 50%;
  font-size: 14px;
  font-weight: 500;
  color: #303133;
}

.day-date.today {
  background: #409eff;
  color: #fff;
}

.day-schedules {
  flex: 1;
  padding: 8px;
  overflow-y: auto;
}

.schedule-item {
  margin-bottom: 8px;
}

.schedule-item :deep(.el-card__body) {
  padding: 10px;
}

.schedule-time {
  font-size: 12px;
  color: #909399;
  margin-bottom: 4px;
}

.schedule-title {
  font-size: 13px;
  color: #303133;
  font-weight: 500;
}

.no-schedule {
  text-align: center;
  color: #c0c4cc;
  font-size: 12px;
  padding: 20px 0;
}
</style>