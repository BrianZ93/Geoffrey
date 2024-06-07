<template>
  <q-layout>
    <navbar-menu />
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import NavbarMenu from 'components/NavbarMenu.vue';

defineOptions({
  name: 'MainLayout',
});
</script>
