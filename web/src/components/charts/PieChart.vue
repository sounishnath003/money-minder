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
            <apexchart type="pie" width="380" :options="chartOptions" :series="series"></apexchart>
        </div>
    </div>
</template>

<script setup>

const series = [44, 55, 13, 43, 22];
const chartOptions = {
    chart: {
        width: 480,
        type: 'pie',
    },
    labels: ['Team A', 'Team B', 'Team C', 'Team D', 'Team E'],
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
    }]
};
</script>