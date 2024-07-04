<template>
  <div v-bind="$attrs">
    <h1 class="my-6 mb-8 text-xl font-bold">{{ t('title') }}</h1>
    <hr />
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 py-3 w-full">
      <div class="border border-gray-300 p-3">
        <p class="border-b border-b-gray-300 pb-1.5 mb-3">
          <span class="rounded-full inline-flex justify-center font-bold text-sm text-white bg-primary-500 w-5 mr-3">
            1
          </span>
          {{ t('basic.title') }}
        </p>
        {{ t('basic.description') }} <sup>1</sup>
      </div>
      <div class="border border-gray-300 p-3">
        <p class="border-b border-b-gray-300 pb-1.5 mb-3">
          <span class="rounded-full inline-flex justify-center font-bold text-sm text-white bg-primary-500 w-5 mr-3">
            2
          </span>
          {{ t('comfort.title') }}
        </p>
        {{ t('comfort.description') }}
      </div>
      <div class="border border-gray-300 p-3">
        <p class="border-b border-b-gray-300 pb-1.5 mb-3">
          <span class="rounded-full inline-flex justify-center font-bold text-sm text-white bg-primary-500 w-5 mr-3">
            3
          </span>
          {{ t('advanced.title') }}
        </p>
        {{ t('advanced.description') }}
        <p class="border-t border-t-gray-300 pt-1.5 mt-3 font-bold">
          {{ t('advanced.option1') }}
        </p>
      </div>
      <div class="border border-gray-300 p-3">
        <p class="border-b border-b-gray-300 pb-1.5 mb-3">
          <span class="rounded-full inline-flex justify-center font-bold text-sm text-white bg-primary-500 w-5 mr-3">
            4
          </span>
          {{ t('ultimate.title') }}
        </p>
        {{ t('ultimate.description') }}
        <p class="border-t border-t-gray-300 pt-1.5 mt-3 font-bold">
          {{ t('ultimate.option1') }}
        </p>
        <p class="border-t border-t-gray-300 pt-1.5 mt-3 font-bold">
          {{ t('ultimate.option2') }}
        </p>
      </div>
    </div>
    <BaseButton
        class="mt-6 mb-1"
        :text="t('buttonText')"
        @click.prevent="store.toggleForm()"
        v-if="!store.isShowForm && (typeof isShowButton === 'undefined' || !isShowButton)"
    />
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useCalculatorStore } from '@/store';
import BaseButton from '@/components/atoms/BaseButton.vue';

defineProps<{
  isShowButton?: boolean;
}>();

const { t } = useI18n();
const store = useCalculatorStore();
</script>

<i18n>
fr:
  title: Choix entre 4 forfaits BMW Service Inclusive
  basic:
    title: BASIC
    description: Tous les entretiens réguliers
  comfort:
    title: COMFORT
    description: Entretien, réparations et remplacement des pièces soumises à l’usure
  advanced:
    title: ADVANCED
    description: Entretien, réparations et remplacement des pièces soumises à l’usure
    option1: + Voiture de remplacement
  ultimate:
    title: ULTIMATE
    description: Entretien, réparations et remplacement des pièces soumises à l’usure
    option1: + Voiture de remplacement
    option2: + Pneus été et/ou hiver étoilés MINI avec garantie, montage, stockage et permutation inclus
  buttonText: Calculez le prix de votre contrat d'entretien
nl:
  title: Keuze uit 4 BMW Service Inclusive pakketten
  basic:
    title: BASIC
    description: Alle reguliere onderhoudswerkzaamheden
  comfort:
    title: COMFORT
    description: Onderhoud, herstellingen en vervanging van slijtage-onderdelen
  advanced:
    title: ADVANCED
    description: Onderhoud, herstellingen en vervanging van slijtage-onderdelen
    option1: + Vervangwagen
  ultimate:
    title: ULTIMATE
    description: Onderhoud, herstellingen en vervanging van slijtage-onderdelen
    option1: + Vervangwagen
    option2: + MINI Zomer en/of winter sterbanden met garantie, inclusief plaatsing, stockage en wissel
  buttonText: Bereken de prijs van uw pakket
de:
  title: Keuze uit 2 BMW Service Inclusive pakketten
  basic:
    title: BASIC
    description: Alle reguliere onderhoudswerkzaamheden
  comfort:
    title: COMFORT
    description: Onderhoud, herstellingen en vervanging van slijtage-onderdelen
  buttonText: Bereken de prijs van uw pakket
</i18n>