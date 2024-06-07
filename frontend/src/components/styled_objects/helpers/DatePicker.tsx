import { Ref } from 'vue';

export const handleUpdateDates = (
  date: { from: string; to: string } | string,
  startDate: Ref<string | null>,
  endDate: Ref<string | null>
) => {
  if (typeof date === 'string') {
    startDate.value = date;
    endDate.value = date;
  } else {
    startDate.value = date.from;
    endDate.value = date.to;
  }
};
