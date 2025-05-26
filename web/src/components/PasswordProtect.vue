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
                <p class="text-sm text-gray-600 dark:text-gray-400">Enter your password to access your financial
                    dashboard</p>
            </div>

            <div class="space-y-4">
                <div class="relative">
                    <input type="password" v-model="inputPassword" placeholder="Enter your password" class="w-full px-4 py-3 rounded-lg border border-gray-300 dark:border-gray-600 
                               dark:bg-gray-700 dark:text-white 
                               focus:ring-2 focus:ring-blue-500 focus:border-blue-500
                               placeholder-gray-400 dark:placeholder-gray-500" />
                </div>
                <p class="text-xs text-gray-500 dark:text-gray-400 text-center">
                    Your data is protected with end-to-end encryption
                </p>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const emit = defineEmits(['inputPasswordEvent']);
const inputPassword = ref('');

watch(inputPassword, (newValue) => {
    emit('inputPasswordEvent', newValue);
});
</script>