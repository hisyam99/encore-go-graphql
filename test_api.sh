#!/bin/bash

# 🚀 Portfolio GraphQL API - cURL Testing Scripts
# Make sure your server is running on http://localhost:4001

# Set GraphQL endpoint
GRAPHQL_URL="http://localhost:4001/graphql"

echo "🚀 Portfolio GraphQL API Testing Scripts"
echo "========================================="

# 1. Test GraphQL endpoint availability
echo -e "\n📡 Testing GraphQL endpoint availability..."
curl -s -o /dev/null -w "%{http_code}" -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{"query": "{ __schema { types { name } } }"}' | {
  read status
  if [ "$status" = "200" ]; then
    echo "✅ GraphQL endpoint is available"
  else
    echo "❌ GraphQL endpoint is not available (HTTP $status)"
    exit 1
  fi
}

# 2. Create a new user
echo -e "\n👤 Creating a new user..."
USER_RESPONSE=$(curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreateUser($input: CreateUserInput!) { createUser(input: $input) { id name email createdAt updatedAt } }",
    "variables": {
      "input": {
        "name": "John Doe",
        "email": "john.doe@example.com"
      }
    }
  }')

echo "📝 Create User Response:"
echo "$USER_RESPONSE" | jq '.'

# Extract user ID for further operations
USER_ID=$(echo "$USER_RESPONSE" | jq -r '.data.createUser.id // empty')
if [ -n "$USER_ID" ]; then
  echo "✅ User created with ID: $USER_ID"
else
  echo "❌ Failed to create user"
  USER_ID="1"  # fallback for testing
fi

# 3. Get all users
echo -e "\n👥 Getting all users..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetUsers($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) { users(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) { data { id name email createdAt updatedAt } pagination { page pageSize total totalPages } } }",
    "variables": {
      "page": 1,
      "pageSize": 10,
      "sortBy": "createdAt",
      "sortDirection": "DESC"
    }
  }' | jq '.'

# 4. Create a category
echo -e "\n📁 Creating a new category..."
CATEGORY_RESPONSE=$(curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreateCategory($input: CreateCategoryInput!) { createCategory(input: $input) { id name description createdAt updatedAt } }",
    "variables": {
      "input": {
        "name": "Web Development",
        "description": "Front-end and back-end web development skills"
      }
    }
  }')

echo "📝 Create Category Response:"
echo "$CATEGORY_RESPONSE" | jq '.'

# Extract category ID
CATEGORY_ID=$(echo "$CATEGORY_RESPONSE" | jq -r '.data.createCategory.id // empty')
if [ -n "$CATEGORY_ID" ]; then
  echo "✅ Category created with ID: $CATEGORY_ID"
else
  echo "❌ Failed to create category"
  CATEGORY_ID="1"  # fallback for testing
fi

# 5. Create resume content
echo -e "\n📄 Creating resume content..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateResumeContent(\$input: CreateResumeContentInput!) { createResumeContent(input: \$input) { id title description detail categoryId category { id name description } createdAt updatedAt } }\",
    \"variables\": {
      \"input\": {
        \"title\": \"React.js Development\",
        \"description\": \"Advanced React.js development skills\",
        \"detail\": \"Experienced in building scalable React applications with hooks, context, and state management libraries like Redux and Zustand.\",
        \"categoryId\": \"$CATEGORY_ID\"
      }
    }
  }" | jq '.'

# 6. Create a project
echo -e "\n🚀 Creating a new project..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateProject(\$input: CreateProjectInput!) { createProject(input: \$input) { id title description userId user { id name email } createdAt updatedAt } }\",
    \"variables\": {
      \"input\": {
        \"title\": \"E-commerce Platform\",
        \"description\": \"Full-stack e-commerce application built with React and Node.js\",
        \"userId\": \"$USER_ID\"
      }
    }
  }" | jq '.'

# 7. Create a blog post
echo -e "\n📝 Creating a blog post..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreateBlog($input: CreateBlogInput!) { createBlog(input: $input) { id title content summary slug author publishedAt status tags metaDescription createdAt updatedAt } }",
    "variables": {
      "input": {
        "title": "Getting Started with GraphQL",
        "content": "GraphQL is a powerful query language for APIs and a runtime for executing those queries with your existing data. Unlike REST APIs that expose multiple endpoints for different resources, GraphQL provides a single endpoint that allows clients to request exactly the data they need. This reduces over-fetching and under-fetching of data, making applications more efficient.",
        "summary": "An introduction to GraphQL basics and how to get started",
        "slug": "getting-started-with-graphql",
        "author": "John Doe",
        "status": "DRAFT",
        "tags": ["GraphQL", "API", "Tutorial", "Backend"],
        "metaDescription": "Learn the basics of GraphQL and how to implement it in your applications"
      }
    }
  }' | jq '.'

# 8. Get all categories with resume contents
echo -e "\n📁 Getting all categories with resume contents..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetCategories($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) { categories(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) { data { id name description createdAt updatedAt resumeContents { id title description detail } } pagination { page pageSize total totalPages } } }",
    "variables": {
      "page": 1,
      "pageSize": 10,
      "sortBy": "name",
      "sortDirection": "ASC"
    }
  }' | jq '.'

# 9. Get projects by user
echo -e "\n🚀 Getting projects by user..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"query GetProjectsByUser(\$userId: ID!, \$page: Int, \$pageSize: Int, \$sortBy: String, \$sortDirection: SortDirection) { projectsByUser(userId: \$userId, page: \$page, pageSize: \$pageSize, sortBy: \$sortBy, sortDirection: \$sortDirection) { data { id title description userId user { id name email } createdAt updatedAt } pagination { page pageSize total totalPages } } }\",
    \"variables\": {
      \"userId\": \"$USER_ID\",
      \"page\": 1,
      \"pageSize\": 10,
      \"sortBy\": \"createdAt\",
      \"sortDirection\": \"DESC\"
    }
  }" | jq '.'

# 10. Get all blogs
echo -e "\n📝 Getting all blogs..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetBlogs($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) { blogs(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) { data { id title summary slug author status tags createdAt updatedAt } pagination { page pageSize total totalPages } } }",
    "variables": {
      "page": 1,
      "pageSize": 10,
      "sortBy": "createdAt",
      "sortDirection": "DESC"
    }
  }' | jq '.'

# 11. Get blog by slug
echo -e "\n📝 Getting blog by slug..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetBlogBySlug($slug: String!) { blogBySlug(slug: $slug) { id title content summary slug author publishedAt status tags metaDescription createdAt updatedAt } }",
    "variables": {
      "slug": "getting-started-with-graphql"
    }
  }' | jq '.'

# 12. Complex query - Complete portfolio data
echo -e "\n🔥 Getting complete portfolio data..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"query GetCompletePortfolioData { user(id: \\\"$USER_ID\\\") { id name email projects { id title description } } categories { data { id name description resumeContents { id title description detail } } } publishedBlogs(page: 1, pageSize: 5) { data { id title summary slug publishedAt tags } pagination { total totalPages } } }\"
  }" | jq '.'

# 13. Update user
echo -e "\n👤 Updating user..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation UpdateUser(\$id: ID!, \$input: UpdateUserInput!) { updateUser(id: \$id, input: \$input) { id name email createdAt updatedAt } }\",
    \"variables\": {
      \"id\": \"$USER_ID\",
      \"input\": {
        \"name\": \"John Smith\",
        \"email\": \"john.smith@example.com\"
      }
    }
  }" | jq '.'

# 14. Blog dashboard data
echo -e "\n📊 Getting blog dashboard data..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query BlogDashboard { blogs { data { id title status createdAt updatedAt } pagination { total totalPages } } blogsByStatus(status: DRAFT) { pagination { total } } blogsByStatus(status: PUBLISHED) { pagination { total } } }"
  }' | jq '.'

# 15. Test introspection query
echo -e "\n🔍 Testing GraphQL introspection..."
curl -s -X POST "$GRAPHQL_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query IntrospectionQuery { __schema { queryType { name } mutationType { name } types { name kind description } } }"
  }' | jq '.data.__schema.types | length as $count | "Available types: \($count)"'

echo -e "\n✅ All tests completed!"
echo "========================================="
echo "📖 Check GraphQL_API_Documentation.md for complete API documentation"
echo "📦 Import insomnia_collection.json into Insomnia for GUI testing"
echo "🌐 Access GraphQL Playground at: http://localhost:4001/graphql"