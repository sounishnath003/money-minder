<template>
    <div class="flex items-center justify-between w-full gap-5">
        <RouterLink to="/" class="font-medium lg:text-2xl font-sans dark:text-white text-blue-600"> Money Minder </RouterLink>
        <div class="flex-1"></div>
        <RouterLink to="/"> Home </RouterLink>
        <RouterLink to="/analytics"> Analytics </RouterLink>
    </div>
</template>

<script setup>
import { RouterLink } from 'vue-router';


</script>