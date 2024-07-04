<template>
  <Calculator />
</template>

<script setup lang="ts">
import Calculator from '@/components/pages/Calculator.vue';
import { useI18n } from 'vue-i18n';
import { AvailableLanguages, DefaultLanguage } from '@/config/i18n';
import { onMounted } from 'vue';

const { locale } = useI18n();
const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);

if (urlParams.has('language')) {
  const localeValue = urlParams.get('language') || DefaultLanguage;
  locale.value = AvailableLanguages.includes(localeValue) ? localeValue : DefaultLanguage;
}

onMounted(()=> {
  window.parent.postMessage('[iFrameResizerChild]Ready','*');
});
</script>

<!--suppress CssUnusedSymbol -->
<style lang="postcss">
@tailwind base;
@layer base {
  img {
    @apply inline-block;
  }
}
@tailwind components;
@tailwind utilities;

@font-face {
  font-family: 'BMWTypeNext-Light';
  src: url('assets/fonts/BMWTypeNext-Light.woff2') format('woff2');
}

@font-face {
  font-family: 'BMWTypeNext-Bold';
  src: url('assets/fonts/BMWTypeNext-Bold.woff2') format('woff2');
}

html {
  @apply bg-white scroll-smooth md:scroll-auto;
}

#app {
  @apply px-0.5;
}
</style>