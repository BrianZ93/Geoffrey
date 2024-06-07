<template>
  <transition name="fade" mode="out-in"
    ><LoadingPage v-if="loadState == EventsLoadState.loading" key="loading"
  /></transition>

  <transition name="fade" mode="out-in"
    ><HasEvents v-if="loadState == EventsLoadState.hasEvents" key="has-events"
  /></transition>

  <transition name="fade" mode="out-in">
    <NoEvents v-if="loadState == EventsLoadState.hasNoEvents" key="no-events" />
  </transition>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue';
import HasEvents from './../sub_pages/events_views/HasEvents.vue';
import NoEvents from './../sub_pages/events_views/NoEvents.vue';
import LoadingPage from './LoadingPage.vue';
import { getEvents } from '../../api/events/get_events';
import { useEventsStore, EventsLoadState } from '../../stores/events-state';

const eventsStore = useEventsStore();
const loadState = computed(() => eventsStore.loadState);

onMounted(async () => {
  try {
    const fetchedEvents = await getEvents();

    if (typeof fetchedEvents === 'string') {
      eventsStore.loadState = EventsLoadState.hasNoEvents;
      return;
    } else {
      eventsStore.events = fetchedEvents;
    }

    if (fetchedEvents.length > 0) {
      eventsStore.loadState = EventsLoadState.hasEvents;
    } else {
      eventsStore.loadState = EventsLoadState.hasNoEvents;
    }
  } catch (error) {
    console.error('Failed to fetch events:', error);
  }
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
