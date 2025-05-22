<template>
    <div class="flex flex-col gap-4 min-h-screen justify-center items-center">
        <h2 class="font-medium text-xl text-red-600 dark:text-white">Enter password to access</h2>
        <input type="password" v-model="inputPassword" placeholder="Type password here"
            class="w-full p-2 rounded-lg border border-gray-300 dark:border-gray-700 dark:bg-gray-800 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
    </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const emit = defineEmits(['inputPasswordEvent']);
const inputPassword = ref('');

watch(inputPassword, (newValue) => {
    emit('inputPasswordEvent', newValue);
});
</script>