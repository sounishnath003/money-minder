<template>
    <div class="gap-8 rounded-md">
        <div>
            <h2 class="font-medium text-xl text-blue-600 dark:text-white">&bull; Total Balance</h2>
        </div>
        <div v-if="isLoading" class="text-center my-auto font-medium py-4">
            Loading...
        </div>
        <template v-else>
            <div>
                <span class="font-bold cursor-pointer text-2xl md:text-4xl p-2"
                    :class="determineSign == '+' ? 'text-green-600' : 'text-red-600'"> {{ INRRuppe.format(balance)
                    }}</span>
            </div>
            <div class="grid grid-cols-2 gap-5">
                <div class="text-md md:text-xl font-medium" v-for="tx in spendByTransactionTypes"
                    :id="tx.transactionType">
                    {{ tx.transactionType }}&colon; <span class="font-semibold"
                        :class="tx.transactionType === 'Income' ? 'text-green-600' : 'text-red-600'">
                        {{ tx.amount }}</span>
                </div>
            </div>
            <SpendByCategory class="w-full overflow-x-scroll" />
        </template>
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

import SpendByCategory from './SpendByCategory.vue';

const transactionStore = useTransactionStore();
const isLoading = ref(true);

const transactions = computed(() => transactionStore.allTransactions);

// On component mount fetch all transactions
onMounted(async () => {
    try {
        await transactionStore.getTransactions();
    } catch (error) {
        console.error('Failed to load transactions:', error);
    } finally {
        isLoading.value = false;
    }
})

// Computed properties
const balance = computed(() => transactions.value.map((trans, index) => trans.transactionType === 'Income' ? trans.amount : -trans.amount).reduce((acc, curr) => acc + curr, 0).toFixed(2))

// To determine the sign symbol and CSS
const determineSign = computed(() => +balance.value >= 0 ? '+' : '-');

// Group by the transaction type for the visualizations
const spendByTransactionTypes = computed(() => {
    const summary = {
        'Income': 0,
        'Expense': 0
    };

    transactions.value.forEach((tx) => {
        if (tx.transactionType in summary) {
            summary[tx.transactionType] += tx.amount
        }
    });

    return [
        { transactionType: "Income", amount: INRRuppe.format(summary.Income), sign: '+' },
        { transactionType: "Expense", amount: INRRuppe.format(summary.Expense), sign: '-' },
    ]
});

// Utilities
const USDollar = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
});

const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});

</script>