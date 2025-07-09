import { computed, ref, watch } from 'vue';

export function useTransactionFilters(transactions) {
    // Filters state
    const filters = ref({
        startDate: '',
        endDate: '',
        paymentMethod: [],
        categoryID: '' // Added category filter
    });

    // Pagination state
    const itemsPerPage = 10;
    const currentPage = ref(1);

    // Clear filters function
    const clearFilters = () => {
        filters.value = {
            startDate: '',
            endDate: '',
            paymentMethod: [],
            categoryID: '' // Reset category filter
        };
        currentPage.value = 1; // Reset to first page when clearing filters
    };

    // Filtered transactions
    const filteredTransactions = computed(() => {
        return transactions.value.filter(transaction => {
            // Date range filter
            if (filters.value.startDate || filters.value.endDate) {
                const transactionDate = new Date(transaction.createdAt);
                const startDate = filters.value.startDate ? new Date(filters.value.startDate) : null;
                const endDate = filters.value.endDate ? new Date(filters.value.endDate) : null;

                if (startDate && transactionDate < startDate) return false;
                if (endDate && transactionDate > endDate) return false;
            }

            // Payment method filter (multi-select)
            if (filters.value.paymentMethod.length > 0 && !filters.value.paymentMethod.includes(transaction.paymentMethod)) {
                return false;
            }

            // Category filter
            if (filters.value.categoryID && transaction.categoryID !== Number(filters.value.categoryID)) {
                return false;
            }

            return true;
        });
    });

    // Watch for filter changes and reset to first page
    watch(filters, () => {
        currentPage.value = 1;
    }, { deep: true });

    // Pagination computed properties
    const totalPages = computed(() => Math.ceil(filteredTransactions.value.length / itemsPerPage));

    const startIndex = computed(() => (currentPage.value - 1) * itemsPerPage);
    const endIndex = computed(() => Math.min(startIndex.value + itemsPerPage, filteredTransactions.value.length));

    const paginatedTransactions = computed(() => {
        return filteredTransactions.value.slice(startIndex.value, endIndex.value);
    });

    return {
        // State
        filters,
        currentPage,
        itemsPerPage,

        // Computed
        filteredTransactions,
        totalPages,
        startIndex,
        endIndex,
        paginatedTransactions,

        // Methods
        clearFilters
    };
} 