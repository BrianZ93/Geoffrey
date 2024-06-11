<template>
  <div class="flex-center">
    <div class="spinning-object">
      <!-- Replace this with your desired object -->
      <q-icon name="donut_large" size="64px" />
    </div>
    <h6 class="empty-component-description">
      {{ label }}
    </h6>
    <q-btn round color="primary" icon="add" @click="handleOpenDialog" />
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  onClick: () => void;
  label: string;
}>();

const handleOpenDialog = () => {
  props.onClick();
};
</script>

<style scoped>
.flex-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  text-align: center;
  flex-direction: column;
}

.flex-center {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.spinning-object {
  animation: spin 15s linear infinite;
  margin-top: 5vh;
}

.empty-component-description {
  justify-content: center;
  align-items: center;
  text-align: center;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>
