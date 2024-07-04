<template>
  <Form v-if="!store.isShowResults" />
  <Results v-if="store.isShowResults && store.isSet"/>
</template>

<script setup lang="ts">
import { useCalculatorStore } from '@/store';
import Form from '@/components/templates/Form.vue';
import Results from '@/components/organisms/Results.vue';

const store = useCalculatorStore();
</script>