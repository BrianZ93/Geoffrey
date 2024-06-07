<template>
  <div class="q-pa-md q-gutter-sm">
    <q-dialog
      v-model="persistent"
      persistent
      transition-show="scale"
      transition-hide="scale"
    >
      <q-card class="bg-secondary text-white" style="width: 50vw; height: 50vh">
        <q-card-section>
          <q-card-section>
            <div class="text-h6 text-center">Add a New Guest</div>
          </q-card-section>

          <q-separator dark />

          <q-card-section class="q-pt-none">
            <!-- Guest Name -->
            <q-input v-model="guestName" label="Guest Name" />
            <!-- Phone Number -->
            <q-input
              v-model="phoneNumber"
              type="tel"
              label="Telephone number"
            />
            <!-- Email Address -->
            <q-input v-model="email" type="email" label="Email" />
          </q-card-section>
        </q-card-section>

        <q-card-actions align="right" class="bg-primary text-teal">
          <q-btn flat label="Add Guest" @click="onSubmit" />
          <q-btn flat label="Cancel" @click="handleCancel" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { Notify } from 'quasar';
import { ref, computed } from 'vue';
import { useAppStateStore } from '../../stores/main-application-state';
import { useEventsStore, EventsLoadState } from '../../stores/events-state';
import { Guest } from '../../models/events/guest';

import { addGuest } from '../../api/events/add_guest';
import { getEvents } from '../../api/events/get_events';

const appState = useAppStateStore();
const eventsState = useEventsStore();

const persistent = computed(() => appState.addGuestDialogOpen);

const guestName = ref<string | null>('');
const phoneNumber = ref<string | null>(null);
const email = ref<string | null>(null);

const handleAddGuest = async (
  guestName: string,
  phoneNumber: string,
  email: string
) => {
  const guest = new Guest('0', guestName, email, phoneNumber, false, false, '');
  await addGuest(eventsState.activeEvent.id, guest);
};

const onSubmit = async () => {
  if (!guestName.value || !phoneNumber.value || !email.value) {
    let emptyFields = [] as string[];
    if (!guestName.value) {
      emptyFields.push('Guest Name');
    }

    if (!phoneNumber.value) {
      emptyFields.push('Phone Number');
    }

    if (!email.value) {
      emptyFields.push('Email Address');
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
    await handleAddGuest(guestName.value, phoneNumber.value, email.value);

    eventsState.loadState = EventsLoadState.loading;

    Notify.create({
      color: 'green-4',
      textColor: 'white',
      icon: 'cloud_done',
      message: 'Adding Guest',
      type: 'loading',
    });

    const fetchedEvents = await getEvents();

    Notify.create({
      color: 'green-4',
      textColor: 'white',
      icon: 'cloud_done',
      message: 'Guest Added',
    });

    eventsState.loadState = EventsLoadState.hasEvents;
    if (typeof fetchedEvents !== 'string') {
      eventsState.events = fetchedEvents;
    }
  }
};

const handleCancel = () => {
  appState.addGuestDialogOpen = false;
};
</script>
