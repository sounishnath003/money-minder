<template>
    <div class="flex items-center justify-between mt-4">
        <div class="flex items-center">
            <span class="text-sm text-gray-700 dark:text-gray-400">
                Showing {{ startIndex + 1 }} to {{ endIndex }} of {{ totalItems }} entries
            </span>
        </div>
        <div class="flex items-center space-x-2">
            <button @click="$emit('previous-page')" :disabled="currentPage === 1"
                class="cursor-pointer px-3 py-1 rounded-md bg-gray-200 dark:bg-gray-700 disabled:opacity-50 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors">
                Previous
            </button>
            <span class="text-sm text-gray-700 dark:text-gray-400">
                Page {{ currentPage }} of {{ totalPages }}
            </span>
            <button @click="$emit('next-page')" :disabled="currentPage === totalPages"
                class="cursor-pointer px-3 py-1 rounded-md bg-gray-200 dark:bg-gray-700 disabled:opacity-50 hover:bg-gray-300 dark:hover:bg-gray-600 transition-colors">
                Next
            </button>
        </div>
    </div>
</template>

<script setup>
defineProps({
    currentPage: {
        type: Number,
        required: true
    },
    totalPages: {
        type: Number,
        required: true
    },
    startIndex: {
        type: Number,
        required: true
    },
    endIndex: {
        type: Number,
        required: true
    },
    totalItems: {
        type: Number,
        required: true
    }
});

defineEmits(['previous-page', 'next-page']);
</script>