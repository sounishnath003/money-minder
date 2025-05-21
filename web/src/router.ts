import { createRouter, createWebHistory } from "vue-router";
import Analytics from "./components/Analytics.vue";
import Home from "./components/Home.vue";


const routes = [
    { path: '/', component: Home },
    { path: '/analytics', component: Analytics },
];

export const router = createRouter({
    history: createWebHistory('/'),
    routes
})