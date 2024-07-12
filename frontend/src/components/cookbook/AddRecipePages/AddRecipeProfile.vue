<template>
  <q-card>
    <q-card-section>
      <StringInput
        label="Recipe Name"
        :onValueChange="handleRecipeNameChange"
        :initialValue="recipeName"
      />

      <q-select
        outlined
        v-model="selectedCountry"
        :options="filteredOptions"
        label="Country"
        use-input
        input-debounce="0"
        emit-value
        map-options
        @filter="filterCountries"
        @update:modelValue="handleUpdateCountry"
      />
    </q-card-section>

    <q-card-section>
      <q-item v-for="(key, index) in flavorKeys" :key="key">
        <q-item-section avatar>
          <q-icon :color="colors[index % colors.length]" :name="icons[index]" />
        </q-item-section>
        <q-item-section>
          <q-item-label class="text-center" style="margin-right: 2.75vw">
            {{ capitalize(key) }} ({{ flavorProfile[key] }} of 5)
          </q-item-label>
          <q-slider
            v-model="flavorProfile[key]"
            :min="0"
            :max="5"
            :step="1"
            snap
            label
            label-always
            label-text-color="primary"
            markers
            :color="colors[index % colors.length]"
            :inner-track-color="colors[index % colors.length]"
            @update:modelValue="
              (value) => handleUpdateFlavorProfile(key, value)
            "
          />
        </q-item-section>
      </q-item>
    </q-card-section>
  </q-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import StringInput from '../../styled_objects/StringInput.vue';
import { handleOnValueChange } from '../../styled_objects/helpers/StringInput';
import { Country } from '../../../models/main';
import { useCookbookStore } from '../../../stores/cookbook-state';

const cookbookState = useCookbookStore();

// Set relevant refs
const recipeName = ref<string | null>('');
recipeName.value = cookbookState.recipeAdding.name;
const selectedCountry = ref<string | null>(null);
for (const [key, value] of Object.entries(Country)) {
  if (value === cookbookState.recipeAdding.country) {
    selectedCountry.value = key;
    break;
  }
}
type FlavorKey = 'sweet' | 'salty' | 'spicy' | 'sour' | 'savory';
const flavorProfile = ref<Record<FlavorKey, number>>({
  sweet: 0,
  salty: 0,
  spicy: 0,
  sour: 0,
  savory: 0,
});
const flavorKeys = Object.keys(flavorProfile.value) as FlavorKey[];
flavorProfile.value = cookbookState.recipeAdding.flavorProfile;

const checkIfAllFieldsAreComplete = () => {
  if (recipeName.value === '' || recipeName.value === null) {
    return;
  }

  if (selectedCountry.value === null) {
    return;
  }

  cookbookState.recipeProfileComplete = true;
  console.log(cookbookState.recipeProfileComplete);
};

const handleRecipeNameChange = (value: string | null) => {
  handleOnValueChange(value, recipeName);
  checkIfAllFieldsAreComplete();
  if (recipeName.value !== null) {
    cookbookState.recipeAdding.name = recipeName.value;
    console.log(cookbookState.recipeAdding);
  }
};

const handleUpdateCountry = (value: number | null) => {
  if (value !== null) {
    cookbookState.recipeAdding.country = value;
    checkIfAllFieldsAreComplete();
  }
};

// Transform the Country enum into a format suitable for q-select options
const countryOptions = Object.keys(Country)
  .filter((key) => isNaN(Number(key))) // Filter out numeric keys
  .map((key) => {
    return {
      label: key.replace(/([A-Z])/g, ' $1').trim(), // Convert camel case to normal text
      value: key.replace(/([A-Z])/g, ' $1').trim(), // Use the same text as value
    };
  });

const filteredOptions = ref(countryOptions);

const filterCountries = (
  val: string,
  update: (callback: () => void) => void
) => {
  update(() => {
    if (val === '') {
      filteredOptions.value = countryOptions;
    } else {
      const needle = val.toLowerCase();
      filteredOptions.value = countryOptions.filter((option) =>
        option.label.toLowerCase().includes(needle)
      );
    }
  });
};

const colors = ['pink', 'white', 'red', 'green', 'purple'];
const icons = ['cake', 'grain', 'local_fire_department', 'egg', 'restaurant'];

// Method to handle flavor profile updates
const handleUpdateFlavorProfile = (key: FlavorKey, value: number | null) => {
  if (value !== null) {
    flavorProfile.value[key] = value;
    cookbookState.recipeAdding.flavorProfile = flavorProfile.value; // Corrected access using bracket notation
    console.log(cookbookState.recipeAdding);
  }
  checkIfAllFieldsAreComplete();
};

// Utility function to capitalize the first letter of a string
const capitalize = (s: string) => s.charAt(0).toUpperCase() + s.slice(1);
</script>
