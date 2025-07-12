import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { router } from './router'
import { createPinia } from 'pinia'

// Dynamically import ApexCharts to reduce initial bundle size
const loadApexCharts = async () => {
    const { default: VueApexCharts } = await import('vue3-apexcharts');
    return VueApexCharts;
};

const app = createApp(App)
// Register a central store for the vue application
app.use(createPinia());
// Register the routes into the application
app.use(router);

// Load ApexCharts dynamically
loadApexCharts().then(VueApexCharts => {
    app.component('apexchart', VueApexCharts);
});

// Mount the application with the DOM #id
app.mount('#app');
