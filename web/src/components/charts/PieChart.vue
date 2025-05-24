<template>
    <div class="flex flex-col gap-3 items-center justify-center">
        <div class="text-center">
            <div class="text-xl text-blue-700 dark:text-gray-200 font-semibold">&bull; Category-wise Spending
                Distribution (Current month)
            </div>
            <div class="text-sm">Visualize how your expenses are distributed across different categories to identify
                major spending areas
            </div>
        </div>
        <div id="chart">
            <div v-if="!series?.length || !allLabels?.length" class="flex justify-center items-center h-[350px]">
                <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
            </div>
            <apexchart v-else type="pie" width="480" :options="chartOptions" :series="series"></apexchart>
        </div>
    </div>
</template>

<script setup>
import { computed, onMounted } from 'vue';
import { useTransactionStore } from '../../store/transaction.store';

const transactionStore = useTransactionStore();

// Fetch data on component mount
onMounted(async () => {
    await transactionStore.getAllSpendsByCategories();
});

const series = computed(() => transactionStore.allSpendsByCategories.map(val => val.totalAmount));
const allLabels = computed(() => transactionStore.allSpendsByCategories.map((val) => val.category));

const chartOptions = computed(() => ({
    chart: {
        width: 480,
        type: 'pie',
        toolbar: {
            show: true,
            tools: {
                download: true,
                selection: true,
                zoom: true,
                zoomin: true,
                zoomout: true,
                pan: true,
                reset: true
            }
        }
    },
    labels: allLabels.value,
    responsive: [{
        breakpoint: 480,
        options: {
            chart: {
                width: 200
            },
            legend: {
                position: 'bottom'
            }
        }
    }],
    tooltip: {
        y: {
            formatter: function (val) {
                return INRRuppe.format(val);
            }
        }
    }
}));

// Utilities
const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});
</script>