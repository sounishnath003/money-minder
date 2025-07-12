<template>
    <header class="w-full mx-auto sm:px-1 dark:bg-gray-900 dark:border-gray-700">
        <div class="flex justify-between items-center h-16 md:h-10">
            <!-- Logo - visible on all screen sizes -->
            <RouterLink to="/" class="flex items-center">
                <span class="font-bold text-2xl sm:text-2xl font-sans dark:text-white text-blue-600">
                    Money Minder
                </span>
            </RouterLink>

            <!-- Desktop Navigation -->
            <nav class="hidden md:flex items-center space-x-8">
                <RouterLink v-for="route in routings" :key="route.path" :to="route.path" :class="[
                    'font-medium transition-colors duration-200',
                    $route.path === route.path
                        ? 'text-blue-600 dark:text-blue-400'
                        : 'text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400'
                ]">
                    {{ route.name }}
                </RouterLink>

                <button @click="auth.logout"
                    class="bg-red-600 hover:bg-red-700 px-4 py-2 rounded-md text-sm font-medium text-white transition-colors duration-200 cursor-pointer">
                    Logout
                </button>
            </nav>

            <!-- Mobile menu button -->
            <div class="md:hidden flex items-center space-x-2">
                <button @click="auth.logout"
                    class="bg-red-600 hover:bg-red-700 px-3 py-1.5 rounded text-sm font-medium text-white transition-colors duration-200 cursor-pointer">
                    Logout
                </button>

                <button @click="isMobileMenuOpen = !isMobileMenuOpen"
                    class="inline-flex items-center justify-center p-2 rounded-md text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 hover:bg-gray-100 dark:hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-blue-500"
                    aria-expanded="false">
                    <span class="sr-only">Open main menu</span>
                    <!-- Icon when menu is closed -->
                    <svg v-if="!isMobileMenuOpen" class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none"
                        viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16" />
                    </svg>
                    <!-- Icon when menu is open -->
                    <svg v-else class="block h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>
        </div>

        <!-- Mobile Navigation Menu -->
        <div v-show="isMobileMenuOpen" class="md:hidden border-t border-gray-200 dark:border-gray-700">
            <div class="px-2 pt-2 pb-3 space-y-1">
                <RouterLink v-for="route in routings" :key="route.path" :to="route.path"
                    @click="isMobileMenuOpen = false" :class="[
                        'block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200',
                        $route.path === route.path
                            ? 'text-blue-600 dark:text-blue-400 bg-blue-50 dark:bg-blue-900/20'
                            : 'text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 hover:bg-gray-50 dark:hover:bg-gray-800'
                    ]">
                    {{ route.name }}
                </RouterLink>
            </div>
        </div>
    </header>
</template>

<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router';
import { useRoute } from 'vue-router';
import { useAuthStore } from '../store/auth.store';

const route = useRoute();
const auth = useAuthStore();
const isMobileMenuOpen = ref(false);

const routings = [
    { name: 'Home', path: '/' },
    { name: 'Analytics', path: '/analytics' },
    { name: 'Everything', path: '/everthing' },
    { name: 'Advisory', path: '/finance-advisory' },
];
</script>