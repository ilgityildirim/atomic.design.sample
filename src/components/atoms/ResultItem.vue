<template>
  <div class="flex flex-col sm:grid sm:grid-cols-8 gap-4 py-5 border-b border-b-gray-200">
    <div class="font-bold sm:font-normal sm:col-span-4 text-left text-lg md:text-base md:text-center" v-html="title" />
    <div class="sm:text-center">
      <span class="sm:hidden w-16 inline-block mr-4 font-bold">Basic</span>
      <img v-if="typeof basic === 'boolean'" :src="check" class="w-4 h-4 inline-block" alt="basic" />
      <span v-else>{{ basic }}</span>
    </div>
    <div class="sm:text-center">
      <span class="sm:hidden w-16 inline-block mr-4 font-bold">Comfort</span>
      <img v-if="typeof comfort === 'boolean'" :src="check" class="w-4 h-4 inline-block" alt="basic" />
      <span v-else>{{ comfort }}</span>
    </div>
    <div class="sm:text-center">
      <span class="sm:hidden w-16 inline-block mr-4 font-bold">Advanced</span>
      <img v-if="typeof advanced === 'boolean'" :src="check" class="w-4 h-4 inline-block" alt="basic" />
      <span v-else>{{ advanced }}</span>
    </div>
    <div class="sm:text-center">
      <span class="sm:hidden w-16 inline-block mr-4 font-bold">Ultimate</span>
      <img v-if="typeof ultimate === 'boolean'" :src="check" class="w-4 h-4 inline-block" alt="basic" />
      <span v-else>{{ ultimate }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';
import check from '@/assets/tick-bmw.svg';

defineProps<{
  title: string;
  basic: string|boolean;
  comfort: string|boolean;
  advanced: string|boolean;
  ultimate: string|boolean;
}>();
</script>
