<template>
  <q-input
    v-model="displayValue"
    :label="label"
    @update:model-value="updateInputValue"
    @blur="formatOnBlur"
    @keypress="restrictToNumericInput"
    clearable
    outlined
    class="form-input-container"
    prefix="$"
    type="text"
  />
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';

const props = defineProps<{
  label: string;
  onValueChange: (value: number | null) => void;
  initialValue?: number | null;
}>();

const numericValue = ref<number | null>(null); // Store the numeric value
const displayValue = ref<string>(''); // Store the value being displayed in the input

// Set initial value on mount
onMounted(() => {
  if (props.initialValue !== undefined && props.initialValue !== null) {
    numericValue.value = props.initialValue;
    displayValue.value = formatToUSD(props.initialValue); // Initially show the formatted value
  }
});

// Function to format number to USD currency string
const formatToUSD = (value: number) => {
  return value
    .toLocaleString('en-US', {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2,
    })
    .replace('$', ''); // Remove $ since it is handled by the prefix
};

// Function to parse the input string back to a number
const parseUSD = (value: string): number | null => {
  const numericValue = parseFloat(value.replace(/[^0-9.-]/g, ''));
  return isNaN(numericValue) ? null : numericValue;
};

// Handler for input updates
const updateInputValue = (value: string | number | null) => {
  if (typeof value === 'string') {
    // Only allow valid characters (numbers, dot, dash) while typing
    displayValue.value = value.replace(/[^0-9.-]/g, '');

    // Parse string input into a number
    const parsedValue = parseUSD(displayValue.value);
    numericValue.value = parsedValue; // Keep track of the numeric value
    props.onValueChange(parsedValue); // Emit the parsed number value to the parent
  } else if (value === null) {
    displayValue.value = ''; // Clear display if input is invalid or null
    numericValue.value = null;
    props.onValueChange(null); // Emit null to the parent
  }
};

// Format the value when the user leaves the input (on blur)
const formatOnBlur = () => {
  if (numericValue.value !== null) {
    displayValue.value = formatToUSD(numericValue.value); // Format display value as USD string
  }
};

// Prevent non-numeric characters from being entered
const restrictToNumericInput = (event: KeyboardEvent) => {
  const allowedKeys = [
    '0',
    '1',
    '2',
    '3',
    '4',
    '5',
    '6',
    '7',
    '8',
    '9',
    '.',
    '-',
  ];
  const isAllowedKey = allowedKeys.includes(event.key);

  // Allow control keys like Backspace, Delete, Arrow keys, Tab
  const isControlKey =
    event.ctrlKey ||
    ['Backspace', 'Tab', 'Delete', 'ArrowLeft', 'ArrowRight'].includes(
      event.key
    );

  if (!isAllowedKey && !isControlKey) {
    event.preventDefault(); // Block the key press if it's not a valid character
  }
};
</script>
