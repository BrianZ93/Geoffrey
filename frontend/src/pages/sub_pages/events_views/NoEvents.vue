<template>
  <div class="q-pa-md flex-container">
    <CreateEventsDialog />

    <div class="flex-center">
      <div class="spinning-object">
        <!-- Replace this with your desired object -->
        <q-icon name="donut_large" size="64px" />
      </div>
      <h4>No events yet, click below to create an event</h4>
      <q-btn
        round
        color="secondary"
        icon="add"
        @click="handleCreateEventDialogOpen"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import CreateEventsDialog from './../../../components/events/CreateEventDialog.vue';
import { useAppStateStore } from 'src/stores/main-application-state';

const appState = useAppStateStore();

const handleCreateEventDialogOpen = () => {
  appState.addEventDialogOpen = true;
};
</script>

<style scoped>
.flex-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  text-align: center;
  flex-direction: column;
}

.flex-center {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.spinning-object {
  animation: spin 15s linear infinite;
  margin-bottom: 20px;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
