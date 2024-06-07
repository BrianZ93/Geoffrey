<template>
  <div class="q-pa-md">
    <create-events-dialog />
    <add-guest-dialog />

    <div class="row section events-header-row">
      <!-- First row content -->
      <q-toolbar class="bg-secondary">
        <q-avatar><q-icon name="celebration" size="md"></q-icon></q-avatar>
        <drop-down-list :listItems="events" />

        <q-toolbar-title class="bg-secondary text-white"
          >Events</q-toolbar-title
        >
        <q-btn
          glossy
          dense
          icon="add"
          color="primary"
          label="Add Event"
          @click="handleCreateEventDialogOpen"
        />
      </q-toolbar>
    </div>
    <div class="row section events-main-content-row q-pa-md">
      <div class="event-boxes-container">
        <div class="event-box" v-for="box in boxes" :key="box.caption">
          <q-icon :name="box.icon" size="64px" />
          <div class="big-number">{{ box.number }}</div>
          <div class="caption">{{ box.caption }}</div>
        </div>
      </div>
      <div class="events-table-container">
        <guests-table />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Event } from 'src/models/events/event';
import { ref, computed } from 'vue';
import { useEventsStore } from 'src/stores/events-state';
import GuestsTable from '../../../components/events/GuestsTable.vue';
import CreateEventsDialog from './../../../components/events/CreateEventDialog.vue';

import {
  DropDownList,
  ListItem,
} from '../../../components/styled_objects/DropdownList.vue';

import AddGuestDialog from '../../../components/events/AddGuestDialog.vue';
import { useAppStateStore } from 'src/stores/main-application-state';

const appState = useAppStateStore();
const eventsState = useEventsStore();

const listItems = ref<ListItem[]>([]);
const events = computed(() => eventsState.events as Event[]);

const selectedEvent = ref(null);
const eventOptions = ref<{ label: string; value: string }[]>([]);

const boxes = ref([
  { icon: 'schedule_send', number: 10, caption: 'Remaining RSVPs' },
  { icon: 'group', number: 50, caption: 'Confirmed Guests' },
  { icon: 'location_on', number: 2, caption: 'Venues' },
  { icon: 'schedule', number: 20, caption: 'Days Until Event' },
]);

const handleCreateEventDialogOpen = () => {
  appState.addEventDialogOpen = true;
};

const filterFn = (val: string, update: (callback: () => void) => void) => {
  update(() => {
    if (val === '') {
      eventOptions.value = events.value.map((event) => ({
        label: event.title,
        value: event.id,
      }));
    } else {
      const needle = val.toLowerCase();
      eventOptions.value = events.value
        .filter((event) => event.title.toLowerCase().includes(needle))
        .map((event) => ({ label: event.title, value: event.id }));
    }

    const selected =
      events.value.find((event) => event.id === selectedEvent.value) || null;

    if (selected !== null) {
      eventsState.activeEvent = selected;
      console.log(selected);
    }
  });
};
</script>
