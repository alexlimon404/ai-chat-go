#!/bin/bash

# Test script for AI Chat Go API

BASE_URL="http://localhost:8080"

echo "Testing AI Chat Go API..."
echo ""

# Test GET /go-api/models
echo "1. GET /go-api/models - Get all active AI models"
curl -s -X GET "$BASE_URL/go-api/models" | jq .
echo ""
echo "---"
echo ""

# Health check (if implemented)
echo "2. Server health check"
curl -s -o /dev/null -w "HTTP Status: %{http_code}\n" "$BASE_URL/go-api/models"
echo ""
