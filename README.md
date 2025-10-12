# Open-Telemorph-Prime

A simplified, single-binary observability platform designed for home users and developers. Open-Telemorph-Prime eliminates the complexity of enterprise observability platforms while maintaining core functionality for ingesting, storing, and querying OpenTelemetry signals (traces, metrics, logs).

## 🚀 Quick Start

### Option 1: Docker Compose (Recommended)

```bash
# Clone the repository
git clone https://github.com/your-org/open-telemorph-prime.git
cd open-telemorph-prime

# Start the service
docker-compose up -d

# Open your browser
open http://localhost:8080
```

### Option 2: Direct Binary

```bash
# Download the latest release
wget https://github.com/your-org/open-telemorph-prime/releases/latest/download/open-telemorph-prime-linux-amd64
chmod +x open-telemorph-prime-linux-amd64

# Run
./open-telemorph-prime-linux-amd64

# Open your browser
open http://localhost:8080
```

### Option 3: Build from Source

```bash
# Prerequisites: Go 1.21+
git clone https://github.com/your-org/open-telemorph-prime.git
cd open-telemorph-prime

# Install dependencies
go mod tidy

# Build
go build -o open-telemorph-prime .

# Run
./open-telemorph-prime
```

## 📊 Features

- **Single Binary**: One executable, zero configuration
- **Minimal Resource Usage**: Runs on any modern machine (<2GB RAM)
- **OTLP Support**: Ingest traces, metrics, and logs via HTTP/gRPC
- **Web UI**: Simple, responsive interface for data exploration
- **SQLite Storage**: Lightweight, file-based storage
- **REST API**: Query your data programmatically
- **Health Checks**: Built-in monitoring endpoints

## 🔧 Configuration

Open-Telemorph-Prime uses a simple YAML configuration file:

```yaml
server:
  port: 8080
  environment: "development"

storage:
  type: "sqlite"
  path: "./data/telemorph.db"
  retention_days: 30

ingestion:
  grpc_port: 4317
  http_port: 4318

web:
  enabled: true
  title: "Open-Telemorph-Prime"
```

## 📡 Sending Data

### HTTP Endpoint

```bash
# Send traces
curl -X POST http://localhost:8080/v1/traces \
  -H "Content-Type: application/json" \
  -d '{"resourceSpans": [...]}'

# Send metrics
curl -X POST http://localhost:8080/v1/metrics \
  -H "Content-Type: application/json" \
  -d '{"resourceMetrics": [...]}'

# Send logs
curl -X POST http://localhost:8080/v1/logs \
  -H "Content-Type: application/json" \
  -d '{"resourceLogs": [...]}'
```

### OpenTelemetry SDK Integration

```go
// Go example
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/http"
)

exporter, err := otlptracehttp.New(
    context.Background(),
    otlptracehttp.WithEndpoint("http://localhost:8080"),
    otlptracehttp.WithInsecure(),
)
```

## 🔍 API Endpoints

### Health
- `GET /health` - Health check
- `GET /ready` - Readiness check

### Data
- `GET /api/v1/metrics` - List metrics
- `GET /api/v1/traces` - List traces
- `GET /api/v1/logs` - List logs
- `GET /api/v1/services` - List services
- `POST /api/v1/query` - Generic query endpoint

### Web UI
- `GET /` - Home page
- `GET /dashboard` - Dashboard
- `GET /metrics` - Metrics explorer
- `GET /traces` - Traces explorer
- `GET /logs` - Logs viewer

## 🏗️ Architecture

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
└─────────────────────────────────────────────────────────────────┘
```

## 🛠️ Development

### Prerequisites
- Go 1.21+
- Docker (optional)
- SQLite3 (for local development)

### Building

```bash
# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build binary
go build -o open-telemorph-prime .

# Run in development mode
go run main.go -config config.yaml
```

### Project Structure

```
open-telemorph-prime/
├── main.go                 # Entry point
├── internal/
│   ├── config/            # Configuration management
│   ├── ingestion/         # OTLP receivers
│   ├── storage/           # SQLite storage
│   └── web/               # Web UI and API
├── web/                   # Static web assets
│   ├── index.html
│   └── static/
├── config.yaml            # Default configuration
├── Dockerfile
└── docker-compose.yml
```

## 📈 Performance

### Simple Mode (Default)
- **Ingestion**: 1K+ spans/sec, 10K+ metrics/sec
- **Query Latency**: <100ms for simple queries
- **Storage**: <1GB for 30 days of data
- **Memory Usage**: <512MB RAM
- **Startup Time**: <5 seconds

## 🔒 Security

- CORS enabled for cross-origin requests
- Input validation on all endpoints
- SQL injection protection via prepared statements
- Rate limiting (configurable)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [OpenTelemetry](https://opentelemetry.io/) for the telemetry standards
- [Gin](https://gin-gonic.com/) for the web framework
- [SQLite](https://sqlite.org/) for the embedded database

---

**Open-Telemorph-Prime**: Observability made simple. 🚀

*Perfect for home labs, development environments, and anyone who wants observability without the complexity.*