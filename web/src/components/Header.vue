<template>
    <div class="flex items-center justify-around w-full gap-5">
        <RouterLink to="/" class="font-medium lg:text-2xl font-sans dark:text-white text-blue-600 hidden sm:inline">
            Money Minder
        </RouterLink>
        <div class="flex-1 hidden sm:block"></div>
        <RouterLink v-for="route in routings" :to="route.path" :class="[
            $route.path === route.path
                ? 'text-blue-600 dark:text-blue-400 font-medium'
                : 'hover:text-blue-600 dark:hover:text-blue-400'
        ]">
            {{ route.name }}
        </RouterLink>

        <button @click="auth.logout"
            class="bg-red-600 hover:bg-red-700 px-2 cursor-pointer py-1 rounded text-white font-medium">
            Logout
        </button>
    </div>
</template>

<script setup>
import { RouterLink } from 'vue-router';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../store/auth.store';

const route = useRoute();

const auth = useAuthStore();

const routings = [
    { name: 'Home', path: '/' },
    { name: 'Analytics', path: '/analytics' },
    { name: 'Everything', path: '/everthing' },
    { name: 'Advisory', path: '/finance-advisory' },
]
</script>