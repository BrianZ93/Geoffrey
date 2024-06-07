<template>
  <div>
    <q-input
      v-model="inputValue"
      :label="label"
      class="date-picker-input date-picker-container"
      @blur="handleBlur"
      clearable
      outlined
      ref="inputField"
    >
      <template v-slot:append>
        <q-btn flat round icon="event" @click="focusInput"></q-btn>
      </template>
    </q-input>

    <q-popup-proxy
      v-model="showDatePicker"
      transition-show="scale"
      transition-hide="scale"
    >
      <q-date
        v-model="model"
        @update:model-value="handleDateUpdate"
        mask="YYYY-MM-DD"
        range
        class="date-picker-range"
      />
    </q-popup-proxy>
  </div>
</template>

<script setup lang="ts">
import { ref, Ref } from 'vue';

const props = defineProps<{
  label: string;
  onUpdate: (date: string) => void;
}>();

const model = ref({ from: '2024/06/02', to: '2024/06/05' });
const inputValue = ref('');
const showDatePicker = ref(false);
const inputField: Ref<HTMLInputElement | null> = ref(null);

const formatDateRange = (date: { from: string; to: string } | string) => {
  if (typeof date === 'string') {
    return date;
  } else {
    return `${date.from} - ${date.to}`;
  }
};

const handleDateUpdate = (date: { from: string; to: string } | string) => {
  inputValue.value = formatDateRange(date);
  props.onUpdate(formatDateRange(date));
  closeDatePicker();
};

const closeDatePicker = () => {
  showDatePicker.value = false;
};

const focusInput = () => {
  inputField.value?.focus();
};

const handleBlur = () => {
  if (!showDatePicker.value) {
    closeDatePicker();
  }
};
</script>
