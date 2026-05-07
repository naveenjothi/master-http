# 📝 Build Your Own HTTP Server in Go

This project is about learning **how HTTP works on top of TCP** by implementing a minimal HTTP server from scratch and progressing towards a mini framework.

---

## ✅ Understanding

- Browser (client) wraps a request into **HTTP format** → sends over **TCP**.
- Server listens on a **TCP socket** → unwraps/parses the HTTP request.
- Routes the request → generates an HTTP response.
- Each client has its **own TCP connection** (concurrency with goroutines).
- Learnings: parsing overhead, latency, concurrency limits, framework design patterns.

---

## 🚀 Project Roadmap

### **Phase 1 – Barebones TCP Server**

- [x] Open a TCP socket on a given port (e.g., `:8080`).
- [x] Accept connections in a loop.
- [x] Spin off a goroutine per connection for concurrency.
- [x] Read raw bytes sent by client.
- [x] Send back a static response (`Hello World`).

---

### **Phase 1.5 – HTTP Fundamentals**

- [ ] Manually craft HTTP requests with `telnet localhost 8080`
- [ ] Study raw HTTP request/response format byte-by-byte
- [ ] Implement a simple "HTTP response builder" function
- [ ] Test with browsers vs curl (observe different headers)
- [ ] Understand HTTP/1.1 vs HTTP/1.0 differences

---

### **Phase 2 – Enhanced HTTP Parsing**

- [ ] Parse the **request line** → method, path, version (e.g., `GET /hello HTTP/1.1`).
- [ ] Parse **headers** into a map (handle edge cases: folded headers, case-insensitivity).
- [ ] **Parse query parameters** (`/search?q=golang&page=2`)
- [ ] **Handle malformed requests** gracefully (return `400 Bad Request`)
- [ ] **Parse cookies** from headers
- [ ] Ignore body for now (only GET requests).
- [ ] Respond with valid HTTP format (status line + headers + body).

---

### **Phase 3 – Basic Routing System**

- [ ] Build a simple router:
  - Map of `path → handler function`.
- [ ] Define a `HandlerFunc(w, r)` signature (like `net/http`).
- [ ] Support at least `/` and `/about`.
- [ ] If path not found → return `404 Not Found`.

---

### **Phase 3.5 – Advanced Routing**

- [ ] Support **path parameters** (`/users/:id`, `/posts/:slug`)
- [ ] **HTTP method-specific routing** (`GET /users` vs `POST /users`)
- [ ] **Route groups/prefixes** (`/api/v1/*`)
- [ ] **Wildcard routes** (`/files/*filepath`)
- [ ] Route priority and matching order

---

### **Phase 4 – Enhanced Connection Management**

- [ ] Confirm each client gets its own goroutine.
- [ ] Add logging of connection ID and request path.
- [ ] **Add timeout handling** (read/write timeouts)
- [ ] **Add rate limiting** per client IP
- [ ] **Graceful shutdown** (finish existing requests before closing)
- [ ] Measure throughput with multiple concurrent `curl` requests.
- [ ] Discuss limitations: memory use, goroutine overhead, parsing costs.

---

### **Phase 5 – Handling Request Body**

- [ ] Support POST requests with body.
- [ ] Parse `Content-Length` header to read exact body size.
- [ ] Handle `Transfer-Encoding: chunked`
- [ ] Echo back the request body (like a simple API).
- [ ] Support different content types (JSON, form data, multipart)

---

### **Phase 6 – HTTP Protocol Improvements**

- [ ] Add support for `Keep-Alive` (multiple requests on same connection).
- [ ] Add basic error handling (bad request → `400 Bad Request`).
- [ ] Return proper Content-Type (text/plain, application/json).
- [ ] Implement HTTP status codes properly (200, 201, 400, 404, 500)
- [ ] Add response compression (gzip)

---

### **Phase 6.5 – Framework Essentials**

- [ ] **Context object** (request ID, user data, cancellation, deadlines)
- [ ] **Response writer interface** (JSON, HTML, file responses)
- [ ] **Error handling patterns** (panic recovery, error middleware)
- [ ] **Request/response lifecycle hooks** (before/after middleware)
- [ ] **Middleware chain** implementation (logging, auth, CORS)

---

### **Phase 7 – Mini Framework Features**

- [ ] **Template rendering** (HTML templates with data binding)
- [ ] **JSON API helpers** (bind request JSON, respond with JSON)
- [ ] **Session management** (in-memory, cookie-based)
- [ ] **CORS handling** (preflight requests, headers)
- [ ] **Request validation** (required fields, types, custom validators)
- [ ] **Static file serving** with caching headers
- [ ] **Database integration patterns** (connection per request, pooling)
- [ ] **Authentication middleware** (JWT, session-based)

---

### **Phase 8 – Performance & Production**

- [ ] **Memory profiling** (heap analysis, goroutine leak detection)
- [ ] **CPU profiling** under load testing
- [ ] **Load testing** with realistic payloads using `ab`, `wrk`, or `hey`
- [ ] **Optimization techniques**: string pooling, buffer reuse, object pooling
- [ ] **Benchmarking** against popular frameworks (Gin, Echo, Fiber)
- [ ] **Security hardening**: input sanitization, CSRF protection, headers
- [ ] **Observability**: metrics (requests/sec, latency percentiles), structured logging
- [ ] **Configuration management** (environment variables, config files)

---

## 🎯 Learning Goals

By completing all phases, you'll:

1. **Understand HTTP internals**: How browsers and servers communicate at the protocol level
2. **Master Go networking**: TCP sockets, goroutines, and concurrent programming patterns
3. **Build a mini framework**: Core concepts used in production frameworks like Gin, Echo, Fiber
4. **Performance awareness**: Bottlenecks, optimization techniques, and scaling considerations
5. **Production readiness**: Error handling, logging, monitoring, and security basics

## 📚 Additional Learning Resources

- **Study existing frameworks**: Examine Gin's routing, Echo's context handling, Fiber's performance optimizations
- **Protocol deep-dive**: Implement HTTP/2 features, WebSocket upgrades
- **Architecture patterns**: Middleware chains, dependency injection, plugin systems
- **Testing strategies**: Unit tests, integration tests, load testing, chaos engineering

## 🔧 Implementation Tips

1. Start simple and iterate - don't try to implement everything at once
2. Write tests early, especially for parsing logic
3. Use Go's built-in profiling tools (`go tool pprof`)
4. Read the HTTP/1.1 RFC for edge cases
5. Study how popular Go web frameworks solve common problems
6. Measure performance at each phase to understand trade-offs

---

**Goal**: Build a production-ready HTTP server and mini framework while gaining deep insights into web server internals, performance optimization, and modern web development patterns.
