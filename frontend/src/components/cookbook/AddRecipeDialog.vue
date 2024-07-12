<template>
  <q-dialog v-model="dialogOpen" full-width full-height>
    <q-card class="q-pa-md bg-primary" style="width: 100%; height: 100%">
      <q-btn
        label="Reset"
        push
        color="secondary"
        text-color="tertiary"
        @click="step = 1"
        class="q-mb-md"
      />

      <q-btn
        label="Cancel"
        push
        color="secondary"
        text-color="tertiary"
        @click="cookbookState.addRecipeDialogOpen = false"
        class="q-mb-md"
        style="margin-left: 1vh"
      />

      <q-stepper
        v-model="step"
        header-nav
        ref="stepper"
        text-color="white"
        animated
      >
        <q-step
          :name="1"
          title="Create Recipe Profile"
          icon="settings"
          :done="step > 1"
          :header-nav="step > 1"
          color="white"
        >
          <AddRecipeProfile />

          <q-stepper-navigation>
            <q-btn
              @click="
                () => {
                  step = 2;
                }
              "
              color="secondary"
              label="Continue"
              :disable="!recipeProfileComplete"
            />
          </q-stepper-navigation>
        </q-step>

        <q-step
          :name="2"
          title="Add Steps"
          icon="local_dining"
          :done="step > 2"
          :header-nav="step > 2"
          color="white"
        >
          <AddSteps />

          <q-stepper-navigation>
            <q-btn
              @click="
                () => {
                  step = 3;
                }
              "
              color="secondary"
              label="Continue"
            />
            <q-btn
              @click="step = 1"
              color="primary"
              label="Back"
              class="q-ml-sm"
            />
          </q-stepper-navigation>
        </q-step>

        <q-step
          :name="3"
          title="Create an ad"
          icon="add_comment"
          :header-nav="step > 3"
          color="white"
        >
          <AddRecipeConfirmation />
          <q-stepper-navigation>
            <q-btn color="primary" label="Finish" />
            <q-btn
              flat
              @click="step = 2"
              color="primary"
              label="Back"
              class="q-ml-sm"
            />
          </q-stepper-navigation>
        </q-step>
      </q-stepper>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useCookbookStore } from '../../stores/cookbook-state';
import AddRecipeProfile from '../cookbook/AddRecipePages/AddRecipeProfile.vue';
import AddSteps from '../cookbook/AddRecipePages/AddSteps.vue';
import AddRecipeConfirmation from '../cookbook/AddRecipePages/AddRecipeConfirmation.vue';

// Triggers the dialog to open - set this to computed variable from store
const cookbookState = useCookbookStore();
const dialogOpen = computed(() => cookbookState.addRecipeDialogOpen);
const recipeProfileComplete = computed(
  () => cookbookState.recipeProfileComplete
);
const step = ref<number>(1);
</script>
