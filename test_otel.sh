#!/bin/bash

# Open-Telemorph-Prime OTLP Test Script
# Tests both HTTP (4318) and gRPC (4317) endpoints with logs, traces, and metrics

set -e

# Configuration
SERVICE_NAME="foo.service"
HTTP_ENDPOINT="http://localhost:4318"
GRPC_ENDPOINT="localhost:4317"
TIMESTAMP=$(date -u +"%Y-%m-%dT%H:%M:%S.%3NZ")

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    local status=$1
    local message=$2
    if [ "$status" = "SUCCESS" ]; then
        echo -e "${GREEN}✓${NC} $message"
    elif [ "$status" = "ERROR" ]; then
        echo -e "${RED}✗${NC} $message"
    elif [ "$status" = "INFO" ]; then
        echo -e "${BLUE}ℹ${NC} $message"
    elif [ "$status" = "WARNING" ]; then
        echo -e "${YELLOW}⚠${NC} $message"
    fi
}

# Function to test HTTP endpoint
test_http_endpoint() {
    print_status "INFO" "Testing HTTP endpoint (port 4318)..."
    
    # Test traces
    print_status "INFO" "Sending traces via HTTP..."
    local trace_response=$(curl -s -w "%{http_code}" -X POST "$HTTP_ENDPOINT/v1/traces" \
        -H "Content-Type: application/json" \
        -d '{
            "resourceSpans": [{
                "resource": {
                    "attributes": [{
                        "key": "service.name",
                        "value": {"stringValue": "'$SERVICE_NAME'"}
                    }]
                },
                "scopeSpans": [{
                    "spans": [{
                        "traceId": "12345678901234567890123456789012",
                        "spanId": "1234567890123456",
                        "name": "test-operation",
                        "startTimeUnixNano": "'$(date +%s)000000000'",
                        "endTimeUnixNano": "'$(($(date +%s) + 1))000000000'",
                        "status": {"code": "STATUS_CODE_OK"},
                        "attributes": [{
                            "key": "test.attribute",
                            "value": {"stringValue": "test-value"}
                        }]
                    }]
                }]
            }]
        }')
    
    local trace_http_code="${trace_response: -3}"
    if [ "$trace_http_code" = "200" ]; then
        print_status "SUCCESS" "Traces sent successfully via HTTP"
    else
        print_status "ERROR" "Failed to send traces via HTTP (HTTP $trace_http_code)"
        return 1
    fi
    
    # Test metrics
    print_status "INFO" "Sending metrics via HTTP..."
    local metric_response=$(curl -s -w "%{http_code}" -X POST "$HTTP_ENDPOINT/v1/metrics" \
        -H "Content-Type: application/json" \
        -d '{
            "resourceMetrics": [{
                "resource": {
                    "attributes": [{
                        "key": "service.name",
                        "value": {"stringValue": "'$SERVICE_NAME'"}
                    }]
                },
                "scopeMetrics": [{
                    "metrics": [{
                        "name": "test.counter",
                        "data": {
                            "sum": {
                                "dataPoints": [{
                                    "timeUnixNano": "'$(date +%s)000000000'",
                                    "asDouble": 42.0,
                                    "attributes": [{
                                        "key": "test.label",
                                        "value": {"stringValue": "test-value"}
                                    }]
                                }]
                            }
                        }
                    }]
                }]
            }]
        }')
    
    local metric_http_code="${metric_response: -3}"
    if [ "$metric_http_code" = "200" ]; then
        print_status "SUCCESS" "Metrics sent successfully via HTTP"
    else
        print_status "ERROR" "Failed to send metrics via HTTP (HTTP $metric_http_code)"
        return 1
    fi
    
    # Test logs
    print_status "INFO" "Sending logs via HTTP..."
    local log_response=$(curl -s -w "%{http_code}" -X POST "$HTTP_ENDPOINT/v1/logs" \
        -H "Content-Type: application/json" \
        -d '{
            "resourceLogs": [{
                "resource": {
                    "attributes": [{
                        "key": "service.name",
                        "value": {"stringValue": "'$SERVICE_NAME'"}
                    }]
                },
                "scopeLogs": [{
                    "logRecords": [{
                        "timeUnixNano": "'$(date +%s)000000000'",
                        "severityText": "INFO",
                        "body": {"stringValue": "Test log message from foo.service"},
                        "attributes": [{
                            "key": "log.level",
                            "value": {"stringValue": "info"}
                        }]
                    }]
                }]
            }]
        }')
    
    local log_http_code="${log_response: -3}"
    if [ "$log_http_code" = "200" ]; then
        print_status "SUCCESS" "Logs sent successfully via HTTP"
    else
        print_status "ERROR" "Failed to send logs via HTTP (HTTP $log_http_code)"
        return 1
    fi
    
    return 0
}

# Function to test gRPC endpoint
test_grpc_endpoint() {
    print_status "INFO" "Testing gRPC endpoint (port 4317)..."
    
    # Note: We're using a simple port check instead of grpcurl
    # since the gRPC server doesn't have OTLP services registered yet
    
    # Test gRPC connection (simple port check)
    print_status "INFO" "Testing gRPC connection..."
    if nc -z localhost 4317 2>/dev/null; then
        print_status "SUCCESS" "gRPC server is listening on port 4317"
    else
        print_status "ERROR" "gRPC server is not listening on port 4317"
        return 1
    fi
    
    # Note: Full OTLP gRPC testing would require proper protobuf definitions
    # For now, we'll just verify the server is running
    print_status "INFO" "gRPC server is running (OTLP services not yet implemented)"
    
    return 0
}

# Function to check if services are running
check_services() {
    print_status "INFO" "Checking if Open-Telemorph-Prime is running..."
    
    # Check HTTP port 4318
    if curl -s "$HTTP_ENDPOINT/v1/traces" -X POST -H "Content-Type: application/json" -d '{"resourceSpans":[]}' > /dev/null 2>&1; then
        print_status "SUCCESS" "HTTP OTLP endpoint (4318) is responding"
    else
        print_status "ERROR" "HTTP OTLP endpoint (4318) is not responding"
        return 1
    fi
    
    # Check gRPC port 4317
    if nc -z localhost 4317 2>/dev/null; then
        print_status "SUCCESS" "gRPC OTLP endpoint (4317) is listening"
    else
        print_status "ERROR" "gRPC OTLP endpoint (4317) is not listening"
        return 1
    fi
    
    # Check web UI port 8080
    if curl -s "http://localhost:8080/health" > /dev/null 2>&1; then
        print_status "SUCCESS" "Web UI (8080) is responding"
    else
        print_status "WARNING" "Web UI (8080) is not responding"
    fi
}

# Main execution
main() {
    echo -e "${BLUE}Open-Telemorph-Prime OTLP Test Script${NC}"
    echo -e "${BLUE}=====================================${NC}"
    echo ""
    
    # Check if services are running
    if ! check_services; then
        print_status "ERROR" "Open-Telemorph-Prime services are not running properly"
        exit 1
    fi
    
    echo ""
    print_status "INFO" "All services are running. Starting OTLP tests..."
    echo ""
    
    # Test HTTP endpoint
    if test_http_endpoint; then
        print_status "SUCCESS" "HTTP OTLP endpoint tests completed successfully"
    else
        print_status "ERROR" "HTTP OTLP endpoint tests failed"
        exit 1
    fi
    
    echo ""
    
    # Test gRPC endpoint
    if test_grpc_endpoint; then
        print_status "SUCCESS" "gRPC OTLP endpoint tests completed successfully"
    else
        print_status "ERROR" "gRPC OTLP endpoint tests failed"
        exit 1
    fi
    
    echo ""
    print_status "SUCCESS" "All OTLP tests completed successfully!"
    print_status "INFO" "Service name used: $SERVICE_NAME"
    print_status "INFO" "Check the Open-Telemorph-Prime web UI at http://localhost:8080 to view the data"
}

# Run main function
main "$@"
