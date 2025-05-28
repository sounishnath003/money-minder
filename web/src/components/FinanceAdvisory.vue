<template>
    <div class="p-6 max-w-xl mx-auto text-white">
        <div class="mb-6 text-center dark:text-white">
            <h1 class="text-3xl font-bold text-blue-700 dark:text-white">Finance Advisor</h1>
            <p class="text-sm text-gray-900">Make smarter financial decisions with personalized insights</p>
        </div>

        <div class="space-y-4">
            <component :is="steps[currentStep]" @next="currentStep++" @prev="currentStep--" :form-data="formData"
                @update:form-data="formData = $event" />
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import IncomeForm from './advisory/IncomeForm.vue'
import GoalForm from './advisory/GoalForm.vue'
import ResultView from './advisory/ResultForm.vue'

const currentStep = ref(0)
const steps = [IncomeForm, GoalForm, ResultView]

const formData = ref({
    income: 10000,
    savingsPct: 20,
    hoursPerDay: 8,
    daysPerWeek: 5,

    goalName: '',
    goalCost: 0,
    goalType: 'Product',
    goalYears: 3,
    impactLevel: 3,
})
</script>
