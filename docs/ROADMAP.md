# Open-Telemorph-Prime - Simplified Observability Platform

## Executive Summary

A simplified, single-binary observability platform designed for home users and developers. Open-Telemorph-Prime eliminates the complexity of enterprise observability platforms while maintaining core functionality for ingesting, storing, and querying OpenTelemetry signals (traces, metrics, logs).

---

## Philosophy: Simplicity First

Open-Telemorph-Prime prioritizes:
- **Single Binary Deployment** - One executable, zero configuration
- **Minimal Resource Usage** - Runs on any modern machine (<2GB RAM)
- **Progressive Complexity** - Start simple, scale up when needed
- **Developer-Friendly** - Easy to understand, modify, and extend

---

## Architecture

### Simplified System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    Data Sources                                  │
│  Applications with OTEL SDKs → Direct Ingestion                 │
└────────────────────┬────────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│              Unified Service (Single Binary)                    │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │ gRPC/HTTP    │  │   Storage    │  │   Query      │          │
│  │ Receivers    │  │   Engine     │  │   Engine     │          │
│  │ (OTLP)       │  │ (SQLite/DB)  │  │ (Built-in)   │          │
│  └──────────────┘  └──────────────┘  └──────────────┘          │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │   Web UI     │  │   REST API   │  │   Health     │          │
│  │  (Embedded)  │  │  Endpoints   │  │   Checks     │          │
│  └──────────────┘  └──────────────┘  └──────────────┘          │
└────────────────────┬────────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│              Storage Backend (Configurable)                     │
│  • SQLite (default) - for single user                          │
│  • PostgreSQL - for multi-user                                 │
│  • File-based - for minimal setup                              │
└─────────────────────────────────────────────────────────────────┘
```

### Data Flow (Simplified)

```
1. Application → OTEL SDK → Direct HTTP/gRPC
                                ↓
2. Unified Service → Parse & Validate → Storage
                                ↓
3. Query Engine → Filter & Aggregate → Web UI/API
```

---

## Technology Stack

### Core Platform
- **Language**: Go 1.21+
- **OTLP Handling**: `go.opentelemetry.io/collector` (receivers only)
- **Storage**: SQLite (default), PostgreSQL (optional)
- **Web UI**: Embedded HTML/CSS/JavaScript (no build step)
- **Query Engine**: Built-in SQL with basic PromQL support

### Optional Advanced Features
- **Message Queue**: Apache Kafka (when enabled)
- **Stream Processing**: Apache Flink (when enabled)
- **Advanced Storage**: Druid + Hudi (when enabled)
- **Distributed Query**: Trino (when enabled)

---

## Project Phases

### Phase 1: Minimal Ingestion and Storage (Weeks 1-3) ✅ COMPLETED
**Goal**: Create a single binary that can ingest and store telemetry data

#### Deliverables:
- ✅ Single Go binary with embedded web UI
- ✅ OTLP gRPC and HTTP receivers (ports 4317/4318)
- ✅ SQLite storage with basic schema (using modernc.org/sqlite)
- ✅ Simple REST API for data retrieval
- ✅ Basic web interface for data exploration
- ✅ Health check endpoints
- ✅ Configuration file support
- ✅ Docker support with parametric Go versions
- ✅ Dogfood feature for self-monitoring
- ✅ Dynamic system metrics in Admin UI
- ✅ Services page with real-time data

#### Architecture:
```
open-telemorph-prime
├── main.go                 # Single entry point
├── internal/
│   ├── ingestion/         # OTLP receivers
│   ├── storage/           # SQLite/PostgreSQL interface
│   ├── query/             # Basic query engine
│   ├── web/               # Embedded web UI
│   └── config/            # Configuration management
├── web/                   # Static web assets
│   ├── index.html
│   ├── dashboard.js
│   └── styles.css
└── config.yaml            # Configuration file
```

#### Storage Schema (SQLite):
```sql
-- Metrics table
CREATE TABLE metrics (
    id INTEGER PRIMARY KEY,
    timestamp INTEGER NOT NULL,
    metric_name TEXT NOT NULL,
    value REAL NOT NULL,
    labels TEXT, -- JSON
    service_name TEXT,
    created_at INTEGER DEFAULT (strftime('%s', 'now'))
);

-- Traces table
CREATE TABLE traces (
    id INTEGER PRIMARY KEY,
    trace_id TEXT NOT NULL,
    span_id TEXT NOT NULL,
    parent_span_id TEXT,
    service_name TEXT,
    operation_name TEXT,
    start_time INTEGER NOT NULL,
    duration_nanos INTEGER NOT NULL,
    attributes TEXT, -- JSON
    status_code TEXT,
    created_at INTEGER DEFAULT (strftime('%s', 'now'))
);

-- Logs table
CREATE TABLE logs (
    id INTEGER PRIMARY KEY,
    timestamp INTEGER NOT NULL,
    service_name TEXT,
    level TEXT,
    message TEXT,
    attributes TEXT, -- JSON
    trace_id TEXT,
    span_id TEXT,
    created_at INTEGER DEFAULT (strftime('%s', 'now'))
);

-- Indexes for performance
CREATE INDEX idx_metrics_timestamp ON metrics(timestamp);
CREATE INDEX idx_metrics_service ON metrics(service_name);
CREATE INDEX idx_traces_trace_id ON traces(trace_id);
CREATE INDEX idx_traces_service ON traces(service_name);
CREATE INDEX idx_logs_timestamp ON logs(timestamp);
CREATE INDEX idx_logs_service ON logs(service_name);
```

#### API Endpoints:
```
GET  /health                    # Health check
GET  /ready                     # Readiness check
GET  /api/v1/metrics            # List metrics
GET  /api/v1/traces             # List traces
GET  /api/v1/logs               # List logs
GET  /api/v1/services           # List services
POST /api/v1/query              # Generic query endpoint
GET  /                          # Web UI
```

#### Configuration:
```yaml
# config.yaml
server:
  port: 8080
  grpc_port: 4317
  http_port: 4318

storage:
  type: "sqlite"  # sqlite, postgres, file
  path: "./data/telemorph.db"
  retention_days: 30

query:
  max_results: 10000
  timeout_seconds: 30

web:
  enabled: true
  title: "Open-Telemorph-Prime"

logging:
  level: "info"
  format: "json"
```

#### Testing:
- Sample applications generating telemetry
- Load testing with basic tools
- Web UI functionality verification

---

### Phase 2: Inline Query API (Weeks 4-5) 🚧 IN PROGRESS
**Goal**: Add comprehensive querying capabilities

#### Deliverables:
- 🚧 Basic PromQL support for metrics
- 📋 Simple log query language
- 📋 Trace filtering and search
- 📋 Query result caching
- 📋 Export functionality (JSON, CSV)

#### Implementation Status:
- 🚧 **PromQL Parser**: Starting implementation
- 📋 **Log Query Language**: Planned
- 📋 **Trace Filtering**: Planned
- 📋 **Query Caching**: Planned
- 📋 **Export Functions**: Planned

#### Query Language Support:

**Metrics (PromQL-like)**:
```promql
# Simple metric query
http_requests_total

# Rate calculation
rate(http_requests_total[5m])

# Filtering
http_requests_total{service="api"}

# Aggregation
sum(http_requests_total) by (service)
```

**Logs (Simple Query Language)**:
```sql
-- Service filtering
service:api-gateway

-- Level filtering
level:ERROR

-- Text search
message:error

-- Combined
service:api-gateway level:ERROR message:timeout
```

**Traces (Simple Filtering)**:
```sql
-- Service filtering
service:api-gateway

-- Duration filtering
duration:>1s

-- Status filtering
status:error

-- Combined
service:api-gateway duration:>1s status:error
```

#### Query API Enhancements:
```
POST /api/v1/query/metrics      # PromQL queries
POST /api/v1/query/logs         # Log queries
POST /api/v1/query/traces       # Trace queries
GET  /api/v1/query/export       # Export results
```

---

### Phase 3: Basic Frontend (Weeks 6-7)
**Goal**: Create a simple but functional web interface

#### Deliverables:
- Dashboard with key metrics
- Metrics explorer with basic charts
- Logs viewer with filtering
- Traces explorer with timeline view
- Service overview page
- Basic query builder

#### Web UI Features:

**Dashboard**:
- System overview metrics
- Recent errors and warnings
- Service health status
- Quick query interface

**Metrics Explorer**:
- Time-series line charts
- Metric selection dropdown
- Time range picker
- Basic aggregation controls

**Logs Viewer**:
- Log stream with auto-refresh
- Filter by service, level, time
- Search functionality
- Log detail view

**Traces Explorer**:
- Trace list with filters
- Timeline view for spans
- Trace detail panel
- Service dependency hints

**Query Builder**:
- Form-based query construction
- Query history
- Saved queries
- Export options

#### UI Technology:
- Vanilla HTML/CSS/JavaScript (no build step)
- Chart.js for visualizations
- Simple CSS framework (Tailwind-like)
- Responsive design

---

### Phase 4: Optional Advanced Features (Weeks 8-10)
**Goal**: Enable enterprise features for power users

#### Deliverables:
- Configuration-driven complexity
- Kafka integration (optional)
- PostgreSQL support
- Basic alerting
- Data retention policies
- Performance optimizations

#### Advanced Configuration:
```yaml
# config.yaml - Advanced mode
mode: "advanced"  # simple, advanced

storage:
  type: "postgres"
  host: "localhost"
  port: 5432
  database: "telemorph"
  username: "telemorph"
  password: "password"

messaging:
  kafka:
    enabled: true
    brokers: ["localhost:9092"]
    topics:
      metrics: "otel.metrics"
      traces: "otel.traces"
      logs: "otel.logs"

stream_processing:
  flink:
    enabled: true
    jobmanager: "localhost:8081"

advanced_storage:
  druid:
    enabled: true
    broker_url: "http://localhost:8082"
  hudi:
    enabled: true
    warehouse_path: "s3://telemorph-data"

alerting:
  enabled: true
  rules:
    - name: "High Error Rate"
      query: "rate(http_requests_total{status=~'5..'}[5m]) > 0.05"
      duration: "2m"
      severity: "warning"
```

#### Advanced Features:
- **Kafka Mode**: Enable for high-throughput scenarios
- **PostgreSQL**: Multi-user support with better performance
- **Alerting**: Basic rule-based alerting
- **Data Retention**: Automatic cleanup of old data
- **Performance**: Query optimization and caching

---

## Deployment Options

### Single Binary (Default)
```bash
# Download and run
wget https://github.com/your-org/open-telemorph-prime/releases/latest/download/open-telemorph-prime-linux-amd64
chmod +x open-telemorph-prime-linux-amd64
./open-telemorph-prime-linux-amd64

# Or with configuration
./open-telemorph-prime-linux-amd64 -config config.yaml
```

### Docker Compose (Simple)
```yaml
# docker-compose.yml
version: '3.8'
services:
  open-telemorph-prime:
    image: open-telemorph-prime:latest
    ports:
      - "8080:8080"
      - "4317:4317"
      - "4318:4318"
    volumes:
      - ./data:/app/data
      - ./config.yaml:/app/config.yaml
    environment:
      - CONFIG_PATH=/app/config.yaml
```

### Docker Compose (Advanced)
```yaml
# docker-compose.advanced.yml
version: '3.8'
services:
  open-telemorph-prime:
    image: open-telemorph-prime:latest
    ports:
      - "8080:8080"
    environment:
      - MODE=advanced
    depends_on:
      - postgres
      - kafka

  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: telemorph
      POSTGRES_USER: telemorph
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data

  kafka:
    image: confluentinc/cp-kafka:7.5.0
    environment:
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
    depends_on:
      - zookeeper

  zookeeper:
    image: confluentinc/cp-zookeeper:7.5.0
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
```

---

## Performance Targets

### Simple Mode:
- **Ingestion**: 1K+ spans/sec, 10K+ metrics/sec
- **Query Latency**: <100ms for simple queries
- **Storage**: <1GB for 30 days of data
- **Memory Usage**: <512MB RAM
- **Startup Time**: <5 seconds

### Advanced Mode:
- **Ingestion**: 10K+ spans/sec, 100K+ metrics/sec
- **Query Latency**: <500ms for complex queries
- **Storage**: <10GB for 30 days of data
- **Memory Usage**: <2GB RAM
- **Startup Time**: <30 seconds

---

## Resource Requirements

### Minimum (Simple Mode):
- **CPU**: 1 core
- **RAM**: 512MB
- **Storage**: 1GB
- **OS**: Linux, macOS, Windows

### Recommended (Advanced Mode):
- **CPU**: 2 cores
- **RAM**: 2GB
- **Storage**: 10GB
- **OS**: Linux (Docker)

---

## Success Metrics

### User Experience:
- **Setup Time**: <5 minutes from download to running
- **Learning Curve**: <1 hour to understand basic features
- **Resource Usage**: Runs on any modern machine
- **Reliability**: 99%+ uptime for single-user scenarios

### Technical:
- **Code Complexity**: <10K lines of Go code
- **Dependencies**: <20 external packages
- **Binary Size**: <50MB
- **Build Time**: <2 minutes

---

## Comparison with Telemorph-Prime

| Feature | Telemorph-Prime | Open-Telemorph-Prime |
|---------|-----------------|----------------------|
| **Complexity** | High (20+ services) | Low (1-3 services) |
| **Resource Usage** | 1TB+ RAM, 50+ nodes | <2GB RAM, 1 machine |
| **Setup Time** | Hours/Days | Minutes |
| **Learning Curve** | Weeks | Hours |
| **Target Users** | Enterprise | Home/Dev |
| **Scalability** | 1000+ services | 10-100 services |
| **Features** | Full enterprise | Essential features |

---

## Development Timeline

### Week 1-2: Core Platform
- Basic Go binary structure
- OTLP receivers (gRPC/HTTP)
- SQLite storage implementation
- Simple REST API

### Week 3: Web Interface
- Embedded web UI
- Basic dashboard
- Configuration management

### Week 4-5: Query Engine
- PromQL-like query support
- Log and trace filtering
- Export functionality

### Week 6-7: Frontend Polish
- Complete web interface
- Query builder
- Visualization improvements

### Week 8-10: Advanced Features
- Optional Kafka integration
- PostgreSQL support
- Basic alerting
- Performance optimizations

---

## Getting Started

### Quick Start (5 minutes):
```bash
# 1. Download binary
curl -L https://github.com/your-org/open-telemorph-prime/releases/latest/download/open-telemorph-prime-linux-amd64 -o open-telemorph-prime
chmod +x open-telemorph-prime

# 2. Run
./open-telemorph-prime

# 3. Open browser
open http://localhost:8080
```

### With Docker:
```bash
# 1. Create config
cat > config.yaml << EOF
server:
  port: 8080
storage:
  type: sqlite
  path: ./data/telemorph.db
EOF

# 2. Run with Docker
docker run -p 8080:8080 -v $(pwd)/data:/app/data -v $(pwd)/config.yaml:/app/config.yaml open-telemorph-prime:latest
```

---

## Contributing

Open-Telemorph-Prime is designed to be simple and approachable for contributors:

1. **Fork the repository**
2. **Create a feature branch**
3. **Make your changes** (keep it simple!)
4. **Add tests** (if applicable)
5. **Submit a pull request**

### Development Setup:
```bash
# Clone repository
git clone https://github.com/your-org/open-telemorph-prime.git
cd open-telemorph-prime

# Install dependencies
go mod tidy

# Run in development mode
go run main.go -config config.yaml

# Build binary
go build -o open-telemorph-prime .
```

---

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

---

## Roadmap Status

- ✅ **Phase 1**: Minimal Ingestion and Storage (COMPLETED)
- 🚧 **Phase 2**: Inline Query API (IN PROGRESS)
- 📋 **Phase 3**: Basic Frontend (Planned)
- 📋 **Phase 4**: Optional Advanced Features (Planned)

### Current Focus: Phase 2 - Query API Implementation

**Immediate Next Steps:**
1. **PromQL Parser Implementation** - Core metrics query language
2. **Log Query Language** - Service, level, and text filtering
3. **Trace Filtering** - Duration, status, and service filtering
4. **Query API Endpoints** - RESTful query interfaces
5. **Export Functionality** - JSON/CSV data export

**Files to Create:**
```
internal/query/
├── promql/
│   ├── parser.go          # PromQL query parser
│   ├── evaluator.go       # Query evaluation
│   └── functions.go       # Rate, sum, avg functions
├── logs/
│   ├── parser.go          # Log query parser
│   └── evaluator.go       # Log query evaluation
├── traces/
│   ├── parser.go          # Trace query parser
│   └── evaluator.go       # Trace query evaluation
└── export/
    ├── json.go            # JSON export
    └── csv.go             # CSV export
```

---

**Open-Telemorph-Prime**: Observability made simple. 🚀

*Perfect for home labs, development environments, and anyone who wants observability without the complexity.*
