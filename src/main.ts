import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { setupI18n } from './config/i18n';
import App from './App.vue';

const i18n = setupI18n({
  legacy: false,
});

const store = createPinia();

createApp(App)
  .use(i18n)
  .use(store)
  .mount('#app')
;
