<template>
    <footer
        class="py-4 px-6 text-center text-sm text-gray-600 dark:text-gray-400 border-t border-gray-200 dark:border-gray-700">
        <div class="flex flex-col md:flex-row justify-center items-center gap-2 md:gap-4">
            <span>© 2025 Money Minder</span>
            <span class="hidden md:inline">&bull;</span>
            <span>Version: <a :href="`https://github.com/sounishnath003/money-minder/commit/${version}`" target="_blank"
                    class="text-blue-600 dark:text-blue-400 hover:underline">{{ version }}</a></span>
            <span class="hidden md:inline">&bull;</span>
            <span>By <a href="https://github.com/sounishnath003" target="_blank"
                    class="text-blue-600 dark:text-blue-400 hover:underline font-medium">Sounish Nath</a></span>
            <span class="hidden md:inline">&bull;</span>
            <span>Made with <span class="text-red-500">❤️</span> using <b> Vue, Golang & Google Cloud</b> </span>
        </div>
    </footer>
</template>

<script setup>
import { ref } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

const transactionStore = useTransactionStore();

const version = ref(transactionStore.commitHash);
</script>