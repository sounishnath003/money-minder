<template>
  <div v-if="inputPassword === sitePassword" class="h-1 bg-blue-600"></div>
  <div class="flex flex-col gap-5 p-4 rounded-lg dark:bg-gray-900 bg-neutral-50 min-h-screen">
    <PasswordProtect v-if="inputPassword !== sitePassword" @inputPasswordEvent="updateInputPassword" />
    <div v-else class="flex flex-col gap-4">
      <!-- header -->
      <Header />
      <!-- header -->
      <!-- body -->
      <main class="container mx-auto">
        <RouterView />
      </main>
      <!-- body -->
      <!-- footer -->
      <Footer />
      <!-- footer -->
    </div>
  </div>

</template>


<script setup>
import { ref } from 'vue';
import { RouterView } from 'vue-router';
import Header from './components/Header.vue';
import Home from './components/Home.vue';
import Footer from './components/Footer.vue';
import { useTransactionStore } from './store/transaction.store';

import PasswordProtect from './components/PasswordProtect.vue';

// define the store
const transactionStore = useTransactionStore();
const inputPassword = ref('');

// Add a simple password input to see all transactions and modules
const sitePassword = transactionStore.sitePassword;

const updateInputPassword = (password) => {
  inputPassword.value = password;
}

</script>