<template>
    <div class="w-full bg-white dark:bg-gray-800 p-6 rounded-lg shadow">
        <h3 class="text-lg font-semibold text-gray-700 dark:text-gray-300 mb-4">&bull; Payment Methods Analysis</h3>
        <p class="text-sm text-gray-600 dark:text-gray-400 mb-4">Track your spending patterns across different payment
            methods. Understand which payment methods you use most frequently and their impact on your overall finances.
        </p>

        <!-- Summary Cards -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            <div v-for="kpi in paymentMethodKPIs" :key="kpi.method" class="bg-gray-50 dark:bg-gray-700 p-4 rounded-lg">
                <div class="flex justify-between items-center mb-3">
                    <h4 class="text-sm font-medium text-gray-600 dark:text-gray-300">{{ kpi.method }}</h4>
                    <span class="text-md px-2 py-1 rounded-full"
                        :class="kpi.total >= 0 ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'">
                        {{ kpi.percentage }}%
                    </span>
                </div>
                <div class="space-y-2">
                    <div class="flex justify-between items-center">
                        <span class="text-xs text-gray-500">Net Amount:</span>
                        <span class="text-sm font-medium" :class="kpi.total >= 0 ? 'text-green-600' : 'text-red-600'">
                            {{ INRRuppe.format(kpi.total) }}
                        </span>
                    </div>
                    <div class="flex justify-between items-center">
                        <span class="text-xs text-gray-500">Income:</span>
                        <span class="text-sm font-medium text-green-600">{{ INRRuppe.format(kpi.income) }}</span>
                    </div>
                    <div class="flex justify-between items-center">
                        <span class="text-xs text-gray-500">Expense:</span>
                        <span class="text-sm font-medium text-red-600">{{ INRRuppe.format(kpi.expense) }}</span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Detailed Metrics -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Transaction Metrics -->
            <div class="bg-gray-50 dark:bg-gray-700 p-4 rounded-lg">
                <h4 class="text-sm font-medium text-gray-600 dark:text-gray-300 mb-3">Transaction Metrics</h4>
                <div class="space-y-3">
                    <div v-for="kpi in paymentMethodKPIs" :key="kpi.method" class="space-y-2">
                        <div class="flex justify-between items-center">
                            <span class="text-sm text-gray-600">{{ kpi.method }}</span>
                            <span class="text-xs px-2 py-1 rounded-full bg-blue-100 text-blue-800">
                                {{ kpi.transactionFrequency }}% of total
                            </span>
                        </div>
                        <div class="grid grid-cols-2 gap-2 text-xs">
                            <div class="flex justify-between">
                                <span class="text-gray-500">Count:</span>
                                <span class="font-medium">{{ kpi.count }}</span>
                            </div>
                            <div class="flex justify-between">
                                <span class="text-gray-500">Avg:</span>
                                <span class="font-medium">{{ INRRuppe.format(kpi.average) }}</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Utilization Metrics -->
            <div class="bg-gray-50 dark:bg-gray-700 p-4 rounded-lg">
                <h4 class="text-sm font-medium text-gray-600 dark:text-gray-300 mb-3">Utilization Metrics</h4>
                <div class="space-y-3">
                    <div v-for="kpi in paymentMethodKPIs" :key="kpi.method" class="space-y-2">
                        <div class="flex justify-between items-center">
                            <span class="text-sm text-gray-600">{{ kpi.method }}</span>
                            <span class="text-xs px-2 py-1 rounded-full"
                                :class="kpi.utilizationRate > 50 ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800'">
                                {{ kpi.utilizationRate }}% utilized
                            </span>
                        </div>
                        <div class="grid grid-cols-2 gap-2 text-xs">
                            <div class="flex justify-between">
                                <span class="text-gray-500">Income Ratio:</span>
                                <span class="font-medium">{{ kpi.incomeRatio }}%</span>
                            </div>
                            <div class="flex justify-between">
                                <span class="text-gray-500">Expense Ratio:</span>
                                <span class="font-medium">{{ kpi.expenseRatio }}%</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
const props = defineProps({
    paymentMethodKPIs: {
        type: Array,
        required: true,
        default: []
    }
});

// Utilities
const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});
</script>