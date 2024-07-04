<template>
  <Welcome class="mb-20" :isShowButton="false" />
  <Box id="results" class="break-after-avoid-page">
    <h2 class="font-bold text-lg mb-6">
      {{ t('title', { series: store.car.series, fuelType: store.car.fuelType.toLowerCase() }) }}
    </h2>
    <p>
      {{ t('description', { duration: store.premium.duration, distance: store.premium.distance.toString().replace(/\B(?=(\d{3})+(?!\d))/g, '.') }) }}
    </p>
    <PremiumDetailsForm
        optionsWrapperClass="flex flex-wrap space-x-2.5"
        optionClass="w-72"
        noTitle
        autoSelect

    />
    <ResultsListing class="mt-10" />
    <div class="flex flex-wrap items-end justify-start space-y-5 sm:justify-center md:justify-end" data-html2canvas-ignore="true">
      <BaseAnchor
          :text="t('buttonContactDealer')"
          :href="dealerLink"
          class="mr-4"
          target="_blank"
      />
      <BaseButton
          :text="t('buttonDownload')"
          class="mr-4"
          @click.prevent="export2PDF"
      />
      <BaseButton
          :text="t('buttonNewCalculation')"
          @click="store.reset()"
      />
    </div>
    <p class="text-sm mt-10">
      {{ t('footer', { vat: vat }) }}
    </p>
    <div class="flex mt-6 mx-3 p-3 border border-gray-300">
      <h3 class="text-base font-bold w-1/2">
        <sup>1</sup>&nbsp;{{ t('maintenanceTitle') }}
      </h3>
      <ul class="block w-1/2">
        <li class="grid">
          <span class="text-sm mt-3 font-bold">{{ t('maintenance') }}</span>
          <span class="text-sm mt-3">{{ t('maintenanceDescription') }}</span>
        </li>
        <li class="grid">
          <span class="text-sm mt-3 font-bold">{{ t('electricMaintenanceTitle') }}</span>
          <span class="text-sm mt-3">{{ t('electricMaintenanceDescription') }}</span>
        </li>
      </ul>
    </div>
  </Box>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useI18n } from 'vue-i18n';
import html2pdf from 'html2pdf.js';
import { French } from '@/config/i18n';
import { useCalculatorStore } from '@/store';
import Box from '@/components/atoms/Box.vue';
import BaseButton from '@/components/atoms/BaseButton.vue';
import BaseAnchor from '@/components/atoms/BaseAnchor.vue';
import Welcome from '@/components/molecules/Welcome.vue';
import ResultsListing from '@/components/molecules/Results.vue';
import PremiumDetailsForm from '@/components/molecules/PremiumDetailsForm.vue';
import {COUNTRY, LUXEMBOURG} from "@/config";

const { t, locale } = useI18n();
const store = useCalculatorStore();

const dealerLink = computed(() => {
  if (locale.value === French) {
    return 'https://www.bmw.be/fr/fastlane/dealer-locator.html';
  }
  return 'https://www.bmw.be/nl/fastlane/dealer-locator.html';
});

const export2PDF = () => {
  const el = document.getElementById('results');
  if (!el) {
    return;
  }

  html2pdf(el, {
    margin: [ 0, 0, 0, 0 ],
    filename: `BMW_${toTitleCase(store.car.series)}_${store.car.fuelType.toLowerCase()}.pdf`,
    enableLinks: false,
    image: { type: 'jpeg', quality: 0.98 },
    html2canvas: { scale: 4 },
    jsPDF: { unit: 'px', format: [el.offsetWidth, el.offsetHeight], orientation: 'portrait' },
  });
};

const toTitleCase = (str: string): string => {
  return str.replace(/\w\S*/g, (txt) => txt.charAt(0).toUpperCase() + txt.substring(1).toLowerCase());
};

const vat = computed(() => {
  if (COUNTRY === LUXEMBOURG) {
    return 17;
  }
  return 21;
});
</script>

<i18n>
fr:
  title: 'Résultat du calcul de votre prix pour une {series}, BMW avec type de carburant {fuelType}:'
  description: >
    Les prix sont indicatifs et peuvent varier en fonction de la motorisation de votre véhicule. Contactez votre partenaire BMW pour obtenir un devis personnalisé pour votre véhicule.
  footer: >
    Les prix mentionnés sont valables à partir du 01/03/2024 ({vat} % TVAC). Spécifications et les prix sont sous réserve d'erreurs et de modifications.  Pour plus d'infos sur nos tarifs, veuillez vous adresser à votre Partenaire BMW Service.
  buttonContactDealer: Contactez votre partenaire BMW
  buttonDownload: Télécharger le calcul
  buttonNewCalculation: Réaliser un nouveau calcul
  maintenanceTitle: Qu'est-ce qui est inclus dans l'entretien ?
  maintenance: Entretien
  maintenanceExplanation: Entretien suivant les directives de BMW
  maintenanceDescription: >
    Changement d'huile moteur, changement du filtre à huile, remplacement du ou des filtre(s) à air, remplacement du ou des filtre(s) intérieur(s), changement de filtre(s) à carburant, remplacement des bougies d'allumage, changement du liquide de frein.
  electricMaintenanceTitle: Entretien véhicule électrique suivant les directives de BMW
  electricMaintenanceDescription: Remplacement du filtre d’habitacle et du liquide de frein selon les directives du constructeur.
nl:
  title: 'Resultaat prijsberekening voor een BMW {series}, met brandstoftype {fuelType}:'
  description: >
    Prijzen zijn indicatief en kunnen variëren afhankelijk van de motorisatie van uw voertuig. Contacteer uw BMW Partner voor een gepersonaliseerde offerte voor uw voertuig.
  footer: >
    Vermelde prijzen zijn geldig vanaf 01/03/2024 ({vat}% BTWi). Specificaties en prijzen zijn onder voorbehoud van fouten en wijzigingen. Voor meer informatie over de tarieven kan u steeds terecht bij uw BMW (Service) Partner.
  buttonContactDealer: Contacteer uw BMW Partner
  buttonDownload: Download berekening
  buttonNewCalculation: Maak nieuwe berekening
  maintenanceTitle: Wat zit er allemaal in een onderhoud?
  maintenance: Onderhoud
  maintenanceExplanation: Entretien suivant les directives de BMW
  maintenanceDescription: >
    Onderhoudswerkzaamheden volgens de richtlijnen van de constructeur: Motorolie verversen, vervangen van oliefilter, brandstoffilter, luchtfilter, interieurfilter, ontstekingsbougies en remvloeistof.
  electricMaintenanceTitle: Onderhoud voor elektrische modellen
  electricMaintenanceDescription: Vervangen van interieurfilter en remvloeistof volgens de richtlijnen van de constructeur.
de:
  title: 'Resultaat prijsberekening voor een BMW {series}, met brandstoftype {fuelType}:'
  description: >
    De prijs is gebaseerd op een duurtijd van {duration} maanden en {distance} km te rekenen vanaf de startdatum van uw contract.
  footer: >
    Vermelde prijzen zijn geldig vanaf 01/03/2024 ({vat}% BTWi). Specificaties en prijzen zijn onder voorbehoud van fouten en wijzigingen. Voor meer informatie over de tarieven kan u steeds terecht bij uw BMW (Service) Partner.
  buttonContactDealer: Contacteer uw BMW Partner
  buttonDownload: Download berekening
  buttonNewCalculation: Maak nieuwe berekening
  maintenanceTitle: Wat zit er allemaal in een onderhoud?
  maintenance: Onderhoud
  maintenanceExplanation: Entretien suivant les directives de BMW
  maintenanceDescription: >
    Onderhoudswerkzaamheden volgens de richtlijnen van de constructeur: Motorolie verversen, vervangen van oliefilter, brandstoffilter, luchtfilter, interieurfilter, ontstekingsbougies en remvloeistof.
  electricMaintenanceTitle: Onderhoud voor elektrische modellen
  electricMaintenanceDescription: Vervangen van interieurfilter en remvloeistof volgens de richtlijnen van de constructeur.
</i18n>