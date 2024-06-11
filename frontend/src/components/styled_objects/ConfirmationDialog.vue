<template>
  <q-dialog v-model="isOpen">
    <q-card>
      <q-card-section class="row items-center">
        <q-avatar icon="warning" color="orange" />
        <span class="q-ml-md">Are you sure you want to {{ message }}?</span>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat label="Delete" color="negative" @click="confirm" />
        <q-btn flat label="Cancel" color="secondary" @click="cancel" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useAppStateStore } from 'src/stores/main-application-state';

const props = defineProps<{
  cancel: () => void;
  confirm: () => void;
  message: string;
  modelValue: string;
}>();

const appState = useAppStateStore();

const isOpen = computed(
  () => appState.confirmationDialogOpen[props.modelValue]
);

const cancel = () => {
  props.cancel();
};

const confirm = () => {
  props.confirm();
};
</script>
