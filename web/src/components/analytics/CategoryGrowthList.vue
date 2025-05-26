<template>
    <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
        <div class="text-lg font-medium mb-4">Category Growth Trends</div>
        <div class="space-y-3">
            <div v-for="category in categoryGrowth" :key="category.name" class="flex justify-between items-center">
                <span class="text-gray-600 dark:text-gray-400">{{ category.name }}</span>
                <span :class="[
                    'font-medium',
                    parseFloat(category.growth) > 0 ? 'text-red-600' : 'text-green-600'
                ]">
                    {{ parseFloat(category.growth) > 0 ? '+' : '' }}{{ category.growth }}%
                </span>
            </div>
            <div v-if="!categoryGrowth.length" class="text-center text-gray-500 py-4">
                No category growth data available
            </div>
        </div>
    </div>
</template>

<script setup>
defineProps({
    categoryGrowth: {
        type: Array,
        required: true,
        default: () => []
    }
});
</script>
