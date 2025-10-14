# Release Notes v0.2.1

**Release Date:** October 14, 2025  
**Release Type:** Pre-release  
**Status:** Ready for Testing

## 🚀 What's New

### OTLP Ingestion Improvements
- **Fixed OTLP Port Configuration**: Now uses standard OpenTelemetry Collector ports
  - **Port 4317**: OTLP gRPC endpoint for traces, metrics, and logs
  - **Port 4318**: OTLP HTTP endpoint for traces, metrics, and logs
  - **Port 8080**: Web UI and REST API (unchanged)
- **Separate HTTP/gRPC Servers**: Implemented dedicated servers for each OTLP protocol
- **Configurable Endpoints**: Added enable/disable checkboxes for each OTLP port in admin panel

### User Interface Enhancements
- **Dynamic Version Display**: Version now updates automatically across all pages
- **Client-Side Filtering**: Implemented filtering for traces, logs, and metrics pages
- **Dynamic Service Loading**: Service dropdowns now populate from actual API data
- **Consistent Page Layout**: Standardized metrics page to match traces and logs format
- **Fixed HTML Structure**: Resolved table rendering issues in data display pages

### Admin Panel Updates
- **Separate Endpoint Configuration**: 
  - HTTP Endpoint: `0.0.0.0:4318` (configurable)
  - gRPC Endpoint: `0.0.0.0:4317` (configurable)
- **Enable/Disable Controls**: Individual checkboxes for each OTLP port
- **Real-time Configuration**: Changes take effect immediately

### Docker & Deployment
- **Optimized Dockerfile**: Reduced build context with `.dockerignore`
- **Improved Health Checks**: Added wget dependency for reliable health monitoring
- **Security Enhancements**: Config file mounted as read-only
- **Better Documentation**: Updated port mappings and usage examples

### Documentation & Organization
- **Organized Documentation**: Moved all markdown files to `docs/` folder
- **Updated README**: Corrected port information and usage examples
- **Test Scripts**: Added comprehensive OTLP endpoint testing tools
- **Clean Project Structure**: Removed temporary files and build artifacts

## 🐛 Bug Fixes

- **UI Data Display**: Fixed traces and logs not appearing in web interface
- **JavaScript Execution**: Resolved timing issues with data loading
- **HTML Structure**: Fixed invalid table markup causing display errors
- **Port Listening**: Ensured both OTLP ports (4317/4318) are properly listening
- **Filter Functionality**: Fixed "Apply Filters" button not working on data pages

## 🔧 Technical Improvements

- **Backend Architecture**: Separated OTLP ingestion from main web server
- **Template System**: Added dynamic version passing to all HTML templates
- **API Consistency**: Standardized response formats across all endpoints
- **Error Handling**: Improved error messages and debugging information
- **Code Organization**: Better separation of concerns between services

## 📋 Testing

### New Test Scripts
- `test_otel_simple.sh`: Simple HTTP endpoint testing
- `test_otel.sh`: Comprehensive OTLP endpoint validation

### Test Coverage
- ✅ OTLP HTTP endpoint (port 4318)
- ✅ OTLP gRPC endpoint (port 4317) 
- ✅ Web UI functionality
- ✅ Admin panel configuration
- ✅ Data filtering and display
- ✅ Docker containerization

## 🚨 Breaking Changes

- **Port Changes**: OTLP ingestion moved from port 8080 to dedicated ports 4317/4318
- **Configuration**: New `grpc_enabled` and `http_enabled` flags in config.yaml
- **File Structure**: Documentation moved to `docs/` folder

## 📦 Installation & Upgrade

### New Installation
```bash
# Clone and build
git clone <repository-url>
cd open-telemorph-prime
go build -o open-telemorph-prime .

# Run with default configuration
./open-telemorph-prime
```

### Docker Deployment
```bash
# Build and run with docker-compose
docker-compose up -d

# Or build directly
docker build -t open-telemorph-prime .
```

### Configuration Update
If upgrading from v0.2.0, update your `config.yaml`:
```yaml
ingestion:
  grpc_port: 4317
  http_port: 4318
  grpc_enabled: true    # New
  http_enabled: true    # New
```

## 🔗 Endpoints

### OTLP Ingestion
- **HTTP**: `http://localhost:4318/v1/{traces,metrics,logs}`
- **gRPC**: `localhost:4317` (traces, metrics, logs)

### Web Interface
- **Main UI**: `http://localhost:8080`
- **API**: `http://localhost:8080/api/v1/{metrics,traces,logs,services}`

## 🎯 Next Steps

- **Full OTLP gRPC Implementation**: Complete protobuf service registration
- **Server-Side Filtering**: Implement backend filtering for better performance
- **Time Range Filtering**: Add date/time range selection for logs
- **Export Functionality**: Implement data export features
- **Performance Optimization**: Add caching and query optimization

## 📞 Support

For issues, questions, or contributions:
- **Issues**: GitHub Issues
- **Documentation**: `docs/` folder
- **Testing**: Use provided test scripts

---

**Note**: This is a pre-release version. Please test thoroughly before using in production environments.
