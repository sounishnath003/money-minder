import { createRouter, createWebHistory } from "vue-router";
import Analytics from "./components/Analytics.vue";
import Home from "./components/Home.vue";
import FinanceAdvisory from "./components/FinanceAdvisory.vue";


const routes = [
    { path: '/', component: Home },
    { path: '/analytics', component: Analytics },
    { path: '/finance-advisory', component: FinanceAdvisory },
];

export const router = createRouter({
    history: createWebHistory('/'),
    routes
})