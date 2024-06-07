<template>
  <div class="q-pa-md">
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
          <q-td key="id" :props="props">
            {{ props.row.id }}
          </q-td>
          <q-td key="name" :props="props">
            {{ props.row.name }}
            <q-popup-edit v-model="props.row.name" v-slot="scope">
              <q-input
                v-model="scope.value"
                dense
                autofocus
                counter
                @keyup.enter="scope.set"
              />
            </q-popup-edit>
          </q-td>
          <q-td key="email" :props="props">
            {{ props.row.email }}
            <q-popup-edit v-model="props.row.email" v-slot="scope">
              <q-input
                v-model="scope.value"
                dense
                autofocus
                counter
                @keyup.enter="scope.set"
              />
            </q-popup-edit>
          </q-td>
          <q-td key="phoneNumber" :props="props">
            {{ props.row.phoneNumber }}
            <q-popup-edit v-model="props.row.phoneNumber" v-slot="scope">
              <q-input
                v-model="scope.value"
                dense
                autofocus
                counter
                @keyup.enter="scope.set"
              />
            </q-popup-edit>
          </q-td>
          <q-td key="attending" :props="props">
            {{ props.row.attending ? 'Yes' : 'No' }}
          </q-td>
          <q-td key="rsvpReceived" :props="props">
            {{ props.row.rsvpReceived ? 'Yes' : 'No' }}
          </q-td>
          <q-td key="note" :props="props">
            {{ props.row.note }}
            <q-popup-edit v-model="props.row.note" v-slot="scope">
              <q-input
                v-model="scope.value"
                dense
                autofocus
                counter
                @keyup.enter="scope.set"
              />
            </q-popup-edit>
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
import { Guest } from 'src/models/events/guest';
import { useEventsStore } from 'src/stores/events-state';
import { useAppStateStore } from 'src/stores/main-application-state';

const appState = useAppStateStore();
const eventStore = useEventsStore();
const events = computed(() => eventStore.events);

const columns = ref([
  {
    name: 'id',
    required: true,
    label: 'ID',
    align: 'left' as const,
    field: (row: Guest) => row.id,
    format: (val: string) => `${val}`,
    sortable: true,
  },
  {
    name: 'name',
    label: 'Name',
    field: (row: Guest) => row.name,
    sortable: true,
  },
  {
    name: 'email',
    label: 'Email',
    field: (row: Guest) => row.email,
    sortable: true,
  },
  {
    name: 'phoneNumber',
    label: 'Phone Number',
    field: (row: Guest) => row.phoneNumber,
    sortable: true,
  },
  {
    name: 'attending',
    label: 'Attending',
    field: (row: Guest) => (row.attending ? 'Yes' : 'No'),
    sortable: true,
  },
  {
    name: 'rsvpReceived',
    label: 'RSVP Received',
    field: (row: Guest) => (row.rsvpReceived ? 'Yes' : 'No'),
    sortable: true,
  },
  {
    name: 'note',
    label: 'Note',
    field: (row: Guest) => row.note,
    sortable: true,
  },
]);

const filter = ref('');

const rows = computed(() => {
  // Extract guests from events and handle cases where there are no guests
  return events.value.flatMap((event) => event.guests || []);
});

const handleOpenAddGuestDialog = () => {
  appState.addGuestDialogOpen = true;
  console.log('clicked');
};
</script>
