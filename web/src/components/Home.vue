<template>
    <div class="grid grid-rows-2 lg:grid-cols-2 gap-5 my-6">
        <AddTransaction />
        <Balance class="flex flex-col justify-start items-start md:items-center  order-first lg:order-last" />
    </div>
</template>

<script setup>
import AddTransaction from './AddTransaction.vue';
import Balance from './Balance.vue';


</script>