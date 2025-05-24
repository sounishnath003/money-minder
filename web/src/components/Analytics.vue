<template>
    <div class="text-3xl">Analytics Section</div>

    <div id="chart">
        <apexchart type="area" height="350" :options="chartOptions" :series="series"></apexchart>
    </div>
</template>

<script setup>

const dates = [
    [1716537600000, 10],
    [1716624000000, 20],
    [1716710400000, 30],
    [1716796800000, 40],
    [1716883200000, 50],
    [1716969600000, 60],
    [1717056000000, 70],
]
const series = [{
    name: 'XYZ MOTORS',
    data: dates
}];


const chartOptions = {
    chart: {
        type: 'area',
        stacked: false,
        height: 350,
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
        text: 'Stock Price Movement',
        align: 'left'
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
                return val.toFixed(0); // (val / 1000000).toFixed(0);
            },
        },
        title: {
            text: 'Price'
        },
    },
    xaxis: {
        type: 'datetime',
    },
    tooltip: {
        shared: false,
        y: {
            formatter: function (val) {
                return (val / 1000000).toFixed(0)
            }
        }
    }
};

</script>