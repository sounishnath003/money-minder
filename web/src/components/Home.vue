<template>
    <div v-if="inputPassword === sitePassword" class="flex flex-col gap-4">
        <div class="grid lg:grid-cols-2 gap-8 lg:gap-20 my-6">
            <AddTransaction />
            <Balance class="flex flex-col justify-start items-center order-first lg:order-last" />
        </div>
        <ListTransactions class="my-20" />
    </div>
    <PasswordProtect v-else @inputPasswordEvent="updateInputPassword" />
</template>

<script setup>
import AddTransaction from './AddTransaction.vue';
import Balance from './Balance.vue';
import ListTransactions from './ListTransactions.vue';
import { ref } from 'vue';
import { useTransactionStore } from '../store/transaction.store';
import PasswordProtect from './PasswordProtect.vue';

// define the store
const transactionStore = useTransactionStore();
const inputPassword = ref('');

// Add a simple password input to see all transactions and modules
const sitePassword = transactionStore.sitePassword;

const updateInputPassword = (password) => {
    inputPassword.value = password;
}
</script>