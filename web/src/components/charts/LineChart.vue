<template>
    <div class="flex flex-col gap-3 items-center justify-center w-full">
        <div class="text-center">
            <div class="text-xl text-blue-700 dark:text-gray-200 font-semibold">&bull; Current month spends
            </div>
            <div class="text-sm">Track your daily expenses to identify patterns and make informed financial decisions.
            </div>
        </div>

        <div id="chart" class="w-full">
            <div v-if="!props.data?.length" class="flex justify-center items-center" :style="{ height: props.height }">
                <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
            </div>
            <apexchart v-else type="area" :height="props.height" :options="chartOptions" :series="series"></apexchart>
        </div>
    </div>
</template>

<script setup>

const props = defineProps({
    name: {
        required: true,
        type: String
    },
    data: {
        required: true,
        type: Array,
    },
    height: {
        required: true,
        type: String,
    },
    width: {
        required: true,
        type: String,
    },
    xtext: {
        requireed: true,
        type: String,
    },
    ytext: {
        requireed: true,
        type: String,
    }
})
const series = [{
    name: props.name,
    data: props.data,
}];


const chartOptions = {
    chart: {
        type: 'area',
        stacked: false,
        height: props.height,
        zoom: {
            type: 'x',
            enabled: true,
            autoScaleYaxis: true
        },
        toolbar: {
            autoSelected: 'zoom'
        }
    },
    dataLabels: {
        enabled: false
    },
    markers: {
        size: 0,
    },
    title: {
        text: props.xtext,
        align: 'right'
    },
    fill: {
        type: 'gradient',
        gradient: {
            shadeIntensity: 1,
            inverseColors: false,
            opacityFrom: 0.5,
            opacityTo: 0,
            stops: [0, 90, 100]
        },
    },
    yaxis: {
        labels: {
            formatter: function (val) {
                return (val).toFixed(0);
            },
        },
        title: {
            text: props.ytext,
        },
    },
    xaxis: {
        type: 'datetime',
    },
    tooltip: {
        shared: false,
        y: {
            formatter: function (val) {
                return (val).toFixed(0)
            }
        }
    }
};

</script>