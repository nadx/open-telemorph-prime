// Open-Telemorph-Prime Web App JavaScript

class TelemorphApp {
    constructor() {
        this.apiBase = '/api/v1';
        this.init();
    }

    init() {
        this.setupEventListeners();
        this.loadInitialData();
    }

    setupEventListeners() {
        // Add any global event listeners here
        console.log('Open-Telemorph-Prime initialized');
    }

    async loadInitialData() {
        try {
            // Load services for navigation
            const services = await this.fetchServices();
            this.updateServicesList(services);
        } catch (error) {
            console.error('Failed to load initial data:', error);
        }
    }

    async fetchServices() {
        const response = await fetch(`${this.apiBase}/services`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data.services || [];
    }

    updateServicesList(services) {
        // Update any service-related UI elements
        console.log('Loaded services:', services);
    }

    // Generic API call method
    async apiCall(endpoint, options = {}) {
        const url = `${this.apiBase}${endpoint}`;
        const defaultOptions = {
            headers: {
                'Content-Type': 'application/json',
            },
        };

        const response = await fetch(url, { ...defaultOptions, ...options });
        
        if (!response.ok) {
            throw new Error(`API call failed: ${response.status} ${response.statusText}`);
        }

        return await response.json();
    }

    // Metrics methods
    async getMetrics(limit = 100, offset = 0) {
        return await this.apiCall(`/metrics?limit=${limit}&offset=${offset}`);
    }

    // Traces methods
    async getTraces(limit = 100, offset = 0) {
        return await this.apiCall(`/traces?limit=${limit}&offset=${offset}`);
    }

    // Logs methods
    async getLogs(limit = 100, offset = 0) {
        return await this.apiCall(`/logs?limit=${limit}&offset=${offset}`);
    }

    // Query method
    async query(type, query, limit = 100, offset = 0) {
        return await this.apiCall('/query', {
            method: 'POST',
            body: JSON.stringify({
                type,
                query,
                limit,
                offset
            })
        });
    }

    // Utility methods
    formatTimestamp(timestamp) {
        return new Date(timestamp).toLocaleString();
    }

    formatDuration(nanoseconds) {
        if (nanoseconds < 1000) {
            return `${nanoseconds}ns`;
        } else if (nanoseconds < 1000000) {
            return `${(nanoseconds / 1000).toFixed(2)}μs`;
        } else if (nanoseconds < 1000000000) {
            return `${(nanoseconds / 1000000).toFixed(2)}ms`;
        } else {
            return `${(nanoseconds / 1000000000).toFixed(2)}s`;
        }
    }

    showLoading(element) {
        if (element) {
            element.innerHTML = '<div class="loading"></div>';
        }
    }

    showError(element, message) {
        if (element) {
            element.innerHTML = `<div class="error">Error: ${message}</div>`;
        }
    }
}

// Initialize the app when DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
    window.telemorphApp = new TelemorphApp();
});

// Export for use in other scripts
if (typeof module !== 'undefined' && module.exports) {
    module.exports = TelemorphApp;
}
