<template>
  <Box class="mb-10" v-bind="$attrs">
    <ResultHeader />
    <ResultItem
        :title="t('monthlyPayment')"
        :basic="t('notAvailable')"
        :comfort="comfortMonthlyFee !== '-' ? comfortMonthlyFee : t('notAvailable')"
        :advanced="advancedMonthlyFee !== '-' ? advancedMonthlyFee : t('notAvailable')"
        :ultimate="ultimateMonthlyFee !== '-' ? ultimateMonthlyFee : t('notAvailable')"
    />
    <ResultItem
        :title="t('singlePayment')"
        :basic="hasBasic ? basicFee : t('notAvailable')"
        :comfort="comfortUpfrontFee !== '-' ? comfortUpfrontFee : t('notAvailable')"
        :advanced="advancedUpfrontFee !== '-' ? advancedUpfrontFee : t('notAvailable')"
        :ultimate="ultimateUpfrontFee !== '-' ? ultimateUpfrontFee : t('notAvailable')"
    />
    <ResultItem
        :title="t('duration')"
        :basic="hasBasic ? duration : t('notAvailable')"
        :comfort="comfortUpfrontFee !== '-' ? duration : t('notAvailable')"
        :advanced="advancedUpfrontFee !== '-' ? duration : t('notAvailable')"
        :ultimate="ultimateUpfrontFee !== '-' ? duration : t('notAvailable')"
    />
    <ResultItem
        :title="t('distance')"
        :basic="hasBasic ? distance : t('notAvailable')"
        :comfort="comfortUpfrontFee !== '-' ? distance : t('notAvailable')"
        :advanced="advancedUpfrontFee !== '-' ? distance : t('notAvailable')"
        :ultimate="ultimateUpfrontFee !== '-' ? distance : t('notAvailable')"
    />
    <ResultItem
        :title="`${t('maintenance')} <sup>1</sup>`"
        :basic="hasBasic ? true : t('notAvailable')"
        :comfort="comfortUpfrontFee !== '-' ? true : t('notAvailable')"
        :advanced="advancedUpfrontFee !== '-' ? true : t('notAvailable')"
        :ultimate="ultimateUpfrontFee !== '-' ? true : t('notAvailable')"
    />
    <ResultItem
        :title="t('wearReplacement')"
        :basic="t('notIncluded')"
        :comfort="comfortUpfrontFee !== '-' ? true : t('notAvailable')"
        :advanced="advancedUpfrontFee !== '-' ? true : t('notAvailable')"
        :ultimate="ultimateUpfrontFee !== '-' ? true : t('notAvailable')"
    />
    <ResultItem
        :title="t('vehiculeReplacement')"
        :basic="t('notIncluded')"
        :comfort="t('notIncluded')"
        :advanced="advancedUpfrontFee !== '-' ? true : t('notAvailable')"
        :ultimate="ultimateUpfrontFee !== '-' ? true : t('notAvailable')"
    />
    <ResultItem
        :title="t('tires')"
        :basic="t('notIncluded')"
        :comfort="t('notIncluded')"
        :advanced="t('notIncluded')"
        :ultimate="ultimateUpfrontFee !== '-' ? true : t('notAvailable')"
    />
  </Box>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import Box from '@/components/atoms/Box.vue';
import ResultHeader from '@/components/atoms/ResultHeader.vue';
import ResultItem from '@/components/atoms/ResultItem.vue';
import { computed } from 'vue';
import { useCalculatorStore } from '@/store';
import { data, Unlimited } from '@/config';
import {
  AdvancedPackage,
  BasicPackage,
  ComfortPackage,
  PackageLanguages,
  PackageName,
  UltimatePackage
} from '@/config/types';

const { t, locale } = useI18n();
const store = useCalculatorStore();

const basic = computed<BasicPackage>(() => {
  const localeVal = locale.value as PackageLanguages;
  return data[PackageName.Basic].filter((item) => {
    return item.model === store.car.model && item.series[localeVal] === store.car.series &&
        item.fuelType[localeVal] === store.car.fuelType &&
        typeof item.price[`${store.premium.duration}/${store.premium.distance}`] !== 'undefined';
  })?.[0] || {};
});
const comfort = computed<ComfortPackage>(() => {
  const localeVal = locale.value as PackageLanguages;
  return data[PackageName.Comfort].filter((item) => {
    return item.model === store.car.model && item.series[localeVal] === store.car.series &&
        item.fuelType[localeVal] === store.car.fuelType &&
        typeof item.monthly[`${store.premium.duration}/${store.premium.distance}`] !== 'undefined';
  })?.[0] || {};
});
const advanced = computed<AdvancedPackage>(() => {
  const localeVal = locale.value as PackageLanguages;
  return data[PackageName.Advanced].filter((item) => {
    return item.model === store.car.model && item.series[localeVal] === store.car.series &&
        item.fuelType[localeVal] === store.car.fuelType &&
        typeof item.monthly[`${store.premium.duration}/${store.premium.distance}`] !== 'undefined';
  })?.[0] || {};
});
const ultimate = computed<UltimatePackage>(() => {
  const localeVal = locale.value as PackageLanguages;
  return data[PackageName.Ultimate].filter((item) => {
    return item.model === store.car.model && item.series[localeVal] === store.car.series &&
        item.fuelType[localeVal] === store.car.fuelType &&
        typeof item.monthly[`${store.premium.duration}/${store.premium.distance}`] !== 'undefined';
  })?.[0] || {};
});
const duration = computed(() => {
  return t('months', { months: store.premium.duration });
});
const distance = computed(() => {
  if (store.premium.distance === Unlimited){
    return t('unlimited');
  }
  return t('km', { distance: store.premium.distance.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.') });
});
const comfortMonthlyFee = computed(() => {
  return comfort.value.monthly?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || '-';
});
const comfortUpfrontFee = computed(() => {
  return comfort.value.upfront?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || '-';
});
const advancedMonthlyFee = computed(() => {
  return advanced.value.monthly?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || '-';
});
const advancedUpfrontFee = computed(() => {
  return advanced.value.upfront?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || '-';
});
const ultimateMonthlyFee = computed(() => {
  return ultimate.value.monthly?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || '-';
});
const ultimateUpfrontFee = computed(() => {
  return ultimate.value.upfront?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || '-';
});
const basicFee = computed(() => {
  return basic.value.price?.[`${store.premium.duration}/${store.premium.distance}`]?.toLocaleString(locale.value, { style: 'currency', currency: 'EUR'}) || t('notAvailable');
});
const hasBasic = computed(() => {
  return basic.value.price?.[`${store.premium.duration}/${store.premium.distance}`] > 0 || false;
});
</script>

<i18n>
fr:
  monthlyPayment: Paiement mensuel
  singlePayment: Paiement unique
  duration: Durée
  distance: Kilométrage
  maintenance: Entretien
  wearReplacement: Remplacement des pièces soumises à l'usure
  vehiculeReplacement: Voiture de remplacement
  tires: Pneumatiques
  notIncluded: Non compris
  notAvailable: Non disponible
  months: '{months} mois'
  km: '{distance} km'
  unlimited: Illimité
nl:
  monthlyPayment: Maandelijkse betaling
  singlePayment: Eenmalige betaling
  duration: Looptijd
  distance: Kilometers
  maintenance: Onderhoud
  wearReplacement: Vervanging slijtage onderdelen
  vehiculeReplacement: Vervangwagen
  tires: Banden
  notIncluded: Niet inbegrepen
  notAvailable: Niet beschikbaar
  months: '{months} maanden'
  km: '{distance} km'
  unlimited: Onbeperkt
de:
  monthlyPayment: Maandelijkse betaling
  singlePayment: Eenmalige betaling
  duration: Looptijd
  distance: Kilometers
  maintenance: Onderhoud
  wearReplacement: Vervanging slijtage onderdelen
  repairAfterWarranty: Herstellingen na garantieperiode
  notIncluded: Niet inbegrepen
  notAvailable: Niet beschikbaar
  months: '{months} maanden'
  km: '{distance} km'
  unlimited: Onbeperkt
</i18n>