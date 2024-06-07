<template>
  <q-card-section class="card-actions-container">
    <q-card-actions align="right" class="bg-primary text-teal">
      <q-btn flat label="Create Event" @click="onSubmit" />
      <q-btn flat label="Cancel" @click="onCancel" />
    </q-card-actions>
  </q-card-section>
</template>

<script setup lang="ts">
defineProps<{
  onSubmit: () => void;
  onCancel: () => void;
}>();
</script>
