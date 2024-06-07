<template>
  <div class="q-pa-md">
    <q-btn-dropdown split color="cyan" push no-caps @click="onMainClick">
      <template v-slot:label>
        <div class="row items-center no-wrap">
          <q-icon left name="map" />
          <div class="text-center">Custom<br />Content</div>
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
            <q-item-label>{{ item.name }}</q-item-label>
            <q-item-label caption>{{ item.date }}</q-item-label>
          </q-item-section>
          <q-item-section side>
            <q-icon name="info" color="amber" />
          </q-item-section>
        </q-item>
      </q-list>
    </q-btn-dropdown>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

// Define the type for list items
interface ListItem {
  id: number;
  name: string;
  date: string;
}

const props = defineProps<{
  listItems: ListItem[];
}>();

const onMainClick = () => {
  console.log('Clicked on main button');
};

const onItemClick = (item: ListItem) => {
  console.log('Click on an item', item);
};
</script>
