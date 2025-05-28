<template>
    <div class="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-xl border border-gray-200 dark:border-gray-700">
        <div class="text-6xl mb-4 animate-bounce">
            <span v-if="score < 30">😞</span>
            <span v-else-if="score < 70">🤔</span>
            <span v-else>🎉</span>
        </div>

        <h2 class="text-4xl font-bold mb-4"
            :class="score < 30 ? 'text-red-700' : score < 70 ? 'text-yellow-700' : 'text-green-700'">
            {{ resultMessage }}
        </h2>

        <div class="bg-gray-50 dark:bg-gray-700 p-4 rounded-lg mb-6">
            <h3 class="text-xl font-semibold mb-2 text-gray-900 dark:text-white">{{ formData.goalName }}</h3>
            <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">₹{{ formData.goalCost.toLocaleString() }}</p>
        </div>

        <div class="space-y-4 text-left">
            <div class="bg-gray-50 dark:bg-gray-700 p-4 rounded-lg">
                <h4 class="font-semibold text-gray-900 dark:text-white mb-2">Time to Save</h4>
                <div class="grid grid-cols-3 gap-2 text-center">
                    <div>
                        <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ years.toFixed(1) }}</p>
                        <p class="text-sm text-gray-600 dark:text-gray-400">Years</p>
                    </div>
                    <div>
                        <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ months }}</p>
                        <p class="text-sm text-gray-600 dark:text-gray-400">Months</p>
                    </div>
                    <div>
                        <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ weeks.toFixed(1) }}</p>
                        <p class="text-sm text-gray-600 dark:text-gray-400">Weeks</p>
                    </div>
                </div>
            </div>

            <div class="bg-gray-50 dark:bg-gray-700 p-4 rounded-lg">
                <h4 class="font-semibold text-gray-900 dark:text-white mb-2">Hourly Effort</h4>
                <p class="text-2xl font-bold text-blue-600 dark:text-blue-400">{{ hours.toFixed(1) }} hours</p>
            </div>
        </div>

        <div class="mt-8">
            <button @click="$emit('prev')"
                class="w-full bg-gray-100 hover:bg-gray-200 dark:bg-gray-600 dark:hover:bg-gray-700 px-6 py-3 rounded-lg font-medium text-gray-900 dark:text-white transition-colors flex items-center justify-center gap-2">
                <span class="text-lg">←</span>
                Back to Goal
            </button>
        </div>
    </div>
</template>

<script setup>
import { reactive } from 'vue';

const props = defineProps(['formData'])

const { income, savingsPct, hoursPerDay, daysPerWeek, goalCost, goalYears, impactLevel } = props.formData

const monthlySavings = (income * savingsPct) / 100
const totalMonths = goalCost / monthlySavings
const hours = totalMonths * 4.33 * daysPerWeek * hoursPerDay
const days = totalMonths * 30
const weeks = totalMonths * 4.33
const months = Math.ceil(totalMonths)
const years = Math.round(months / 12)

const score = Math.max(0, 100 - (years / goalYears) * 50 + (impactLevel * 10))
const resultMessage = score < 30 ? "Nay, Not the best choice" : score < 70 ? "Think twice, Savings is real" : "Good deal, Go for it!"
</script>
