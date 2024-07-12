<template>
  <transition name="fade" mode="out-in"
    ><LoadingPage v-if="loadState == FinancesLoadState.loading" key="loading"
  /></transition>

  <transition name="fade" mode="out-in"
    ><AssetsPage key="has-assets"
  /></transition>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue';
import AssetsPage from './../sub_pages/finances_views/AssetsPage.vue';
import LoadingPage from './LoadingPage.vue';
import {
  useFinancesStore,
  FinancesLoadState,
} from '../../stores/finances-state';

const financesState = useFinancesStore();
const loadState = computed(() => financesState.loadState);

onMounted(() => {
  financesState.loadState = FinancesLoadState.hasAssets;
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
