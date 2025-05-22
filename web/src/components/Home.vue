<template>
    <div>
        <div v-if="inputPassword === sitePassword" class="flex flex-col gap-4">
            <div class="grid lg:grid-cols-2 gap-8 lg:gap-20 my-6">
                <AddTransaction />
                <Balance class="flex flex-col justify-start items-center order-first lg:order-last" />
            </div>
            <ListTransactions class="my-20" />
        </div>
        <div v-else class="flex flex-col gap-4 min-h-screen justify-center items-center">
            <h2 class="font-medium text-xl text-red-600 dark:text-white">Enter password to access</h2>
            <input type="password" v-model="inputPassword" placeholder="Type password here"
                class="w-full p-2 rounded-lg border border-gray-300 dark:border-gray-700 dark:bg-gray-800 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
        </div>
    </div>
</template>

<script setup>
import AddTransaction from './AddTransaction.vue';
import Balance from './Balance.vue';
import ListTransactions from './ListTransactions.vue';
import { ref } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

// define the store
const transactionStore = useTransactionStore();
const inputPassword = ref('');

// Add a simple password input to see all transactions and modules
const sitePassword = transactionStore.sitePassword;
</script>