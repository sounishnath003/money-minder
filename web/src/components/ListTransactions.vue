<template>
    <div>
        <div class="my-5">
            <h2 class="font-medium text-xl lg:text-3xl text-blue-600 dark:text-white">&bull; All transactions ({{
                transactions.length
                }})</h2>
        </div>
        <div class="w-full overflow-x-auto">
            <div class="min-w-full inline-block align-middle">
                <div class="overflow-hidden">
                    <table class="min-w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                            <tr>
                                <th scope="col" class="px-6 py-3">
                                    Created At
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    Category
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    Name
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    Payment Method
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    Transaction Type
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    Amount
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="transaction in paginatedTransactions" :id="transaction.id"
                                class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200">
                                <th scope="row"
                                    class="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {{ transaction.createdAt.toLocaleString().split('T')[0] }}
                                </th>
                                <th scope="row"
                                    class="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {{ getCategoryName(transaction.categoryID) }}
                                </th>
                                <th scope="row"
                                    class="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {{ transaction.name }}
                                </th>
                                <th scope="row" class="px-6 py-2 font-medium text-gray-900 whitespace-nowrap">
                                    <span class="px-3 py-1 rounded-full text-sm"
                                        :class="getPaymentMethodColor(transaction.paymentMethod)">
                                        {{ transaction.paymentMethod }}
                                    </span>
                                </th>
                                <th scope="row" class="px-6 py-2 font-medium text-gray-900 whitespace-nowrap">
                                    <span class="px-3 py-1 rounded-full text-sm"
                                        :class="transaction.transactionType === 'Income' ? 'bg-green-50 text-green-600' : 'bg-red-50 text-red-600'">
                                        {{ transaction.transactionType }}
                                    </span>
                                </th>
                                <th scope="row" class="px-6 py-2 font-medium whitespace-nowrap"
                                    :class="transaction.transactionType === 'Income' ? 'text-green-600' : 'text-red-600'">
                                    {{ transaction.transactionType === 'Income' ? '+' : '-' }}{{
                                        INRRuppe.format(transaction.amount)
                                    }}
                                </th>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>

        <!-- Pagination -->
        <div class="flex items-center justify-between mt-4">
            <div class="flex items-center">
                <span class="text-sm text-gray-700 dark:text-gray-400">
                    Showing {{ startIndex + 1 }} to {{ endIndex }} of {{ transactions.length }} entries
                </span>
            </div>
            <div class="flex items-center space-x-2">
                <button @click="currentPage--" :disabled="currentPage === 1"
                    class="cursor-pointer px-3 py-1 rounded-md bg-gray-200 dark:bg-gray-700 disabled:opacity-50">
                    Previous
                </button>
                <span class="text-sm text-gray-700 dark:text-gray-400">
                    Page {{ currentPage }} of {{ totalPages }}
                </span>
                <button @click="currentPage++" :disabled="currentPage === totalPages"
                    class="cursor-pointer px-3 py-1 rounded-md bg-gray-200 dark:bg-gray-700 disabled:opacity-50">
                    Next
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

const transactionStore = useTransactionStore();
const transactions = computed(() => transactionStore.allTransactions);

// Pagination
const itemsPerPage = 10;
const currentPage = ref(1);

const totalPages = computed(() => Math.ceil(transactions.value.length / itemsPerPage));

const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage);
const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage, transactions.value.length));

const paginatedTransactions = computed(() => {
    return transactions.value.slice(startIndex.value, endIndex.value);
});

const getCategoryName = (categoryId) => {
    const category = transactionStore.allCategories.find(cat => cat.id === Number(categoryId));
    return category ? category.category : 'Unknown';
};

// Utilities

const getPaymentMethodColor = (paymentMethod) => {
    switch (paymentMethod) {
        case 'UPI':
            return 'bg-blue-50 text-blue-600 dark:bg-blue-900 dark:text-blue-300';
        case 'Cash':
            return 'bg-green-50 text-green-600 dark:bg-green-900 dark:text-green-300';
        case 'Bank Deposit':
            return 'bg-purple-50 text-purple-600 dark:bg-purple-900 dark:text-purple-300';
        case 'Credit Card':
            return 'bg-red-50 text-red-600 dark:bg-red-900 dark:text-red-300';
        default:
            return 'bg-gray-50 text-gray-600 dark:bg-gray-700 dark:text-gray-300';
    }
};


const USDollar = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
});

const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});
</script>