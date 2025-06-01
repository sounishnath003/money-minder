<template>
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 w-full">
        <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Average Daily Spend</h3>
            <p class="text-2xl font-bold text-blue-600">{{ INRRuppe.format(averageDailySpend) }}</p>
        </div>
        <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Highest Spending Category</h3>
            <p class="text-xl font-bold"
                :class="highestSpendingCategory.category === 'Salary' ? 'text-green-600' : 'text-red-600'">{{
                    highestSpendingCategory.category }}</p>
            <p class="text-sm text-gray-600">
                {{ INRRuppe.format(highestSpendingCategory.amount) }}
                <span class="text-xs">(excl. salary, house rents and mutual funds categories)</span>
            </p>
        </div>
        <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Total Investments</h3>
            <p class="text-2xl font-bold text-blue-600">{{ INRRuppe.format(investments) }}</p>
        </div>
        <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
            <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Savings Rate</h3>
            <p class="text-2xl font-bold" :class="savingsRate >= 0 ? 'text-green-600' : 'text-red-600'">
                {{ savingsRate }}%
            </p>
        </div>
    </div>
</template>

<script setup>
defineProps({
    averageDailySpend: {
        type: Number,
        required: true,
        default: 0
    },
    highestSpendingCategory: {
        type: Object,
        required: true,
        default: () => ({ category: 'N/A', amount: 0 })
    },
    savingsRate: {
        type: [Number, String],
        required: true,
        default: '0'
    },
    investments: {
        type: Number,
        required: true,
        default: 0
    }
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
