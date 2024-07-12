import { defineStore } from 'pinia';
import { FlavorProfile, Recipe } from '../models/cookbook/recipes';

export enum CookbookLoadState {
  hasRecipes = 0,
  hasNoRecipes = 1,
  loading = 2,
}

export const useCookbookStore = defineStore('cookbook-state-store', {
  state: () => ({
    loadState: CookbookLoadState.hasNoRecipes,
    recipes: [] as Recipe[],
    // Dialog Box States
    addRecipeDialogOpen: false,
    recipeAdding: new Recipe(
      '0',
      '',
      [],
      0,
      0,
      new FlavorProfile(0, 0, 0, 0, 0),
      193,
      [],
      ''
    ),
    recipeModifying: new Recipe(
      '0',
      '',
      [],
      0,
      0,
      new FlavorProfile(0, 0, 0, 0, 0),
      193,
      [],
      ''
    ),
    recipeProfileComplete: false,
  }),
  getters: {
    getLoadState: (state) => state.loadState,
  },
});
