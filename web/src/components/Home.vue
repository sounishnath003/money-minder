<template>
    <div class="grid grid-rows-2 lg:grid-cols-2 gap-2 md:gap-8 md:gap-20 my-6 mx-3">
        <AddTransaction />
        <Balance class="flex flex-col justify-start items-center order-first lg:order-last" />
    </div>
</template>

<script setup>
import AddTransaction from './AddTransaction.vue';
import Balance from './Balance.vue';


</script>