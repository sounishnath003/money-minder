import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './router'
import { createPinia } from 'pinia'

import VueApexCharts from 'vue3-apexcharts';

const app = createApp(App)
// Register a central store for the vue application
app.use(createPinia());
// Register the routes into the application
app.use(router);
// Define the apex charts into the vue components
// app.use(VueApexCharts);
app.component('apexchart', VueApexCharts)
// Mount the application with the DOM #id
app.mount('#app');
