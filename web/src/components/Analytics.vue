<template>
    <div class="text-2xl font-medium mb-3 lg:text-left text-center">
        <b class="text-blue-600">Analytics &amp; Insights: </b>Understand your spends
    </div>
    <div v-if="isLoading" class="flex justify-center items-center h-[350px] font-medium">
        Loading...
    </div>
    <div v-else class="flex flex-col gap-6 mx-auto items-center justify-center w-full">
        <LineChart name="Total spend" :data="dailyTotalSpends" height="350" width="500" xtext="Timeline"
            ytext="Stock" />
        <BarChart :categories="spendOnCategoriesMonthOnMonth.months" :series="spendOnCategoriesMonthOnMonth.series" />
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue';
import BarChart from './charts/BarChart.vue';
import LineChart from './charts/LineChart.vue';
import { useTransactionStore } from '../store/transaction.store';

// Define store
const transactionStore = useTransactionStore();

// Loading states
const isLoading = ref(true);

// On Mount
onMounted(async () => {
    try {
        await Promise.all([
            transactionStore.getDailyTotalSpendByTimeframe(),
            transactionStore.getSpendOnCategoriesMonthOnMonth()
        ]);
    } catch (error) {
        console.error('Failed to load analytics data:', error);
    } finally {
        isLoading.value = false;
    }
});

// ref
const dailyTotalSpends = computed(() => {
    if (!transactionStore.totalDailySpends?.length) return [];
    return transactionStore.totalDailySpends.map(curr => [curr.unixMiliseconds, curr.amount]);
});

/**
```
{
    name: 'Net Profit',
    data: [44, 55, 57, 56, 61, 58, 63, 60, 66]
}, {
    name: 'Revenue',
    data: [76, 85, 101, 98, 87, 105, 91, 114, 94]
}, {
    name: 'Free Cash Flow',
    data: [35, 41, 36, 26, 45, 48, 52, 53, 41]
}
```
 */
const spendOnCategoriesMonthOnMonth = computed(() => {
    if (!transactionStore.spendOnCategoriesMonthOnMonth?.length) {
        return { months: [], series: [] };
    }

    // Get unique months for X-axis
    const months = [...new Set(transactionStore.spendOnCategoriesMonthOnMonth.map(item => item.month))];

    // Group data by category
    const groupedByCategory = transactionStore.spendOnCategoriesMonthOnMonth.reduce((acc, curr) => {
        if (!acc[curr.category]) {
            acc[curr.category] = {
                name: curr.category,
                data: new Array(months.length).fill(0)
            };
        }
        const monthIndex = months.indexOf(curr.month);
        acc[curr.category].data[monthIndex] = curr.totalSpendAmount;
        return acc;
    }, {});

    return {
        months,
        series: Object.values(groupedByCategory)
    };
});

</script>