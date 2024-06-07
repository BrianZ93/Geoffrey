<template>
  <div class="q-pa-md q-gutter-sm">
    <q-dialog
      v-model="persistent"
      persistent
      transition-show="scale"
      transition-hide="scale"
    >
      <q-card
        class="bg-secondary text-white dialog-card-layout"
        style="width: 50vw; height: 50vh"
      >
        <q-card-section>
          <q-card-section>
            <div class="text-h6 text-center">Add a New Event</div>
          </q-card-section>

          <q-separator dark />

          <q-card-section class="q-pt-none">
            <!-- Event Title -->
            <string-input
              label="Event Title"
              :onValueChange="handleEventNameChange"
            />
            <!-- Event Date -->
            <date-picker label="Event Date" :onUpdate="handleDateUpdate" />
          </q-card-section>
        </q-card-section>

        <form-actions :onSubmit="onSubmit" :onCancel="handleCancel" />
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { Notify } from 'quasar';
import { ref, computed } from 'vue';
import { useAppStateStore } from 'src/stores/main-application-state';
import { useEventsStore, EventsLoadState } from 'src/stores/events-state';
import { Event } from 'src/models/events/event';

import { createEvent } from '../../api/events/create_table';
import { getEvents } from '../../api/events/get_events';
import DatePicker from '../../components/styled_objects/DatePicker.vue';
import StringInput from '../../components/styled_objects/StringInput.vue';
import FormActions from '../../components/styled_objects/FormActions.vue';
import { handleUpdateDates } from '../../components/styled_objects/helpers/DatePicker';
import { handleOnValueChange } from '../../components/styled_objects/helpers/StringInput';

const appState = useAppStateStore();
const eventsState = useEventsStore();

const persistent = computed(() => appState.addEventDialogOpen);

const eventTitle = ref<string | null>('');
const startDate = ref<string | null>(null);
const endDate = ref<string | null>(null);

const handleEventNameChange = (value: string | null) => {
  handleOnValueChange(value, eventTitle);
};

const handleDateUpdate = (date: { from: string; to: string } | string) => {
  handleUpdateDates(date, startDate, endDate);
};

const handleCreateEvent = async (
  title: string,
  startDate: string,
  endDate: string
) => {
  const event = new Event('0', title, startDate, endDate, [], []);
  await createEvent(event);
};

const onSubmit = async () => {
  console.log(startDate.value);
  console.log(typeof startDate.value);
  if (!eventTitle.value || !startDate.value || !endDate.value) {
    let emptyFields = [] as string[];
    if (!eventTitle.value) {
      emptyFields.push('Event Title');
    }

    if (!startDate.value) {
      emptyFields.push('Start Date');
    }

    if (!endDate.value) {
      emptyFields.push('End Date');
    }

    for (let field of emptyFields) {
      Notify.create({
        message: `You need to provide a(n) ${field} before submitting`,
        icon: 'warning',
        textColor: 'white',
        color: 'red-5',
      });
    }
  } else {
    const start = new Date(startDate.value);
    const end = new Date(endDate.value);

    if (start > end) {
      Notify.create({
        message: 'Start Date must be earlier than End Date',
        icon: 'warning',
        textColor: 'white',
        color: 'red-5',
      });
    } else {
      await handleCreateEvent(eventTitle.value, startDate.value, endDate.value);

      eventsState.loadState = EventsLoadState.loading;

      Notify.create({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Creating Event',
        type: 'loading',
      });

      const fetchedEvents = await getEvents();

      Notify.create({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Event Created',
      });

      eventsState.loadState = EventsLoadState.hasEvents;
      if (typeof fetchedEvents !== 'string') {
        eventsState.events = fetchedEvents;
      }
    }
  }
};

const handleCancel = () => {
  appState.addEventDialogOpen = false;
};
</script>
