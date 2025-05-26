<template>
    <div class="text-2xl font-medium mb-3 lg:text-left text-center">
        <b class="text-blue-600">Money Minder Analytics: </b>Track & Optimize Your Finances
    </div>
    <div class="text-sm text-gray-600 dark:text-gray-400 mb-6 lg:text-left text-center">
        Visualize your spending patterns and make smarter financial decisions
    </div>
    <div v-if="isLoading" class="flex justify-center items-center h-[350px] font-medium">
        Loading...
    </div>
    <div v-else class="flex flex-col gap-6 mx-auto items-center justify-center w-full">
        <!-- Summary Cards -->
        <AnalyticsSummaryCards :dailyTotalSpends="dailyTotalSpends" :highestSpendingCategory="highestSpendingCategory"
            :savingsRate="savingsRate" :averageDailySpend="averageDailySpend" />

        <!-- Main Charts -->
        <div class="flex flex-col gap-4 lg:flex-row justify-between items-center w-full">
            <div class="w-full lg:w-1/2">
                <BarChart :categories="spendOnCategoriesMonthOnMonth.months"
                    :series="spendOnCategoriesMonthOnMonth.series" />
            </div>
            <div class="w-full lg:w-1/2">
                <PieChart />
            </div>
        </div>

        <!-- Additional Analytics -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 w-full">
            <SpendingTrendsChart :dailyTotalSpends="dailyTotalSpends" />
            <CategoryGrowthList :categoryGrowth="categoryGrowth" />
        </div>
    </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue';
import BarChart from './charts/BarChart.vue';
import LineChart from './charts/LineChart.vue';
import { useTransactionStore } from '../store/transaction.store';
import PieChart from './charts/PieChart.vue';
import AnalyticsSummaryCards from './analytics/AnalyticsSummaryCards.vue';
import CategoryGrowthList from './analytics/CategoryGrowthList.vue';
import SpendingTrendsChart from './analytics/SpendingTrendsChart.vue';

// Define store
const transactionStore = useTransactionStore();

// Loading states
const isLoading = ref(true);

// Currency formatter
const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});

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

// Computed properties
const dailyTotalSpends = computed(() => {
    if (!transactionStore.totalDailySpends?.length) return [];
    return transactionStore.totalDailySpends.map(curr => [curr.unixMiliseconds, curr.amount]);
});

const averageDailySpend = computed(() => {
    if (!transactionStore.totalDailySpends?.length) return 0;
    const total = transactionStore.totalDailySpends.reduce((sum, curr) => sum + curr.amount, 0);
    return total / transactionStore.totalDailySpends.length;
});

const highestSpendingCategory = computed(() => {
    if (!transactionStore.spendOnCategoriesMonthOnMonth?.length) return { category: 'N/A', amount: 0 };
    const categories = transactionStore.spendOnCategoriesMonthOnMonth.filter(tx => (tx.category !== 'Salary' && tx.category !== 'House Rents')).reduce((acc, curr) => {
        if (!acc[curr.category]) acc[curr.category] = 0;
        acc[curr.category] += curr.totalSpendAmount;
        return acc;
    }, {});

    const maxCategory = Object.entries(categories).reduce((max, [category, amount]) =>
        amount > max.amount ? { category, amount } : max,
        { category: '', amount: 0 }
    );

    return maxCategory;
});

const savingsRate = computed(() => {
    if (!transactionStore.allTransactions?.length) return 0;

    // Get current month's transactions
    const currentMonth = new Date().getMonth();
    const currentYear = new Date().getFullYear();

    const monthlyTransactions = transactionStore.allTransactions.filter(tx => {
        const txDate = tx.createdAt;
        return txDate.getMonth() === currentMonth && txDate.getFullYear() === currentYear;
    });

    const totalIncome = monthlyTransactions
        .filter(tx => tx.transactionType === 'Income')
        .reduce((sum, curr) => sum + curr.amount, 0);


    const totalExpenses = monthlyTransactions
        .filter(tx => tx.transactionType === 'Expense')
        .reduce((sum, curr) => sum + curr.amount, 0);

    return totalIncome ? ((totalIncome - totalExpenses) / totalIncome * 100).toFixed(1) : 0;
});

const categoryGrowth = computed(() => {
    if (!transactionStore.spendOnCategoriesMonthOnMonth?.length) return [];

    const categories = transactionStore.spendOnCategoriesMonthOnMonth.reduce((acc, curr) => {
        if (!acc[curr.category]) {
            acc[curr.category] = {
                firstMonth: curr.totalSpendAmount,
                lastMonth: curr.totalSpendAmount
            };
        } else {
            acc[curr.category].lastMonth = curr.totalSpendAmount;
        }
        return acc;
    }, {});

    return Object.entries(categories).map(([name, data]) => ({
        name,
        growth: data.firstMonth ?
            ((data.lastMonth - data.firstMonth) / data.firstMonth * 100).toFixed(1) : 0
    }));
});

const spendOnCategoriesMonthOnMonth = computed(() => {
    if (!transactionStore.spendOnCategoriesMonthOnMonth?.length) {
        return { months: [], series: [] };
    }

    const months = [...new Set(transactionStore.spendOnCategoriesMonthOnMonth.map(item => item.month))];
    const groupedByCategory = transactionStore.spendOnCategoriesMonthOnMonth.filter(tx => (tx.category !== 'Salary' && tx.category !== 'House Rents')).reduce((acc, curr) => {
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