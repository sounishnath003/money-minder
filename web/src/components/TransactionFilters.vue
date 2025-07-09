<template>
    <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 mb-6">
        <!-- Header with Toggle Button -->
        <div class="p-4 border-b border-gray-200 dark:border-gray-700">
            <div class="flex items-center justify-between">
                <h3 class="text-lg font-medium text-gray-900 dark:text-white">Filters</h3>
                <button @click="isCollapsed = !isCollapsed"
                    class="flex items-center space-x-2 px-3 py-2 text-sm font-medium text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md transition-colors">
                    <span>{{ isCollapsed ? 'Show' : 'Hide' }} Filters</span>
                    <svg class="w-4 h-4 transition-transform duration-200" :class="{ 'rotate-180': !isCollapsed }"
                        fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                    </svg>
                </button>
            </div>
        </div>

        <!-- Collapsible Content -->
        <div v-show="!isCollapsed" class="p-4">
            <div class="flex flex-col lg:flex-row gap-4">
                <!-- Date Range Filter -->
                <div class="flex-1">
                    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Date Range</label>
                    <div class="flex flex-col lg:flex-row gap-2">
                        <input type="date" v-model="filters.startDate"
                            class="flex-1 px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                        <span class="flex items-center text-gray-500 dark:text-gray-400">to</span>
                        <input type="date" v-model="filters.endDate"
                            class="flex-1 px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                    </div>
                </div>

                <!-- Category Filter -->
                <div class="flex-1">
                    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Category</label>
                    <select v-model="filters.categoryID"
                        class="w-full px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-700 text-gray-900 dark:text-white text-sm focus:ring-2 focus:ring-blue-500 focus:border-blue-500">
                        <option value="">All Categories</option>
                        <option v-for="category in categories" :key="category.id" :value="category.id">
                            {{ category.category }}
                        </option>
                    </select>
                </div>

                <!-- Payment Method Filter -->
                <div class="flex-1">
                    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Payment
                        Method</label>
                    <div class="space-y-2 max-h-32 overflow-y-auto">
                        <label class="flex items-center space-x-2 cursor-pointer">
                            <input type="checkbox" :value="'UPI'" v-model="filters.paymentMethod"
                                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700">
                            <span class="text-sm text-gray-700 dark:text-gray-300">UPI</span>
                        </label>
                        <label class="flex items-center space-x-2 cursor-pointer">
                            <input type="checkbox" :value="'Cash'" v-model="filters.paymentMethod"
                                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700">
                            <span class="text-sm text-gray-700 dark:text-gray-300">Cash</span>
                        </label>
                        <label class="flex items-center space-x-2 cursor-pointer">
                            <input type="checkbox" :value="'Bank Deposit'" v-model="filters.paymentMethod"
                                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700">
                            <span class="text-sm text-gray-700 dark:text-gray-300">Bank Deposit</span>
                        </label>
                        <label class="flex items-center space-x-2 cursor-pointer">
                            <input type="checkbox" :value="'Credit Card'" v-model="filters.paymentMethod"
                                class="rounded border-gray-300 text-blue-600 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700">
                            <span class="text-sm text-gray-700 dark:text-gray-300">Credit Card</span>
                        </label>
                    </div>
                    <div v-if="filters.paymentMethod.length > 0" class="mt-2 text-xs text-gray-500 dark:text-gray-400">
                        Selected: {{ filters.paymentMethod.join(', ') }}
                    </div>
                </div>

                <!-- Clear Filters Button -->
                <div class="flex items-end">
                    <button @click="$emit('clear-filters')"
                        class="px-4 py-2 bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300 rounded-md hover:bg-gray-200 dark:hover:bg-gray-600 transition-colors text-sm font-medium">
                        Clear Filters
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

defineProps({
    filters: {
        type: Object,
        required: true
    },
    categories: {
        type: Array,
        required: true
    }
});

defineEmits(['clear-filters']);

// Collapsible state
const isCollapsed = ref(true);
</script>