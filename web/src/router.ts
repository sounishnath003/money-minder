import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        path: '/',
        component: () => import("./components/Home.vue")
    },
    {
        path: '/analytics',
        component: () => import("./components/Analytics.vue")
    },
    {
        path: '/everthing',
        component: () => import("./components/view360/View360.vue")
    },
    {
        path: '/finance-advisory',
        component: () => import("./components/FinanceAdvisory.vue")
    },
];

export const router = createRouter({
    history: createWebHistory('/'),
    routes
})