# OTLP gRPC Implementation Documentation

## Overview

This document details the complete implementation of OpenTelemetry Protocol (OTLP) gRPC services in Open-Telemorph-Prime. The implementation provides full support for receiving telemetry data (traces, metrics, and logs) via the standard OTLP gRPC protocol.

## Implementation Summary

### ✅ Completed Features

1. **Complete OTLP gRPC Protocol Support** - All three OTLP services are fully implemented
2. **Protobuf Service Registration** - Services are properly registered using official OTLP protobuf definitions
3. **Data Conversion** - Proper conversion from OTLP protobuf format to internal storage format
4. **Error Handling** - Comprehensive error handling and logging
5. **gRPC Reflection** - Enabled for debugging and service discovery
6. **Backward Compatibility** - HTTP OTLP endpoints continue to work alongside gRPC

### 🏗️ Architecture

The implementation follows a modular architecture with clear separation of concerns:

```
internal/grpc/
├── server.go      # gRPC server setup and service registration
├── traces.go      # Trace service implementation
├── metrics.go     # Metrics service implementation
└── logs.go        # Logs service implementation
```

### 🔧 Technical Implementation

#### Dependencies Added

- `go.opentelemetry.io/proto/otlp v1.2.0` - Official OTLP protobuf definitions
- `google.golang.org/grpc v1.76.0` - gRPC framework (already present)

#### Service Registration

All OTLP services are registered with the gRPC server:

```go
// Services registered:
coltracepb.RegisterTraceServiceServer(grpcServer, traceService)
colmetricspb.RegisterMetricsServiceServer(grpcServer, metricsService)
collogspb.RegisterLogsServiceServer(grpcServer, logsService)
```

#### Port Configuration

- **gRPC Server**: Port 4317 (standard OTLP gRPC port)
- **HTTP Server**: Port 8080 (main web interface)
- **OTLP HTTP**: Port 4318 (HTTP-based OTLP ingestion)

### 📊 Service Details

#### 1. Trace Service (`internal/grpc/traces.go`)

**Purpose**: Handles distributed tracing data ingestion

**Key Features**:
- Processes `ResourceSpans` from OTLP trace requests
- Extracts service names from resource attributes
- Converts protobuf spans to internal trace format
- Handles parent-child span relationships
- Supports trace and span ID conversion
- Comprehensive attribute processing

**Data Flow**:
```
OTLP Trace Request → ResourceSpans → ScopeSpans → Spans → Internal Trace Format
```

#### 2. Metrics Service (`internal/grpc/metrics.go`)

**Purpose**: Handles metrics data ingestion

**Supported Metric Types**:
- **Gauge Metrics** - Point-in-time measurements
- **Sum Metrics** - Cumulative measurements
- **Histogram Metrics** - Distribution measurements with buckets
- **Exponential Histogram Metrics** - Exponential distribution measurements
- **Summary Metrics** - Statistical summaries with quantiles

**Key Features**:
- Processes all OTLP metric types
- Converts numeric values (double/int) appropriately
- Handles metric attributes and labels
- Supports bucket and quantile processing for histograms/summaries
- Service name extraction from resource attributes

**Data Flow**:
```
OTLP Metrics Request → ResourceMetrics → ScopeMetrics → Metrics → Internal Metric Format
```

#### 3. Logs Service (`internal/grpc/logs.go`)

**Purpose**: Handles log data ingestion

**Key Features**:
- Processes `ResourceLogs` from OTLP log requests
- Extracts log body content from various value types
- Handles trace and span ID associations
- Converts severity levels appropriately
- Comprehensive attribute processing
- Service name extraction from resource attributes

**Data Flow**:
```
OTLP Logs Request → ResourceLogs → ScopeLogs → LogRecords → Internal Log Format
```

### 🚀 Server Configuration

#### gRPC Server Options

```go
grpcServer := grpc.NewServer(
    grpc.MaxRecvMsgSize(4*1024*1024), // 4MB max message size
    grpc.MaxSendMsgSize(4*1024*1024),  // 4MB max message size
)
```

#### Service Registration

```go
func NewServer(storage storage.Storage, port int) *Server {
    // Create gRPC server with options
    grpcServer := grpc.NewServer(...)
    
    // Create service instances
    traceService := NewTraceService(storage)
    metricsService := NewMetricsService(storage)
    logsService := NewLogsService(storage)
    
    // Register services
    coltracepb.RegisterTraceServiceServer(grpcServer, traceService)
    colmetricspb.RegisterMetricsServiceServer(grpcServer, metricsService)
    collogspb.RegisterLogsServiceServer(grpcServer, logsService)
    
    // Enable gRPC reflection
    reflection.Register(grpcServer)
    
    return &Server{...}
}
```

### 🔍 Error Handling

Each service implements comprehensive error handling:

- **Input Validation**: Request validation and error responses
- **Data Processing**: Graceful handling of malformed data
- **Storage Errors**: Proper error logging for storage failures
- **Partial Success**: Processing continues even if individual items fail

### 📈 Performance Considerations

- **Message Size Limits**: 4MB maximum message size for large telemetry data
- **Concurrent Processing**: Services handle multiple requests concurrently
- **Memory Efficiency**: Proper cleanup of protobuf objects
- **Graceful Shutdown**: Clean resource cleanup on server shutdown

### 🧪 Testing and Validation

#### Server Startup Verification

When the server starts successfully, you should see:

```
2025/10/16 22:33:42 Starting OTLP gRPC server on port 4317
2025/10/16 22:33:42 Registered services:
2025/10/16 22:33:42   - TraceService (opentelemetry.proto.collector.trace.v1.TraceService)
2025/10/16 22:33:42   - MetricsService (opentelemetry.proto.collector.metrics.v1.MetricsService)
2025/10/16 22:33:42   - LogsService (opentelemetry.proto.collector.logs.v1.LogsService)
2025/10/16 22:33:42   - gRPC Reflection enabled
```

#### Client Connection

Clients can connect to the gRPC server at:
- **Endpoint**: `localhost:4317`
- **Protocol**: gRPC
- **Services**: All three OTLP services available

### 🔧 Configuration

The gRPC server configuration is managed through the existing configuration system:

```yaml
ingestion:
  grpc_port: 4317
  grpc_enabled: true
  batch_size: 1000
  flush_interval: "5s"
```

### 📚 API Reference

#### Trace Service API

```protobuf
service TraceService {
  rpc Export(ExportTraceServiceRequest) returns (ExportTraceServiceResponse);
}
```

#### Metrics Service API

```protobuf
service MetricsService {
  rpc Export(ExportMetricsServiceRequest) returns (ExportMetricsServiceResponse);
}
```

#### Logs Service API

```protobuf
service LogsService {
  rpc Export(ExportLogsServiceRequest) returns (ExportLogsServiceResponse);
}
```

### 🚀 Usage Examples

#### Connecting with OpenTelemetry SDK

```go
// Example client configuration
exporter, err := otlptracegrpc.New(
    context.Background(),
    otlptracegrpc.WithEndpoint("localhost:4317"),
    otlptracegrpc.WithInsecure(),
)
```

#### gRPC Reflection

The server supports gRPC reflection, allowing tools like `grpcurl` to discover services:

```bash
# List all services
grpcurl -plaintext localhost:4317 list

# List methods for a specific service
grpcurl -plaintext localhost:4317 list opentelemetry.proto.collector.trace.v1.TraceService
```

### 🔄 Backward Compatibility

The implementation maintains full backward compatibility:

- **HTTP OTLP endpoints** continue to work on port 4318
- **Web interface** remains available on port 8080
- **Existing API endpoints** are unchanged
- **Configuration format** remains the same

### 📋 Future Enhancements

Potential areas for future improvement:

1. **Authentication/Authorization** - Add security mechanisms
2. **Rate Limiting** - Implement request rate limiting
3. **Metrics Collection** - Add server-side metrics
4. **Health Checks** - Enhanced health check endpoints
5. **TLS Support** - Add secure gRPC connections
6. **Load Balancing** - Support for multiple server instances

### 🐛 Troubleshooting

#### Common Issues

1. **Port Already in Use**: Ensure no other services are using ports 4317, 4318, or 8080
2. **Permission Denied**: Check file permissions for the database and log files
3. **Memory Issues**: Monitor memory usage with large telemetry volumes

#### Debugging

- Enable gRPC reflection for service discovery
- Check server logs for detailed error messages
- Use `grpcurl` for manual testing of gRPC services

### 📖 References

- [OpenTelemetry Protocol Specification](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/otlp.md)
- [OTLP gRPC Exporter](https://opentelemetry.io/docs/languages/go/exporters/#otlp-grpc-exporter)
- [gRPC Go Documentation](https://grpc.io/docs/languages/go/)

---

**Implementation Date**: October 16, 2025  
**Version**: 0.2.1  
**Status**: Complete and Production Ready
