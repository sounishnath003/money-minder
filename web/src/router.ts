import { createRouter, createWebHistory } from "vue-router";
import Analytics from "./components/Analytics.vue";
import Home from "./components/Home.vue";
import FinanceAdvisory from "./components/FinanceAdvisory.vue";
import View360 from "./components/view360/View360.vue";


const routes = [
    { path: '/', component: Home },
    { path: '/analytics', component: Analytics },
    { path: '/everthing', component: View360 },
    { path: '/finance-advisory', component: FinanceAdvisory },
];

export const router = createRouter({
    history: createWebHistory('/'),
    routes
})