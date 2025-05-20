import { defineStore } from 'pinia';
import { ref } from 'vue';

const API_BASE_URL = `http://localhost:3000`;

export interface Transaction {
    id: number;
    name: string;
    transactionType: string;
    catagoryID: number;
    amount: number;
    userID: number;
    createdAt: Date;
}


export const useTransactionStore = defineStore('transactionState', {
    state: () => ({
        addTransactionForm: {
            name: ref(''),
            transactionType: ref('Select transaction type'),
            categoryID: ref('Select category'),
            amount: ref(0)
        },
    }),
    actions: {
        async getTransactions(fromDate: Date, toDate: Date): Promise<Transaction[]> {
            try {
                const fromDateStr = fromDate.toISOString().slice(0, 10);
                const toDateStr = toDate.toISOString().slice(0, 10);

                const url = new URL(`${API_BASE_URL}/api/v1/transactions`);
                url.searchParams.append('from', fromDateStr);
                url.searchParams.append('to', toDateStr);

                const resp = await window.fetch(url.toString(), {
                    method: "GET",
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });

                if (!resp.ok) {
                    throw new Error(`HTTP error! status: ${resp.status}`);
                }

                const data = await resp.json();
                return data.data as Transaction[];
            } catch (error) {
                console.error('Failed to fetch transactions:', error);
                throw error;
            }
        },
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
    },
})