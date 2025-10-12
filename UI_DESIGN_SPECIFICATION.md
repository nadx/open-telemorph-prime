# UI Design Specification: Matching Telemorph-Frontend Aesthetics

## Overview

This document outlines the design system and layout requirements to match the telemorph-frontend aesthetics in open-telemorph-prime. The frontend uses a modern, professional design system based on shadcn/ui components with Tailwind CSS.

## Design System Analysis

### 1. **Layout Structure**

The telemorph-frontend uses a **sidebar + main content** layout:

```
┌─────────────────────────────────────────────────────────┐
│                    Header (Top Bar)                      │
│  ┌─────────┐  ┌─────────────────────────────────────┐   │
│  │ Search  │  │ Theme │ Notifications │ User Menu │   │
│  └─────────┘  └─────────────────────────────────────┘   │
├─────────────┬───────────────────────────────────────────┤
│             │                                           │
│   Sidebar   │              Main Content                 │
│   (Fixed)   │              (Scrollable)                 │
│             │                                           │
│  ┌─────────┐│  ┌─────────────────────────────────────┐   │
│  │  Logo   ││  │        Page Header                  │   │
│  │ Telemorph││  │  Title + Actions + Time Range     │   │
│  └─────────┘│  └─────────────────────────────────────┘   │
│             │                                           │
│  ┌─────────┐│  ┌─────────────────────────────────────┐   │
│  │Dashboard││  │                                     │   │
│  │ Metrics ││  │        Page Content                 │   │
│  │ Traces  ││  │     (Cards, Tables, Charts)        │   │
│  │  Logs   ││  │                                     │   │
│  │Services ││  │                                     │   │
│  │ Alerts  ││  │                                     │   │
│  │ Query   ││  │                                     │   │
│  │ Admin   ││  │                                     │   │
│  └─────────┘│  └─────────────────────────────────────┘   │
│             │                                           │
│  ┌─────────┐│                                           │
│  │Version  ││                                           │
│  │ v0.1.0  ││                                           │
│  └─────────┘│                                           │
└─────────────┴───────────────────────────────────────────┘
```

**Key Layout Features:**
- **Sidebar**: Fixed 256px width (`w-64`)
- **Header**: Fixed 64px height (`h-16`)
- **Main Content**: Flexible, scrollable area
- **Full Height**: Uses `h-screen` for full viewport height
- **Dark Mode Support**: Automatic theme switching

### 2. **Color System**

The design uses a comprehensive color system with CSS custom properties:

#### Light Theme Colors:
```css
--background: 255, 255, 255;        /* White background */
--foreground: 0, 0, 0;              /* Black text */
--primary: 59, 130, 246;            /* Blue (#3B82F6) */
--secondary: 107, 114, 128;         /* Gray (#6B7280) */
--muted: 229, 231, 235;             /* Light gray (#E5E7EB) */
--accent: 16, 185, 129;             /* Green (#10B981) */
--destructive: 239, 68, 68;         /* Red (#EF4444) */
--card: 255, 255, 255;              /* White cards */
--border: 229, 231, 235;            /* Light gray borders */
```

#### Dark Theme Colors:
```css
--background: 17, 24, 39;           /* Dark blue-gray (#111827) */
--foreground: 255, 255, 255;        /* White text */
--primary: 96, 165, 250;            /* Light blue (#60A5FA) */
--secondary: 156, 163, 175;         /* Light gray (#9CA3AF) */
--muted: 75, 85, 99;                /* Dark gray (#4B5563) */
--accent: 52, 211, 153;             /* Light green (#34D399) */
--destructive: 248, 113, 113;       /* Light red (#F87171) */
--card: 31, 41, 55;                 /* Dark cards (#1F2937) */
--border: 75, 85, 99;               /* Dark gray borders */
```

### 3. **Typography System**

**Font Stack:**
```css
font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
```

**Text Sizes:**
- **H1**: `text-3xl font-bold` (30px, bold)
- **H2**: `text-2xl font-semibold` (24px, semibold)
- **H3**: `text-xl font-bold` (20px, bold)
- **Body**: `text-sm` (14px, regular)
- **Small**: `text-xs` (12px, regular)
- **Muted**: `text-muted-foreground` (secondary color)

### 4. **Component System**

#### **Cards**
```css
.card {
  background: rgb(var(--card));
  color: rgb(var(--card-foreground));
  border: 1px solid rgb(var(--border));
  border-radius: 0.5rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.card-header {
  padding: 1.5rem;
  border-bottom: 1px solid rgb(var(--border));
}

.card-content {
  padding: 1.5rem;
}
```

#### **Buttons**
```css
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  font-weight: 500;
  transition: all 0.2s;
  cursor: pointer;
}

.btn-primary {
  background: rgb(var(--primary));
  color: rgb(var(--primary-foreground));
  border: none;
}

.btn-secondary {
  background: rgb(var(--secondary));
  color: rgb(var(--secondary-foreground));
  border: none;
}

.btn-outline {
  border: 1px solid rgb(var(--border));
  background: transparent;
  color: rgb(var(--foreground));
}
```

#### **Badges**
```css
.badge {
  display: inline-flex;
  align-items: center;
  border-radius: 9999px;
  padding: 0.25rem 0.625rem;
  font-size: 0.75rem;
  font-weight: 600;
}

.badge-primary {
  background: rgb(var(--primary));
  color: rgb(var(--primary-foreground));
}

.badge-destructive {
  background: rgb(var(--destructive));
  color: rgb(var(--destructive-foreground));
}
```

## Required Implementation for Open-Telemorph-Prime

### 1. **Layout Structure Updates**

#### **Current Issues:**
- Single-column layout instead of sidebar + main
- No fixed header
- Missing proper navigation structure

#### **Required Changes:**

**A. Update HTML Structure:**
```html
<div class="app-layout">
  <header class="app-header">
    <!-- Header content -->
  </header>
  <div class="app-body">
    <aside class="app-sidebar">
      <!-- Sidebar navigation -->
    </aside>
    <main class="app-main">
      <!-- Main content -->
    </main>
  </div>
</div>
```

**B. CSS Layout Classes:**
```css
.app-layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: rgb(var(--background));
  color: rgb(var(--foreground));
}

.app-header {
  height: 4rem; /* 64px */
  border-bottom: 1px solid rgb(var(--border));
  background: rgb(var(--card));
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 1.5rem;
}

.app-body {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.app-sidebar {
  width: 16rem; /* 256px */
  border-right: 1px solid rgb(var(--border));
  background: rgb(var(--card));
  display: flex;
  flex-direction: column;
}

.app-main {
  flex: 1;
  overflow-y: auto;
  padding: 1.5rem;
}
```

### 2. **Sidebar Navigation**

#### **Required Features:**
- **Logo/Brand**: Activity icon + "Telemorph" text
- **Navigation Menu**: Dashboard, Metrics, Traces, Logs, Services, Alerts, Query Builder, Admin
- **Icons**: Lucide React icons for each menu item
- **Active States**: Highlighted current page
- **Version Info**: Footer with version number

#### **Navigation Structure:**
```html
<nav class="sidebar-nav">
  <div class="sidebar-header">
    <div class="brand">
      <Activity class="brand-icon" />
      <div class="brand-text">
        <h1>Telemorph</h1>
        <p>Observability Platform</p>
      </div>
    </div>
  </div>
  
  <div class="sidebar-menu">
    <a href="/" class="nav-item active">
      <LayoutDashboard class="nav-icon" />
      <span>Dashboard</span>
    </a>
    <a href="/metrics" class="nav-item">
      <BarChart3 class="nav-icon" />
      <span>Metrics</span>
    </a>
    <!-- ... more items -->
  </div>
  
  <div class="sidebar-footer">
    <div class="version-info">
      <div>Phase 4 Frontend</div>
      <div>v0.1.0</div>
    </div>
  </div>
</nav>
```

### 3. **Header Component**

#### **Required Features:**
- **Search Bar**: Global search with search icon
- **Theme Toggle**: Dark/light mode switch
- **Notifications**: Bell icon with badge
- **User Menu**: User icon/profile

#### **Header Structure:**
```html
<header class="app-header">
  <div class="header-left">
    <div class="search-container">
      <Search class="search-icon" />
      <input type="text" placeholder="Search metrics, traces, logs..." class="search-input" />
    </div>
  </div>
  
  <div class="header-right">
    <button class="theme-toggle">
      <Sun class="theme-icon" />
    </button>
    <button class="notifications">
      <Bell class="notification-icon" />
      <span class="notification-badge">3</span>
    </button>
    <button class="user-menu">
      <User class="user-icon" />
    </button>
  </div>
</header>
```

### 4. **Page Layouts**

#### **Dashboard Page:**
- **Page Header**: Title + description + action buttons
- **Metrics Grid**: 4-column responsive grid of metric cards
- **Two-Column Layout**: Recent alerts + service health
- **Quick Actions**: 4-button grid for common tasks

#### **Metrics Explorer:**
- **Page Header**: Title + time range + refresh buttons
- **Three-Column Layout**: Query editor (2 cols) + metrics browser (1 col)
- **Query Editor**: Input field + execute/export buttons
- **Chart Area**: Placeholder for visualizations
- **Metrics Browser**: Available metrics list
- **Query History**: Recent queries

#### **Traces Explorer:**
- **Page Header**: Title + time range + filter buttons
- **Three-Column Layout**: Search/filters (1 col) + trace list (2 cols)
- **Search Panel**: Query input + service filter + status filter
- **Trace List**: Trace cards with status indicators
- **Trace Details**: Flame graph + span details (when selected)

#### **Logs Viewer:**
- **Page Header**: Title + time range + filter buttons
- **Two-Column Layout**: Filters (1 col) + log entries (1 col)
- **Filter Panel**: Service + level + search filters
- **Log Entries**: Structured log display with levels

### 5. **Component Styling**

#### **Card Components:**
```css
.card {
  background: rgb(var(--card));
  color: rgb(var(--card-foreground));
  border: 1px solid rgb(var(--border));
  border-radius: 0.5rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.card-header {
  padding: 1.5rem;
  border-bottom: 1px solid rgb(var(--border));
}

.card-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.card-description {
  font-size: 0.875rem;
  color: rgb(var(--muted-foreground));
}

.card-content {
  padding: 1.5rem;
}
```

#### **Button Components:**
```css
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  font-weight: 500;
  padding: 0.5rem 1rem;
  transition: all 0.2s;
  cursor: pointer;
  border: none;
}

.btn-primary {
  background: rgb(var(--primary));
  color: rgb(var(--primary-foreground));
}

.btn-secondary {
  background: rgb(var(--secondary));
  color: rgb(var(--secondary-foreground));
}

.btn-outline {
  border: 1px solid rgb(var(--border));
  background: transparent;
  color: rgb(var(--foreground));
}

.btn-ghost {
  background: transparent;
  color: rgb(var(--foreground));
}

.btn-sm {
  height: 2.25rem;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
}
```

#### **Badge Components:**
```css
.badge {
  display: inline-flex;
  align-items: center;
  border-radius: 9999px;
  padding: 0.25rem 0.625rem;
  font-size: 0.75rem;
  font-weight: 600;
  border: 1px solid transparent;
}

.badge-default {
  background: rgb(var(--primary));
  color: rgb(var(--primary-foreground));
}

.badge-secondary {
  background: rgb(var(--secondary));
  color: rgb(var(--secondary-foreground));
}

.badge-destructive {
  background: rgb(var(--destructive));
  color: rgb(var(--destructive-foreground));
}

.badge-outline {
  border: 1px solid rgb(var(--border));
  background: transparent;
  color: rgb(var(--foreground));
}
```

### 6. **Responsive Design**

#### **Breakpoints:**
- **Mobile**: < 768px
- **Tablet**: 768px - 1024px
- **Desktop**: > 1024px

#### **Mobile Adaptations:**
- **Sidebar**: Collapsible/hidden on mobile
- **Header**: Compact layout
- **Grids**: Single column on mobile
- **Cards**: Full width on mobile

### 7. **Dark Mode Implementation**

#### **Theme Toggle:**
```javascript
// Theme toggle functionality
function toggleTheme() {
  const isDark = document.documentElement.classList.contains('dark');
  if (isDark) {
    document.documentElement.classList.remove('dark');
    localStorage.setItem('theme', 'light');
  } else {
    document.documentElement.classList.add('dark');
    localStorage.setItem('theme', 'dark');
  }
}

// Load saved theme
const savedTheme = localStorage.getItem('theme');
if (savedTheme === 'dark') {
  document.documentElement.classList.add('dark');
}
```

#### **CSS Custom Properties:**
```css
:root {
  /* Light theme variables */
  --background: 255, 255, 255;
  --foreground: 0, 0, 0;
  /* ... other light theme colors */
}

.dark {
  /* Dark theme variables */
  --background: 17, 24, 39;
  --foreground: 255, 255, 255;
  /* ... other dark theme colors */
}
```

### 8. **Icon System**

#### **Required Icons (Lucide React):**
- `Activity` - Brand/logo
- `LayoutDashboard` - Dashboard
- `BarChart3` - Metrics
- `GitBranch` - Traces
- `FileText` - Logs
- `Network` - Services
- `Bell` - Alerts
- `Search` - Query Builder
- `Settings` - Admin
- `Sun` - Light mode
- `Moon` - Dark mode
- `User` - User menu
- `Clock` - Time range
- `RefreshCw` - Refresh
- `Play` - Execute
- `Download` - Export
- `Filter` - Filters
- `AlertTriangle` - Alerts
- `CheckCircle` - Success
- `XCircle` - Error

### 9. **Implementation Priority**

#### **Phase 1: Core Layout**
1. ✅ Update HTML structure for sidebar + main layout
2. ✅ Implement CSS custom properties for theming
3. ✅ Create sidebar navigation component
4. ✅ Create header component
5. ✅ Update all page layouts

#### **Phase 2: Components**
1. ✅ Implement card components
2. ✅ Implement button components
3. ✅ Implement badge components
4. ✅ Add icon system
5. ✅ Implement form components

#### **Phase 3: Features**
1. ✅ Add dark mode toggle
2. ✅ Implement responsive design
3. ✅ Add search functionality
4. ✅ Add notification system
5. ✅ Polish animations and transitions

### 10. **File Structure Updates**

#### **Required Files:**
```
web/
├── index.html              # Updated with new layout
├── dashboard.html          # Updated with new layout
├── metrics.html            # Updated with new layout
├── traces.html             # Updated with new layout
├── logs.html               # Updated with new layout
└── static/
    ├── styles.css          # Complete design system
    ├── app.js              # Enhanced functionality
    └── icons.js            # Icon definitions
```

## Conclusion

This specification provides a comprehensive guide for implementing the telemorph-frontend design system in open-telemorph-prime. The key focus areas are:

1. **Layout Structure**: Sidebar + main content layout
2. **Design System**: Consistent colors, typography, and components
3. **Responsive Design**: Mobile-first approach
4. **Dark Mode**: Complete theme switching
5. **Component Library**: Reusable UI components
6. **Icon System**: Consistent iconography

Following this specification will ensure that open-telemorph-prime has the same professional, modern appearance as telemorph-frontend while maintaining its simplified architecture.
