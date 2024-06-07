import { defineStore } from 'pinia';
import { Event, Venue } from 'src/models/events/event';
import { Guest } from 'src/models/events/guest';

export enum EventsLoadState {
  loading = 0,
  hasEvents = 1,
  hasNoEvents = 2,
}

export const useEventsStore = defineStore('events-state-store', {
  state: () => ({
    loadState: EventsLoadState.loading,
    events: [] as Event[],
    activeEvent: new Event(
      '0',
      '',
      '',
      '',
      [] as Guest[],
      [] as Venue[]
    ) as Event,
  }),
  getters: {},
  actions: {},
});
