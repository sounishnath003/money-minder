import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    tailwindcss()
  ],
  build: {
    // Increase chunk size warning limit
    chunkSizeWarningLimit: 1000,
    rollupOptions: {
      output: {
        // Manual chunk splitting for better caching
        manualChunks: {
          // Vendor chunks
          'vue-vendor': ['vue', 'vue-router', 'pinia'],
          'charts-vendor': ['vue3-apexcharts', 'apexcharts'],
          // Separate analytics components
          'analytics': [
            './src/components/Analytics.vue',
            './src/components/analytics/AnalyticsSummaryCards.vue',
            './src/components/analytics/SpendingTrendsChart.vue',
            './src/components/analytics/CategoryGrowthList.vue',
            './src/components/analytics/PaymentMethodsKPI.vue'
          ],
          // Separate chart components
          'charts': [
            './src/components/charts/BarChart.vue',
            './src/components/charts/LineChart.vue',
            './src/components/charts/PieChart.vue'
          ]
        }
      }
    },
    // Enable source maps for debugging (optional)
    sourcemap: false,
    // Minify CSS
    cssMinify: true,
    // Target modern browsers for smaller bundles
    target: 'es2015'
  },
  // Optimize dependencies
  optimizeDeps: {
    include: ['vue', 'vue-router', 'pinia']
  }
})