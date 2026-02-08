# Implementation Plan: Phase 1 MVP

## Goal
Build a functional simplified clone of Jira with "Project" and "Issue" entities, capable of moving issues across statuses on a board.

## Tech Stack
-   **Frontend**: Next.js 14+ (App Router), TypeScript, Tailwind CSS, Shadcn/UI (if generic components needed).
-   **Backend**: Go 1.22+, `net/http` or `chi` router, `pgx` for Postgres.
-   **Database**: PostgreSQL 16.
-   **Infrastructure**: Docker Compose for local dev.

## Step-by-Step

### 1. Scaffolding (Current Task)
-   Create `jira-next-gen` root.
-   `frontend/`: Next.js app.
-   `backend/`: Go module.
-   `docker-compose.yml`: DB definition.

### 2. Backend Core
-   **Models**:
    -   `User` (ID, Email, Name)
    -   `Project` (ID, Key, Name)
    -   `Issue` (ID, Key, Summary, Status, Type)
-   **API**:
    -   `GET /api/projects`
    -   `POST /api/issues`
    -   `PATCH /api/issues/{id}` (Move status)

### 3. Frontend Core
-   **layouts**: AppShell with Sidebar.
-   **pages**:
    -   `/projects`: List of projects.
    -   `/projects/[key]/board`: Kanban board.
-   **components**:
    -   `IssueCard`: display summary, priority.
    -   `BoardColumn`: drag and drop zone (using `dnd-kit`).

## Validation
-   Can I create a project?
-   Can I create an issue?
-   Can I move an issue from "Todo" to "Done"?
