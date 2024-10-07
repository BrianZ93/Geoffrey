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
        :range="isRange"
        class="date-picker-range"
      />
    </q-popup-proxy>
  </div>
</template>

<script setup lang="ts">
import { ref, Ref } from 'vue';

const props = defineProps<{
  label: string;
  onUpdate: (date: { from: string; to: string } | string) => void;
  initialDate?: { from: string; to: string } | string;
  isRange: boolean; // New optional prop to control if the date picker is a range or single date
}>();

const isRange = props.isRange !== undefined ? props.isRange : true; // Default to true (range)

// Set up the initial model based on whether it's a range or single date
const model = ref(
  typeof props.initialDate === 'string'
    ? props.initialDate
    : { from: props.initialDate?.from || '', to: props.initialDate?.to || '' }
);

// Input value shown in the input field
const inputValue = ref('');

// Set initial value if provided
if (props.initialDate !== undefined) {
  if (typeof props.initialDate === 'string') {
    model.value = props.initialDate;
    inputValue.value = `${model.value}`;
  } else {
    model.value = { from: props.initialDate.from, to: props.initialDate.to };
    inputValue.value = `${model.value.from} - ${model.value.to}`;
  }
}

const showDatePicker = ref(false);
const inputField: Ref<HTMLInputElement | null> = ref(null);

// Helper to format the date or date range for display
const formatDate = (date: { from: string; to: string } | string) => {
  if (typeof date === 'string') {
    return date;
  } else {
    return `${date.from} - ${date.to}`;
  }
};

// Handle updates from the date picker
const handleDateUpdate = (date: { from: string; to: string } | string) => {
  inputValue.value = formatDate(date);
  props.onUpdate(date); // Emit the selected date(s) to the parent component
  closeDatePicker();
};

// Close the date picker
const closeDatePicker = () => {
  showDatePicker.value = false;
};

// Focus the input field
const focusInput = () => {
  inputField.value?.focus();
};

// Close the popup when input field loses focus
const handleBlur = () => {
  if (!showDatePicker.value) {
    closeDatePicker();
  }
};
</script>
