# AI Agent Knowledge Transfer & Project Modifications Summary

*This file serves as a knowledge base and handover document for other AI agents working on the `SIGAP SULTAN` project.*

## 1. Project Overview & Architecture
- **Backend (Golang / Fiber):** Exposes APIs on port `8080`. Connects to a PostgreSQL database (`dev` schema by default, `sigap_sultan` DB).
- **Frontend (Next.js):** Runs on port `3000`. Uses Client-Side Fetching (Axios + React Query).
- **CMS (Laravel):** Runs on port `8000`. Connects to the same PostgreSQL database.

## 2. Recent Modifications & Bug Fixes

### A. Database Sync & NULL Pointers in Golang
- **Issue:** The database dump provided by the vendor had missing foreign keys or `NULL` values (e.g., `unit_id` in `tm_commodity`, `assets_relation_id` in `tm_city`). The original Golang code used `INNER JOIN` and strictly mapped to non-pointer Go types, causing `pgx.ErrNoRows` or panics when records didn't match.
- **Fix Applied:** 
  - Changed `INNER JOIN` to `LEFT JOIN` in multiple query files (e.g., `tm_city_queries.go`, `tm_commodity_queries.go`, `login_repository.go`).
  - Updated Go struct models (`models.TmCity`, `models.TmCommodity`) to use pointer types (e.g., `*int32`) so they can gracefully receive `NULL` from the database.

### B. Admin Authentication (Plaintext vs SHA256)
- **Issue:** The `/login` API hashes incoming passwords using `SHA256` before checking the `tm_user` table. However, the database dump had the admin password (`admin@sigapsultan.com`) stored in plain text (`"tes"`). This caused login to fail and return "Data is not found".
- **Fix Applied:** 
  - Manually hashed `"tes"` to its `SHA256` equivalent (`ce0f6c28b5869ff166714da5fe08554c70c731a335ff9702e38b00f81ad348c6`) and updated the database record for `admin@sigapsultan.com`.

### C. Fiber Panic Recovery Middleware
- **Issue:** When a panic occurred inside the Fiber app (due to unhandled data types or `nil` pointers), the entire Go process would crash and cause Docker containers to exit with code `255`, resulting in connection drops.
- **Fix Applied:** 
  - Added `recover.New()` middleware in `main.go` to ensure panics are caught and translated into `500 Internal Server Error` responses without killing the server.

### D. Frontend API Proxy (Next.js)
- **Issue:** `NEXT_PUBLIC_BASE_API_URL` was hardcoded to `http://localhost:8080` in `docker-compose.yml`. This caused network failures if the frontend was accessed from outside the host machine (e.g., via LAN or public IP).
- **Fix Applied:** 
  - Modified `docker-compose.yml` to set `NEXT_PUBLIC_BASE_API_URL=/api`.
  - Added an async `rewrites()` function in `sigap-sultan-fe/next.config.js` to proxy any request starting with `/api` to the backend container (`http://sigap_backend:8080`).

## 3. How to Work on This Repository
- **Docker Flow:** Always use `docker-compose up -d` in the root folder (`sigap_sultan`) to bring up all services.
- **Testing APIs:** You can hit `http://localhost:8080/city` or `http://localhost:8080/commodities` to verify backend data.
- **Frontend State:** Most data is public. Protected routes (like `/dashboard/profile`) rely on the `useAuth` hook and a JWT token.
- **Avoid Hardcoding Endpoints:** Use relative paths and Next.js rewrites when communicating with the backend.
