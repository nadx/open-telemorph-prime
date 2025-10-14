#!/bin/bash

# Simple Open-Telemorph-Prime OTLP Test Script
# Tests HTTP endpoint (4318) with logs, traces, and metrics

set -e

# Configuration
SERVICE_NAME="foo.service"
HTTP_ENDPOINT="http://localhost:4318"

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
    fi
}

# Function to send OTLP data
send_otlp_data() {
    local data_type=$1
    local data=$2
    
    local response=$(curl -s -w "%{http_code}" -X POST "$HTTP_ENDPOINT/v1/$data_type" \
        -H "Content-Type: application/json" \
        -d "$data")
    
    local http_code="${response: -3}"
    if [ "$http_code" = "200" ]; then
        print_status "SUCCESS" "$data_type sent successfully"
        return 0
    else
        print_status "ERROR" "Failed to send $data_type (HTTP $http_code)"
        return 1
    fi
}

# Main execution
main() {
    echo -e "${BLUE}Simple Open-Telemorph-Prime OTLP Test${NC}"
    echo -e "${BLUE}=====================================${NC}"
    echo ""
    
    # Check if service is running
    print_status "INFO" "Checking if Open-Telemorph-Prime is running..."
    if ! curl -s "$HTTP_ENDPOINT/v1/traces" -X POST -H "Content-Type: application/json" -d '{"resourceSpans":[]}' > /dev/null 2>&1; then
        print_status "ERROR" "Open-Telemorph-Prime is not running on port 4318"
        exit 1
    fi
    print_status "SUCCESS" "Open-Telemorph-Prime is running"
    echo ""
    
    # Generate timestamp
    local timestamp=$(date +%s)000000000
    local end_timestamp=$((timestamp + 1000000000))
    
    # Test traces
    print_status "INFO" "Sending traces..."
    send_otlp_data "traces" '{
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
                    "startTimeUnixNano": "'$timestamp'",
                    "endTimeUnixNano": "'$end_timestamp'",
                    "status": {"code": "STATUS_CODE_OK"},
                    "attributes": [{
                        "key": "test.attribute",
                        "value": {"stringValue": "test-value"}
                    }]
                }]
            }]
        }]
    }'
    
    # Test metrics
    print_status "INFO" "Sending metrics..."
    send_otlp_data "metrics" '{
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
                                "timeUnixNano": "'$timestamp'",
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
    }'
    
    # Test logs
    print_status "INFO" "Sending logs..."
    send_otlp_data "logs" '{
        "resourceLogs": [{
            "resource": {
                "attributes": [{
                    "key": "service.name",
                    "value": {"stringValue": "'$SERVICE_NAME'"}
                }]
            },
            "scopeLogs": [{
                "logRecords": [{
                    "timeUnixNano": "'$timestamp'",
                    "severityText": "INFO",
                    "body": {"stringValue": "Test log message from foo.service"},
                    "attributes": [{
                        "key": "log.level",
                        "value": {"stringValue": "info"}
                    }]
                }]
            }]
        }]
    }'
    
    echo ""
    print_status "SUCCESS" "All OTLP tests completed!"
    print_status "INFO" "Service name: $SERVICE_NAME"
    print_status "INFO" "Check the web UI at http://localhost:8080 to view the data"
}

# Run main function
main "$@"
