<template>
  <div>
    <PasswordProtect v-if="!auth.token" />
    <div v-else class="flex flex-col gap-5 p-4 rounded-lg dark:bg-gray-900 bg-neutral-50 min-h-screen">
      <!-- header -->
      <Header />
      <!-- body -->
      <main class="container w-full">
        <RouterView />
      </main>
      <!-- footer -->
      <Footer />
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router';
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
import { useAuthStore } from './store/auth.store';
import PasswordProtect from './components/PasswordProtect.vue';

const auth = useAuthStore();
</script>