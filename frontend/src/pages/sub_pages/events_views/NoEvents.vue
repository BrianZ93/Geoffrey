<template>
  <div class="q-pa-md flex-container">
    <CreateEventsDialog />

    <EmptyPage
      :onClick="handleCreateEventDialogOpen"
      label="No events yet, click below to create an event."
    />
  </div>
</template>

<script setup lang="ts">
import CreateEventsDialog from './../../../components/events/CreateEventDialog.vue';
import { useAppStateStore } from 'src/stores/main-application-state';
import EmptyPage from './../../../components/styled_objects/EmptyPage.vue';

const appState = useAppStateStore();

const handleCreateEventDialogOpen = () => {
  appState.addEventDialogOpen = true;
};
</script>
