<template>
  <div class="q-pa-md">
    <div class="row section header-row">
      <!-- Top Row -->
      <q-toolbar class="bg-primary glossy text-white">
        <q-avatar icon="account_balance" size="50px" color="secondary">
        </q-avatar>

        <q-toolbar-title>Portfolio Manager</q-toolbar-title>

        <q-tabs v-model="tab" shrink>
          <q-tab name="overview" label="Overview" />
          <q-tab name="tab1" label="Real Estate" />
          <q-tab name="tab2" label="Equities" />
          <q-tab name="tab3" label="Crypto" />
          <q-tab name="tab3" label="Markets" />
        </q-tabs>

        <q-btn flat round dense icon="whatshot" />
      </q-toolbar>
    </div>
    <!-- First Section -->
    <div class="finances-main-content-row">
      <OverviewTab v-if="tab === 'overview'" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import OverviewTab from './tabs/OverviewTab.vue';

const tab = ref<string>('overview');
</script>
