import { defineStore } from 'pinia';
import { ref } from 'vue';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || `http://localhost:3000`;

export interface User {
    userId: number;
    name: string;
}

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: ref<string | null>(localStorage.getItem('jwt_token')),
        user: ref<User | null>(null),
        loading: ref(false),
        error: ref<string | null>(null),
    }),
    actions: {
        async login(password: string) {
            this.loading = true;
            this.error = null;
            try {
                const resp = await fetch(`${API_BASE_URL}/api/v1/auth/login`, {
                    method: 'POST',
                    credentials: 'same-origin',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ password })
                });
                const data = await resp.json();
                if (!resp.ok) throw new Error(data.error?.errorMessage || 'Login failed');
                this.token = data.data.token;
                this.user = data.data.user;
                localStorage.setItem('jwt_token', this.token!);
                return true;
            } catch (e: any) {
                this.error = e.message;
                this.token = null;
                this.user = null;
                localStorage.removeItem('jwt_token');
                return false;
            } finally {
                this.loading = false;
            }
        },
        logout() {
            this.token = null;
            this.user = null;
            window.localStorage.removeItem('jwt_token');
            // For get the double surerity check
            window.location.replace('/')
        },
        loadFromStorage() {
            this.token = localStorage.getItem('jwt_token');
        }
    }
}); 