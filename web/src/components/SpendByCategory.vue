<template>
    <div class="text-xl text-center my-2 font-medium text-blue-600 dark:text-white"><span>&bull; Total spend in current
            month (<span class="font-semibold">{{ currentMonth }}</span>)</span>
    </div>
    <div class="w-full overflow-x-auto">
        <div class="min-w-full">
            <div class="overflow-x-auto overflow-y-auto max-h-[300px]">
                <table class="min-w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead
                        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 sticky top-0">
                        <tr>
                            <th scope="col" class="px-6 py-3">
                                Category
                            </th>
                            <th scope="col" class="px-6 py-3">
                                Total spend
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="spnd in spendByCategories" :id="spnd.category"
                            class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200 whitespace-normal md:whitespace-nowrap">
                            <th scope="row"
                                class="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                {{ spnd.category }}
                            </th>
                            <td class="px-6 py-2 font-semibold">
                                {{ INRRuppe.format(spnd.totalAmount) }}
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useTransactionStore } from '../store/transaction.store';
import { computed } from 'vue';

const transactionStore = useTransactionStore();

const spendByCategories = computed(() => transactionStore.allSpendsByCategories);
const currentMonth = computed(() => transactionStore.currentMonth);

// Utilities
const USDollar = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
});

const INRRuppe = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
});

</script>