import { defineStore } from 'pinia';

export enum EducationLoadState {
  hasEducation = 0,
  loading = 1,
}

export const useEducationStore = defineStore('education-state-store', {
  state: () => ({
    loadState: EducationLoadState.loading,
  }),
  getters: {},
  actions: {},
});
