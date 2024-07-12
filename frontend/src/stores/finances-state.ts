import { defineStore } from 'pinia';
import { Portfolio } from '../models/finances/portfolio';

export enum FinancesLoadState {
  loading = 0,
  hasAssets = 1,
  hasNoAssets = 2,
}

export const useFinancesStore = defineStore('finances-state-store', {
  state: () => ({
    loadState: FinancesLoadState.loading,
    activePortfolio: Portfolio,
    // Dialog Boxes
    addPropertyDialogBoxOpen: false,
  }),
  getters: {},
  actions: {},
});
