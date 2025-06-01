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
            :savingsRate="savingsRate" :averageDailySpend="averageDailySpend" :investments="investments" />

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

        <!-- paymentMethod KPI -->
        <PaymentMethodsKPI :paymentMethodKPIs="paymentMethodKPIs" />
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
import PaymentMethodsKPI from './analytics/PaymentMethodsKPI.vue';

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

const investments = computed(() => {
    if (!transactionStore.allTransactions?.length) return 0;
    return transactionStore.allTransactions
        .filter(tx => tx.transactionType === 'Investments')
        .reduce((sum, curr) => sum + curr.amount, 0);
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

    const totalInvestments = monthlyTransactions
        .filter(tx => tx.transactionType === 'Investment')
        .reduce((sum, curr) => sum + curr.amount, 0);

    return totalIncome ? ((totalIncome - totalExpenses - totalInvestments) / totalIncome * 100).toFixed(1) : 0;
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

const paymentMethodKPIs = computed(() => {
    if (!transactionStore.allTransactions?.length) return [];

    // Separate income, expense and investment transactions
    const incomeTransactions = transactionStore.allTransactions.filter(tx => tx.transactionType === 'Income');
    const expenseTransactions = transactionStore.allTransactions.filter(tx => tx.transactionType === 'Expense');
    const investmentTransactions = transactionStore.allTransactions.filter(tx => tx.transactionType === 'Investment');

    const paymentMethodTotals = transactionStore.allTransactions.reduce((acc, curr) => {
        if (!acc[curr.paymentMethod]) {
            acc[curr.paymentMethod] = {
                total: 0,
                income: 0,
                expense: 0,
                investment: 0,
                count: 0,
                incomeCount: 0,
                expenseCount: 0,
                investmentCount: 0,
                average: 0,
                incomeAverage: 0,
                expenseAverage: 0,
                investmentAverage: 0,
                utilizationRate: 0
            };
        }

        if (curr.transactionType === 'Income') {
            acc[curr.paymentMethod].income += curr.amount;
            acc[curr.paymentMethod].incomeCount += 1;
        } else if (curr.transactionType === 'Expense') {
            acc[curr.paymentMethod].expense += curr.amount;
            acc[curr.paymentMethod].expenseCount += 1;
        } else if (curr.transactionType === 'Investment') {
            acc[curr.paymentMethod].investment += curr.amount;
            acc[curr.paymentMethod].investmentCount += 1;
        }

        acc[curr.paymentMethod].total = acc[curr.paymentMethod].income - acc[curr.paymentMethod].expense - acc[curr.paymentMethod].investment;
        acc[curr.paymentMethod].count = acc[curr.paymentMethod].incomeCount + acc[curr.paymentMethod].expenseCount + acc[curr.paymentMethod].investmentCount;

        // Calculate averages
        acc[curr.paymentMethod].incomeAverage = acc[curr.paymentMethod].incomeCount ?
            acc[curr.paymentMethod].income / acc[curr.paymentMethod].incomeCount : 0;
        acc[curr.paymentMethod].expenseAverage = acc[curr.paymentMethod].expenseCount ?
            acc[curr.paymentMethod].expense / acc[curr.paymentMethod].expenseCount : 0;
        acc[curr.paymentMethod].average = acc[curr.paymentMethod].count ?
            acc[curr.paymentMethod].total / acc[curr.paymentMethod].count : 0;

        // Calculate utilization rate (expense/income ratio)
        acc[curr.paymentMethod].utilizationRate = acc[curr.paymentMethod].income ?
            (acc[curr.paymentMethod].expense / acc[curr.paymentMethod].income * 100).toFixed(1) : 0;

        return acc;
    }, {});

    const totalNetAmount = Object.values(paymentMethodTotals).reduce((sum, data) => sum + data.total, 0);

    return Object.entries(paymentMethodTotals).map(([method, data]) => ({
        method,
        total: data.total,
        income: data.income,
        expense: data.expense,
        count: data.count,
        incomeCount: data.incomeCount,
        expenseCount: data.expenseCount,
        average: data.average,
        incomeAverage: data.incomeAverage,
        expenseAverage: data.expenseAverage,
        utilizationRate: data.utilizationRate,
        percentage: totalNetAmount ? (data.total / totalNetAmount * 100).toFixed(1) : 0,
        // Additional KPIs
        transactionFrequency: (data.count / transactionStore.allTransactions.length * 100).toFixed(1),
        expenseRatio: data.expense ? (data.expense / transactionStore.allTransactions
            .filter(tx => tx.transactionType === 'Expense')
            .reduce((sum, tx) => sum + tx.amount, 0) * 100).toFixed(1) : 0,
        incomeRatio: data.income ? (data.income / transactionStore.allTransactions
            .filter(tx => tx.transactionType === 'Income')
            .reduce((sum, tx) => sum + tx.amount, 0) * 100).toFixed(1) : 0
    }));
});


</script>