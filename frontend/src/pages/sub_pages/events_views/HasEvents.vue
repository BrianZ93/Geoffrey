<template>
  <div class="q-pa-md">
    <create-events-dialog v-if="appState.addEventDialogOpen" />
    <modify-events-dialog v-if="appState.modifyEventDialogOpen" />
    <add-guest-dialog v-if="appState.addGuestDialogOpen" />

    <div class="row section header-row">
      <!-- First row content -->
      <q-toolbar class="bg-secondary">
        <drop-down-list
          :listItems="events"
          icon="celebration"
          :onSelect="handleSelectEvent"
          :onModify="handleOpenModifyEventDialog"
          :activeItem="convertActiveEventToListItem"
          width="25"
        />

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
          <div class="big-number">
            {{ box.number }}
          </div>
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
import ModifyEventsDialog from './../../../components/events/ModifyEventDialog.vue';

import DropDownList from '../../../components/styled_objects/DropdownList.vue';
import { ListItem } from 'src/components/styled_objects/helpers/DropdownList';

import AddGuestDialog from '../../../components/events/AddGuestDialog.vue';
import { useAppStateStore } from 'src/stores/main-application-state';

const appState = useAppStateStore();
const eventsState = useEventsStore();

const remainingRSVPs = ref<number>(0);
const confirmedGuests = ref<number>(0);
const venues = ref<number>(0);
const daysUntilEvent = ref<number>(0);

const events = computed(() => {
  let listItems = [] as ListItem[];
  const events = eventsState.events as Event[];
  for (let event of events) {
    const new_item = new ListItem(event.id, event.title, 'caption');
    listItems.push(new_item);
  }
  return listItems;
});

if (eventsState.events.length > 0) {
  eventsState.activeEvent = eventsState.events[0];
}

const boxes = ref([
  {
    icon: 'schedule_send',
    number: remainingRSVPs.value,
    caption: 'Remaining RSVPs',
  },
  { icon: 'group', number: confirmedGuests.value, caption: 'Confirmed Guests' },
  { icon: 'location_on', number: venues.value, caption: 'Venues' },
  {
    icon: 'schedule',
    number: daysUntilEvent.value,
    caption: 'Days Until Event',
  },
]);

const handleCreateEventDialogOpen = () => {
  appState.addEventDialogOpen = true;
};

const handleOpenModifyEventDialog = () => {
  appState.modifyEventDialogOpen = true;
};

const calculateEventValues = (selectedEvent: Event) => {
  let remainingRsvpsValue = 0;
  let confirmedGuestsValue = 0;

  const eventDate = new Date(selectedEvent.startTime);
  const currentDate = new Date();

  // Calculate the difference in time (in milliseconds)
  const timeDifference = eventDate.getTime() - currentDate.getTime();

  daysUntilEvent.value = Math.floor(timeDifference / (1000 * 60 * 60 * 24));
  let daysUntilEventText = '';

  if (daysUntilEvent.value === 0) {
    daysUntilEventText = 'The Day is Here!';
  } else if (daysUntilEvent.value < 0) {
    daysUntilEventText = 'The Event Date Has Passed';
  }

  // Calculate guest data
  if (selectedEvent.guests) {
    for (let guest of selectedEvent.guests) {
      if (guest.attending) {
        confirmedGuestsValue += 1;
      }

      if (!guest.rsvpReceived) {
        remainingRsvpsValue += 1;
      }
    }
  }

  remainingRSVPs.value = remainingRsvpsValue;
  confirmedGuests.value = confirmedGuestsValue;

  // Update boxes without re-wrapping in ref
  boxes.value = [
    {
      icon: 'schedule_send',
      number: remainingRSVPs.value,
      caption: 'Remaining RSVPs',
    },
    {
      icon: 'group',
      number: confirmedGuests.value,
      caption: 'Confirmed Guests',
    },
    { icon: 'location_on', number: venues.value, caption: 'Venues' },
    {
      icon: 'schedule',
      number: daysUntilEvent.value,
      caption: daysUntilEventText,
    },
  ];
};

if (eventsState.activeEvent) {
  calculateEventValues(eventsState.activeEvent);
}

const handleSelectEvent = (val: ListItem) => {
  const selectedEventId = val.id;
  const selectedEvent = eventsState.events.find(
    (event: Event) => event.id === selectedEventId
  );
  if (selectedEvent) {
    eventsState.activeEvent = selectedEvent;
    console.log('Selected Event:', selectedEvent);
    calculateEventValues(selectedEvent);
  } else {
    console.log('Event not found');
  }
};

const convertActiveEventToListItem = computed(() => {
  const activeEvent = eventsState.activeEvent;
  const convertedListItem = new ListItem(
    activeEvent.id,
    activeEvent.title,
    'caption'
  );

  return convertedListItem;
});
</script>
