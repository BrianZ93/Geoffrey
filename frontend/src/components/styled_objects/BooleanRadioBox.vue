<template>
  <q-dialog v-model="model" persistent>
    <q-card class="q-pa-md">
      <div class="update-boolean-header">
        <span
          >Guest: <strong>{{ headerLabel }}</strong></span
        >
      </div>
      <div class="q-gutter-sm">
        <q-radio
          v-model="color"
          :val="trueLabel"
          :label="trueLabel"
          color="green"
        />
        <q-radio
          v-model="color"
          :val="falseLabel"
          :label="falseLabel"
          color="red"
        />
      </div>
      <div class="update-boolean-button">
        <q-btn label="Update" color="secondary" glossy @click="handleSubmit" />
      </div>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
  headerLabel: string;
  trueLabel: string;
  falseLabel: string;
  onSubmit: (selectedValue: boolean) => void;
}>();

const color = ref<string>(props.trueLabel);
const model = ref<boolean>(true); // NEVER EDIT THIS

const handleSubmit = () => {
  let state = false;
  if (color.value === props.trueLabel) {
    state = true;
  }
  props.onSubmit(state);
};
</script>
