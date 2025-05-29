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

// Calculate worthiness score using multiple factors
const score = computed(() => {
    // Base score starts at 50
    let score = 50

    // Factor 1: Time to save vs value years (up to ±25 points)
    const timeValueRatio = years.value / goalYears
    score += (1 - timeValueRatio) * 25

    // Factor 2: Impact level (up to ±15 points)
    score += (impactLevel - 3) * 5

    // Factor 3: Cost relative to income (up to ±10 points)
    const costIncomeRatio = goalCost / (income * 12)
    score -= costIncomeRatio * 10

    // Factor 4: Goal type bonus (up to ±10 points)
    score += goalType === 'Experience' ? 10 : 0

    // Factor 5: Savings rate adjustment (up to ±10 points)
    score += (savingsPct - 20) / 8

    // Ensure score stays within 0-100 range
    return Math.max(0, Math.min(100, score))
})

// Enhanced result messages with more context
const result = computed(() => {
    const defaultResult = {
        verdict: 'Calculating...',
        message: 'Please wait while we analyze your goal.',
        color: 'text-gray-600 dark:text-gray-400'
    }

    if (!score.value && score.value !== 0) return defaultResult

    if (score.value < 30) {
        return {
            verdict: `Your wallet is screaming in agony, ${goalName}!`,
            message: "I've seen better financial decisions in a horror movie. Your savings account will haunt you like a ghost if you proceed. Maybe consult a financial exorcist first?",
            color: "text-red-600 dark:text-red-400"
        }
    } else if (score.value < 70) {
        return {
            verdict: `${goalName}, are you sure you want to dance with the devil?`,
            message: "Your bank account is giving me serious 'The Shining' vibes. It's not all bad, but proceed with caution - your future self might be watching you through the financial mirror.",
            color: "text-yellow-600 dark:text-yellow-400"
        }
    } else {
        return {
            verdict: `${goalName}, the financial stars have aligned!`,
            message: "Your wallet is purring like a satisfied cat. This goal is so good, it's almost suspicious. Are you sure you're not a financial wizard in disguise?",
            color: "text-green-600 dark:text-green-400"
        }
    }
})
</script>
