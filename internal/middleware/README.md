| Concept                | Purpose                                                        |
| ---------------------- | -------------------------------------------------------------- |
| `responseWriter`       | A wrapper that records status code and bytes written           |
| `WriteHeader` override | Captures HTTP status                                           |
| `Write` override       | Captures body size and defaults to 200 OK                      |
| Used in                | Logging / metrics middleware                                   |
| Benefit                | Lets you log and monitor responses *without changing handlers* |

---

🗂️ It has two main files:

### 1️⃣ response_writer.go

**What it does:**
Keeps track of what your handler writes in the response.

**Why:**
The default `http.ResponseWriter` in Go doesn’t tell you what status code or how many bytes were sent.
This wrapper records:

* `status` → the HTTP code (like 200, 404, 500)
* `size` → how many bytes were written in the body

**So after a handler runs, middleware can log things like:**

> "200 OK – 52 bytes – 3ms"

---

### 2️⃣ logging.go

**What it does:**
Logs every incoming request and its outcome.

**How it works (step by step):**

* When a request comes in, it logs details like:
    > `method=POST path=/add/int remote=127.0.0.1`

* It wraps the normal `ResponseWriter` with our custom one (from `response_writer.go`).

* The request goes to your handler (like `/add/int`).

* After the handler finishes, it logs:
    * The status code (from the wrapper)
    * How many bytes were sent
    * How long the request took

So you get clean logs for every request without writing logging code in each handler.

---

### Request Flow

```text
Incoming Request
      ↓
Logging Middleware (records start time)
      ↓
responseWriter (tracks status & size)
      ↓
Your Handler (does the business logic)
      ↓
Logging Middleware (records end time, logs status, duration)
      ↓
Response Sent