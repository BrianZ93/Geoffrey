import { defineStore } from 'pinia';

enum ActivePage {
  Dashboard = 0,
  Finances = 1,
  Events = 2,
  Calendar = 3,
  HomeMaintenance = 4,
}

export const useAppStateStore = defineStore('app-state-store', {
  state: () => ({
    // Main Page States
    sidemenuOpen: false,
    activePage: ActivePage.Dashboard,
    // Events Page States
    addEventDialogOpen: false,
    addGuestDialogOpen: false,
  }),
  getters: {
    ActivePage: (state) => state.activePage,
  },
  actions: {
    setActivePage(page: ActivePage) {
      this.activePage = page;
    },
  },
});
