# Open-Telemorph-Prime Test Scripts

This directory contains test scripts to verify that Open-Telemorph-Prime is working correctly with OTLP (OpenTelemetry Protocol) data ingestion.

## Available Scripts

### 1. `test_otel_simple.sh` (Recommended)
A simple test script that only requires `curl` and tests the HTTP OTLP endpoint (port 4318).

**Features:**
- Tests traces, metrics, and logs via HTTP
- Uses service name `foo.service`
- Provides colored output with success/error indicators
- No external dependencies beyond `curl`

**Usage:**
```bash
./test_otel_simple.sh
```

### 2. `test_otel.sh` (Comprehensive)
A comprehensive test script that tests both HTTP and gRPC endpoints.

**Features:**
- Tests HTTP OTLP endpoint (port 4318) with traces, metrics, and logs
- Tests gRPC endpoint (port 4317) connectivity
- Checks all service ports (4317, 4318, 8080)
- Provides detailed status reporting
- Uses service name `foo.service`

**Usage:**
```bash
./test_otel.sh
```

## Prerequisites

1. **Open-Telemorph-Prime must be running** on the default ports:
   - Port 4317: OTLP gRPC endpoint
   - Port 4318: OTLP HTTP endpoint  
   - Port 8080: Web UI

2. **Required tools:**
   - `curl` (for HTTP requests)
   - `nc` (netcat, for port checking)

## Test Data

Both scripts send sample OTLP data with the following characteristics:

- **Service Name:** `foo.service`
- **Traces:** Single span with trace ID `12345678901234567890123456789012`
- **Metrics:** Counter metric `test.counter` with value `42.0`
- **Logs:** INFO level log message "Test log message from foo.service"

## Expected Output

### Successful Test
```
✓ Open-Telemorph-Prime is running
✓ traces sent successfully
✓ metrics sent successfully  
✓ logs sent successfully
✓ All OTLP tests completed!
ℹ Service name: foo.service
ℹ Check the web UI at http://localhost:8080 to view the data
```

### Failed Test
```
✗ Open-Telemorph-Prime is not running on port 4318
```

## Troubleshooting

1. **"Open-Telemorph-Prime is not running"**
   - Start the application: `./open-telemorph-prime`
   - Check if ports are listening: `netstat -an | grep -E "(4317|4318|8080)"`

2. **"Failed to send [data type]"**
   - Check if the application is receiving data by looking at the logs
   - Verify the OTLP endpoint is enabled in the admin panel

3. **"gRPC server is not listening"**
   - Ensure gRPC is enabled in the configuration
   - Check the application logs for gRPC startup messages

## Viewing Test Data

After running the test scripts, you can view the ingested data in the Open-Telemorph-Prime web UI:

1. Open http://localhost:8080 in your browser
2. Navigate to the appropriate sections:
   - **Traces:** View the test trace with operation name "test-operation"
   - **Metrics:** View the test counter metric "test.counter"
   - **Logs:** View the test log message from "foo.service"
   - **Services:** See "foo.service" listed in the services

## Customization

You can modify the test scripts to:
- Change the service name by editing the `SERVICE_NAME` variable
- Add more test data (additional spans, metrics, logs)
- Test different OTLP endpoints or ports
- Add custom attributes or labels

## Notes

- The gRPC endpoint (4317) currently only tests connectivity since OTLP gRPC services are not yet fully implemented
- All test data uses valid OTLP JSON format
- The scripts are designed to be idempotent - you can run them multiple times safely
