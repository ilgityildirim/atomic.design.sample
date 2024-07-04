<template>
  <div class="pt-6">
    <Welcome class="mb-2" />
    <Form v-if="store.isShowForm" />
  </div>
</template>

<script setup lang="ts">
import Welcome from '@/components/molecules/Welcome.vue';
import Form from '@/components/organisms/Form.vue';
import { useCalculatorStore } from '@/store';

const store = useCalculatorStore();
</script>