# 📖 GraphQL API Documentation - Portfolio Backend

Dokumentasi lengkap untuk semua operasi GraphQL yang tersedia di API Portfolio Backend.

## 🎯 Base URL
```
http://localhost:4001/graphql
```

## 📋 Table of Contents
1. [Queries](#queries)
   - [User Queries](#user-queries)
   - [Category Queries](#category-queries)
   - [Resume Content Queries](#resume-content-queries)
   - [Project Queries](#project-queries)
   - [Blog Queries](#blog-queries)
2. [Mutations](#mutations)
   - [User Mutations](#user-mutations)
   - [Category Mutations](#category-mutations)
   - [Resume Content Mutations](#resume-content-mutations)
   - [Project Mutations](#project-mutations)
   - [Blog Mutations](#blog-mutations)
3. [Types](#types)
4. [Examples](#examples)

---

## 🔍 Queries

### User Queries

#### 1. Get All Users (Paginated)
```graphql
query GetUsers($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  users(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      name
      email
      createdAt
      updatedAt
      projects {
        id
        title
        description
      }
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

**Variables:**
```json
{
  "page": 1,
  "pageSize": 10,
  "sortBy": "createdAt",
  "sortDirection": "DESC"
}
```

#### 2. Get User by ID
```graphql
query GetUser($id: ID!) {
  user(id: $id) {
    id
    name
    email
    createdAt
    updatedAt
    projects {
      id
      title
      description
      createdAt
      updatedAt
    }
  }
}
```

**Variables:**
```json
{
  "id": "1"
}
```

### Category Queries

#### 3. Get All Categories (Paginated)
```graphql
query GetCategories($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  categories(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data { 
      id
      name
      description
      createdAt
      updatedAt
      resumeContents {
        id
        title
        description
        detail
      }
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

#### 4. Get Category by ID
```graphql
query GetCategory($id: ID!) {
  category(id: $id) {
    id
    name
    description
    createdAt
    updatedAt
    resumeContents {
      id
      title
      description
      detail
      createdAt
      updatedAt
    }
  }
}
```

### Resume Content Queries

#### 5. Get All Resume Contents (Paginated)
```graphql
query GetResumeContents($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  resumeContents(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      description
      detail
      categoryId
      category {
        id
        name
        description
      }
      createdAt
      updatedAt
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

#### 6. Get Resume Content by ID
```graphql
query GetResumeContent($id: ID!) {
  resumeContent(id: $id) {
    id
    title
    description
    detail
    categoryId
    category {
      id
      name
      description
    }
    createdAt
    updatedAt
  }
}
```

#### 7. Get Resume Contents by Category
```graphql
query GetResumeContentsByCategory($categoryId: ID!, $page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  resumeContentsByCategory(categoryId: $categoryId, page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      description
      detail
      categoryId
      category {
        id
        name
        description
      }
      createdAt
      updatedAt
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

### Project Queries

#### 8. Get All Projects (Paginated)
```graphql
query GetProjects($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  projects(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      description
      userId
      user {
        id
        name
        email
      }
      createdAt
      updatedAt
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

#### 9. Get Project by ID
```graphql
query GetProject($id: ID!) {
  project(id: $id) {
    id
    title
    description
    userId
    user {
      id
      name
      email
    }
    createdAt
    updatedAt
  }
}
```

#### 10. Get Projects by User
```graphql
query GetProjectsByUser($userId: ID!, $page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  projectsByUser(userId: $userId, page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      description
      userId
      user {
        id
        name
        email
      }
      createdAt
      updatedAt
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

### Blog Queries

#### 11. Get All Blogs (Paginated)
```graphql
query GetBlogs($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  blogs(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      content
      summary
      slug
      author
      publishedAt
      status
      tags
      metaDescription
      createdAt
      updatedAt
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

#### 12. Get Blog by ID
```graphql
query GetBlog($id: ID!) {
  blog(id: $id) {
    id
    title
    content
    summary
    slug
    author
    publishedAt
    status
    tags
    metaDescription
    createdAt
    updatedAt
  }
}
```

#### 13. Get Blog by Slug
```graphql
query GetBlogBySlug($slug: String!) {
  blogBySlug(slug: $slug) {
    id
    title
    content
    summary
    slug
    author
    publishedAt
    status
    tags
    metaDescription
    createdAt
    updatedAt
  }
}
```

#### 14. Get Published Blogs
```graphql
query GetPublishedBlogs($page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  publishedBlogs(page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      content
      summary
      slug
      author
      publishedAt
      status
      tags
      metaDescription
      createdAt
      updatedAt
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

#### 15. Get Blogs by Status
```graphql
query GetBlogsByStatus($status: BlogStatus!, $page: Int, $pageSize: Int, $sortBy: String, $sortDirection: SortDirection) {
  blogsByStatus(status: $status, page: $page, pageSize: $pageSize, sortBy: $sortBy, sortDirection: $sortDirection) {
    data {
      id
      title
      content
      summary
      slug
      author
      publishedAt
      status
      tags
      metaDescription
      createdAt
      updatedAt
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

---

## ✏️ Mutations

### User Mutations

#### 1. Create User
```graphql
mutation CreateUser($input: CreateUserInput!) {
  createUser(input: $input) {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

**Variables:**
```json
{
  "input": {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
}
```

#### 2. Update User
```graphql
mutation UpdateUser($id: ID!, $input: UpdateUserInput!) {
  updateUser(id: $id, input: $input) {
    id
    name
    email
    createdAt
    updatedAt
  }
}
```

**Variables:**
```json
{
  "id": "1",
  "input": {
    "name": "John Smith",
    "email": "john.smith@example.com"
  }
}
```

#### 3. Delete User
```graphql
mutation DeleteUser($id: ID!) {
  deleteUser(id: $id)
}
```

**Variables:**
```json
{
  "id": "1"
}
```

### Category Mutations

#### 4. Create Category
```graphql
mutation CreateCategory($input: CreateCategoryInput!) {
  createCategory(input: $input) {
    id
    name
    description
    createdAt
    updatedAt
  }
}
```

**Variables:**
```json
{
  "input": {
    "name": "Web Development",
    "description": "Front-end and back-end web development skills"
  }
}
```

#### 5. Update Category
```graphql
mutation UpdateCategory($id: ID!, $input: UpdateCategoryInput!) {
  updateCategory(id: $id, input: $input) {
    id
    name
    description
    createdAt
    updatedAt
  }
}
```

#### 6. Delete Category
```graphql
mutation DeleteCategory($id: ID!) {
  deleteCategory(id: $id)
}
```

### Resume Content Mutations

#### 7. Create Resume Content
```graphql
mutation CreateResumeContent($input: CreateResumeContentInput!) {
  createResumeContent(input: $input) {
    id
    title
    description
    detail
    categoryId
    category {
      id
      name
      description
    }
    createdAt
    updatedAt
  }
}
```

**Variables:**
```json
{
  "input": {
    "title": "React.js Development",
    "description": "Advanced React.js development skills",
    "detail": "Experienced in building scalable React applications with hooks, context, and state management",
    "categoryId": "1"
  }
}
```

#### 8. Update Resume Content
```graphql
mutation UpdateResumeContent($id: ID!, $input: UpdateResumeContentInput!) {
  updateResumeContent(id: $id, input: $input) {
    id
    title
    description
    detail
    categoryId
    category {
      id
      name
      description
    }
    createdAt
    updatedAt
  }
}
```

#### 9. Delete Resume Content
```graphql
mutation DeleteResumeContent($id: ID!) {
  deleteResumeContent(id: $id)
}
```

### Project Mutations

#### 10. Create Project
```graphql
mutation CreateProject($input: CreateProjectInput!) {
  createProject(input: $input) {
    id
    title
    description
    userId
    user {
      id
      name
      email
    }
    createdAt
    updatedAt
  }
}
```

**Variables:**
```json
{
  "input": {
    "title": "E-commerce Platform",
    "description": "Full-stack e-commerce application built with React and Node.js",
    "userId": "1"
  }
}
```

#### 11. Update Project
```graphql
mutation UpdateProject($id: ID!, $input: UpdateProjectInput!) {
  updateProject(id: $id, input: $input) {
    id
    title
    description
    userId
    user {
      id
      name
      email
    }
    createdAt
    updatedAt
  }
}
```

#### 12. Delete Project
```graphql
mutation DeleteProject($id: ID!) {
  deleteProject(id: $id)
}
```

### Blog Mutations

#### 13. Create Blog
```graphql
mutation CreateBlog($input: CreateBlogInput!) {
  createBlog(input: $input) {
    id
    title
    content
    summary
    slug
    author
    publishedAt
    status
    tags
    metaDescription
    createdAt
    updatedAt
  }
}
```

**Variables:**
```json
{
  "input": {
    "title": "Getting Started with GraphQL",
    "content": "GraphQL is a powerful query language...",
    "summary": "An introduction to GraphQL basics",
    "slug": "getting-started-with-graphql",
    "author": "John Doe",
    "status": "DRAFT",
    "tags": ["GraphQL", "API", "Tutorial"],
    "metaDescription": "Learn the basics of GraphQL"
  }
}
```

#### 14. Update Blog
```graphql
mutation UpdateBlog($id: ID!, $input: UpdateBlogInput!) {
  updateBlog(id: $id, input: $input) {
    id
    title
    content
    summary
    slug
    author
    publishedAt
    status
    tags
    metaDescription
    createdAt
    updatedAt
  }
}
```

#### 15. Delete Blog
```graphql
mutation DeleteBlog($id: ID!) {
  deleteBlog(id: $id)
}
```

---

## 📊 Types & Enums

### BlogStatus Enum
```graphql
enum BlogStatus {
  DRAFT
  PUBLISHED
}
```

### SortDirection Enum
```graphql
enum SortDirection {
  ASC
  DESC
}
```

### Main Types

#### User Type
```graphql
type User {
  id: ID!
  name: String!
  email: String!
  createdAt: String!
  updatedAt: String!
  projects: [Project!]!
}
```

#### Category Type
```graphql
type Category {
  id: ID!
  name: String!
  description: String
  createdAt: String!
  updatedAt: String!
  resumeContents: [ResumeContent!]!
}
```

#### ResumeContent Type
```graphql
type ResumeContent {
  id: ID!
  title: String!
  description: String
  detail: String
  categoryId: ID!
  category: Category!
  createdAt: String!
  updatedAt: String!
}
```

#### Project Type
```graphql
type Project {
  id: ID!
  title: String!
  description: String
  userId: ID
  user: User
  createdAt: String!
  updatedAt: String!
}
```

#### Blog Type
```graphql
type Blog {
  id: ID!
  title: String!
  content: String!
  summary: String
  slug: String!
  author: String
  publishedAt: String
  status: BlogStatus!
  tags: [String!]!
  metaDescription: String
  createdAt: String!
  updatedAt: String!
}
```

### Pagination Type
```graphql
type PaginationInfo {
  page: Int!
  pageSize: Int!
  total: Int!
  totalPages: Int!
}
```

---

## 🔧 Examples & Use Cases

### 1. Complete Portfolio Data Fetch
```graphql
query GetCompletePortfolioData {
  # Get user info
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
  
  # Get all categories with resume contents
  categories {
    data {
      id
      name
      description
      resumeContents {
        id
        title
        description
        detail
      }
    }
  }
  
  # Get published blogs
  publishedBlogs(page: 1, pageSize: 5) {
    data {
      id
      title
      summary
      slug
      publishedAt
      tags
    }
    pagination {
      total
      totalPages
    }
  }
}
```

### 2. Blog Management Dashboard
```graphql
query BlogDashboard {
  # All blogs with status
  blogs {
    data {
      id
      title
      status
      createdAt
      updatedAt
    }
    pagination {
      total
      totalPages
    }
  }
  
  # Draft blogs count
  blogsByStatus(status: DRAFT) {
    pagination {
      total
    }
  }
  
  # Published blogs count
  blogsByStatus(status: PUBLISHED) {
    pagination {
      total
    }
  }
}
```

### 3. User Portfolio with Projects
```graphql
query UserPortfolio($userId: ID!) {
  user(id: $userId) {
    id
    name
    email
    createdAt
  }
  
  projectsByUser(userId: $userId) {
    data {
      id
      title
      description
      createdAt
    }
    pagination {
      total
    }
  }
}
```

### 4. Category-based Resume Content
```graphql
query CategoryResumeContents($categoryId: ID!) {
  category(id: $categoryId) {
    id
    name
    description
  }
  
  resumeContentsByCategory(categoryId: $categoryId) {
    data {
      id
      title
      description
      detail
      createdAt
    }
    pagination {
      total
    }
  }
}
```

---

## 🚀 Advanced Features

### Pagination Parameters
- `page`: Page number (default: 1)
- `pageSize`: Items per page (default: 10)
- `sortBy`: Field to sort by (varies per type)
- `sortDirection`: ASC or DESC (default varies)

### Default Sort Fields
- **Users**: `createdAt` (DESC)
- **Categories**: `name` (ASC)
- **Resume Contents**: `createdAt` (DESC)
- **Projects**: `createdAt` (DESC)
- **Blogs**: `createdAt` (DESC)
- **Published Blogs**: `publishedAt` (DESC)

### Error Handling
All mutations and queries include proper error handling:
- Input validation errors
- Not found errors (404)
- Database constraint errors
- Authentication/authorization errors (if implemented)

### Response Format
All paginated queries return data in this format:
```json
{
  "data": [...],
  "pagination": {
    "page": 1,
    "pageSize": 10,
    "total": 50,
    "totalPages": 5
  }
}
```

---

## 📱 Testing with GraphQL Playground

Akses GraphQL Playground di: `http://localhost:4001/graphql`

### Sample Test Queries

1. **Test Basic User Query:**
```graphql
{
  users {
    data {
      id
      name
      email
    }
  }
}
```

2. **Test Create and Read Operations:**
```graphql
# First, create a user
mutation {
  createUser(input: { name: "Test User", email: "test@example.com" }) {
    id
    name
    email
  }
}

# Then fetch the user
query {
  user(id: "1") {
    id
    name
    email
    createdAt
  }
}
```

3. **Test Pagination:**
```graphql
{
  blogs(page: 1, pageSize: 5, sortBy: "createdAt", sortDirection: DESC) {
    data {
      id
      title
      status
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

---

## 🎯 Summary

API ini menyediakan **15 Query operations** dan **15 Mutation operations** untuk mengelola:

- **Users** (CRUD + list)
- **Categories** (CRUD + list)
- **Resume Contents** (CRUD + list + filter by category)
- **Projects** (CRUD + list + filter by user)
- **Blogs** (CRUD + list + filter by status + get by slug)

Setiap operasi mendukung:
- ✅ **Pagination** dengan `page`, `pageSize`
- ✅ **Sorting** dengan `sortBy`, `sortDirection`
- ✅ **Proper error handling**
- ✅ **Input validation**
- ✅ **Relationship loading**
- ✅ **Type safety**

API siap digunakan untuk frontend portfolio application! 🚀