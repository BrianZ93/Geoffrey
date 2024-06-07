<template>
  <q-list>
    <q-item
      v-for="(item, index) in menuItems"
      :key="item.label"
      clickable
      v-ripple
      :class="{ 'sidebar-tab-active': isActivePage(index) }"
      @click="setActivePage(index)"
    >
      <q-item-section avatar>
        <q-icon :name="item.icon" />
      </q-item-section>
      <q-item-section>
        {{ item.label }}
      </q-item-section>
    </q-item>
  </q-list>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAppStateStore } from 'src/stores/main-application-state';

const appState = useAppStateStore();

const menuItems = ref([
  { label: 'Dashboard', icon: 'schema' },
  { label: 'Finances', icon: 'paid' },
  { label: 'Events', icon: 'celebration' },
  { label: 'Calendar', icon: 'calendar_month' },
  {
    label: 'Home Maintenance',
    icon: 'construction',
  },
]);

const setActivePage = (index: number) => {
  appState.activePage = index;
};

const isActivePage = (index: number) => {
  return appState.activePage === index;
};
</script>

<style lang="scss">
/* Add your custom styles here */
.bg-grey-3 {
  background-color: #e0e0e0;
}
</style>
