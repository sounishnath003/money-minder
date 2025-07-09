<template>
    <!-- Friendly Heading -->
    <div class="text-2xl font-medium mb-3 lg:text-left text-center">
        <b class="text-blue-600">Money Minder Analytics: </b>Your Year in Review
    </div>
    <div class="text-sm text-gray-600 dark:text-gray-400 mb-6 lg:text-left text-center">
        Look back at your income and expenses by year,
        month, and category. Where did your money go?
    </div>
    <div class="w-full mx-auto p-4 flex flex-col gap-10">

        <!-- Drilldown Controls -->
        <div class="flex flex-col md:flex-row md:items-center gap-4 mb-2">
            <label class="text-gray-600 dark:text-gray-300 font-medium">Year:</label>
            <select v-model="selectedYear" :class="formInputCss">
                <option v-for="year in allYears" :key="year" :value="year">{{ year }}</option>
            </select>
            <label class="text-gray-600 dark:text-gray-300 font-medium ml-2">Month:</label>
            <select v-model="selectedMonth" :class="formInputCss">
                <option value="">All</option>
                <option v-for="(month, idx) in months" :key="month" :value="idx + 1">{{ month }}</option>
            </select>
            <label class="text-gray-600 dark:text-gray-300 font-medium ml-2">Category:</label>
            <select v-model="selectedCategory" :class="formInputCss">
                <option value="">All Categories</option>
                <option v-for="cat in allCategories" :key="cat" :value="cat">{{ cat }}</option>
            </select>
        </div>

        <!-- Summary Cards -->
        <div class="text-lg md:text-2xl font-semibold text-gray-700 dark:text-gray-200">
            &bull; Yearly & Monthly Financial Summary
        </div>
        <div class="grid grid-cols-1 md:grid-cols-5 gap-4 mb-8 w-full">
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Total Income</h3>
                <p class="text-2xl font-bold text-green-700">{{ INRRuppe.format(totalIncome) }}</p>
            </div>
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Total Expense</h3>
                <p class="text-2xl font-bold text-red-600">{{ INRRuppe.format(totalSpend) }}</p>
            </div>
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Net. Savings</h3>
                <p class="text-2xl font-bold text-blue-600">{{ INRRuppe.format(netSavings) }}</p>
            </div>
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Avg. Monthly Spend</h3>
                <p class="text-2xl font-bold text-blue-600">{{ INRRuppe.format(avgMonthlySpend) }}</p>
            </div>
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
                <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300">Avg. Daily Spend</h3>
                <p class="text-2xl font-bold text-blue-600">{{ INRRuppe.format(avgMonthlySpend / 30) }}</p>
            </div>
        </div>

        <div class="text-lg md:text-2xl font-semibold text-gray-700 dark:text-gray-200">
            &bull; Look at your expense summaries
        </div>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
            <div
                class="bg-gradient-to-br from-red-100 via-white to-red-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 rounded-2xl p-3 shadow-lg flex flex-col items-center border border-red-200 dark:border-red-700 hover:scale-105 transition-transform duration-200">
                <div class="flex items-center gap-2 mb-1">
                    <span class="text-2xl text-red-500 dark:text-red-400">🔥</span>
                    <span class="text-base font-semibold text-gray-800 dark:text-gray-100">Top Expense Category</span>
                </div>
                <div class="text-2xl font-semibold text-red-600 dark:text-red-400 mb-0.5 tracking-tight">{{
                    topCategory.category || 'N/A' }}</div>
                <div class="text-sm text-gray-600 dark:text-gray-300">{{ INRRuppe.format(topCategory.totalAmount || 0)
                    }}</div>
            </div>
            <div
                class="bg-gradient-to-br from-blue-100 via-white to-blue-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 rounded-2xl p-3 shadow-lg flex flex-col items-center border border-blue-200 dark:border-blue-700 hover:scale-105 transition-transform duration-200">
                <div class="flex items-center gap-2 mb-1">
                    <span class="text-2xl text-blue-500 dark:text-blue-400">💧</span>
                    <span class="text-base font-semibold text-gray-800 dark:text-gray-100">Lowest Expense
                        Category</span>
                </div>
                <div class="text-2xl font-semibold text-blue-600 dark:text-blue-400 mb-0.5 tracking-tight">{{
                    lowestCategory.category || 'N/A' }}</div>
                <div class="text-sm text-gray-600 dark:text-gray-300">{{ INRRuppe.format(lowestCategory.totalAmount ||
                    0) }}</div>
            </div>
            <div
                class="bg-gradient-to-br from-green-100 via-white to-green-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 rounded-2xl p-3 shadow-lg flex flex-col items-center border border-green-200 dark:border-green-700 hover:scale-105 transition-transform duration-200">
                <div class="flex items-center gap-2 mb-1">
                    <span class="text-2xl text-green-500 dark:text-green-400">📊</span>
                    <span class="text-base font-semibold text-gray-800 dark:text-gray-100">Most Consistent
                        Expense</span>
                </div>
                <div class="text-2xl font-semibold text-green-600 dark:text-green-400 mb-0.5 tracking-tight">{{
                    mostConsistentCategory.category || 'N/A' }}</div>
                <div class="text-sm text-gray-600 dark:text-gray-300">Std: {{ mostConsistentCategory.std ?
                    mostConsistentCategory.std.toFixed(2) : '-' }}</div>
            </div>
            <div
                class="bg-gradient-to-br from-purple-100 via-white to-purple-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 rounded-2xl p-3 shadow-lg flex flex-col items-center border border-purple-200 dark:border-purple-700 hover:scale-105 transition-transform duration-200">
                <div class="flex items-center gap-2 mb-1">
                    <span class="text-2xl text-purple-500 dark:text-purple-400">🏆</span>
                    <span class="text-base font-semibold text-gray-800 dark:text-gray-100">Highest Single Month
                        Expense</span>
                </div>
                <div class="text-2xl font-semibold text-purple-600 dark:text-purple-400 mb-0.5 tracking-tight">{{
                    highestSingleMonth.category || 'N/A' }}</div>
                <div class="text-sm text-gray-600 dark:text-gray-300">
                    <span>{{ highestSingleMonth.monthYearStr || '-' }}</span>
                    <br>
                    <span>{{ INRRuppe.format(highestSingleMonth.totalAmount || 0) }}</span>
                </div>
            </div>
        </div>

        <!-- Excluded Categories Section (Income & Portfolio) -->
        <div v-if="excludedCategorySummaries.length" class="mt-8">
            <h2 class="text-2xl font-semibold text-blue-600 dark:text-white mb-4">Income & Portfolio Overview</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div v-for="cat in excludedCategorySummaries" :key="cat.category"
                    class="bg-white dark:bg-gray-800 rounded-xl p-8 shadow flex flex-col items-center border border-gray-200 dark:border-gray-700">
                    <div class="flex items-center gap-2 mb-2">
                        <span v-if="cat.category === 'Salary'"
                            class="text-green-600 dark:text-green-400 text-2xl">💰</span>
                        <span v-else-if="cat.category.toLowerCase().includes('mutual')"
                            class="text-purple-600 dark:text-purple-400 text-2xl">📈</span>
                        <span v-else class="text-blue-600 dark:text-blue-400 text-2xl">📊</span>
                        <span class="text-lg font-semibold text-gray-800 dark:text-gray-200">{{ cat.category }}</span>
                    </div>
                    <div class="text-3xl font-bold mb-2"
                        :class="cat.category === 'Salary' ? 'text-green-700 dark:text-green-400' : 'text-purple-700 dark:text-purple-400'">
                        {{ INRRuppe.format(cat.total) }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-300 mb-2">Total for selected period</div>
                    <div v-if="cat.timeline.length" class="w-full">
                        <div class="text-xs text-gray-500 dark:text-gray-300 mb-1">Timeline:</div>
                        <ul class="text-sm text-gray-700 dark:text-gray-200">
                            <li v-for="item in cat.timeline" :key="item.monthYearStr">
                                <span class="font-semibold">{{ item.monthYearStr }}:</span> {{
                                    INRRuppe.format(item.totalAmount) }}
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>

        <!-- Expense Bar Chart(s) -->
        <div v-if="selectedCategory === ''" class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow cursor-pointer">
            <div class="mb-2 text-lg font-semibold text-blue-600 dark:text-white">{{ `Expense by Category
                (${INCOME_CATEGORIES.join(', ')})` }}
            </div>
            <BarChart :categories="barCategories" :series="barSeriesAll" />
        </div>
        <div v-else class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow cursor-pointer">
            <div class="mb-2 text-lg font-medium text-blue-600 dark:text-white">{{ selectedCategory }} - Expense Over
                Time</div>
            <BarChart :categories="barMonths" :series="barSeriesCategory" />
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import BarChart from '../charts/BarChart.vue';
import { useTransactionStore } from '../../store/transaction.store';


// formInput box css
const formInputCss = `w-full px-4 py-3 rounded-lg border border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white  focus:ring-2 focus:ring-blue-500 focus:border-blue-500 placeholder-gray-400 dark:placeholder-gray-500`;

// Configurable list of income categories to exclude from expenses
const INCOME_CATEGORIES = ['Salary', 'Mutual Funds'];

const transactionStore = useTransactionStore();
const INRRuppe = new Intl.NumberFormat('en-IN', { style: 'currency', currency: 'INR' });

const months = [
    'Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
    'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'
];

const selectedYear = ref('');
const selectedMonth = ref('');
const selectedCategory = ref('');
const allCategories = ref([]);
const allYears = ref([]);

onMounted(async () => {
    await transactionStore.getSpendOnCategoriesByAllYearMonthAggregated();
    // Extract unique categories and years, excluding income categories
    const cats = Array.from(new Set(transactionStore.spendOnCategoriesByAllYearMonthAggregated.map(d => d.category)));
    allCategories.value = cats.filter(cat => !INCOME_CATEGORIES.includes(cat));
    const years = Array.from(new Set(transactionStore.spendOnCategoriesByAllYearMonthAggregated.map(d => d.year)));
    allYears.value = years.sort((a, b) => b - a);
    selectedYear.value = allYears.value[0] || '';
});

// Filtered data for expenses (excluding income categories)
const filteredData = computed(() => {
    let data = transactionStore.spendOnCategoriesByAllYearMonthAggregated;
    if (selectedYear.value) {
        data = data.filter(d => d.year === Number(selectedYear.value));
    }
    if (selectedMonth.value) {
        data = data.filter(d => d.month === Number(selectedMonth.value));
    }
    if (selectedCategory.value) {
        data = data.filter(d => d.category === selectedCategory.value);
    }
    // Exclude income categories from expenses
    return data.filter(d => !INCOME_CATEGORIES.includes(d.category));
});

// Income data (all income categories)
const incomeData = computed(() => {
    let data = transactionStore.spendOnCategoriesByAllYearMonthAggregated.filter(d => INCOME_CATEGORIES.includes(d.category));
    if (selectedYear.value) {
        data = data.filter(d => d.year === Number(selectedYear.value));
    }
    if (selectedMonth.value) {
        data = data.filter(d => d.month === Number(selectedMonth.value));
    }
    return data;
});

const incomeTimeline = computed(() => incomeData.value);

const totalIncome = computed(() => incomeData.value.reduce((sum, d) => sum + d.totalAmount, 0));
const totalSpend = computed(() => filteredData.value.reduce((sum, d) => sum + d.totalAmount, 0));
const netSavings = computed(() => totalIncome.value - totalSpend.value);

const filteredCategories = computed(() => {
    // Only categories present in filtered data
    return Array.from(new Set(filteredData.value.map(d => d.category)));
});

const avgMonthlySpend = computed(() => {
    // Group by monthYearStr
    const months = Array.from(new Set(filteredData.value.map(d => d.monthYearStr)));
    if (!months.length) return 0;
    return totalSpend.value / months.length;
});

const topCategory = computed(() => {
    if (!filteredData.value.length) return { category: '', totalAmount: 0 };
    // Group by category
    const grouped = {};
    filteredData.value.forEach(d => {
        if (!grouped[d.category]) grouped[d.category] = 0;
        grouped[d.category] += d.totalAmount;
    });
    const arr = Object.entries(grouped).map(([category, totalAmount]) => ({ category, totalAmount }));
    return arr.reduce((max, curr) => curr.totalAmount > max.totalAmount ? curr : max, arr[0]);
});

const lowestCategory = computed(() => {
    if (!filteredData.value.length) return { category: '', totalAmount: 0 };
    // Group by category
    const grouped = {};
    filteredData.value.forEach(d => {
        if (!grouped[d.category]) grouped[d.category] = 0;
        grouped[d.category] += d.totalAmount;
    });
    const arr = Object.entries(grouped).map(([category, totalAmount]) => ({ category, totalAmount }));
    return arr.reduce((min, curr) => curr.totalAmount < min.totalAmount ? curr : min, arr[0]);
});

const mostConsistentCategory = computed(() => {
    if (!filteredData.value.length) return { category: '', std: null };

    // Group data by category and by month
    const byCatMonth = {};
    filteredData.value.forEach(d => {
        if (!byCatMonth[d.category]) byCatMonth[d.category] = {};
        byCatMonth[d.category][d.monthYearStr] = d.totalAmount;
    });

    // Only consider categories that have data for all months in the filtered set
    const allMonths = Array.from(new Set(filteredData.value.map(d => d.monthYearStr)));
    let minStd = Infinity;
    let minCat = '';

    for (const cat in byCatMonth) {
        const catMonths = byCatMonth[cat];
        // Fill missing months with 0
        const vals = allMonths.map(month => catMonths[month] ?? 0);
        // If all values are zero, skip
        if (vals.every(v => v === 0)) continue;
        const mean = vals.reduce((a, b) => a + b, 0) / vals.length;
        const std = Math.sqrt(vals.reduce((a, b) => a + Math.pow(b - mean, 2), 0) / vals.length);
        if (std < minStd) {
            minStd = std;
            minCat = cat;
        }
    }

    // If no valid category found, return empty
    if (!minCat) return { category: '', std: null };
    return { category: minCat, std: minStd };
});

const highestSingleMonth = computed(() => {
    if (!filteredData.value.length) return { category: '', totalAmount: 0, monthYearStr: '' };
    return filteredData.value.reduce((max, curr) => curr.totalAmount > max.totalAmount ? curr : max, filteredData.value[0]);
});

const activeMonths = computed(() => {
    return Array.from(new Set(filteredData.value.map(d => d.monthYearStr))).length;
});

// Bar chart for all categories (grouped by month)
const barCategories = computed(() => {
    // X-axis: months (e.g. '2024-01', ...)
    const months = Array.from(new Set(filteredData.value.map(d => d.monthYearStr)));
    return months.sort();
});
const barSeriesAll = computed(() => {
    // One series per category
    const cats = filteredCategories.value;
    const months = barCategories.value;
    return cats.map(cat => {
        return {
            name: cat,
            data: months.map(monthStr => {
                const found = filteredData.value.find(d => d.category === cat && d.monthYearStr === monthStr);
                return found ? found.totalAmount : 0;
            })
        };
    });
});

// Bar chart for selected category (spend over months)
const barMonths = computed(() => {
    // X-axis: months for selected category
    if (!selectedCategory.value) return [];
    const months = filteredData.value.filter(d => d.category === selectedCategory.value).map(d => d.monthYearStr);
    return months.sort();
});
const barSeriesCategory = computed(() => {
    if (!selectedCategory.value) return [];
    const months = barMonths.value;
    return [{
        name: selectedCategory.value,
        data: months.map(monthStr => {
            const found = filteredData.value.find(d => d.category === selectedCategory.value && d.monthYearStr === monthStr);
            return found ? found.totalAmount : 0;
        })
    }];
});

// New computed property for excluded category summaries
const excludedCategorySummaries = computed(() => {
    const summaries = [];
    INCOME_CATEGORIES.forEach(category => {
        const data = incomeData.value.filter(d => d.category === category);
        const total = data.reduce((sum, d) => sum + d.totalAmount, 0);
        const timeline = data.map(d => ({ monthYearStr: d.monthYearStr, totalAmount: d.totalAmount }));
        summaries.push({ category, total, timeline });
    });
    return summaries;
});
</script>