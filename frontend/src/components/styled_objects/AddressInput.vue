<template>
  <q-card-section class="address-input-container">
    <div class="top-row">
      <q-input
        v-model="address.streetNumber"
        label="Street Number"
        @update:model-value="updateStreetNumber"
        clearable
        outlined
        class="form-input"
        :rules="[(val) => !!val || 'Street Number is required']"
      />
      <q-input
        v-model="address.streetName"
        label="Street Name"
        @update:model-value="updateStreetName"
        clearable
        outlined
        class="form-input"
        :rules="[(val) => !!val || 'Street Name is required']"
      />
    </div>
    <div class="bottom-row">
      <q-input
        v-model="address.city"
        label="City"
        @update:model-value="updateCity"
        clearable
        outlined
        class="form-input"
        :rules="[(val) => !!val || 'City is required']"
      />
      <q-input
        v-model="address.state"
        label="State"
        @update:model-value="updateState"
        clearable
        outlined
        class="form-input"
        :rules="[(val) => !!val || 'State is required']"
      />
      <q-input
        v-model="address.zip"
        label="Zip Code"
        @update:model-value="updateZip"
        clearable
        outlined
        class="form-input"
        :rules="[
          (val) => !!val || 'Zip Code is required',
          (val) => !isNaN(val) || 'Zip Code must be a number',
        ]"
      />
    </div>
  </q-card-section>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { Address } from '../../models/finances/property';

const props = defineProps<{
  onAddressChange: (address: Address | null) => void;
  initialAddress?: Address;
}>();

const address = ref<Address>({
  id: '',
  streetNumber: 0,
  streetName: '',
  city: '',
  state: '',
  zip: 0,
});

if (props.initialAddress) {
  address.value = { ...props.initialAddress };
}

const updateStreetNumber = (value: string | number | null) => {
  if (typeof value === 'number') {
    address.value.streetNumber = parseInt(value.toString(), 10);
  } else {
    address.value.streetNumber = value ? parseInt(value, 10) : 0;
  }
  props.onAddressChange(address.value);
};

const updateStreetName = (value: string | number | null) => {
  address.value.streetName = value?.toString() ?? '';
  props.onAddressChange(address.value);
};

const updateCity = (value: string | number | null) => {
  address.value.city = value?.toString() ?? '';
  props.onAddressChange(address.value);
};

const updateState = (value: string | number | null) => {
  address.value.state = value?.toString() ?? '';
  props.onAddressChange(address.value);
};

const updateZip = (value: string | number | null) => {
  if (typeof value === 'number') {
    address.value.zip = parseInt(value.toString(), 10);
  } else {
    address.value.zip = value ? parseInt(value, 10) : 0;
  }
  props.onAddressChange(address.value);
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
</style>
