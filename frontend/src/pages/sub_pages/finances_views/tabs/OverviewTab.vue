<template>
  <div class="row">
    <!-- Dialog Boxes -->
    <AddPropertyDialogBox />
    <!-- Top Row -->
    <div class="finances-overview-top-row">
      <!-- Properties Section -->
      <div class="finances-overview-properties-container section">
        <EmptyComponent
          :onClick="handleAddPropertyDialogOpen"
          label="No properties have been added, click below to add one"
        />
      </div>
      <!-- Equities Section -->
      <div class="finances-overview-equities-container section">
        <EmptyComponent
          :onClick="handleAddEquityDialogOpen"
          label="No equities have been added, click below to get started"
        />
      </div>
    </div>
    <div class="finances-overview-bottom-row section">
      <!-- Section One -->
      <div class="finances-overview-bottom-container-one">
        <EmptyComponent
          :onClick="handleAddDebtDialogOpen"
          label="No debts have been added, click below to get started"
        />
      </div>
      <!-- Section Two -->
      <div class="finances-overview-bottom-container-two">
        <EmptyComponent
          :onClick="handleAddProjectionDialogOpen"
          label="No projections have been added, click below to get started"
        />
      </div>
      <!-- Section Three -->
      <div class="finances-overview-bottom-container-three">
        <EmptyComponent
          :onClick="handleMarketInsightsDialogOpen"
          label="No market insights have been added, click below to get started"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// import CreatePortfolioDialog from '../../../components/finances/CreatePortfolioDialog.vue';
// import { useAppStateStore } from 'src/stores/main-application-state';
import { useFinancesStore } from 'src/stores/finances-state';
// import { computed } from 'vue';
import EmptyComponent from '../../../../components/styled_objects/EmptyComponent.vue';
// Dialog Box Imports
import AddPropertyDialogBox from '../../../../components/finances/AddPropertyDialog.vue';

// const appState = useAppStateStore();
const financesState = useFinancesStore();

// const portfolio = computed(() => financesState.activePortfolio);

const handleAddPropertyDialogOpen = () => {
  console.log('add property function to be placed here');
  financesState.addPropertyDialogBoxOpen = true;
};

const handleAddEquityDialogOpen = () => {
  console.log('add equity function to be placed here');
};

const handleAddDebtDialogOpen = () => {
  console.log('add debt function to be placed here');
};

const handleAddProjectionDialogOpen = () => {
  console.log('add projection function to be placed here');
};

const handleMarketInsightsDialogOpen = () => {
  console.log('add market insights function to be placed here');
};
</script>
