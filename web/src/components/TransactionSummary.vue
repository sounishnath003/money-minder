<template>
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-4 mb-6">
        <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">Transaction Summary</h3>

        <div class="flex flex-col lg:flex-row gap-4">
            <!-- Income Card -->
            <div
                class="flex-1 bg-green-50 dark:bg-green-900/20 rounded-lg p-4 border border-green-200 dark:border-green-800">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm font-medium text-green-600 dark:text-green-400">Total Income</p>
                        <p class="text-2xl font-bold text-green-700 dark:text-green-300">
                            {{ INRRuppe.format(totalIncome) }}
                        </p>
                    </div>
                    <div class="w-12 h-12 bg-green-100 dark:bg-green-800 rounded-full flex items-center justify-center">
                        <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="none" stroke="currentColor"
                            viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4">
                            </path>
                        </svg>
                    </div>
                </div>
            </div>

            <!-- Expense Card -->
            <div class="flex-1 bg-red-50 dark:bg-red-900/20 rounded-lg p-4 border border-red-200 dark:border-red-800">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm font-medium text-red-600 dark:text-red-400">Total Expense</p>
                        <p class="text-2xl font-bold text-red-700 dark:text-red-300">
                            {{ INRRuppe.format(totalExpense) }}
                        </p>
                    </div>
                    <div class="w-12 h-12 bg-red-100 dark:bg-red-800 rounded-full flex items-center justify-center">
                        <svg class="w-6 h-6 text-red-600 dark:text-red-400" fill="none" stroke="currentColor"
                            viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4"></path>
                        </svg>
                    </div>
                </div>
            </div>

            <!-- Net Amount Card -->
            <div
                class="flex-1 bg-blue-50 dark:bg-blue-900/20 rounded-lg p-4 border border-blue-200 dark:border-blue-800">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm font-medium text-blue-600 dark:text-blue-400">Net Amount</p>
                        <p class="text-2xl font-bold"
                            :class="netAmount >= 0 ? 'text-blue-700 dark:text-blue-300' : 'text-red-700 dark:text-red-300'">
                            {{ netAmount >= 0 ? '+' : '' }}{{ INRRuppe.format(netAmount) }}
                        </p>
                    </div>
                    <div class="w-12 h-12 rounded-full flex items-center justify-center"
                        :class="netAmount >= 0 ? 'bg-blue-100 dark:bg-blue-800' : 'bg-red-100 dark:bg-red-800'">
                        <svg class="w-6 h-6"
                            :class="netAmount >= 0 ? 'text-blue-600 dark:text-blue-400' : 'text-red-600 dark:text-red-400'"
                            fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path>
                        </svg>
                    </div>
                </div>
            </div>

            <!-- Transaction Count Card -->
            <div class="flex-1 bg-gray-50 dark:bg-gray-700 rounded-lg p-4 border border-gray-200 dark:border-gray-600">
                <div class="flex items-center justify-between">
                    <div>
                        <p class="text-sm font-medium text-gray-600 dark:text-gray-400">Transactions</p>
                        <p class="text-2xl font-bold text-gray-700 dark:text-gray-300">
                            {{ transactionCount }}
                        </p>
                    </div>
                    <div class="w-12 h-12 bg-gray-100 dark:bg-gray-600 rounded-full flex items-center justify-center">
                        <svg class="w-6 h-6 text-gray-600 dark:text-gray-400" fill="none" stroke="currentColor"
                            viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2">
                            </path>
                        </svg>
                    </div>
                </div>
            </div>
        </div>

        <!-- Progress Bar for Income vs Expense -->
        <div class="mt-6">
            <div class="flex items-center justify-between mb-2">
                <span class="text-sm font-medium text-gray-700 dark:text-gray-300">Income vs Expense Ratio</span>
                <span class="text-sm text-gray-500 dark:text-gray-400">
                    {{ totalIncome > 0 ? Math.round((totalIncome / (totalIncome + totalExpense)) * 100) : 0 }}% Income
                </span>
            </div>
            <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-3">
                <div class="bg-gradient-to-r from-green-500 to-red-500 h-3 rounded-full transition-all duration-300"
                    :style="{ width: `${totalIncome > 0 ? Math.round((totalIncome / (totalIncome + totalExpense)) * 100) : 0}%` }">
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
    transactions: {
        type: Array,
        required: true
    }
});

// Currency formatter
const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});

// Computed properties for summary
const totalIncome = computed(() => {
    return props.transactions
        .filter(t => t.transactionType === 'Income')
        .reduce((sum, t) => sum + t.amount, 0);
});

const totalExpense = computed(() => {
    return props.transactions
        .filter(t => t.transactionType === 'Expense')
        .reduce((sum, t) => sum + t.amount, 0);
});

const netAmount = computed(() => {
    return totalIncome.value - totalExpense.value;
});

const transactionCount = computed(() => {
    return props.transactions.length;
});
</script>