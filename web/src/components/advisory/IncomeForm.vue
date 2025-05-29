<template>
    <div class="bg-white dark:bg-gray-800 p-6 rounded-xl shadow border border-gray-200 dark:border-gray-700">
        <h2 class="text-2xl font-bold text-gray-700 dark:text-blue-400 mb-6 flex items-center">
            <span class="mr-2">💰</span>
            Income Details
        </h2>

        <div class="space-y-6">
            <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Monthly Income (₹)</label>
                <div class="relative">
                    <span class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 dark:text-gray-400">₹</span>
                    <input type="number" v-model="localData.income" required
                        class="w-full rounded-lg bg-gray-50 dark:bg-gray-700 p-3 pl-8 text-gray-900 dark:text-white border border-gray-300 dark:border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors" />
                </div>
            </div>

            <div class="space-y-2">
                <div class="flex justify-between items-center">
                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Savings Percentage</label>
                    <span class="text-blue-600 dark:text-blue-400 font-medium">{{ localData.savingsPct }}%</span>
                </div>
                <input type="range" v-model="localData.savingsPct" min="0" max="100" required
                    class="w-full h-2 bg-gray-200 dark:bg-gray-700 rounded-lg appearance-none cursor-pointer accent-blue-500" />
                <div class="dark:text-white text-gray-800 text-right font-semibold">
                    <p> {{ INRRuppe.format((localData.income * localData.savingsPct) / 100) }} </p>
                </div>
            </div>

            <div class="grid grid-cols-2 gap-6">
                <div class="space-y-2">
                    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Hours/Day</label>
                    <input type="number" v-model="localData.hoursPerDay"
                        class="w-full rounded-lg bg-gray-50 dark:bg-gray-700 p-3 text-gray-900 dark:text-white border border-gray-300 dark:border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors" />
                </div>
                <div class="space-y-2">
                    <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Days/Week</label>
                    <input type="number" v-model="localData.daysPerWeek"
                        class="w-full rounded-lg bg-gray-50 dark:bg-gray-700 p-3 text-gray-900 dark:text-white border border-gray-300 dark:border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors" />
                </div>
            </div>
        </div>

        <button @click="emitNext"
            class="mt-8 w-full bg-blue-700 hover:bg-blue-800 p-3 rounded-lg font-medium text-white transition-colors flex items-center justify-center gap-2">
            Next
            <span class="text-lg">→</span>
        </button>
    </div>
</template>

<script setup>
import { reactive } from 'vue';

const props = defineProps(['formData'])
const emit = defineEmits(['next', 'update:form-data'])

const localData = reactive({ ...props.formData })

const emitNext = () => {
    emit('update:form-data', { ...localData })
    emit('next')
}

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
