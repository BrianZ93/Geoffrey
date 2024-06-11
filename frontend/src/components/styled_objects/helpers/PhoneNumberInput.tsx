import { Ref } from 'vue';

export const handleOnPhoneNumberValueChange = (
  inputValue: string | null,
  outputValue: Ref<string | null>
) => {
  outputValue.value = inputValue;
};
