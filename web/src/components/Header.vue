<template>
    <div class="flex items-center justify-between w-full gap-5">
        <RouterLink to="/" class="font-medium lg:text-2xl font-sans dark:text-white text-blue-600"> Money Minder
        </RouterLink>
        <div class="flex-1"></div>
        <RouterLink v-for="route in routings" :to="route.path" :class="[
            $route.path === route.path
                ? 'text-blue-600 dark:text-blue-400 font-medium'
                : 'hover:text-blue-600 dark:hover:text-blue-400'
        ]">
            {{ route.name }}
        </RouterLink>
    </div>
</template>

<script setup>
import { RouterLink } from 'vue-router';
import { useRoute } from 'vue-router';

const route = useRoute();

const routings = [
    { name: 'Home', path: '/' },
    { name: 'Analytics', path: '/analytics' },
    { name: 'Finance Advisory', path: '/finance-advisory' },
]
</script>