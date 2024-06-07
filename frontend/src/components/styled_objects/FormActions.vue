<template>
  <q-card-section class="card-actions-container">
    <q-card-actions align="right" class="bg-primary text-teal">
      <q-btn flat :label="submitText" @click="onSubmit" v-if="onSubmit" />
      <q-btn flat label="Modify" @click="onModify" v-if="onModify" />
      <q-btn flat label="Delete" @click="onDelete" v-if="onDelete" />
      <q-btn flat label="Cancel" @click="onCancel" />
    </q-card-actions>
  </q-card-section>
</template>

<script setup lang="ts">
defineProps<{
  onSubmit?: () => void;
  submitText?: string;
  onCancel: () => void;
  onDelete?: () => void;
  onModify?: () => void;
}>();
</script>
