<template>
    <div class="text-center px-6 py-2 bg-gray-100 dark:bg-gray-800 rounded-lg shadow text-sm">
        <div class="flex flex-col gap-1">
            <div>
                <span class="font-medium text-blue-700 dark:text-blue-300">Today's Spending: </span>
                <span class="font-semibold text-lg"
                    :class="todaysSpend > yesterdaysSpend * 0.80 ? 'text-red-600' : 'text-gray-800 dark:text-white'">
                    {{ INRRuppe.format(todaysSpend) }}
                </span>
            </div>
            <div v-if="spendComparison" class="text-xs font-medium">
                <span :class="spendComparison.percentage >= 0 ? 'text-red-600' : 'text-green-600'">
                    {{ spendComparison.percentage >= 0 ? '&uarr;' : '&darr;' }}
                    {{ Math.abs(spendComparison.percentage).toFixed(1) }}%
                    {{ spendComparison.percentage >= 0 ? 'more' : 'less' }} than yesterday
                </span>
                <span class="text-gray-500 ml-2">
                    (Yesterday: {{ INRRuppe.format(spendComparison.yesterdayAmount) }})
                </span>
            </div>
        </div>
    </div>

</template>


<script setup>
import { ref, computed } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

const transactionStore = useTransactionStore();
const transactions = computed(() => transactionStore.allTransactions);

const todaysSpend = computed(() => transactionStore.allTransactions.filter((tx) => (tx.createdAt.toDateString() === new Date().toDateString() && tx.transactionType === 'Expense')).reduce((prev, tx) => prev + tx.amount, 0))

// Get yesterday's date
const getYesterdayDate = () => {
    const yesterday = new Date();
    yesterday.setDate(yesterday.getDate() - 1);
    return yesterday;
}

// Calculate yesterday's spending
const yesterdaysSpend = computed(() => {
    const yesterday = getYesterdayDate();
    return transactionStore.allTransactions
        .filter((tx) => tx.createdAt.toDateString() === yesterday.toDateString())
        .reduce((prev, tx) => prev + tx.amount, 0);
});

// Calculate spending comparison
const spendComparison = computed(() => {
    if (yesterdaysSpend.value === 0) return null;

    const percentage = ((todaysSpend.value - yesterdaysSpend.value) / yesterdaysSpend.value) * 100;

    return {
        percentage,
        yesterdayAmount: yesterdaysSpend.value
    };
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