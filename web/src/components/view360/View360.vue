<template>
    <div v-if="isLoading" class="text-center my-auto font-medium py-4">
        Loading...
    </div>

    <template v-else>
        <!-- Friendly Heading -->
        <div class="text-2xl font-medium mb-3 lg:text-left text-center">
            <b class="text-blue-600">Money Minder Analytics: </b>Your Year in Review
        </div>
        <div class="text-sm text-gray-600 dark:text-gray-400 mb-6 lg:text-left text-center">
            Look back at your income and expenses by year,
            month, category, and payment method. Where did your money go?
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
                <label class="text-gray-600 dark:text-gray-300 font-medium ml-2">Payment Method:</label>
                <select v-model="selectedPaymentMethod" :class="formInputCss">
                    <option value="">All Payment Methods</option>
                    <option v-for="pay in allPaymentMethods" :key="pay" :value="pay">{{ pay }}</option>
                </select>
            </div>

            <!-- Summary Cards -->
            <div class="text-lg md:text-2xl font-semibold text-blue-600 dark:text-gray-200">
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

            <div class="text-lg md:text-2xl font-semibold text-blue-600 dark:text-gray-200">
                &bull; Look at your expense summaries
            </div>
            <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
                <div
                    class="bg-gradient-to-br from-red-100 via-white to-red-50 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 rounded-2xl p-3 shadow-lg flex flex-col items-center border border-red-200 dark:border-red-700 hover:scale-105 transition-transform duration-200">
                    <div class="flex items-center gap-2 mb-1">
                        <span class="text-2xl text-red-500 dark:text-red-400">🔥</span>
                        <span class="text-base font-semibold text-gray-800 dark:text-gray-100">Top Expense
                            Category</span>
                    </div>
                    <div class="text-2xl font-semibold text-red-600 dark:text-red-400 mb-0.5 tracking-tight">{{
                        topCategory.category || 'N/A' }}</div>
                    <div class="text-sm text-gray-600 dark:text-gray-300">{{ INRRuppe.format(topCategory.totalAmount ||
                        0)
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
                    <div class="text-sm text-gray-600 dark:text-gray-300">{{ INRRuppe.format(lowestCategory.totalAmount
                        ||
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

            <div class="text-lg md:text-2xl font-semibold text-blue-600 dark:text-gray-200">
                &bull; {{ `Expense by Category
                (${INCOME_CATEGORIES.join(', ')})` }}
            </div>
            <!-- Expense Bar Chart(s) -->
            <div v-if="selectedCategory === ''" class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow cursor-pointer">
                <BarChart :categories="barCategories" :series="barSeriesAll" />
            </div>
            <div v-else class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow cursor-pointer">
                <div class="mb-2 text-lg font-medium text-blue-600 dark:text-white">{{ selectedCategory }} - Expense
                    Over
                    Time</div>
                <BarChart :categories="barMonths" :series="barSeriesCategory" />
            </div>

            <!-- Comparison Section: Spend Growth/Decrease -->
            <div class="text-lg md:text-2xl font-semibold text-blue-600 dark:text-gray-200">
                &bull; Compare your spend growth/decrease by category
            </div>
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow mb-8">
                <div class="flex flex-col md:flex-row gap-4 mb-4 items-center">
                    <div>
                        <label class="font-medium text-gray-700 dark:text-gray-300 mr-2">Year:</label>
                        <select v-model="compareYear" :class="formInputCss">
                            <option v-for="y in allYearsForComparison" :key="y" :value="y">{{ y }}</option>
                        </select>
                    </div>
                    <div>
                        <label class="font-medium text-gray-700 dark:text-gray-300 mr-2">Older Month:</label>
                        <select v-model="compareMonth1" :class="formInputCss">
                            <option v-for="m in monthsForYear" :key="m" :value="m">{{ getMonthName(m) }}</option>
                        </select>
                    </div>
                    <div>
                        <label class="font-medium text-gray-700 dark:text-gray-300 mr-2">Newer Month:</label>
                        <select v-model="compareMonth2" :class="formInputCss">
                            <option v-for="m in monthsForYear" :key="m" :value="m">{{ getMonthName(m) }}</option>
                        </select>
                    </div>
                </div>
                <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                        <thead>
                            <tr>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Category</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    {{ getMonthName(compareMonth1) }}</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    {{ getMonthName(compareMonth2) }}</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Growth/Decrease (%)</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="row in comparisonTable" :key="row.category">
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200 font-medium">{{
                                    row.category }}</td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{
                                    INRRuppe.format(row.month1) }}</td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{
                                    INRRuppe.format(row.month2) }}</td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200 font-medium">
                                    <span v-if="row.percent === null">-</span>
                                    <span v-else
                                        :class="row.percent > 0 ? 'text-green-600' : row.percent < 0 ? 'text-red-600' : ''">
                                        {{ row.percent > 0 ? '+' : '' }}{{ row.percent.toFixed(2) }}%
                                        <span v-if="row.percent > 0">&uparrow;</span>
                                        <span v-else-if="row.percent < 0">&downarrow;</span>
                                        <span v-else>&nbsp;</span>
                                    </span>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div class="text-lg md:text-2xl font-semibold text-blue-600 dark:text-gray-200">
                &bull; All Transactions in the selected period ({{ filteredData.length }})
            </div>
            <!-- Table View of Aggregated Data -->
            <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow overflow-x-auto">
                <div class="mb-2 text-lg font-semibold text-blue-600 dark:text-white">Detailed Data Table</div>
                <div class="max-h-[350px] overflow-y-auto">
                    <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                        <thead class="sticky top-0 bg-white dark:bg-gray-800 z-10">
                            <tr>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Year</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Month</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Category</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Payment Method</th>
                                <th
                                    class="px-4 py-2 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Total Amount</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="row in filteredData"
                                :key="`${row.year}-${row.month}-${row.category}-${row.paymentMethod}`"
                                class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200">
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{ row.year }}</td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{ row.monthYearStr }}
                                </td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{ row.category }}</td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{ row.paymentMethod ||
                                    '-' }}</td>
                                <td class="px-4 py-2 text-sm text-gray-700 dark:text-gray-200">{{
                                    INRRuppe.format(row.totalAmount) }}</td>
                            </tr>
                            <tr v-if="filteredData.length === 0">
                                <td colspan="5" class="px-4 py-2 text-center text-gray-400">No data available for the
                                    selected filters.</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <!-- Excluded Categories Section (Income & Portfolio) -->
            <div v-if="excludedCategorySummaries.length" class="mt-8">
                <h2 class="text-2xl font-semibold text-blue-600 dark:text-white mb-4">&bull; Income & Portfolio Overview
                </h2>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div v-for="cat in excludedCategorySummaries" :key="cat.category"
                        class="bg-gradient-to-br from-green-50 via-white to-green-100 dark:from-gray-900 dark:via-gray-800 dark:to-gray-900 rounded-2xl p-4 shadow-lg flex flex-col items-center border border-green-200 dark:border-green-700 hover:scale-105 transition-transform duration-200">
                        <div class="flex items-center gap-2 mb-1">
                            <span v-if="cat.category === 'Salary'"
                                class="text-green-600 dark:text-green-400 text-2xl">💰</span>
                            <span v-else-if="cat.category.toLowerCase().includes('mutual')"
                                class="text-purple-600 dark:text-purple-400 text-2xl">📈</span>
                            <span v-else class="text-blue-600 dark:text-blue-400 text-2xl">📊</span>
                            <span class="text-base font-semibold text-gray-800 dark:text-gray-100">{{ cat.category
                                }}</span>
                        </div>
                        <div class="text-xl md:text-3xl font-semibold mb-1"
                            :class="cat.category === 'Salary' ? 'text-green-700 dark:text-green-400' : 'text-purple-700 dark:text-purple-400'">
                            {{ INRRuppe.format(cat.total) }}
                        </div>
                        <div class="text-xs text-gray-500 dark:text-gray-300 mb-1">Total for selected period</div>
                        <div class="flex flex-col items-center w-full">
                            <div v-if="cat.timeline.length" class="text-xs text-gray-500 dark:text-gray-300">
                                Last month:
                                <span class="font-semibold text-gray-700 dark:text-gray-200">
                                    {{
                                        (() => {
                                            const last = cat.timeline[cat.timeline.length - 1];
                                            return last ? INRRuppe.format(last.totalAmount) : '-';
                                        })()
                                    }}
                                </span>
                            </div>
                            <div v-else class="text-xs text-gray-400 dark:text-gray-500">No monthly data</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </template>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue';
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
const selectedPaymentMethod = ref('');
const allCategories = ref([]);
const allPaymentMethods = ref([]);
const allYears = ref([]);
const isLoading = ref(true);

onMounted(async () => {
    await transactionStore.getSpendOnCategoriesByAllYearMonthAggregated();
    // Extract unique categories and years, excluding income categories
    const cats = Array.from(new Set(transactionStore.spendOnCategoriesByAllYearMonthAggregated.map(d => d.category)));
    allCategories.value = cats.filter(cat => !INCOME_CATEGORIES.includes(cat));
    const pays = Array.from(new Set(transactionStore.spendOnCategoriesByAllYearMonthAggregated.map(d => d.paymentMethod)));
    allPaymentMethods.value = pays;
    const years = Array.from(new Set(transactionStore.spendOnCategoriesByAllYearMonthAggregated.map(d => d.year)));
    allYears.value = years.sort((a, b) => b - a);
    selectedYear.value = allYears.value[0] || '';
    isLoading.value = false;
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
    if (selectedPaymentMethod.value) {
        data = data.filter(d => d.paymentMethod === selectedPaymentMethod.value);
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
    if (selectedPaymentMethod.value) {
        data = data.filter(d => d.paymentMethod === selectedPaymentMethod.value);
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

// --- Comparison Section ---
const allYearsForComparison = computed(() => {
    // All unique years in filteredData
    return Array.from(new Set(filteredData.value.map(d => d.year))).sort((a, b) => b - a);
});
const compareYear = ref('');

// Auto-select current year by default for compareYear
watch(
    allYearsForComparison,
    (years) => {
        if (years.length && !compareYear.value) {
            const currentYear = new Date().getFullYear();
            compareYear.value = years.includes(currentYear) ? currentYear : years[0];
        }
    },
    { immediate: true }
);

const monthsForYear = computed(() => {
    if (!compareYear.value) return [];
    // All unique monthYearStrs for the selected year
    return Array.from(new Set(filteredData.value.filter(d => d.year === Number(compareYear.value)).map(d => d.monthYearStr))).sort();
});

const compareMonth1 = ref('');
const compareMonth2 = ref('');

// Set defaults after data loads or year changes
watch(
    [() => monthsForYear.value, () => compareYear.value],
    ([months]) => {
        if (months.length && (!compareMonth1.value || !months.includes(compareMonth1.value))) {
            compareMonth1.value = months[1]; // Older month
        }
        if (months.length && (!compareMonth2.value || !months.includes(compareMonth2.value))) {
            compareMonth2.value = months[0]; // Newer month
        }
    },
    { immediate: true }
);

const comparisonTable = computed(() => {
    // For each unique category, get spend in both months and compute % change
    const categories = Array.from(new Set(filteredData.value.filter(d => d.year === Number(compareYear.value)).map(d => d.category)));
    return categories.map(cat => {
        const m1 = filteredData.value.find(d => d.category === cat && d.monthYearStr === compareMonth1.value && d.year === Number(compareYear.value));
        const m2 = filteredData.value.find(d => d.category === cat && d.monthYearStr === compareMonth2.value && d.year === Number(compareYear.value));
        const val1 = m1 ? m1.totalAmount : 0;
        const val2 = m2 ? m2.totalAmount : 0;
        let percent = null;
        if (val1 === 0 && val2 === 0) percent = null;
        else if (val1 === 0) percent = 100;
        else percent = ((val2 - val1) / Math.abs(val1)) * 100;
        return {
            category: cat,
            month1: val1,
            month2: val2,
            percent
        };
    });
});

function getMonthName(monthYearStr) {
    // monthYearStr is like '2024-01', '2024-02', etc.
    const parts = monthYearStr.split('-');
    if (parts.length !== 2) return monthYearStr;
    const monthIdx = Number(parts[1]) - 1;
    return months[monthIdx] || monthYearStr;
}

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