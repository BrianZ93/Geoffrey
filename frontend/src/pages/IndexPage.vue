<template>
  <q-page>
    <div class="main-section">
      <div class="sidebar-section section">
        <!-- Column 1 content goes here -->
        <div class="q-pa-md">
          <sidebar-menu />
        </div>
      </div>
      <div class="main-content-section section">
        <!-- Column 2 content goes here -->
        <div>
          <dashboard-page v-if="activePage == 0" />
          <suspense>
            <finances-page v-if="activePage == 1" />
          </suspense>
          <suspense>
            <events-page v-if="activePage == 2" />
          </suspense>
          <suspense>
            <cookbook-page v-if="activePage == 5" />
          </suspense>
          <suspense>
            <education-page v-if="activePage == 6" />
          </suspense>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
// import { Loading } from 'quasar';
import { computed } from 'vue';
import SidebarMenu from './../components/SidebarMenu.vue';
// Page Imports
import DashboardPage from './../pages/sub_pages/Dashboard.vue';
import EventsPage from './../pages/sub_pages/EventsPage.vue';
import FinancesPage from './../pages/sub_pages/FinancesPage.vue';
import CookbookPage from './../pages/sub_pages/CookbookPage.vue';
import EducationPage from './../pages/sub_pages/EducationPage.vue';

// Other Imports
import { useAppStateStore } from 'src/stores/main-application-state';

const appState = useAppStateStore();
const activePage = computed(() => appState.activePage);

defineOptions({
  name: 'IndexPage',
});
</script>
