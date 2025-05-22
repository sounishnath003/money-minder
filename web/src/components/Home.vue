<template>
    <div class="grid lg:grid-cols-2 gap-8 lg:gap-20 my-6">
        <AddTransaction />
        <Balance class="flex flex-col justify-start items-center order-first lg:order-last" />
    </div>
    <ListTransactions class="my-20"/>
</template>

<script setup>
import AddTransaction from './AddTransaction.vue';
import Balance from './Balance.vue';
import ListTransactions from './ListTransactions.vue';

</script>