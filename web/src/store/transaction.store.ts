import { defineStore } from 'pinia';
import { ref } from 'vue';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || `http://localhost:3000`;

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
        startDate: ref(new Date(new Date().getFullYear(), new Date().getMonth(), 2)),
        endDate: ref(new Date(new Date().getFullYear(), new Date().getMonth() + 1, 0)),
        addTransactionForm: {
            name: ref(''),
            transactionType: ref('Select transaction type'),
            categoryID: ref(0),
            amount: ref(0)
        },
        allTransactions: ref<Transaction[]>([]),
        allSpendsByCategories: ref<{ category: string, totalAmount: number }[]>([]),
    }),
    actions: {
        async getAllCategories() {
            const resp = await window.fetch(`${API_BASE_URL}/api/v1/categories`, {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            const data = await resp.json();
            return data.data as { id: number; category: string; }[];
        },
        async getTransactions(): Promise<Transaction[]> {
            try {
                const fromDateStr = this.startDate.toISOString().slice(0, 10);
                const toDateStr = this.endDate.toISOString().slice(0, 10);

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
                this.allTransactions = data.data as Transaction[];
                // Update spend by categories
                await this.getAllSpendsByCategories();
                return data.data as Transaction[];
            } catch (error) {
                console.error('Failed to fetch transactions:', error);
                throw error;
            }
        },
        async getAllSpendsByCategories(): Promise<{ category: string, totalAmount: number }[]> {
            const fromDateStr = this.startDate.toISOString().slice(0, 10);
            const toDateStr = this.endDate.toISOString().slice(0, 10);

            const url = new URL(`${API_BASE_URL}/api/v1/transactions/spendByCategory`);
            url.searchParams.append('from', fromDateStr);
            url.searchParams.append('to', toDateStr);

            const resp = await window.fetch(url.toString(), {
                method: "GET",
                headers: {
                    'Content-Type': 'application/json',
                }
            });

            if (!resp.ok) {
                throw new Error(`HTTP error! status: ${resp.status}`);
            }

            const data = await resp.json();
            this.allSpendsByCategories = data.data as { category: string, totalAmount: number }[];
            return data.data as { category: string, totalAmount: number }[];
        },
        async createTransaction() {
            const newTransaction = {
                id: Date.now(),
                userID: 1,
                name: this.addTransactionForm.name,
                transactionType: this.addTransactionForm.transactionType,
                categoryID: +this.addTransactionForm.categoryID,
                amount: this.addTransactionForm.amount,
                createdAt: new Date()
            };

            const resp = await window.fetch(`${API_BASE_URL}/api/v1/transactions`, {
                method: 'POST',
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(newTransaction)
            });
            if (resp.status !== 200) throw new Error(JSON.stringify(await resp.json()))
            console.log('transaction saved.');
            this.resetTransactionForm();

            // Update the transactions and spends by categories
            Promise.all([this.getTransactions(), this.getAllSpendsByCategories()]).then(() => {
                console.log('transactions and spends by categories updated.');
            }).catch((error) => {
                console.error('Failed to update transactions and spends by categories:', error);
            });
        },
        resetTransactionForm() {
            this.addTransactionForm.name = '';
            this.addTransactionForm.transactionType = 'Select transaction type';
            this.addTransactionForm.categoryID = 0;
            this.addTransactionForm.amount = 0;
        }
    },
})