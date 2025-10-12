// Open-Telemorph-Prime Web App JavaScript

class TelemorphApp {
    constructor() {
        this.apiBase = '/api/v1';
        this.isSidebarOpen = false;
        this.searchTimeout = null;
        this.notifications = [];
        this.init();
    }

    init() {
        this.setupEventListeners();
        this.loadTheme();
        this.loadInitialData();
        this.initializeNotifications();
        this.setupResponsiveHandlers();
    }

    setupEventListeners() {
        // Add any global event listeners here
        console.log('Open-Telemorph-Prime initialized');
        
        // Set active navigation item
        this.setActiveNavItem();
        
        // Initialize search
        this.initializeSearch();
        
        // Setup keyboard shortcuts
        this.setupKeyboardShortcuts();
    }

    // Theme management
    toggleTheme() {
        const isDark = document.documentElement.classList.contains('dark');
        if (isDark) {
            document.documentElement.classList.remove('dark');
            localStorage.setItem('theme', 'light');
            this.updateThemeIcon('sun');
        } else {
            document.documentElement.classList.add('dark');
            localStorage.setItem('theme', 'dark');
            this.updateThemeIcon('moon');
        }
    }

    updateThemeIcon(iconType) {
        const themeIcon = document.querySelector('.theme-icon');
        if (!themeIcon) return;
        
        if (iconType === 'sun') {
            // Sun icon
            themeIcon.innerHTML = `
                <circle cx="12" cy="12" r="5"></circle>
                <line x1="12" y1="1" x2="12" y2="3"></line>
                <line x1="12" y1="21" x2="12" y2="23"></line>
                <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
                <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
                <line x1="1" y1="12" x2="3" y2="12"></line>
                <line x1="21" y1="12" x2="23" y2="12"></line>
                <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
                <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
            `;
        } else {
            // Moon icon
            themeIcon.innerHTML = `
                <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
            `;
        }
    }

    loadTheme() {
        const savedTheme = localStorage.getItem('theme');
        const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        
        if (savedTheme === 'dark' || (!savedTheme && prefersDark)) {
            document.documentElement.classList.add('dark');
            this.updateThemeIcon('moon');
        } else {
            document.documentElement.classList.remove('dark');
            this.updateThemeIcon('sun');
        }
    }

    // Navigation management
    setActiveNavItem() {
        const currentPath = window.location.pathname;
        const navItems = document.querySelectorAll('.nav-item');
        
        navItems.forEach(item => {
            item.classList.remove('active');
            if (item.getAttribute('href') === currentPath) {
                item.classList.add('active');
            }
        });
    }

    // Search functionality
    initializeSearch() {
        const searchInput = document.querySelector('.search-input');
        if (!searchInput) return;
        
        searchInput.addEventListener('keypress', (e) => {
            if (e.key === 'Enter') {
                const query = e.target.value.trim();
                if (query) {
                    this.performSearch(query);
                }
            }
        });
    }

    performSearch(query) {
        console.log('Searching for:', query);
        
        // Clear previous timeout
        if (this.searchTimeout) {
            clearTimeout(this.searchTimeout);
        }
        
        // Debounce search
        this.searchTimeout = setTimeout(() => {
            this.executeSearch(query);
        }, 300);
    }
    
    async executeSearch(query) {
        try {
            const response = await fetch(`${this.apiBase}/query`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    query: query,
                    type: 'search'
                })
            });
            
            if (response.ok) {
                const results = await response.json();
                this.displaySearchResults(results);
            } else {
                console.error('Search failed:', response.statusText);
            }
        } catch (error) {
            console.error('Search error:', error);
        }
    }
    
    displaySearchResults(results) {
        // TODO: Implement search results display
        console.log('Search results:', results);
    }
    
    // Mobile sidebar management
    toggleSidebar() {
        const sidebar = document.getElementById('sidebar');
        const overlay = document.getElementById('sidebarOverlay');
        
        this.isSidebarOpen = !this.isSidebarOpen;
        
        if (this.isSidebarOpen) {
            sidebar.classList.add('open');
            overlay.classList.add('open');
            document.body.style.overflow = 'hidden';
        } else {
            sidebar.classList.remove('open');
            overlay.classList.remove('open');
            document.body.style.overflow = '';
        }
    }
    
    closeSidebar() {
        const sidebar = document.getElementById('sidebar');
        const overlay = document.getElementById('sidebarOverlay');
        
        this.isSidebarOpen = false;
        sidebar.classList.remove('open');
        overlay.classList.remove('open');
        document.body.style.overflow = '';
    }
    
    // Responsive handlers
    setupResponsiveHandlers() {
        window.addEventListener('resize', () => {
            if (window.innerWidth > 768 && this.isSidebarOpen) {
                this.closeSidebar();
            }
        });
    }
    
    // Keyboard shortcuts
    setupKeyboardShortcuts() {
        document.addEventListener('keydown', (e) => {
            // Ctrl/Cmd + K for search
            if ((e.ctrlKey || e.metaKey) && e.key === 'k') {
                e.preventDefault();
                const searchInput = document.querySelector('.search-input');
                if (searchInput) {
                    searchInput.focus();
                }
            }
            
            // Escape to close sidebar
            if (e.key === 'Escape' && this.isSidebarOpen) {
                this.closeSidebar();
            }
            
            // Ctrl/Cmd + D for dark mode toggle
            if ((e.ctrlKey || e.metaKey) && e.key === 'd') {
                e.preventDefault();
                this.toggleTheme();
            }
        });
    }
    
    // Notification system
    initializeNotifications() {
        // Load notifications from localStorage
        const savedNotifications = localStorage.getItem('telemorph-notifications');
        if (savedNotifications) {
            this.notifications = JSON.parse(savedNotifications);
        }
        
        this.updateNotificationBadge();
    }
    
    addNotification(notification) {
        this.notifications.unshift({
            id: Date.now(),
            ...notification,
            timestamp: new Date().toISOString()
        });
        
        // Keep only last 50 notifications
        this.notifications = this.notifications.slice(0, 50);
        
        localStorage.setItem('telemorph-notifications', JSON.stringify(this.notifications));
        this.updateNotificationBadge();
    }
    
    updateNotificationBadge() {
        const badge = document.querySelector('.notification-badge');
        if (badge) {
            const unreadCount = this.notifications.filter(n => !n.read).length;
            badge.textContent = unreadCount;
            badge.style.display = unreadCount > 0 ? 'flex' : 'none';
        }
    }
    
    markNotificationAsRead(notificationId) {
        const notification = this.notifications.find(n => n.id === notificationId);
        if (notification) {
            notification.read = true;
            localStorage.setItem('telemorph-notifications', JSON.stringify(this.notifications));
            this.updateNotificationBadge();
        }
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
    
    // Make functions available globally for onclick handlers
    window.toggleTheme = () => window.telemorphApp.toggleTheme();
    window.toggleSidebar = () => window.telemorphApp.toggleSidebar();
    window.closeSidebar = () => window.telemorphApp.closeSidebar();
});

// Listen for system theme changes
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function(e) {
    if (!localStorage.getItem('theme')) {
        if (e.matches) {
            document.documentElement.classList.add('dark');
            if (window.telemorphApp) {
                window.telemorphApp.updateThemeIcon('moon');
            }
        } else {
            document.documentElement.classList.remove('dark');
            if (window.telemorphApp) {
                window.telemorphApp.updateThemeIcon('sun');
            }
        }
    }
});

// Export for use in other scripts
if (typeof module !== 'undefined' && module.exports) {
    module.exports = TelemorphApp;
}
