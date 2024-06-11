<template>
  <q-input
    v-model="inputValue"
    type="tel"
    :label="label"
    @update:model-value="updateInputValue"
    clearable
    outlined
    class="form-input-container"
    mask="(###) ###-####"
    :rules="[phoneRule]"
    lazyRules
  />
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
  label: string;
  onValueChange: (value: string | null) => void;
  initialValue?: string;
}>();

const inputValue = ref<string | null>('');

if (props.initialValue && props.initialValue !== undefined) {
  inputValue.value = props.initialValue;
}

const updateInputValue = (value: string | number | null) => {
  if (typeof value === 'number') {
    props.onValueChange(value.toString());
  } else {
    props.onValueChange(value);
  }
};

const phoneRule = (val: string) => {
  // Remove all non-numeric characters
  const numericVal = val.replace(/\D/g, '');
  return numericVal.length === 10 || 'Phone number must be 10 digits long';
};
</script>
