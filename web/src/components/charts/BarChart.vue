<template>
    <div class="text-center">
        <div class="text-xl text-blue-700 dark:text-gray-200 font-semibold">&bull; Spend By Categories (MoM)
        </div>
        <div class="text-sm">Your spend by categories month-on-month. Keep an eye on your expenditures.</div>
    </div>
    <div id="chart" class="w-full">
        <div v-if="!props.categories?.length || !props.series?.length"
            class="flex justify-center items-center h-[350px]">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        </div>
        <apexchart v-else type="bar" height="350" :options="chartOptions" :series="props.series"></apexchart>
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
        height: 350,
        zoom: {
            enabled: true,
            type: 'x',
            autoScaleYaxis: true,
            zoomedArea: {
                fill: {
                    color: '#90CAF9',
                    opacity: 0.4
                },
                stroke: {
                    color: '#0D47A1',
                    opacity: 0.4,
                    width: 1
                }
            }
        },
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
            },
            autoSelected: 'zoom'
        }
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