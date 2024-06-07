<template>
  <q-input
    v-model="inputValue"
    :label="label"
    @update:model-value="updateInputValue"
    clearable
    outlined
    class="string-input-container"
  />
</template>

<script setup lang="ts">
import { ref } from 'vue';

const inputValue = ref<string | null>('');

const props = defineProps<{
  label: string;
  onValueChange: (value: string | null) => void;
}>();

const updateInputValue = (value: string | number | null) => {
  if (typeof value === 'number') {
    props.onValueChange(value.toString());
  } else {
    props.onValueChange(value);
  }
};
</script>
