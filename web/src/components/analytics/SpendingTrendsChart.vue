<template>
    <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
        <div class="text-lg font-medium mb-4">Spending Trends</div>

        <LineChart name="Total spend" :data="dailyTotalSpends" height="250" width="500" xtext="Timeline"
            ytext="Stock" />

    </div>
</template>

<script setup>
import LineChart from '../charts/LineChart.vue';

defineProps({
    dailyTotalSpends: {
        type: Array,
        required: true,
        default: () => []
    }
});
</script>
