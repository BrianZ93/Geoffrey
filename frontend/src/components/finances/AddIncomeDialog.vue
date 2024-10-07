<template>
  <q-dialog v-model="fixed">
    <q-card class="bg-secondary">
      <q-card-section>
        <div class="text-h6">Add Income to Your Budget</div>
      </q-card-section>

      <q-separator />

      <q-card-section style="max-height: 50vh" class="scroll">
        <string-input
          label="Income Nickname"
          :onValueChange="handleIncomeNicknameChange"
        />
        <money-input
          label="Income Amount"
          :onValueChange="handleIncomeAmountChange"
        />
        <q-select
          outlined
          v-model="selectedFrequency"
          :options="filteredOptions"
          label="Frequency"
          use-input
          input-debounce="0"
          emit-value
          map-options
          @filter="filterFrequencyOptions"
          @update:modelValue="handleUpdateFrequency"
          clearable
          style="margin-top: 1vh"
        />
        <string-input
          label="Income Category"
          :onValueChange="handleIncomeCategoryChange"
        />

        <!-- Radio buttons to toggle expiration date option -->
        <q-radio
          v-model="hasExpiration"
          val="yes"
          label="Yes, this income has an expiration date"
          @update:model-value="toggleDatePicker"
        />
        <q-radio
          v-model="hasExpiration"
          val="no"
          label="No, this income has no expiration date"
          @update:model-value="toggleDatePicker"
        />

        <!-- Show date picker only if 'Yes' is selected -->
        <date-picker
          v-if="hasExpiration === 'yes'"
          label="Expiration Date"
          :isRange="false"
          :onUpdate="handleDateUpdate"
        />
      </q-card-section>

      <form-actions
        :on-submit="handleCreateIncome"
        submit-text="Add Income"
        :onCancel="handleCancelCreateIncome"
      />
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { useFinancesStore } from '../../stores/finances-state';
import FormActions from '../styled_objects/FormActions.vue';
import StringInput from '../styled_objects/StringInput.vue';
import MoneyInput from '../styled_objects/MoneyInput.vue';
import DatePicker from '../styled_objects/DatePicker.vue';
import { handleUpdateDates } from '../../components/styled_objects/helpers/DatePicker';
import { handleOnValueChange } from '../../components/styled_objects/helpers/StringInput';
import { handleOnMoneyValueChange } from '../../components/styled_objects/helpers/MoneyInput'; // Add the correct path
import { Frequency } from '../../models/finances/budget';

import { computed, ref } from 'vue';

const financesState = useFinancesStore();

const fixed = computed(() => financesState.addIncomeDialogBoxOpen);

const incomeNickname = ref<string | null>('');
const incomeAmount = ref<number | null>(null);
const selectedFrequency = ref<number | null>(null);
const startDate = ref<string | null>(null);
const endDate = ref<string | null>(null);
const hasExpiration = ref<string>('no'); // Default is 'no'

// Transform the Frequency enum into a format suitable for q-select options
const frequencyOptions = Object.keys(Frequency)
  .filter((key) => isNaN(Number(key))) // Filter out numeric keys
  .map((key) => {
    return {
      label: key.replace(/([A-Z])/g, ' $1').trim(), // Convert camel case to normal text
      value: Frequency[key as keyof typeof Frequency], // Use the enum value
    };
  });

const filteredOptions = ref(frequencyOptions); // Copy of options for filtering

// Method to handle filtering options in the q-select
const filterFrequencyOptions = (
  val: string,
  update: (callback: () => void) => void
) => {
  update(() => {
    if (val === '') {
      filteredOptions.value = frequencyOptions;
    } else {
      const needle = val.toLowerCase();
      filteredOptions.value = frequencyOptions.filter((option) =>
        option.label.toLowerCase().includes(needle)
      );
    }
  });
};

// Method to handle frequency selection updates
const handleUpdateFrequency = (value: number | null) => {
  if (value !== null) {
    console.log('Selected frequency:', value);
    selectedFrequency.value = value;
    // Perform any additional logic or state updates here
  }
};

const handleCreateIncome = () => {
  financesState.addIncomeDialogBoxOpen = false;
};

const handleCancelCreateIncome = () => {
  financesState.addIncomeDialogBoxOpen = false;
};

const handleIncomeNicknameChange = (value: string | null) => {
  console.log('Nickname changed to:', value);
  handleOnValueChange(value, incomeNickname);
};

const handleIncomeAmountChange = (value: number | null) => {
  console.log('Income amount changed to:', value);
  handleOnMoneyValueChange(value ? value.toString() : null, incomeAmount); // Using the helper function for money
};

const handleIncomeCategoryChange = (value: string | null) => {
  console.log('Category changed to:', value);
  handleOnValueChange(value, incomeNickname);
};

const handleDateUpdate = (date: { from: string; to: string } | string) => {
  handleUpdateDates(date, startDate, endDate);
};

// Function to toggle visibility of the date picker based on radio button selection
const toggleDatePicker = (value: string) => {
  if (value === 'yes') {
    // Show the date picker
    hasExpiration.value = 'yes';
  } else {
    // Hide the date picker
    hasExpiration.value = 'no';
    startDate.value = null; // Clear date values if "No" is selected
    endDate.value = null;
  }
};
</script>
