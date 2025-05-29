<template>
    <div class="p-6 max-w-xl mx-auto text-white">
        <div class="mb-6 text-center dark:text-white">
            <h1 class="text-3xl font-bold text-blue-700 dark:text-white">Discover Finance Advisory</h1>
            <p class="text-sm text-gray-900 dark:text-white">Evaluate if your next purchase truly deserves your valuable
                time and money</p>
        </div>

        <div class="space-y-4">
            <component :is="steps[currentStep]" @next="currentStep++" @prev="currentStep--" :form-data="formData"
                @update:form-data="formData = $event" />
        </div>
        <div class="text-gray-700 dark:text-white text-sm my-3 text-center">
            <p>Inspired by: <spa><a class="dark:text-white text-blue-700 underline" href="https://worthit.arnav.tech"
                        target="_blank">worthit.arnav.tech</a></spa>
            </p>
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
