# 📋 GraphQL API Quick Reference

## 🚀 Base URL
```
http://localhost:4001/graphql
```

## 🔍 Quick Queries

### Get All Users
```graphql
{ users { data { id name email } } }
```

### Get User with Projects
```graphql
{ user(id: "1") { id name email projects { id title } } }
```

### Get All Categories
```graphql
{ categories { data { id name description } } }
```

### Get Resume Contents by Category
```graphql
{ resumeContentsByCategory(categoryId: "1") { data { id title description } } }
```

### Get All Projects
```graphql
{ projects { data { id title description user { name } } } }
```

### Get All Blogs
```graphql
{ blogs { data { id title summary slug status tags } } }
```

### Get Published Blogs Only
```graphql
{ publishedBlogs { data { id title summary slug publishedAt } } }
```

### Get Blog by Slug
```graphql
{ blogBySlug(slug: "my-blog-post") { id title content author } }
```

## ✏️ Quick Mutations

### Create User
```graphql
mutation {
  createUser(input: { name: "John Doe", email: "john@example.com" }) {
    id name email
  }
}
```

### Create Category
```graphql
mutation {
  createCategory(input: { name: "Web Dev", description: "Web development skills" }) {
    id name description
  }
}
```

### Create Resume Content
```graphql
mutation {
  createResumeContent(input: { 
    title: "React.js", 
    description: "Frontend framework",
    detail: "Advanced React development skills",
    categoryId: "1" 
  }) {
    id title category { name }
  }
}
```

### Create Project
```graphql
mutation {
  createProject(input: { 
    title: "My Portfolio", 
    description: "Personal portfolio website",
    userId: "1" 
  }) {
    id title user { name }
  }
}
```

### Create Blog
```graphql
mutation {
  createBlog(input: { 
    title: "Getting Started", 
    content: "This is my first blog post...",
    slug: "getting-started",
    status: DRAFT,
    tags: ["tutorial", "beginner"]
  }) {
    id title slug status
  }
}
```

## 🔧 Pagination Examples

### With Pagination
```graphql
{
  users(page: 1, pageSize: 5, sortBy: "createdAt", sortDirection: DESC) {
    data { id name email }
    pagination { page pageSize total totalPages }
  }
}
```

### Sort by Name (Ascending)
```graphql
{
  categories(sortBy: "name", sortDirection: ASC) {
    data { id name }
  }
}
```

## 🔍 Complex Queries

### Complete Portfolio Data
```graphql
{
  user(id: "1") {
    id name email
    projects { id title description }
  }
  categories {
    data {
      id name
      resumeContents { id title description }
    }
  }
  publishedBlogs(page: 1, pageSize: 3) {
    data { id title summary slug }
  }
}
```

### Blog Dashboard Stats
```graphql
{
  allBlogs: blogs { pagination { total } }
  draftBlogs: blogsByStatus(status: DRAFT) { pagination { total } }
  publishedBlogs: blogsByStatus(status: PUBLISHED) { pagination { total } }
}
```

## 🚨 Common Errors to Test

### Invalid ID
```graphql
{ user(id: "invalid") { id name } } # Should return error
```

### Missing Required Fields
```graphql
mutation {
  createUser(input: { name: "" }) { id } # Should fail validation
}
```

### Non-existent Resource
```graphql
{ user(id: "999999") { id name } } # Should return null
```

## 📱 Test with cURL

### Basic Query
```bash
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ users { data { id name email } } }"}'
```

### Create User
```bash
curl -X POST http://localhost:4001/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createUser(input: { name: \"Test User\", email: \"test@example.com\" }) { id name email } }"
  }'
```

## 🎯 Available Operations

### Queries (15 total)
- `users` - Get all users with pagination
- `user(id)` - Get user by ID
- `categories` - Get all categories
- `category(id)` - Get category by ID  
- `resumeContents` - Get all resume contents
- `resumeContent(id)` - Get resume content by ID
- `resumeContentsByCategory(categoryId)` - Filter by category
- `projects` - Get all projects
- `project(id)` - Get project by ID
- `projectsByUser(userId)` - Filter by user
- `blogs` - Get all blogs
- `blog(id)` - Get blog by ID
- `blogBySlug(slug)` - Get blog by slug
- `publishedBlogs` - Get only published blogs
- `blogsByStatus(status)` - Filter by status

### Mutations (15 total)
- `createUser`, `updateUser`, `deleteUser`
- `createCategory`, `updateCategory`, `deleteCategory`  
- `createResumeContent`, `updateResumeContent`, `deleteResumeContent`
- `createProject`, `updateProject`, `deleteProject`
- `createBlog`, `updateBlog`, `deleteBlog`

## 🔧 Tools for Testing

1. **GraphQL Playground**: `http://localhost:4001/graphql`
2. **Automated Script**: `./test_api.sh`
3. **Insomnia Collection**: Import `insomnia_collection.json`
4. **cURL Commands**: See examples above

---

**💡 Tip**: Use GraphQL Playground for the best development experience with auto-completion and documentation!