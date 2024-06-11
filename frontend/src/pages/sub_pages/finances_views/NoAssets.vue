<template>
  <div class="q-pa-md flex-container">
    <CreatePortfolioDialog />

    <EmptyPage
      :onClick="handleCreatePortfolioDialogOpen"
      label="No portfolio has been created, click below to create one"
    />
  </div>
</template>

<script setup lang="ts">
import CreatePortfolioDialog from '../../../components/finances/CreatePortfolioDialog.vue';
import { useAppStateStore } from 'src/stores/main-application-state';
import EmptyPage from '../../../components/styled_objects/EmptyPage.vue';

const appState = useAppStateStore();

const handleCreatePortfolioDialogOpen = () => {
  appState.addPortfolioDialogOpen = true;
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
  margin-bottom: 20px;
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
