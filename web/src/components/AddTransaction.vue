<template>
    <div class="flex flex-col gap-4">
        <div class="text-2xl font-medium">Add new transaction</div>
        <div>
            <form method="post" action="/">
                <div v-for="form in addTransactionForm" class="mb-5" :key="form.id">
                    <label for="first_name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{{
                        form.name }}:</label>
                    <input type="text" id="first_name" :class="formInputCss" :placeholder="form.placeholder" required />
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const formInputCss = `.input-box {
  @apply bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus: ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500
}`
const name = ref('');

const addTransactionForm = [
    {
        id: 1,
        name: 'What have you spent for',
        placeholder: 'Lunch and foods',
    },
    {
        id: 2,
        name: 'Transaction Type',
        placeholder: 'Income or Expense',
    },
]

</script>
