<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900">
    <form @submit.prevent="onLogin" class="bg-white dark:bg-gray-800 p-8 rounded-lg shadow-lg w-full max-w-sm flex flex-col gap-4">
      <h2 class="text-2xl font-bold text-center text-blue-600 dark:text-blue-400">Login</h2>
      <input v-model="email" type="email" placeholder="Email" required class="px-4 py-2 rounded border border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white" />
      <input v-model="password" type="password" placeholder="Password" required class="px-4 py-2 rounded border border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white" />
      <button type="submit" :disabled="auth.loading" class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 rounded">
        <span v-if="auth.loading">Logging in...</span>
        <span v-else>Login</span>
      </button>
      <div v-if="auth.error" class="text-red-500 text-center text-sm">{{ auth.error }}</div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import { useAuthStore } from '../store/auth.store';

const emit = defineEmits(['loginSuccess']);
const auth = useAuthStore();
const email = ref('');
const password = ref('');

async function onLogin() {
  const ok = await auth.login(email.value, password.value);
  if (ok) emit('loginSuccess');
}

// Optionally clear error on input
watch([email, password], () => { auth.error = null; });
</script> 