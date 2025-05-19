import { defineStore } from 'pinia';
import { ref } from 'vue';


export const useTransactionStore = defineStore('transactionState', {
    state: () => ({
        addTransactionForm: {
            name: ref(''),
            transactionType: ref('Select transaction type'),
            categoryID: ref('Select category'),
            amount: ref(0)
        }
    }),
    actions: {
        async createTransaction() {
            console.table({
                name: this.addTransactionForm.name,
                transactionType: this.addTransactionForm.transactionType,
                categoryID: this.addTransactionForm.categoryID,
                amount: this.addTransactionForm.amount
            });
            console.log('transaction saved.');
            this.resetTransactionForm();
        },
        resetTransactionForm() {
            this.addTransactionForm.name = '';
            this.addTransactionForm.transactionType = 'Select transaction type';
            this.addTransactionForm.categoryID = 'Select category';
            this.addTransactionForm.amount = 0;
        }
    }
})