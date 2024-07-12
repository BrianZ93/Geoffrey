<template>
  <div class="q-pa-md">
    <div class="row section header-row">
      <!-- Top Header Row -->
      <q-toolbar class="bg-primary glossy text-white">
        <q-avatar
          icon="account_balance"
          size="50px"
          color="secondary"
        ></q-avatar>

        <q-toolbar-title>Cookbook</q-toolbar-title>
      </q-toolbar>
    </div>
    <!-- Second Row -->
    <div class="row section">
      <MainView v-if="pageState == 0" />
      <NoRecipes v-if="pageState == 1" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import NoRecipes from '../sub_pages/cookbook_views/NoRecipes.vue';
import MainView from '../sub_pages/cookbook_views/MainView.vue';
import { useCookbookStore } from '../../stores/cookbook-state';

const cookbookState = useCookbookStore();
const pageState = computed(() => cookbookState.getLoadState);

console.log(pageState.value);
</script>
