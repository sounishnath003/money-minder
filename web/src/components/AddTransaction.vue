<template>
    <div class="flex flex-col gap-4">
        <div>
            <h2 class="font-medium text-xl lg:text-3xl text-blue-600 dark:text-white">&bull; Add new transaction</h2>
        </div>
        <div>
            <form id="addTransactionForm" @submit.prevent="onSubmit">
                <div class="mb-5">
                    <label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">What have you
                        spend for:</label>
                    <input autocomplete="off" v-model="transactionStore.addTransactionForm.name" type="text" id="name"
                        :class="formInputCss" placeholder="Lunch and foods" required />
                </div>
                <div class="mb-5">
                    <label for="transactionType"
                        class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Transaction Type:</label>
                    <select v-model="transactionStore.addTransactionForm.transactionType" name="transactionType"
                        id="transactionType" :class="formInputCss">
                        <option value="Select transaction type">Select transaction type</option>
                        <option value="Income">Income</option>
                        <option value="Expense">Expense</option>
                        <option value="Investments">Investments</option>
                    </select>
                </div>
                <div class="mb-5">
                    <label for="categoryID"
                        class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Category:</label>
                    <select v-model="transactionStore.addTransactionForm.categoryID" name="categoryID" id="categoryID"
                        :class="formInputCss">
                        <option v-for="category in categories" :id="category.id" :value="category.id">{{
                            category.category }}</option>
                    </select>
                </div>
                <div class="mb-5">
                    <label for="paymentMethod"
                        class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Payment Method:</label>
                    <select v-model="transactionStore.addTransactionForm.paymentMethod" name="paymentMethod"
                        id="paymentMethod" :class="formInputCss">
                        <option value="Select payment method">Select payment method</option>
                        <option v-for="paymentMethod in paymentMethods" :id="paymentMethod" :value="paymentMethod">{{
                            paymentMethod }}</option>
                    </select>
                </div>
                <div class="mb-5">
                    <label for="amount"
                        class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Amount:</label>
                    <input autocomplete="off" v-model="transactionStore.addTransactionForm.amount" type="number"
                        id="amount" :class="formInputCss" placeholder="How much have you spent" required />
                </div>
                <div class="w-full block"><button type="submit"
                        :disabled="transactionStore.addTransactionForm.amount <= 0"
                        class="w-full cursor-pointer text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
                        ref="submitBtn">{{
                            submitBtnText }}</button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useTransactionStore } from '../store/transaction.store';

// define the store
const transactionStore = useTransactionStore();

// define categories list
const categories = ref([]);
const paymentMethods = ref([]);
const submitBtnText = ref('Create new transaction');
const submitBtn = ref(null);

// watch the submitBtn
watch(submitBtn, (newVal) => {
    if (newVal) {
        newVal.disabled = true;
    } else {
        newVal.disabled = false;
    }
})

// formInput box css
const formInputCss = `w-full px-4 py-3 rounded-lg border border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white  focus:ring-2 focus:ring-blue-500 focus:border-blue-500 placeholder-gray-400 dark:placeholder-gray-500`;

// On component mount fetch all transactions
onMounted(async () => {
    await fetchCategories();
})

// fetch all categories
async function fetchCategories() {
    categories.value = await transactionStore.getAllCategories();
    paymentMethods.value = await transactionStore.getPaymentMethods();
}


// On submit trasaction save request
const onSubmit = async () => {
    submitBtnText.value = 'Creating...';
    // also make the button disabled
    submitBtn.value.disabled = true;
    await transactionStore.createTransaction();
    submitBtnText.value = 'Create new transaction';
    submitBtn.value.disabled = false;
}
</script>
