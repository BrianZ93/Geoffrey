<template>
  <q-input
    v-model="inputValue"
    type="email"
    :label="label"
    :rules="[emailRule]"
    lazyRules
    @update:model-value="updateInputValue"
    clearable
    outlined
    class="form-input-container"
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

const emailRule = (val: string) => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(val) || 'Invalid email address';
};
</script>
