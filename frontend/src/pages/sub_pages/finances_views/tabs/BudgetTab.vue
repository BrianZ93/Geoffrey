<template>
  <div class="row">
    <!-- Dialog Boxes -->
    <AddIncomeDialogBox />
    <!-- Top row -->
    <div
      class="finances-overview-top-row q-pa-sm section row"
      style="height: 5vh"
    >
      <q-btn
        align="around"
        color="secondary"
        label="Add Income"
        icon="add"
        glossy
        rounded
        @click="handleOpenAddIncomeDialogBox"
      />
      <q-btn
        align="around"
        color="secondary"
        label="Add Expense"
        icon="add"
        style="margin-left: 0.25vw"
        glossy
        rounded
      />
    </div>
    <!-- Bottom Row -->
    <div class="finances-overview-bottom-row"></div>
  </div>
</template>

<script setup lang="ts">
import { useFinancesStore } from 'src/stores/finances-state';
import AddIncomeDialogBox from '../../../../components/finances/AddIncomeDialog.vue';

const financesState = useFinancesStore();

const handleOpenAddIncomeDialogBox = () => {
  console.log('Opening add income dialog box');
  financesState.addIncomeDialogBoxOpen = true;
};
</script>
