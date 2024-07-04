<template>
  <button
      :id="uid"
      :class="[
        'px-10 py-4 bg-primary-500 text-white font-bold hover:opacity-75',
        disabled ? 'opacity-50 cursor-not-allowed hover:opacity-50' : '',
      ]"
      :aria-invalid="disabled ? true : undefined"
      v-bind="$attrs"
  >
    {{ text }}
  </button>
</template>

<!--suppress JSUnusedGlobalSymbols -->
<script setup lang="ts">
import UniqueID from '@/utils/uniqueID';

interface Props {
  disabled?: boolean;
  text?: string;
}

const props = defineProps<Props>();
const uid = UniqueID().getID().toString();
</script>
