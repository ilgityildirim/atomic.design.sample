<template>
  <Box v-bind="$attrs">
    <h2 class="font-bold text-lg mb-6">{{ t('title') }}</h2>
    <p>{{ t('description') }}</p>
    <div class="my-8 grid md:grid-cols-2 gap-4">
      <CarDetailsForm />
      <PremiumDetailsForm />
    </div>
    <BaseButton :text="t('buttonText')" :disabled="!store.isSet" @click="store.toggleResults()" />
  </Box>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useCalculatorStore } from '@/store';
import Box from '@/components/atoms/Box.vue';
import BaseButton from '@/components/atoms/BaseButton.vue';
import CarDetailsForm from '@/components/molecules/CarDetailsForm.vue';
import PremiumDetailsForm from '@/components/molecules/PremiumDetailsForm.vue';

const { t } = useI18n();
const store = useCalculatorStore();
</script>

<i18n>
fr:
  title: Calculez le prix de votre contrat d'entretien.
  description: >
    Déterminez votre contrat d'entretien en fonction de votre usage personnel. Choisissez la durée et le kilométrage souhaités, afin d'obtenir un contrat sur mesure.
  buttonText: Calculez le prix de votre contrat d'entretien
nl:
  title: Bereken de prijs van uw pakket.
  description: Stem uw onderhoudscontract helemaal af op uw persoonlijke rijgedrag. Kies de gewenste looptijd en kilometrage. Zo bekomt u een onderhoudscontract dat perfect op uw maat berekend is.
  buttonText: Bereken de prijs van uw pakket
de:
  title: Bereken de prijs van uw pakket.
  description: Stem uw onderhoudscontract helemaal af op uw persoonlijke rijgedrag. Kies de gewenste looptijd en kilometrage. Zo bekomt u een onderhoudscontract dat perfect op uw maat berekend is.
  buttonText: Bereken de prijs van uw pakket
</i18n>