import { Ref } from 'vue';

export const handleOnValueChange = (
  inputValue: string | null,
  outputValue: Ref<string | null>
) => {
  outputValue.value = inputValue;
};
