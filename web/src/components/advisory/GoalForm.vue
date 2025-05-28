<template>
    <div class="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-xl border border-gray-200 dark:border-gray-700">
        <h2 class="text-2xl font-bold text-blue-600 dark:text-blue-400 mb-6 flex items-center">
            <span class="mr-2">🎯</span>
            Set Your Goal
        </h2>

        <div class="space-y-6">
            <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Goal Name</label>
                <input type="text" v-model="localData.goalName" placeholder="Buy new iPhone, Bike, or Whatever"
                    class="w-full rounded-lg bg-white dark:bg-gray-700 p-3 text-gray-900 dark:text-white border border-gray-300 dark:border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors" />
            </div>

            <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Cost (₹)</label>
                <div class="relative">
                    <span class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 dark:text-gray-400">₹</span>
                    <input type="number" v-model="localData.goalCost"
                        class="w-full rounded-lg bg-white dark:bg-gray-700 p-3 pl-8 text-gray-900 dark:text-white border border-gray-300 dark:border-gray-600 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 transition-colors" />
                </div>
            </div>

            <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Type of Goal</label>
                <div class="flex flex-wrap gap-2">
                    <button v-for="type in ['Product', 'Experience']" :key="type"
                        :class="['flex-1 min-w-[120px] px-4 py-3 rounded-lg font-medium transition-colors',
                            localData.goalType === type
                                ? 'bg-blue-700 text-white'
                                : 'bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-600']" @click="localData.goalType = type">{{
                                    type }}</button>
                </div>
            </div>

            <div class="space-y-2">
                <div class="flex justify-between items-center">
                    <label class="text-sm font-medium text-gray-700 dark:text-gray-300">Years of Value</label>
                    <span class="text-blue-600 dark:text-blue-400 font-medium">{{ localData.goalYears }}
                        years</span>
                </div>
                <input type="range" v-model="localData.goalYears" min="1" max="10"
                    class="w-full h-2 bg-gray-200 dark:bg-gray-700 rounded-lg appearance-none cursor-pointer accent-blue-500" />
            </div>

            <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">How much do you want
                    it?</label>
                <div class="flex flex-wrap gap-2">
                    <button v-for="i in 5" :key="i"
                        :class="['flex-1 min-w-[60px] p-3 rounded-lg font-medium transition-colors border',
                            localData.impactLevel === i
                                ? 'bg-blue-700 text-white'
                                : 'bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-600']"
                        @click="localData.impactLevel = i">{{
                            i }}</button>
                </div>
            </div>
        </div>

        <div class="flex flex-wrap gap-4 mt-8">
            <button @click="emit('prev')"
                class="flex-1 min-w-[120px] bg-white dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 px-6 py-3 rounded-lg font-medium text-gray-700 dark:text-white transition-colors flex items-center justify-center gap-2">
                <span class="text-lg">←</span>
                Back
            </button>
            <button @click="emitNext"
                class="flex-1 min-w-[120px] bg-blue-700 hover:bg-blue-800 px-6 py-3 rounded-lg font-medium text-white transition-colors flex items-center justify-center gap-2">
                Calculate
                <span class="text-lg">→</span>
            </button>
        </div>
    </div>
</template>

<script setup>
import { reactive } from 'vue';

const props = defineProps(['formData'])
const emit = defineEmits(['next', 'prev', 'update:form-data'])

const localData = reactive({
    ...props.formData,
    priority: props.formData.priority || 'Medium',
    targetDate: props.formData.targetDate || new Date().toISOString().split('T')[0]
})

const emitNext = () => {
    emit('update:form-data', { ...localData })
    emit('next')
}
</script>
