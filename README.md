# Portfolio Backend API

A comprehensive GraphQL-based backend API for a portfolio website built with Encore.dev, featuring user management, project showcase, blog posts, and resume sections.

## 🚀 Features

- **User Management**: Create, read, update, and delete user profiles
- **Project Portfolio**: Manage and showcase personal projects with descriptions
- **Blog System**: Create and manage blog posts with content
- **Resume Management**: Organize resume sections by categories
- **GraphQL API**: Modern, type-safe API with automatic schema generation
- **Database Migrations**: Automated database schema management with Atlas
- **PostgreSQL Integration**: Robust database support with GORM ORM

## 🛠️ Technology Stack

- **Backend Framework**: [Encore.dev](https://encore.dev) - Modern Go backend framework
- **API**: GraphQL with [gqlgen](https://gqlgen.dev) for code generation
- **Database**: PostgreSQL with GORM ORM
- **Migrations**: Atlas for database schema management
- **Language**: Go 1.24.6

## 📋 Prerequisites

**Install Encore:**
- **macOS:** `brew install encoredev/tap/encore`
- **Linux:** `curl -L https://encore.dev/install.sh | bash`
- **Windows:** `iwr https://encore.dev/install.ps1 | iex`
  
**Docker:**
1. Install [Docker](https://docker.com)
2. Start Docker

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <your-repo-url>
cd backend
```

### 2. Run Locally

Before running your application, make sure you have Docker installed and running. Then run this command from your application's root folder:

```bash
encore run
```

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to view Encore's [local developer dashboard](https://encore.dev/docs/go/observability/dev-dash).

### 3. Access GraphQL Playground

Open [http://localhost:4000/graphql/playground](http://localhost:4000/graphql/playground) in your browser to interact with the API.

## 📊 Database Schema

The application uses the following main entities:

### Users
- `id`: Primary key
- `name`: User's full name
- `email`: Unique email address
- `created_at`: Timestamp

### Projects
- `id`: Primary key
- `title`: Project title
- `description`: Project description
- `user_id`: Foreign key to users table

### Blogs
- `id`: Primary key
- `title`: Blog post title
- `content`: Blog post content
- `created_at`: Timestamp

### Resumes
- `id`: Primary key
- `title`: Resume section title
- `description`: Resume section description
- `category`: Resume category (e.g., "Experience", "Education", "Skills")

## 🔌 GraphQL API Reference

### Queries

#### Get All Users
```graphql
query {
  users {
    id
    name
    email
    createdAt
    projects {
      id
      title
      description
    }
  }
}
```

#### Get User by ID
```graphql
query {
  user(id: "1") {
    id
    name
    email
    createdAt
    projects {
      id
      title
      description
    }
  }
}
```

#### Get All Projects
```graphql
query {
  projects {
    id
    title
    description
    userID
    user {
      id
      name
      email
    }
  }
}
```

#### Get All Blogs
```graphql
query {
  blogs {
    id
    title
    content
    createdAt
  }
}
```

#### Get All Resumes
```graphql
query {
  resumes {
    id
    title
    description
    category
  }
}
```

### Mutations

#### Create User
```graphql
mutation {
  createUser(input: {
    name: "John Doe"
    email: "john@example.com"
  }) {
    id
    name
    email
    createdAt
  }
}
```

#### Create Project
```graphql
mutation {
  createProject(input: {
    title: "My Awesome Project"
    description: "A detailed description of the project"
    userID: "1"
  }) {
    id
    title
    description
    userID
  }
}
```

#### Create Blog Post
```graphql
mutation {
  createBlog(input: {
    title: "My First Blog Post"
    content: "This is the content of my blog post..."
  }) {
    id
    title
    content
    createdAt
  }
}
```

#### Create Resume Section
```graphql
mutation {
  createResume(input: {
    title: "Software Engineer"
    description: "Experienced software engineer with 5+ years..."
    category: "Experience"
  }) {
    id
    title
    description
    category
  }
}
```

#### Update Operations
All entities support update operations with partial data:

```graphql
mutation {
  updateUser(id: "1", input: {
    name: "Updated Name"
  }) {
    id
    name
    email
  }
}
```

#### Delete Operations
All entities support delete operations:

```graphql
mutation {
  deleteUser(id: "1")
}
```

## 🗄️ Database Migrations

The project uses Atlas for database migrations. Migrations are located in `app/migrations/`.

### Generate New Migration
```bash
# After modifying models in app/models.go
encore db migrate
```

### Apply Migrations
```bash
encore db migrate
```

## 🧪 Testing

Run the test suite:

```bash
encore test ./...
```

## 📁 Project Structure

```
backend/
├── app/
│   ├── app.go              # Main application setup and database configuration
│   ├── models.go           # GORM models (User, Project, Blog, Resume)
│   ├── migrations/         # Database migration files
│   └── scripts/           # Utility scripts
├── graphql/
│   ├── app.graphqls       # GraphQL schema definition
│   ├── app.resolvers.go   # GraphQL resolvers implementation
│   ├── generated/         # Auto-generated GraphQL code
│   ├── model/             # Generated models
│   └── service.go         # GraphQL service setup
├── gqlgen.yml             # GraphQL code generation configuration
├── go.mod                 # Go module dependencies
├── go.sum                 # Go module checksums
└── encore.app             # Encore application configuration
```

## 🚀 Deployment

### Self-hosting

See the [self-hosting instructions](https://encore.dev/docs/go/self-host/docker-build) for how to use `encore build docker` to create a Docker image and configure it.

### Encore Cloud Platform

Deploy your application to a free staging environment in Encore's development cloud using `git push encore`:

```bash
git add -A .
git commit -m 'Commit message'
git push encore
```

You can also open your app in the [Cloud Dashboard](https://app.encore.dev) to integrate with GitHub, or connect your AWS/GCP account, enabling Encore to automatically handle cloud deployments for you.

## 🔗 GitHub Integration

Follow these steps to link your app to GitHub:

1. Create a GitHub repo, commit and push the app.
2. Open your app in the [Cloud Dashboard](https://app.encore.dev).
3. Go to **Settings ➔ GitHub** and click on **Link app to GitHub** to link your app to GitHub and select the repo you just created.
4. To configure Encore to automatically trigger deploys when you push to a specific branch name, go to the **Overview** page for your intended environment. Click on **Settings** and then in the section **Branch Push** configure the **Branch name** and hit **Save**.
5. Commit and push a change to GitHub to trigger a deploy.

[Learn more in the docs](https://encore.dev/docs/platform/integrations/github)

## 📝 Development

### Code Generation

The GraphQL code is auto-generated using gqlgen. When you modify the schema in `graphql/app.graphqls`, regenerate the code:

```bash
go generate ./...
```

### Adding New Features

1. Update the GraphQL schema in `graphql/app.graphqls`
2. Add corresponding models in `app/models.go` if needed
3. Run `go generate ./...` to regenerate GraphQL code
4. Implement resolvers in `graphql/app.resolvers.go`
5. Test your changes using the GraphQL playground

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

If you encounter any issues or have questions:

1. Check the [Encore.dev documentation](https://encore.dev/docs)
2. Review the [GraphQL documentation](https://graphql.org/learn/)
3. Open an issue in this repository

---

Built with ❤️ using [Encore.dev](https://encore.dev) and [GraphQL](https://graphql.org/)