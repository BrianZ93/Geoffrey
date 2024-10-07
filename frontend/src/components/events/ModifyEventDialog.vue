<template>
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
          <div class="text-h6 text-center">Modify Event</div>
        </q-card-section>

        <q-separator dark />

        <q-card-section class="q-pt-none">
          <!-- Event Title -->
          <string-input
            label="Event Title"
            :onValueChange="handleEventNameChange"
            :initialValue="activeEvent.title"
          />
          <!-- Event Date -->
          <date-picker
            label="Event Date"
            :onUpdate="handleDateUpdate"
            :isRange="true"
            :initial-date="initialDateValue"
          />
        </q-card-section>
      </q-card-section>

      <form-actions :onCancel="handleCancel" :onModify="onModify" />
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { Notify } from 'quasar';
import { ref, computed } from 'vue';
import { useAppStateStore } from '../../stores/main-application-state';
import { useEventsStore, EventsLoadState } from '../../stores/events-state';
import { Event } from '../../models/events/event';

import { updateEvent } from '../../api/events/modify_event';
import { getEvents } from '../../api/events/get_events';
import DatePicker from '../../components/styled_objects/DatePicker.vue';
import StringInput from '../../components/styled_objects/StringInput.vue';
import FormActions from '../../components/styled_objects/FormActions.vue';
import { handleUpdateDates } from '../../components/styled_objects/helpers/DatePicker';
import { handleOnValueChange } from '../../components/styled_objects/helpers/StringInput';

const appState = useAppStateStore();
const eventsState = useEventsStore();

const persistent = computed(() => appState.modifyEventDialogOpen);
const activeEvent = eventsState.activeEvent as Event;

const eventTitle = ref<string | null>(activeEvent.title);
const startDate = ref<string | null>(activeEvent.startTime);
const endDate = ref<string | null>(activeEvent.endTime);

const initialDateValue = computed(() => {
  return { from: activeEvent.startTime, to: activeEvent.endTime };
});

const handleEventNameChange = (value: string | null) => {
  handleOnValueChange(value, eventTitle);
};

const handleDateUpdate = (date: { from: string; to: string } | string) => {
  handleUpdateDates(date, startDate, endDate);
};

const handleModifyEvent = async (
  title: string,
  startDate: string,
  endDate: string
) => {
  const event = new Event(
    activeEvent.id,
    title,
    startDate,
    endDate,
    activeEvent.guests,
    activeEvent.venues
  );
  await updateEvent(event);
};

const onModify = async () => {
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
      await handleModifyEvent(eventTitle.value, startDate.value, endDate.value);

      eventsState.loadState = EventsLoadState.loading;

      Notify.create({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Modifying Event',
        type: 'loading',
      });

      const fetchedEvents = await getEvents();

      Notify.create({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Event Modified',
      });

      eventsState.loadState = EventsLoadState.hasEvents;
      if (typeof fetchedEvents !== 'string') {
        eventsState.events = fetchedEvents;
      }
      appState.modifyEventDialogOpen = false;
    }
  }
};

// const handleDeleteEvent = async () => {
//   appState.modifyEventDialogOpen = false;
// };

const handleCancel = () => {
  appState.modifyEventDialogOpen = false;
};
</script>
../../api/events/create_event ../styled_objects/helpers/Input
