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
    confirmationDialogOpen: { delete_guest: false } as Record<string, boolean>,
    // Events Page States
    addEventDialogOpen: false,
    modifyEventDialogOpen: false,
    addGuestDialogOpen: false,
    // Finances Page States
    addPortfolioDialogOpen: false,
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
