<template>
  <div class="row section cookbook-page-container-no-recipes">
    <AddRecipeDialog />
    <EmptyComponent
      :onClick="handleOpenAddRecipeDialog"
      label="No recipes have been added yet, click below to get started."
    />
  </div>
</template>

<script setup lang="ts">
import EmptyComponent from '../../../components/styled_objects/EmptyComponent.vue';
import AddRecipeDialog from '../../../components/cookbook/AddRecipeDialog.vue';
import { useCookbookStore } from '../../../stores/cookbook-state';

const cookbookState = useCookbookStore();

const handleOpenAddRecipeDialog = () => {
  cookbookState.addRecipeDialogOpen = true;
};
</script>
