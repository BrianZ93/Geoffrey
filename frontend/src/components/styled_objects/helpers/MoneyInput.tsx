import { Ref } from 'vue';

// Helper function to handle value change for money input
export const handleOnMoneyValueChange = (
  inputValue: string | null,
  outputValue: Ref<number | null>
) => {
  // Function to parse the input string to a number
  const parseUSD = (value: string): number | null => {
    const numericValue = parseFloat(value.replace(/[^0-9.-]+/g, ''));
    return isNaN(numericValue) ? null : numericValue;
  };

  // Parse the input string and update the outputValue
  if (inputValue !== null) {
    const parsedValue = parseUSD(inputValue);
    outputValue.value = parsedValue; // Update the numeric value (number or null)
  } else {
    outputValue.value = null; // Clear the value if input is null
  }
};
