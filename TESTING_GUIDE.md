# 🧪 Testing Guide - Portfolio GraphQL API

Complete testing guide untuk Portfolio Backend GraphQL API.

## 🚀 Quick Start

### 1. Start the Server
```bash
cd /path/to/your/backend
encore run --port=4001
```

### 2. Test with cURL Script
```bash
# Make script executable
chmod +x test_api.sh

# Run all tests
./test_api.sh
```

### 3. Access GraphQL Playground
```
http://localhost:4001/graphql
```

## 📋 Available Testing Methods

### Method 1: GraphQL Playground (Recommended)
- **URL**: `http://localhost:4001/graphql`
- **Features**: 
  - Interactive query editor
  - Schema documentation
  - Query validation
  - Auto-completion

### Method 2: Insomnia/Postman Collection
- Import `insomnia_collection.json`
- All queries and mutations pre-configured
- Environment variables included

### Method 3: cURL Commands
- Run `./test_api.sh` for automated testing
- Individual curl commands available in script

### Method 4: Frontend Integration
```javascript
// Example with Apollo Client
import { ApolloClient, InMemoryCache, gql } from '@apollo/client';

const client = new ApolloClient({
  uri: 'http://localhost:4001/graphql',
  cache: new InMemoryCache()
});

// Query example
const GET_USERS = gql`
  query GetUsers {
    users {
      data {
        id
        name
        email
      }
    }
  }
`;
```

## 🎯 Testing Scenarios

### Basic CRUD Testing

#### 1. User Management
```bash
# Create User
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createUser(input: { name: \"Test User\", email: \"test@example.com\" }) { id name email } }"
  }'

# Get Users
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ users { data { id name email } } }"
  }'

# Update User
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { updateUser(id: \"1\", input: { name: \"Updated User\" }) { id name email } }"
  }'

# Delete User
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { deleteUser(id: \"1\") }"
  }'
```

#### 2. Category Management
```bash
# Create Category
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createCategory(input: { name: \"Web Development\", description: \"Frontend and backend skills\" }) { id name description } }"
  }'

# Get Categories
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ categories { data { id name description resumeContents { id title } } } }"
  }'
```

#### 3. Resume Content Management
```bash
# Create Resume Content
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createResumeContent(input: { title: \"React.js\", description: \"Frontend framework\", detail: \"Advanced React development\", categoryId: \"1\" }) { id title category { name } } }"
  }'

# Get Resume Contents by Category
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ resumeContentsByCategory(categoryId: \"1\") { data { id title description detail } } }"
  }'
```

### Advanced Testing Scenarios

#### 1. Pagination Testing
```graphql
query TestPagination {
  users(page: 1, pageSize: 5, sortBy: "createdAt", sortDirection: DESC) {
    data {
      id
      name
      email
    }
    pagination {
      page
      pageSize
      total
      totalPages
    }
  }
}
```

#### 2. Relationship Loading
```graphql
query TestRelationships {
  user(id: "1") {
    id
    name
    email
    projects {
      id
      title
      description
    }
  }
}
```

#### 3. Complex Filtering
```graphql
query TestFiltering {
  blogsByStatus(status: PUBLISHED, page: 1, pageSize: 10) {
    data {
      id
      title
      status
      publishedAt
    }
    pagination {
      total
    }
  }
}
```

## 🔍 Validation Testing

### Input Validation Tests

#### 1. Required Fields
```graphql
# This should fail - missing required fields
mutation TestValidation {
  createUser(input: { name: "" }) {
    id
    name
    email
  }
}
```

#### 2. Email Format Validation
```graphql
# This should fail - invalid email format
mutation TestEmailValidation {
  createUser(input: { name: "Test User", email: "invalid-email" }) {
    id
    name
    email
  }
}
```

#### 3. Unique Constraints
```graphql
# This should fail if user already exists
mutation TestUniqueConstraint {
  createUser(input: { name: "Test User", email: "existing@example.com" }) {
    id
    name
    email
  }
}
```

## 📊 Performance Testing

### 1. Load Testing with curl
```bash
# Test multiple concurrent requests
for i in {1..10}; do
  curl -X POST http://localhost:4001/graphql \
    -H "Content-Type: application/json" \
    -d '{"query": "{ users { data { id name email } } }"}' &
done
wait
```

### 2. Large Dataset Testing
```graphql
# Test pagination with large dataset
query TestLargeDataset {
  blogs(page: 1, pageSize: 100) {
    data {
      id
      title
      status
      createdAt
    }
    pagination {
      total
      totalPages
    }
  }
}
```

## 🐛 Error Handling Testing

### 1. Not Found Errors
```graphql
query TestNotFound {
  user(id: "999999") {
    id
    name
    email
  }
}
```

### 2. Invalid ID Format
```graphql
query TestInvalidID {
  user(id: "invalid-id") {
    id
    name
    email
  }
}
```

### 3. Database Constraint Violations
```graphql
mutation TestConstraintViolation {
  createResumeContent(input: { 
    title: "Test", 
    categoryId: "999999"  # Non-existent category
  }) {
    id
    title
  }
}
```

## 📈 Monitoring & Debugging

### 1. Query Performance Monitoring
```graphql
# Use GraphQL Playground's timing feature
query MonitorPerformance {
  users {
    data {
      id
      name
      projects {
        id
        title
      }
    }
  }
}
```

### 2. Database Query Logging
Enable logging in your application to monitor database queries:
```bash
# Check server logs for query execution times
tail -f server.log | grep "database"
```

## 🔧 Test Data Setup

### Sample Test Data Script
```bash
#!/bin/bash
# Create sample data for testing

GRAPHQL_URL="http://localhost:4001/graphql"

# Create users
for i in {1..5}; do
  curl -s -X POST "$GRAPHQL_URL" \
    -H "Content-Type: application/json" \
    -d "{
      \"query\": \"mutation { createUser(input: { name: \\\"User $i\\\", email: \\\"user$i@example.com\\\" }) { id } }\"
    }"
done

# Create categories
categories=("Web Development" "Mobile Development" "Database" "DevOps" "UI/UX Design")
for category in "${categories[@]}"; do
  curl -s -X POST "$GRAPHQL_URL" \
    -H "Content-Type: application/json" \
    -d "{
      \"query\": \"mutation { createCategory(input: { name: \\\"$category\\\", description: \\\"$category skills\\\" }) { id } }\"
    }"
done
```

## 🚨 Common Issues & Troubleshooting

### 1. Server Not Running
```bash
# Check if server is running
curl -s http://localhost:4001/health || echo "Server not running"

# Start server
encore run --port=4001
```

### 2. CORS Issues
If testing from browser, ensure CORS is configured:
```go
// In your GraphQL handler
handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))
handler.Use(extension.Introspection{})
```

### 3. Database Connection Issues
```bash
# Check database status
encore db shell --env=local
```

### 4. GraphQL Syntax Errors
Use GraphQL Playground for:
- Query validation
- Schema introspection
- Auto-completion

### 5. CORS Issues
Test CORS configuration:

```bash
# Test CORS with curl
curl -H "Origin: https://hisyam.tar.my.id" \
     -H "Access-Control-Request-Method: POST" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS \
     http://localhost:4001/graphql

# Test CORS endpoint
curl -H "Origin: https://hisyam.tar.my.id" \
     http://localhost:4001/cors-test

# Test from browser console (on https://hisyam.tar.my.id)
fetch('http://localhost:4001/cors-test', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
  },
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('CORS Error:', error));
```

Expected CORS headers in response:
- `Access-Control-Allow-Origin: https://hisyam.tar.my.id`
- `Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS`
- `Access-Control-Allow-Headers: Accept, Authorization, Content-Type, X-CSRF-Token, X-Requested-With`
- `Access-Control-Allow-Credentials: true`

## 📚 Additional Resources

- **GraphQL Playground**: `http://localhost:4001/graphql`
- **API Documentation**: `GraphQL_API_Documentation.md`
- **Insomnia Collection**: `insomnia_collection.json`
- **Automated Tests**: `test_api.sh`

## 🎯 Test Checklist

### Basic Functionality
- [ ] All queries return expected data
- [ ] All mutations work correctly
- [ ] Pagination works for all list queries
- [ ] Sorting works in both directions
- [ ] Relationships load correctly

### Error Handling
- [ ] Invalid IDs return proper errors
- [ ] Required field validation works
- [ ] Unique constraint violations handled
- [ ] Not found errors return null/error

### Performance
- [ ] Queries execute within acceptable time limits
- [ ] Pagination prevents large data loads
- [ ] N+1 query problems avoided
- [ ] Memory usage reasonable

### Security
- [ ] Input sanitization works
- [ ] SQL injection prevention active
- [ ] Rate limiting (if implemented)
- [ ] Authentication/authorization (if implemented)

---

## 🎉 Happy Testing!

Your Portfolio GraphQL API is ready for comprehensive testing. Use the tools and methods above to ensure everything works perfectly before deployment.

**Remember**: Always test on a development database, never on production data!