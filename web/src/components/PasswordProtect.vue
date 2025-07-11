<template>
    <div class="-mt-6 flex flex-col gap-6 min-h-screen justify-center items-center bg-gray-50 dark:bg-gray-900">
        <div class="w-full max-w-2xl p-8 space-y-6 bg-white dark:bg-gray-800 rounded-xl shadow-lg">
            <div class="text-center space-y-2">
                <h1 class="text-4xl font-semibold text-blue-600 dark:text-blue-400">Money Minder</h1>
                <p class="text-xl text-gray-600 dark:text-gray-400">Your Personal Finance Dashboard</p>
                <p class="text-sm text-gray-500 dark:text-gray-500">Track expenses, analyze spending patterns, and make
                    smarter financial decisions</p>
            </div>

            <div class="text-center space-y-2">
                <h2 class="text-md font-medium text-gray-900 dark:text-white">Welcome back,
                    <a href="https://github.com/sounishnath003" target="_blank"
                        class="text-blue-600 dark:text-blue-400 hover:underline">@github/sounishnath003</a>
                </h2>
                <p class="text-sm text-gray-600 dark:text-gray-400">Enter your email and password to access your
                    financial
                    dashboard</p>
            </div>

            <form class="space-y-4" @submit.prevent="onLogin">
                <div class="relative">
                    <input type="password" v-model="password" placeholder="Enter your password" required class="w-full px-4 py-3 rounded-lg border border-gray-300 dark:border-gray-600 
                               dark:bg-gray-700 dark:text-white 
                               focus:ring-2 focus:ring-blue-500 focus:border-blue-500
                               placeholder-gray-400 dark:placeholder-gray-500" />
                </div>
                <button type="submit" :disabled="auth.loading"
                    class="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 rounded">
                    <span v-if="auth.loading">Logging in...</span>
                    <span v-else>Login</span>
                </button>
                <div v-if="auth.error" class="text-red-500 text-center text-sm">{{ auth.error }}</div>
                <p class="text-xs text-gray-500 dark:text-gray-400 text-center">
                    Your data is protected with end-to-end encryption
                </p>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useAuthStore } from '../store/auth.store';

const auth = useAuthStore();
const password = ref('');

async function onLogin() {
    await auth.login(password.value);
}

// Optionally clear error on input
watch([password], () => { auth.error = null; });
</script>