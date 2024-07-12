<template>
  <q-card>
    <q-card-section
      style="text-align: center; justify-content: center; align-items: center"
    >
      <q-btn
        label="Add Step"
        icon="add"
        color="primary"
        glossy
        rounded
        style="width: 100%"
        @click="handleAddStep"
      ></q-btn
    ></q-card-section>

    <q-list bordered separator>
      <q-scroll-area style="height: 55vh">
        <q-item v-for="(step, index) in steps" :key="step.id" clickable>
          <div>
            <q-icon
              name="restaurant"
              class="step-icon"
              size="lg"
              color="secondary"
            ></q-icon>
          </div>

          <q-item-section>
            <q-item-section>
              <q-input
                standout="bg-teal text-white"
                v-model="steps[index].name"
                label="Step Name"
                dense />
              <q-select
                outlined
                v-model="selectedAppliance[index]"
                :options="filteredApplianceOptions"
                label="Appliance Used"
                use-input
                input-debounce="0"
                emit-value
                map-options
                @filter="filterAppliances"
                @update:modelValue="
                  (value) => handleUpdateAppliance(index, value)
                " />
              <q-select
                outlined
                v-model="selectedCookingAction[index]"
                :options="filteredCookingActionOptions"
                label="Cooking Action"
                use-input
                input-debounce="0"
                emit-value
                map-options
                @filter="filterCookingActionUsed"
                @update:modelValue="
                  (value) => handleUpdateCookingAction(index, value)
                " /></q-item-section
          ></q-item-section>

          <q-item-section>
            <q-input
              v-model="steps[index].description"
              filled
              type="textarea"
              label="Provide the step description here"
              @update:model-value="
                (value) => handleUpdateDescription(index, value)
              "
            ></q-input>
          </q-item-section>

          <q-item-section>
            <q-item-label class="text-center text-h6"
              >Active Cooking Time</q-item-label
            >
            <q-input
              v-model="activeCookingTimesMinutes[index]"
              filled
              type="number"
              label="Minutes"
              dense
              outlined
              @update:model-value="
                (value) => handleChangeActiveCookingTimeMinutes(value, index)
              "
            />
            <q-input
              v-model="activeCookingTimesHours[index]"
              filled
              type="number"
              label="Hours"
              dense
              outlined
              @update:model-value="
                (value) => handleChangeActiveCookingTimeHours(value, index)
              "
            />
            <!-- <q-input v-model="" /> -->
          </q-item-section>

          <q-item-section>
            <q-item-label class="text-center text-h6"
              >Total Cooking Time</q-item-label
            >
            <q-input
              v-model="totalCookingTimesMinutes[index]"
              filled
              type="number"
              label="Minutes"
              dense
              outlined
              @update:model-value="
                (value) => handleChangeTotalCookingTimeMinutes(value, index)
              "
            />
            <q-input
              v-model="totalCookingTimesHours[index]"
              filled
              type="number"
              label="Hours"
              dense
              outlined
              @update:model-value="
                (value) => handleChangeTotalCookingTimeHours(value, index)
              "
            />
            <q-select
              outlined
              v-model="recipeStepStage[index]"
              :options="recipeStepStageUsedOptions"
              label="Recipe Stage"
              use-input
              input-debounce="0"
              emit-value
              map-options
              @filter="filterRecipeStepStages"
              @update:modelValue="
                (value) => handleUpdateRecipeStepStage(index, value)
              "
            />
          </q-item-section>

          <div>
            <q-btn
              :disable="index == 0"
              flat
              dense
              icon="arrow_upward"
              color="red"
              @click="moveStepUp(index)"
              class="delete-icon"
            ></q-btn>
            <q-btn
              :disable="index == steps.length - 1"
              flat
              dense
              icon="arrow_downward"
              color="red"
              @click="moveStepDown(index)"
              class="delete-icon"
            ></q-btn>
            <q-btn
              flat
              dense
              icon="cancel"
              color="red"
              @click="handleDeleteStep(index)"
              class="delete-icon"
            ></q-btn>
          </div>
        </q-item>
      </q-scroll-area>
    </q-list>
  </q-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import {
  RecipeStep,
  CookingAction,
  CookingAppliance,
  RecipeStepStage,
} from '../../../models/cookbook/recipes';
import { useCookbookStore } from '../../../stores/cookbook-state';

const cookbookState = useCookbookStore();

const steps = ref<RecipeStep[]>([]);
steps.value = cookbookState.recipeAdding.steps;
const activeCookingTimesMinutes = ref<number[]>([]);
const activeCookingTimesHours = ref<number[]>([]);

const totalCookingTimesMinutes = ref<number[]>([]);
const totalCookingTimesHours = ref<number[]>([]);

const selectedAppliance = ref<string[]>([]);
const selectedCookingAction = ref<string[]>([]);

const recipeStepStage = ref<RecipeStepStage[]>([]);

console.log(cookbookState.recipeAdding);

// Transform the Appliance enum into a format suitable for q-select options
const applianceUsedOptions = Object.keys(CookingAppliance)
  .filter((key) => isNaN(Number(key))) // Filter out numeric keys
  .map((key) => {
    return {
      label: key.replace(/([A-Z])/g, ' $1').trim(), // Convert camel case to normal text
      value: key.replace(/([A-Z])/g, ' $1').trim(), // Use the same text as value
    };
  });

// Transform the Appliance enum into a format suitable for q-select options
const cookingActionUsedOptions = Object.keys(CookingAction)
  .filter((key) => isNaN(Number(key))) // Filter out numeric keys
  .map((key) => {
    return {
      label: key.replace(/([A-Z])/g, ' $1').trim(), // Convert camel case to normal text
      value: key.replace(/([A-Z])/g, ' $1').trim(), // Use the same text as value
    };
  });

// Transform the RecipeStageStep enum into a format suitable for q-select options
const recipeStepStageUsedOptions = Object.keys(RecipeStepStage)
  .filter((key) => isNaN(Number(key)))
  .map((key) => {
    return {
      label: key.replace(/([A-Z])/g, ' $1').trim(), // Convert camel case to normal text
      value: key.replace(/([A-Z])/g, ' $1').trim(), // Use the same text as value
    };
  });

const filteredApplianceOptions = ref(applianceUsedOptions);
const filteredCookingActionOptions = ref(cookingActionUsedOptions);
const filteredRecipeStepStageOptions = ref(recipeStepStageUsedOptions);

const filterAppliances = (
  val: string,
  update: (callback: () => void) => void
) => {
  update(() => {
    if (val === '') {
      filteredApplianceOptions.value = applianceUsedOptions;
    } else {
      const needle = val.toLowerCase();
      filteredApplianceOptions.value = applianceUsedOptions.filter((option) =>
        option.label.toLowerCase().includes(needle)
      );
    }
  });
};

const filterCookingActionUsed = (
  val: string,
  update: (callback: () => void) => void
) => {
  update(() => {
    if (val === '') {
      filteredCookingActionOptions.value = cookingActionUsedOptions;
    } else {
      const needle = val.toLowerCase();
      filteredCookingActionOptions.value = cookingActionUsedOptions.filter(
        (option) => option.label.toLowerCase().includes(needle)
      );
    }
  });
};

const filterRecipeStepStages = (
  val: string,
  update: (callback: () => void) => void
) => {
  update(() => {
    if (val === '') {
      filteredRecipeStepStageOptions.value = recipeStepStageUsedOptions;
    } else {
      const needle = val.toLowerCase();
      filteredRecipeStepStageOptions.value = recipeStepStageUsedOptions.filter(
        (option) => option.label.toLowerCase().includes(needle)
      );
    }
  });
};

const handleUpdateAppliance = (index: number, value: number | null) => {
  if (value !== null) {
    steps.value[index].applianceUsed = value;
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleUpdateRecipeStepStage = (index: number, value: number | null) => {
  if (value !== null) {
    steps.value[index].recipeStage = value;
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleUpdateCookingAction = (index: number, value: number | null) => {
  if (value !== null) {
    steps.value[index].cookingAction = value;
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleUpdateDescription = (
  index: number,
  value: string | number | null
) => {
  if (typeof value === 'string') {
    steps.value[index].description = value;
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleAddStep = () => {
  const newStep = new RecipeStep('0', '', 0, 0, 0, 0, [], '', 0);
  steps.value.push(newStep);
  activeCookingTimesMinutes.value.push(0);
  activeCookingTimesHours.value.push(0);
  totalCookingTimesHours.value.push(0);
  totalCookingTimesMinutes.value.push(0);
};

const handleDeleteStep = (index: number) => {
  if (index >= 0 && index < steps.value.length) {
    steps.value.splice(index, 1);
  }
};

const covertTimeToSeconds = (minutes: number, hours: number) => {
  let seconds = minutes * 60;
  seconds += hours * 3600;
  return seconds;
};

const handleChangeActiveCookingTimeMinutes = (
  value: string | number | null,
  index: number
) => {
  if (value !== null) {
    activeCookingTimesMinutes.value[index] = Number(value);
    steps.value[index].timeSeconds = covertTimeToSeconds(
      activeCookingTimesMinutes.value[index],
      activeCookingTimesHours.value[index]
    );
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleChangeActiveCookingTimeHours = (
  value: string | number | null,
  index: number
) => {
  if (value !== null) {
    activeCookingTimesHours.value[index] = Number(value);
    steps.value[index].timeSeconds = covertTimeToSeconds(
      activeCookingTimesMinutes.value[index],
      activeCookingTimesHours.value[index]
    );
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleChangeTotalCookingTimeMinutes = (
  value: string | number | null,
  index: number
) => {
  if (value !== null) {
    totalCookingTimesMinutes.value[index] = Number(value);
    steps.value[index].timeSeconds = covertTimeToSeconds(
      totalCookingTimesMinutes.value[index],
      totalCookingTimesHours.value[index]
    );
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const handleChangeTotalCookingTimeHours = (
  value: string | number | null,
  index: number
) => {
  if (value !== null) {
    totalCookingTimesHours.value[index] = Number(value);
    steps.value[index].timeSeconds = covertTimeToSeconds(
      totalCookingTimesMinutes.value[index],
      totalCookingTimesHours.value[index]
    );
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const moveStepUp = (index: number) => {
  if (index > 0) {
    const temp = steps.value[index];
    steps.value[index] = steps.value[index - 1];
    steps.value[index - 1] = temp;
    cookbookState.recipeAdding.steps = steps.value;
  }
};

const moveStepDown = (index: number) => {
  if (index < steps.value.length - 1) {
    const temp = steps.value[index];
    steps.value[index] = steps.value[index + 1];
    steps.value[index + 1] = temp;
    cookbookState.recipeAdding.steps = steps.value;
  }
};
</script>

<style>
.step-icon {
  align-items: center;
  justify-content: center;
  text-align: center;
  height: 100%;
  margin-right: 1vw;
}

.delete-icon {
  align-items: center;
  justify-content: center;
  text-align: center;
  height: 100%;
  margin-left: 1vw;
}
</style>
