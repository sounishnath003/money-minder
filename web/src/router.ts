import { createRouter, createWebHistory } from "vue-router";
import Home from "./components/Home.vue";
import Analytics from "./components/Analytics.vue";


const routes = [
    { path: '/', component: Home },
    { path: '/analytics', component: Analytics },
];

export const router = createRouter({
    history: createWebHistory('/'),
    routes
})