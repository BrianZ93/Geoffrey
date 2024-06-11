<template>
  <q-dialog
    v-model="persistent"
    persistent
    transition-show="scale"
    transition-hide="scale"
  >
    <q-card
      class="bg-secondary text-white dialog-card-layout"
      style="width: 70vw; height: 70vh"
    >
      <q-card-section>
        <q-card-section>
          <div class="text-h6 text-center">Add a New Guest</div>
        </q-card-section>

        <q-separator dark />

        <q-card-section class="q-pt-none">
          <!-- Guest Name -->
          <string-input
            label="Guest Name"
            :onValueChange="handleGuestNameChange"
          />
          <!-- Phone Number -->
          <phone-number-input
            label="Phone Number"
            :onValueChange="handlePhoneNumberChange"
          />
          <!-- Email Address -->
          <email-input label="Guest Email" :onValueChange="handleEmailChange" />

          <q-btn
            label="Import Guests from File"
            @click="importGuestsFromFile"
            color="primary"
            glossy
            disable
          />
        </q-card-section>
      </q-card-section>

      <form-actions
        :on-submit="onSubmit"
        submit-text="Add Guest"
        :onCancel="handleCancel"
      />
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { Notify } from 'quasar';
import { ref, computed } from 'vue';
import { useAppStateStore } from '../../stores/main-application-state';
import { useEventsStore, EventsLoadState } from '../../stores/events-state';
import { Guest } from '../../models/events/guest';
import FormActions from '../styled_objects/FormActions.vue';
import StringInput from '../styled_objects/StringInput.vue';
import { handleOnValueChange } from '../../components/styled_objects/helpers/StringInput';
import PhoneNumberInput from '../styled_objects/PhoneNumberInput.vue';
import { handleOnPhoneNumberValueChange } from '../../components/styled_objects/helpers/PhoneNumberInput';
import EmailInput from '../styled_objects/EmailInput.vue';

import { addGuest } from '../../api/events/add_guest';
import { getEvents } from '../../api/events/get_events';
import { guests } from '../../data/guests';

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

const handleGuestNameChange = (value: string | null) => {
  handleOnValueChange(value, guestName);
};

const handlePhoneNumberChange = (value: string | null) => {
  handleOnPhoneNumberValueChange(value, phoneNumber);
};

const handleEmailChange = (value: string | null) => {
  handleOnValueChange(value, email);
};

const importGuestsFromFile = async () => {
  for (let guest of guests) {
    await addGuest(eventsState.activeEvent.id, guest);
  }
};
</script>
