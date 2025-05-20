<template>
  <div class="h-1 bg-blue-600"></div>
  <div class="flex flex-col gap-5 p-4 rounded-lg dark:bg-gray-900 bg-neutral-50 min-h-screen">
    <Header />
    <main class="container">
      <RouterView />
    </main>
    <footer class="font-medium">
      🚧 work is in progress
    </footer>
  </div>
</template>


<script setup>
import { RouterView } from 'vue-router';
import Header from './components/Header.vue';
import Home from './components/Home.vue';


</script>