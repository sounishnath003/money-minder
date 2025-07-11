import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useAuthStore } from './auth.store';

const COMMIT_HASH = import.meta.env.VITE_COMMIT_HASH || 'default-commit-hash-1.0.0';
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || `http://localhost:3000`;

export interface Transaction {
    id: number;
    name: string;
    transactionType: string;
    catagoryID: number;
    paymentMethod: string;
    amount: number;
    userID: number;
    createdAt: Date;
}

function getAuthHeaders() {
    const authStore = useAuthStore();
    if (!authStore.token) {
        authStore.logout();
        throw new Error('Not authenticated');
    }
    return {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
    };
}

export const useTransactionStore = defineStore('transactionState', {
    state: () => ({
        commitHash: ref(COMMIT_HASH),
        startDate: ref(new Date(new Date().getFullYear(), new Date().getMonth(), 1, 0, 0, 0)),
        endDate: ref(new Date(new Date().getFullYear(), new Date().getMonth() + 1, 0, 23, 59, 0)),
        currentMonth: ref(
            `${new Date().toLocaleString('default', { month: 'long' })}-${new Date().getFullYear()}`
        ),
        sitePassword: ref(import.meta.env.VITE_SECRET_PASSWORD),
        addTransactionForm: {
            name: ref(''),
            transactionType: ref('Select transaction type'),
            paymentMethod: ref('Select payment method'),
            categoryID: ref(0),
            amount: ref(0)
        },
        allTransactions: ref<Transaction[]>([]),
        allSpendsByCategories: ref<{ category: string, totalAmount: number }[]>([]),
        allCategories: ref<{ id: number; category: string; }[]>([]),
        paymentMethods: ref<string[]>([]),
        totalDailySpends: ref<{ unixMiliseconds: number, amount: number }[]>([]),
        spendOnCategoriesMonthOnMonth: ref<{ month: string, category: string, totalSpendAmount: number }[]>([]),
        spendOnCategoriesByAllYearMonthAggregated: ref<{ year: number, month: number, monthYearStr: string, category: string, paymentMethod: string, totalAmount: number }[]>([]),
    }),
    actions: {
        async getAllCategories() {
            try {
                const resp = await window.fetch(`${API_BASE_URL}/api/v1/categories`, {
                    method: "GET",
                    credentials: 'same-origin',
                    headers: getAuthHeaders(),
                });
                const data = await resp.json();
                this.allCategories = data.data as { id: number; category: string; }[];
                return this.allCategories;
            } catch (error) {
                // Reload back to login page.
                window.location.replace('/');
                return;
            }
        },
        async getPaymentMethods() {
            const resp = await window.fetch(`${API_BASE_URL}/api/v1/paymentMethods`, {
                method: "GET",
                credentials: 'same-origin',
                headers: getAuthHeaders(),
            });
            const data = await resp.json();
            this.paymentMethods = data.data as string[];
            return this.paymentMethods;
        },
        async getTransactions(): Promise<Transaction[]> {
            try {
                const fromDateStr = this.startDate.toISOString();
                const toDateStr = this.endDate.toISOString();

                const url = new URL(`${API_BASE_URL}/api/v1/transactions`);
                url.searchParams.append('from', fromDateStr);
                url.searchParams.append('to', toDateStr);

                const resp = await window.fetch(url.toString(), {
                    method: "GET",
                    credentials: 'same-origin',
                    headers: getAuthHeaders(),
                });

                if (!resp.ok) {
                    throw new Error(`HTTP error! status: ${resp.status}`);
                }

                const data = await resp.json() as { data: Transaction[] };
                // Convert string dates to Date objects for proper date handling in the UI
                // This ensures consistent date formatting and enables date operations
                this.allTransactions = data.data.map(tx => ({
                    ...tx,
                    createdAt: new Date(tx.createdAt)
                }));
                // Update spend by categories
                Promise.all([this.getAllSpendsByCategories(), this.getSpendOnCategoriesByAllYearMonthAggregated()])
                return data.data as Transaction[];
            } catch (error) {
                console.error('Failed to fetch transactions:', error);
                throw error;
            }
        },
        async getAllSpendsByCategories(): Promise<{ category: string, totalAmount: number }[]> {
            const fromDateStr = this.startDate.toISOString();
            const toDateStr = this.endDate.toISOString();

            const url = new URL(`${API_BASE_URL}/api/v1/transactions/spendByCategory`);
            url.searchParams.append('from', fromDateStr);
            url.searchParams.append('to', toDateStr);

            const resp = await window.fetch(url.toString(), {
                method: "GET",
                credentials: 'same-origin',
                headers: getAuthHeaders(),
            });

            if (!resp.ok) {
                throw new Error(`HTTP error! status: ${resp.status}`);
            }

            const data = await resp.json();
            this.allSpendsByCategories = data.data as { category: string, totalAmount: number }[];
            return data.data as { category: string, totalAmount: number }[];
        },
        async getDailyTotalSpendByTimeframe(): Promise<{ unixMiliseconds: number, amount: number }[]> {
            const fromDateStr = this.startDate.toISOString();
            const toDateStr = this.endDate.toISOString();

            const url = new URL(`${API_BASE_URL}/api/v1/analytics/getDailySpendByTimeframe`);
            url.searchParams.append('from', fromDateStr);
            url.searchParams.append('to', toDateStr);

            const resp = await window.fetch(url.toString(), {
                method: "GET",
                credentials: 'same-origin',
                headers: getAuthHeaders(),
            });

            if (!resp.ok) {
                throw new Error(`HTTP error! status: ${resp.status}`);
            }

            const data = await resp.json();
            this.totalDailySpends = data.data as { unixMiliseconds: number, amount: number }[];
            return data.data as { unixMiliseconds: number, amount: number }[];
        },
        async getSpendOnCategoriesMonthOnMonth(): Promise<{ month: string, category: string, totalSpendAmount: number }[]> {
            const fromDateStr = this.startDate.toISOString();
            const toDateStr = this.endDate.toISOString();

            const url = new URL(`${API_BASE_URL}/api/v1/analytics/getSpendOnCategoriesMonthOnMonth`);
            url.searchParams.append('from', fromDateStr);
            url.searchParams.append('to', toDateStr);

            const resp = await window.fetch(url.toString(), {
                method: "GET",
                credentials: 'same-origin',
                headers: getAuthHeaders(),
            });

            if (!resp.ok) {
                throw new Error(`HTTP error! status: ${resp.status}`);
            }

            const data = await resp.json();
            this.spendOnCategoriesMonthOnMonth = data.data as { month: string, category: string, totalSpendAmount: number }[];
            return data.data as { month: string, category: string, totalSpendAmount: number }[];
        },
        async getSpendOnCategoriesByAllYearMonthAggregated(): Promise<{ year: number, month: number, monthYearStr: string, category: string, totalAmount: number }[]> {
            const url = new URL(`${API_BASE_URL}/api/v1/analytics/getSpendOnCategoriesByAggregatedByYearMonth`);
            const resp = await window.fetch(url.toString(), {
                method: "GET",
                credentials: 'same-origin',
                headers: getAuthHeaders(),
            });

            if (!resp.ok) {
                throw new Error(`HTTP error! status: ${resp.status}`);
            }

            const data = await resp.json();
            this.spendOnCategoriesByAllYearMonthAggregated = data.data as { year: number, month: number, monthYearStr: string, category: string, paymentMethod: string, totalAmount: number }[];
            return data.data as { year: number, month: number, monthYearStr: string, category: string, totalAmount: number }[];
        },
        async createTransaction() {
            const newTransaction = {
                id: Date.now(),
                userID: 1, // will be set by backend
                name: this.addTransactionForm.name,
                transactionType: this.addTransactionForm.transactionType,
                categoryID: +this.addTransactionForm.categoryID,
                paymentMethod: this.addTransactionForm.paymentMethod,
                amount: this.addTransactionForm.amount,
                createdAt: new Date()
            };

            if (this.addTransactionForm.transactionType === 'Select transaction type') {
                window.alert('Please select a transaction type');
                return;
            }
            if (this.addTransactionForm.paymentMethod === 'Select payment method') {
                window.alert('Please select a payment method');
                return;
            }
            if (this.addTransactionForm.categoryID === 0) {
                window.alert('Please select a category');
                return;
            }
            if (this.addTransactionForm.amount === 0) {
                window.alert('Please provide amount > 0');
                return;
            }

            const resp = await window.fetch(`${API_BASE_URL}/api/v1/transactions`, {
                method: 'POST',
                credentials: 'same-origin',
                headers: getAuthHeaders(),
                body: JSON.stringify(newTransaction)
            });
            if (resp.status !== 200) throw new Error(JSON.stringify(await resp.json()))
            console.log('transaction saved.');
            this.resetTransactionForm();

            // Update the transactions and spends by categories
            Promise.all([
                this.getTransactions(),
                this.getAllSpendsByCategories(),
                this.getDailyTotalSpendByTimeframe(),
            ]).catch((error) => {
                console.error('Failed to update transactions and spends by categories:', error);
            });
        },
        resetTransactionForm() {
            this.addTransactionForm.name = '';
            this.addTransactionForm.transactionType = 'Select transaction type';
            this.addTransactionForm.paymentMethod = 'Select payment method';
            this.addTransactionForm.categoryID = 0;
            this.addTransactionForm.amount = 0;
        }
    },
})