<template>
    <div class="text-center">
        <div class="text-xl text-blue-700 dark:text-gray-200 font-semibold">&bull; Spend By Categories (MoM)
        </div>
        <div class="text-sm">Your spend by categories month-on-month. Keep an eye on your expenditures.</div>
    </div>
    <div id="chart" class="w-full">
        <apexchart type="bar" height="350" :options="chartOptions" :series="props.series"></apexchart>
    </div>
</template>

<script setup>

const props = defineProps({
    categories: {
        type: Array,
        required: true,
        default: []
    },
    series: {
        type: Array,
        required: true,
        default: []
    }

});

const chartOptions = {
    chart: {
        type: 'bar',
        height: 350
    },
    plotOptions: {
        bar: {
            horizontal: false,
            columnWidth: '55%',
            borderRadius: 5,
            borderRadiusApplication: 'end'
        },
    },
    dataLabels: {
        enabled: false
    },
    stroke: {
        show: true,
        width: 2,
        colors: ['transparent']
    },
    xaxis: {
        categories: props.categories,
    },
    yaxis: {
        title: {
            text: 'INR (Indian Rupee)'
        }
    },
    fill: {
        opacity: 1
    },
    tooltip: {
        y: {
            formatter: function (val) {
                return INRRuppe.format(val);
            }
        }
    }
};


// Utilities
const USDollar = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
});

const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});




</script>