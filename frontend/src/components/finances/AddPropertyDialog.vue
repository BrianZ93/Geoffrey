<template>
  <q-dialog v-model="fixed">
    <q-card class="bg-secondary">
      <q-card-section>
        <div class="text-h6">Add a Property to Your Portfolio</div>
      </q-card-section>

      <q-separator />

      <q-card-section style="max-height: 50vh" class="scroll">
        <string-input
          label="Property Nickname"
          :onValueChange="handlePropertyNicknameChange"
        />
      </q-card-section>

      <q-card-section>
        <AddressInput :onAddressChange="handleAddressChange" />
      </q-card-section>

      <q-card-section>
        <q-select
          outlined
          v-model="propertyUseModel"
          :options="propertyUseOptions"
          label="Property Use"
          @update:modelValue="handlePropertyUseChange"
        />
      </q-card-section>

      <q-card-section>
        <q-select
          outlined
          v-model="propertyTypeModel"
          :options="propertyTypeOptions"
          label="Property Type"
          @update:modelValue="handlePropertyTypeChange"
        />
      </q-card-section>

      <q-separator />

      <form-actions
        :on-submit="handleCreateProperty"
        submit-text="Add Property"
        :onCancel="handleCancelCreateProperty"
      />
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { Notify } from 'quasar';
import { ref, computed } from 'vue';
import { useFinancesStore } from '../../stores/finances-state';
import StringInput from '../styled_objects/StringInput.vue';
import FormActions from '../styled_objects/FormActions.vue';
import AddressInput from '../styled_objects/AddressInput.vue';
import { handleOnValueChange } from '../../components/styled_objects/helpers/StringInput';
import { handleOnAddressChange } from '../../components/styled_objects/helpers/AddressInput';
import {
  Address,
  PropertyUse,
  PropertyType,
} from '../../models/finances/property';
import { addProperty } from '../../api/finances/add_property';

// Utility function to extract enum values
function getEnumValues<T extends object>(enumObj: T): string[] {
  return Object.keys(enumObj)
    .filter((key) => isNaN(Number(key)))
    .map((key) => key);
}

// Utility function to convert string to enum value
function getEnumValueFromString<T extends object>(
  enumObj: T,
  value: string
): T[keyof T] | null {
  const keys = Object.keys(enumObj).filter((key) => isNaN(Number(key)));
  const key = keys.find((key) => key === value);
  return key ? enumObj[key as keyof T] : null;
}

const financesState = useFinancesStore();

const fixed = computed(() => financesState.addPropertyDialogBoxOpen);

const propertyUseOptions = getEnumValues(PropertyUse);
const propertyUseModel = ref<string>(propertyUseOptions[0]);
const propertyTypeOptions = getEnumValues(PropertyType);
const propertyTypeModel = ref<string>(propertyTypeOptions[0]);

// Property Values
const propertyNickname = ref<string | null>('');
const address = ref<Address | null>(null);

const handleCreateProperty = async () => {
  // Validate that the property nickname and address have been entered
  if (
    !propertyNickname.value ||
    !address.value ||
    !address.value.streetNumber ||
    !address.value.streetName ||
    !address.value.city ||
    !address.value.state ||
    !address.value.zip
  ) {
    Notify.create({
      message: 'You must provide all fields before adding a property',
      icon: 'warning',
      textColor: 'white',
      color: 'red-5',
    });
    return;
  }

  // Convert string models to enum values
  const propertyUse = getEnumValueFromString(
    PropertyUse,
    propertyUseModel.value
  );
  const propertyType = getEnumValueFromString(
    PropertyType,
    propertyTypeModel.value
  );

  if (propertyUse === null || propertyType === null) {
    Notify.create({
      message: 'Invalid property use or property type',
      icon: 'warning',
      textColor: 'white',
      color: 'red-5',
    });
    return;
  }

  // Prepare the property data for the API call
  const propertyData = {
    nickname: propertyNickname.value,
    address: address.value,
    purchaseDate: new Date().toISOString(), // Convert Date to string
    purchasePrice: 0, // Example placeholder, adjust as needed
    propertyUse: propertyUse,
    propertyType: propertyType,
    mortgages: [], // Example placeholder, adjust as needed
    tenants: [], // Example placeholder, adjust as needed
    income: [], // Example placeholder, adjust as needed
    expenses: [], // Example placeholder, adjust as needed
  };

  try {
    await addProperty(propertyData);
    console.log('Property added successfully');
    financesState.addPropertyDialogBoxOpen = false;
  } catch (error) {
    console.error('Error adding property:', error);
    Notify.create({
      message: 'Failed to add property',
      icon: 'error',
      textColor: 'white',
      color: 'red-5',
    });
  }
};

const handleCancelCreateProperty = () => {
  console.log('Cancel Create Property Function Here');
  financesState.addPropertyDialogBoxOpen = false;
};

const handlePropertyNicknameChange = (value: string | null) => {
  console.log('Nickname changed to:', value);
  handleOnValueChange(value, propertyNickname);
};

const handleAddressChange = (value: Address | null) => {
  console.log(value);
  if (value !== null) {
    handleOnAddressChange(value, address);
  }
};

const handlePropertyUseChange = (value: string) => {
  console.log('Property Use Change Triggered');
  propertyUseModel.value = value;
};

const handlePropertyTypeChange = (value: string) => {
  console.log('Property Type Change triggered');
  propertyTypeModel.value = value;
};
</script>

<style scoped>
.address-input-container {
  display: flex;
  flex-direction: column;
}

.top-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.bottom-row {
  margin-top: 2vh;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.form-input {
  width: 100%;
}
</style>
