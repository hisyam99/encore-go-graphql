#!/bin/bash

# GraphQL N+1 Testing Script
# This script tests various GraphQL queries that would typically cause N+1 problems

GRAPHQL_ENDPOINT="http://127.0.0.1:4002/graphql"

echo "🚀 Testing GraphQL API for N+1 Query Optimization"
echo "================================================"

# Test 1: Users with Projects (potential N+1)
echo -e "\n📋 Test 1: Query users with their projects"
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetUsersWithProjects { users(page: 1, pageSize: 5) { data { id name email projects { id title description } } pagination { total page pageSize } } }"
  }' \
  "$GRAPHQL_ENDPOINT" | jq '.'

# Test 2: Categories with Resume Contents (potential N+1)
echo -e "\n📋 Test 2: Query categories with resume contents"
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetCategoriesWithResumeContents { categories(page: 1, pageSize: 5) { data { id name description resumeContents { id title description } } pagination { total page pageSize } } }"
  }' \
  "$GRAPHQL_ENDPOINT" | jq '.'

# Test 3: Resume Contents with Categories (potential N+1)
echo -e "\n📋 Test 3: Query resume contents with their categories"
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetResumeContentsWithCategories { resumeContents(page: 1, pageSize: 5) { data { id title description category { id name description } } pagination { total page pageSize } } }"
  }' \
  "$GRAPHQL_ENDPOINT" | jq '.'

# Test 4: Projects with Users (potential N+1)
echo -e "\n📋 Test 4: Query projects with their users"
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query GetProjectsWithUsers { projects(page: 1, pageSize: 5) { data { id title description user { id name email } } pagination { total page pageSize } } }"
  }' \
  "$GRAPHQL_ENDPOINT" | jq '.'

# Test 5: Complex nested query (multiple potential N+1s)
echo -e "\n📋 Test 5: Complex nested query with multiple relationships"
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "query": "query ComplexNestedQuery { users(page: 1, pageSize: 3) { data { id name email projects { id title user { id name } } } } categories(page: 1, pageSize: 3) { data { id name resumeContents { id title category { id name } } } } }"
  }' \
  "$GRAPHQL_ENDPOINT" | jq '.'

echo -e "\n✅ All tests completed!"
echo -e "\n📊 DataLoader should batch the following operations:"
echo "   - User->Projects: Single query for all user projects"
echo "   - Category->ResumeContents: Single query for all category resume contents"
echo "   - ResumeContent->Category: Single query for all categories"
echo "   - Project->User: Single query for all project users"
echo -e "\n🎯 Expected Result: No N+1 queries, efficient batched database operations"