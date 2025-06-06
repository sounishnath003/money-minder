<template>
    <div class="gap-4 rounded-md">
        <div>
            <h2 class="font-medium text-xl text-blue-600 dark:text-white text-center">&bull; Total Balance</h2>
            <div class="text-xs text-gray-600 dark:text-gray-400 -mb-3">Note: Hover over amounts to view them</div>
        </div>
        <div v-if="isLoading" class="text-center my-auto font-medium py-4">
            Loading...
        </div>
        <template v-else>
            <div>
                <span class="font-bold cursor-pointer text-3xl md:text-5xl p-2 flex items-center gap-2"
                    :class="determineSign == '+' ? 'text-green-600' : 'text-red-600'"
                    @mouseenter="isBalanceVisible = true" @mouseleave="isBalanceVisible = false">
                    <span v-if="isBalanceVisible" class="transition-opacity duration-600">{{ INRRuppe.format(balance)
                    }}</span>
                    <span v-else class="transition-opacity duration-600">{{ '•'.repeat(INRRuppe.format(balance).length)
                    }}</span>
                </span>
            </div>
            <div class="grid grid-cols-2 md:grid-cols-3 gap-5">
                <div class="text-lg md:text-2xl font-medium" v-for="tx in spendByTransactionTypes"
                    :id="tx.transactionType">
                    {{ tx.transactionType }}&colon; <span class="font-semibold cursor-pointer flex items-center gap-2"
                        :class="getTransactionTypeColor(tx.transactionType)"
                        @mouseenter="toggleAmountVisibility(tx.transactionType, true)"
                        @mouseleave="toggleAmountVisibility(tx.transactionType, false)">
                        <span v-if="isAmountVisible(tx.transactionType)" class="transition-opacity duration-600">{{
                            tx.amount }}</span>
                        <span v-else class="transition-opacity duration-600">{{ 'X'.repeat(tx.amount.length - 4)
                            }}</span>
                    </span>
                </div>
            </div>
            <DailySpendTracker />
            <div class="w-full overflow-x-scroll">
                <SpendByCategory />
            </div>
        </template>
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

import SpendByCategory from './SpendByCategory.vue';
import DailySpendTracker from './DailySpendTracker.vue';

const transactionStore = useTransactionStore();

const isLoading = ref(true);
const transactions = computed(() => transactionStore.allTransactions);

// Visibility states
const isBalanceVisible = ref(false);
const visibleAmounts = ref({
    Income: false,
    Expense: false,
    Investments: false
});

// Toggle functions
const toggleAmountVisibility = (type, isVisible) => {
    visibleAmounts.value[type] = isVisible;
};

const isAmountVisible = (type) => {
    return visibleAmounts.value[type];
};

// Get color based on transaction type
const getTransactionTypeColor = (type) => {
    switch (type) {
        case 'Income':
            return 'text-green-600';
        case 'Expense':
            return 'text-red-600';
        case 'Investments':
            return 'text-blue-600';
        default:
            return '';
    }
};

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
const balance = computed(() => transactions.value.map((trans) => {
    switch (trans.transactionType) {
        case 'Income':
            return trans.amount;
        case 'Expense':
            return -trans.amount;
        case 'Investments':
            return -trans.amount;
        default:
            return 0;
    }
}).reduce((acc, curr) => acc + curr, 0).toFixed(2))

// To determine the sign symbol and CSS
const determineSign = computed(() => +balance.value >= 0 ? '+' : '-');

// Group by the transaction type for the visualizations
const spendByTransactionTypes = computed(() => {
    const summary = {
        'Income': 0,
        'Expense': 0,
        'Investments': 0
    };

    transactions.value.forEach((tx) => {
        if (tx.transactionType in summary) {
            summary[tx.transactionType] += tx.amount
        }
    });

    return [
        { transactionType: "Income", amount: INRRuppe.format(summary.Income), sign: '+' },
        { transactionType: "Expense", amount: INRRuppe.format(summary.Expense), sign: '-' },
        { transactionType: "Investments", amount: INRRuppe.format(summary.Investments), sign: '+' }
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