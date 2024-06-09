<template>
  <q-btn-dropdown
    split
    color="primary"
    push
    no-caps
    @click="onMainClick"
    :style="style"
    glossy
    :disable-main-btn="activeItem.title === ''"
  >
    <template v-slot:label>
      <div class="row items-center no-wrap">
        <q-icon left :name="icon" />
        <div class="text-center">{{ activeItem.title }}</div>
      </div>
    </template>

    <q-list>
      <q-item
        v-for="item in listItems"
        :key="item.id"
        clickable
        v-close-popup
        @click="() => onItemClick(item)"
      >
        <q-item-section avatar>
          <q-avatar icon="folder" color="primary" text-color="white" />
        </q-item-section>
        <q-item-section>
          <q-item-label>{{ item.title }}</q-item-label>
          <q-item-label caption>{{ item.caption }}</q-item-label>
        </q-item-section>
        <q-item-section side>
          <q-icon name="info" color="amber" />
        </q-item-section>
      </q-item>
    </q-list>
  </q-btn-dropdown>
</template>

<script setup lang="ts">
import { ListItem } from './helpers/DropdownList';

const props = defineProps<{
  listItems: ListItem[];
  icon: string;
  onSelect: (listItem: ListItem) => void;
  onModify: () => void;
  activeItem: ListItem;
  width?: string;
  height?: string;
}>();

const style = `height: ${props.height}vh; width: ${props.width}vw;`;

const onMainClick = () => {
  console.log('Clicked on main button');
  props.onModify();
};

const onItemClick = (item: ListItem) => {
  if (item !== null) {
    props.onSelect(item);
  }
};
</script>
