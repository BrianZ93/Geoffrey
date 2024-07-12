import { Ref } from 'vue';
import { Address } from '../../../models/finances/property';

export const handleOnAddressChange = (
  updatedAddress: Address,
  addressRef: Ref<Address | null>
) => {
  if (addressRef !== null) {
    addressRef.value = { ...updatedAddress };
  }
};
