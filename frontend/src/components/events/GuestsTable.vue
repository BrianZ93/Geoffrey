<template>
  <div class="q-pa-md">
    <confirmation-dialog
      :cancel="cancelDeleteGuest"
      :confirm="confirmDeleteGuest"
      :message="confirmationDialogString"
      modelValue="delete_guest"
    />

    <boolean-radio-box
      v-if="rsvpUpdateRadioVisible"
      :headerLabel="activeGuestName"
      trueLabel="Attending"
      falseLabel="Regrets"
      :onSubmit="(state: boolean) => updateAttending(state, activeGuest)"
    />

    <q-table
      flat
      bordered
      title="Guests"
      :rows="rows"
      :columns="columns"
      row-key="id"
      :filter="filter"
      binary-state-sort
      class="events-table"
      dense
    >
      <!-- Search bar slot -->
      <template v-slot:top-right>
        <q-input
          borderless
          dense
          debounce="300"
          v-model="filter"
          placeholder="Search"
        >
          <template v-slot:append>
            <q-icon name="search" />
            <q-btn
              icon="add"
              color="primary"
              label="Add Guest"
              glossy
              @click="handleOpenAddGuestDialog"
            ></q-btn>
          </template>
        </q-input>
      </template>
      <template v-slot:body="props">
        <q-tr :props="props" v-if="props.row">
          <q-td key="name" :props="props" align="left">
            {{ props.row.name }}
            <q-popup-edit
              v-model="props.row.name"
              @before-hide="updateGuest(props.row)"
            >
              <q-input v-model="props.row.name" dense autofocus counter />
            </q-popup-edit>
          </q-td>
          <q-td key="email" :props="props" align="left">
            {{ props.row.email }}
            <q-popup-edit
              v-model="props.row.email"
              @before-hide="updateGuest(props.row)"
            >
              <q-input v-model="props.row.email" dense autofocus counter />
            </q-popup-edit>
          </q-td>
          <q-td key="phoneNumber" :props="props" align="left">
            {{ props.row.phoneNumber }}
            <q-popup-edit
              v-model="props.row.phoneNumber"
              @before-hide="updateGuest(props.row)"
            >
              <q-input
                v-model="props.row.phoneNumber"
                dense
                autofocus
                counter
              />
            </q-popup-edit>
          </q-td>
          <q-td
            key="attending"
            :props="props"
            align="left"
            @click="handleUpdateRSVPDialog(props.row)"
          >
            {{ props.row.attending ? 'Yes' : 'No' }}
          </q-td>
          <q-td key="rsvpReceived" :props="props" align="left">
            {{ props.row.rsvpReceived ? 'Yes' : 'No' }}
          </q-td>
          <q-td key="delete" :props="props" align="center">
            <q-btn
              flat
              dense
              round
              icon="delete_forever"
              color="white"
              @click="openConfirmDeleteGuestDialog(props.row)"
            />
          </q-td>
        </q-tr>
        <q-tr v-if="rows.length === 0">
          <q-td colspan="100%" class="text-center"> No guests available </q-td>
        </q-tr>
      </template>
      <!-- No Data Slot -->
      <template v-slot:no-data="{ icon, message }">
        <div class="full-width row flex-center text-tertiary q-gutter-sm">
          <q-icon size="2em" name="sentiment_dissatisfied" />
          <span>
            No guests have been added to your event yet... {{ message }}
          </span>
          <q-icon size="2em" :name="filter ? 'filter_b_and_w' : icon" />
        </div>
      </template>
    </q-table>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { Notify, QPopupEdit } from 'quasar';
import { Guest } from 'src/models/events/guest';
import { Event } from 'src/models/events/event';
import { useEventsStore } from 'src/stores/events-state';
import { useAppStateStore } from 'src/stores/main-application-state';
import { modifyGuest } from './../../api/events/modify_guest';
import { deleteGuest } from 'src/api/events/delete_guest';
import { getEvents } from 'src/api/events/get_events';
import ConfirmationDialog from '../styled_objects/ConfirmationDialog.vue';
import BooleanRadioBox from '../styled_objects/BooleanRadioBox.vue';

const appState = useAppStateStore();
const eventsState = useEventsStore();
const events = computed(() => eventsState.events);

const columns = ref([
  {
    name: 'name',
    label: 'Name',
    field: (row: Guest) => row.name,
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'email',
    label: 'Email',
    field: (row: Guest) => row.email,
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'phoneNumber',
    label: 'Phone Number',
    field: (row: Guest) => row.phoneNumber,
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'attending',
    label: 'Attending',
    field: (row: Guest) => (row.attending ? 'Yes' : 'No'),
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'rsvpReceived',
    label: 'RSVP Received',
    field: (row: Guest) => (row.rsvpReceived ? 'Yes' : 'No'),
    sortable: true,
    align: 'left' as const,
  },
  {
    name: 'delete',
    label: '',
    field: 'delete',
    align: 'center' as const,
  },
]);

const filter = ref('');

const rows = computed(() => {
  // Extract guests from events and handle cases where there are no guests
  return events.value.flatMap((event: Event) => event.guests || []);
});

const handleOpenAddGuestDialog = () => {
  appState.addGuestDialogOpen = true;
};

const updateGuest = async (guest: Guest) => {
  await modifyGuest(eventsState.activeEvent.id, guest.id, guest);

  getEvents();
};

const confirmationDialogmodelValue = 'delete_guest';
const confirmationDialogString = ref<string>('deleting guest: ');
const confirmationDialogGuest = ref<Guest | null>(null);

const openConfirmDeleteGuestDialog = (guest: Guest) => {
  confirmationDialogString.value = 'delete ' + guest.name;
  confirmationDialogGuest.value = guest;
  appState.confirmationDialogOpen[confirmationDialogmodelValue] = true;
};

const cancelDeleteGuest = () => {
  appState.confirmationDialogOpen[confirmationDialogmodelValue] = false;
};

const confirmDeleteGuest = async () => {
  if (confirmationDialogGuest.value !== null) {
    await deleteGuest(
      eventsState.activeEvent.id,
      confirmationDialogGuest.value.id
    );
    appState.confirmationDialogOpen[confirmationDialogmodelValue] = false;

    Notify.create({
      color: 'green-4',
      textColor: 'white',
      icon: 'cloud_done',
      message: 'Guest Removed',
    });

    const fetchedEvents = await getEvents();
    if (typeof fetchedEvents !== 'string') {
      eventsState.events = fetchedEvents;
    }
  }
};

// RSVP Related Variables
const activeGuestName = ref<string>('');
const activeGuest = ref<Guest | null>(null);

const rsvpUpdateRadioVisible = ref<boolean>(false);

const handleUpdateRSVPDialog = (guest: Guest) => {
  activeGuest.value = guest;
  activeGuestName.value = guest.name;
  rsvpUpdateRadioVisible.value = true;
};

const updateAttending = async (state: boolean, guest: Guest | null) => {
  rsvpUpdateRadioVisible.value = false;
  if (guest) {
    guest.attending = state;
    guest.rsvpReceived = true;
    await modifyGuest(eventsState.activeEvent.id, guest.id, guest);

    const fetchedEvents = await getEvents();
    if (typeof fetchedEvents !== 'string') {
      eventsState.events = fetchedEvents;
    }
  }
};
</script>
