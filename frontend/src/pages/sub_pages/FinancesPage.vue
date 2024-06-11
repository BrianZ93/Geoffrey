<template>
  <transition name="fade" mode="out-in"
    ><LoadingPage v-if="loadState == FinancesLoadState.loading" key="loading"
  /></transition>

  <transition name="fade" mode="out-in"
    ><HasAssets
      v-if="loadState == FinancesLoadState.hasAssets"
      key="has-assets"
  /></transition>

  <transition name="fade" mode="out-in">
    <NoAssets
      v-if="loadState == FinancesLoadState.hasNoAssets"
      key="no-assets"
    />
  </transition>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue';
import HasAssets from './../sub_pages/finances_views/HasAssets.vue';
import NoAssets from './../sub_pages/finances_views/NoAssets.vue';
import LoadingPage from './LoadingPage.vue';
// import { getPortfolio } from '../../api/finances/get_portfolio';
import {
  useFinancesStore,
  FinancesLoadState,
} from '../../stores/finances-state';

const financesState = useFinancesStore();
const loadState = computed(() => financesState.loadState);

onMounted(() => {
  // const fetchedPortfolio = getPortfolio();
  // if (fetchedPortfolio.assets) {
  //   financesState.loadState = FinancesLoadState.hasNoAssets;
  //   return;
  // } else {
  //   financesState.fetchedPortfolio = fetchedPortfolio;
  // }
  // if (fetchedPortfolio.length > 0) {
  //   financesState.loadState = FinancesLoadState.hasAssets;
  // } else {
  //   financesState.loadState = FinancesLoadState.hasNoAssets;
  // }

  financesState.loadState = FinancesLoadState.hasAssets;
  // financesState.loadState = FinancesLoadState.hasNoAssets;
});
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active in <2.1.8 */ {
  opacity: 0;
}
</style>
