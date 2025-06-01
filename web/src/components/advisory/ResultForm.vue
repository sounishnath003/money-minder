<template>
    <div
        class="bg-white dark:bg-gray-800 p-4 rounded-xl shadow border border-gray-200 dark:border-gray-700 max-w-2xl mx-auto">
        <div class="flex flex-col lg:flex-row text-center lg:text-left items-center justify-center gap-4 mb-4">
            <div class="text-4xl animate-bounce">
                <span v-if="score < 30">😞</span>
                <span v-else-if="score < 70">🤔</span>
                <span v-else>🎉</span>
            </div>
            <div class="flex-grow"></div>
            <div v-if="result">
                <h2 class="text-xl md:text-2xl font-bold" :class="result.color">
                    {{ result.verdict }}
                </h2>
                <p class="text-sm text-gray-600 dark:text-gray-400">{{ result.message }}</p>
            </div>
        </div>

        <!-- Rest of the template remains unchanged -->
        <div class="grid grid-cols-2 gap-4 mb-4">
            <div class="bg-gray-50 dark:bg-gray-700 p-3 rounded-lg">
                <h3 class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Goal Details</h3>
                <p class="font-semibold text-gray-900 dark:text-white">{{ formData.goalName }}</p>
                <p class="text-sm text-gray-600 dark:text-gray-400">{{ formData.goalType }}</p>
            </div>
            <div class="bg-gray-50 dark:bg-gray-700 p-3 rounded-lg">
                <h3 class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-1">Cost</h3>
                <p class="text-xl font-bold text-blue-600 dark:text-white">₹{{ formData.goalCost.toLocaleString() }}
                </p>
                <p class="text-sm text-gray-600 dark:text-gray-400">Value for {{ formData.goalYears }} years</p>
            </div>
        </div>

        <div class="grid grid-cols-2 gap-4">
            <div class="bg-gray-50 dark:bg-gray-700 p-3 rounded-lg">
                <h4 class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-2">Time to Save</h4>
                <div class="grid grid-cols-3 gap-2 text-center">
                    <div>
                        <p class="text-lg font-bold text-blue-600 dark:text-white">{{ years.toFixed(1) }}</p>
                        <p class="text-xs text-gray-600 dark:text-gray-400">Years</p>
                    </div>
                    <div>
                        <p class="text-lg font-bold text-blue-600 dark:text-white">{{ months }}</p>
                        <p class="text-xs text-gray-600 dark:text-gray-400">Months</p>
                    </div>
                    <div>
                        <p class="text-lg font-bold text-blue-600 dark:text-white">{{ weeks.toFixed(1) }}</p>
                        <p class="text-xs text-gray-600 dark:text-gray-400">Weeks</p>
                    </div>
                </div>
            </div>

            <div class="bg-gray-50 dark:bg-gray-700 p-3 rounded-lg">
                <h4 class="text-sm font-medium text-gray-600 dark:text-gray-400 mb-2">Effort Required</h4>
                <p class="text-lg font-bold text-blue-600 dark:text-white">{{ hours.toFixed(1) }} hours</p>
                <p class="text-xs text-gray-600 dark:text-gray-400">Based on {{ hoursPerDay }}hrs/day, {{ daysPerWeek }}
                    days/week</p>
            </div>
        </div>

        <div class="mt-4">
            <button @click="$emit('prev')"
                class="w-full bg-gray-100 cursor-pointer hover:bg-gray-200 dark:bg-gray-600 dark:hover:bg-gray-700 px-4 py-2 rounded-lg font-medium text-gray-900 dark:text-white transition-colors flex items-center justify-center gap-2">
                <span class="text-lg">←</span>
                Back to Goal
            </button>
        </div>
    </div>
</template>
<script setup>
import { computed, ref } from 'vue';

const props = defineProps({
    formData: {
        type: Object,
        required: true,
        default: () => ({
            income: 0,
            savingsPct: 0,
            hoursPerDay: 0,
            daysPerWeek: 0,
            goalCost: 0,
            goalYears: 0,
            impactLevel: 0,
            goalType: '',
            goalName: ''
        })
    }
})

const { income, savingsPct, hoursPerDay, daysPerWeek, goalCost, goalYears, impactLevel, goalType, goalName } = props.formData

// Calculate basic metrics
const monthlySavings = computed(() => (income * savingsPct) / 100)
const totalMonths = computed(() => goalCost / monthlySavings.value)
const hours = computed(() => totalMonths.value * 4.33 * daysPerWeek * hoursPerDay)
const days = computed(() => totalMonths.value * 30)
const weeks = computed(() => totalMonths.value * 4.33)
const months = computed(() => Math.ceil(totalMonths.value))
const years = computed(() => Math.round(months.value / 12))

// Calculate goal score using the new algorithm
const score = computed(() => {
    const expToCostRatio = goalYears / years.value
    let goalScore = expToCostRatio * impactLevel

    // Cap the score at 100
    return Math.min(100, goalScore)
})

// Updated result messages based on new scoring system
const result = computed(() => {
    const defaultResult = {
        verdict: 'Calculating...',
        message: 'Please wait while we analyze your goal.',
        color: 'text-gray-600 dark:text-gray-400'
    }

    if (!score.value && score.value !== 0) return defaultResult

    if (score.value <= 50) {
        return {
            verdict: `Worthless 😞`,
            message: `${goalName} might not be worth your time and money right now. Consider if this aligns with your priorities.`,
            color: "text-red-600 dark:text-red-400"
        }
    } else if (score.value <= 75) {
        return {
            verdict: `Whatever 😶`,
            message: `${goalName} is a neutral choice. It's neither great nor terrible - proceed if it makes you happy.`,
            color: "text-yellow-600 dark:text-yellow-400"
        }
    } else if (score.value <= 95) {
        return {
            verdict: `Worth it 😁`,
            message: `${goalName} looks like a solid investment! The value you'll get seems to justify the cost.`,
            color: "text-green-600 dark:text-green-400"
        }
    } else {
        return {
            verdict: `Just do it! 😱`,
            message: `${goalName} is an amazing opportunity! The value-to-cost ratio is exceptional.`,
            color: "text-blue-600 dark:text-blue-400"
        }
    }
})
</script>
