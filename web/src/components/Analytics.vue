<template>
    <div class="text-2xl font-medium my-1 lg:text-left text-center">Understand spends with <b>Simple Analytics</b></div>
    <div class="flex flex-col gap-6 mx-auto items-center justify-center w-full">
        <LineChart name="asd" :data="dates" height="350" width="500" xtext="Timeline" ytext="Stock" />
        <BarChart />
    </div>
</template>

<script setup>
import BarChart from './charts/BarChart.vue';
import LineChart from './charts/LineChart.vue';

const dates = [
    [1716537600000, 10],
    [1716624000000, 20],
    [1716710400000, 30],
    [1716796800000, 40],
    [1716883200000, 50],
    [1716969600000, 60],
    [1717056000000, 70],
]

</script>