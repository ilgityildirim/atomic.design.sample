import { WritableComputedRef } from 'vue';
import { createI18n, I18n, I18nOptions } from 'vue-i18n';
import storage from '@/utils/storage';
import { COUNTRY, LUXEMBOURG, StorageNames } from '@/config/index';

export const French = 'fr';
export const Dutch = 'nl';
export const German = 'de';
export const DefaultLanguage = French;
export const AvailableLanguagesBE = [French, Dutch];
export const AvailableLanguagesLU = [French, German];

const getAvailableLanguages = () => {
  if (COUNTRY?.toLowerCase() === LUXEMBOURG?.toLowerCase()) {
    return AvailableLanguagesLU;
  }
  return AvailableLanguagesBE;
}
export const AvailableLanguages = getAvailableLanguages();

export function setI18nLanguage(i18n: I18n, locale: string) {
  if (i18n.mode === 'legacy') {
    i18n.global.locale = locale;
  } else {
    (i18n.global.locale as WritableComputedRef<string>).value = locale;
  }
  // Store selected locale. This will be used for API calls.
  storage.setItem(StorageNames.language, locale);
  // Setting up language settings for headers. Utils like axios will automatically use the right language
  (document.querySelector('html') as HTMLHtmlElement).setAttribute('lang', locale);
}

export function setupI18n(options: I18nOptions = { locale: DefaultLanguage }): I18n {
  options = {
    ...{ locale: DefaultLanguage },
    ...options,
  }
  const i18n = createI18n(options);
  setI18nLanguage(i18n, options.locale as string);
  return i18n;
}
