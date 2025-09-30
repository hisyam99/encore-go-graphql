# 🌐 CORS Configuration

Konfigurasi CORS yang clean, modular, dan mengikuti best practices untuk Portfolio GraphQL API.

## ✨ Features

- **Environment-based configuration**: Berbeda untuk development vs production
- **Modular design**: Terpisah dalam utility function
- **Secure defaults**: Hanya allow origins yang diperlukan
- **Easy to maintain**: Centralized configuration

## 🔧 Configuration

### Development Environment
```go
AllowedOrigins: []string{
    "http://localhost:3000",
    "http://localhost:3001", 
    "https://localhost:3000",
    "https://localhost:3001",
    "https://hisyam.tar.my.id",
    "http://hisyam.tar.my.id",
    "http://127.0.0.1:3000",
    "http://127.0.0.1:3001",
}
```

### Production Environment
```go
AllowedOrigins: []string{
    "https://hisyam.tar.my.id",
    "https://www.hisyam.tar.my.id",
}
```

## 🚀 Usage

CORS middleware sudah otomatis aktif di:
- `/graphql` - GraphQL endpoint
- `/graphql/playground` - GraphQL Playground (dev only)
- `/cors-test` - CORS test endpoint

## 🧪 Testing

### Test CORS Headers
```bash
# Test preflight request
curl -H "Origin: https://hisyam.tar.my.id" \
     -H "Access-Control-Request-Method: POST" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS \
     http://localhost:4001/graphql -v

# Test CORS endpoint
curl -H "Origin: https://hisyam.tar.my.id" \
     http://localhost:4001/cors-test
```

### Expected Response Headers
```
Access-Control-Allow-Origin: https://hisyam.tar.my.id
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Accept, Authorization, Content-Type, X-CSRF-Token, X-Requested-With
Access-Control-Allow-Credentials: true
Access-Control-Max-Age: 3600
```

## 🔒 Security

- **Origin validation**: Hanya origins yang ada di allowlist yang bisa akses
- **Production restrictive**: Production hanya allow domain production
- **Development permissive**: Development allow localhost dan development domains
- **Credentials support**: Support untuk authenticated requests

## 🛠️ Customization

Untuk menambah origin baru, edit `app/utils/cors.go`:

```go
// Development
AllowedOrigins: []string{
    // ... existing origins
    "http://localhost:3002",  // New origin
}

// Production  
AllowedOrigins: []string{
    // ... existing origins
    "https://admin.hisyam.tar.my.id",  // New origin
}
```

## ✅ Best Practices Applied

1. **Environment-based configuration** - Berbeda setting untuk dev/prod
2. **Principle of least privilege** - Hanya allow yang diperlukan
3. **Centralized configuration** - Satu tempat untuk manage CORS
4. **Clean separation** - CORS logic terpisah dari business logic
5. **Easy testing** - Dedicated endpoint untuk test CORS
6. **Proper error handling** - Handle preflight requests dengan benar

---

**Status**: ✅ CORS configured and working
**Last tested**: October 1, 2025
**Frontend domain**: https://hisyam.tar.my.id