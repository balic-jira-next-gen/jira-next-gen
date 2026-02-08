# Next-Gen Work Management Platform Specification

## 1. Product Vision & Personas

**Vision:** "Enterprise Grade Power with Consumer Grade Speed."
To build a platform that scales to millions of users without the "Jira Tax" – eliminating slowness, complexity bloat, and administrative overhead. A platform where developers love the speed, and executives trust the data.

**Core Value Proposition:**
*   **For Devs:** Keyboard-first, instant UI, deeply integrated with Git/CI. "Flow state" uninterrupted.
*   **For ITSM:** AI-first service desk that deflects 50% of tickets automatically.
*   **For Execs:** Real-time portfolio visibility without manual aggregation.
*   **For Admins:** "Configuration as Code" options, comprehensive audit capability, and granular permissions without the headache.

**Target Personas:**
1.  **Software Teams:** Require high-velocity backlogs, sprint planning, and deep code integration.
2.  **ITSM/Ops:** Require SLA tracking, incident management swarming, and asset linking.
3.  **Business Teams (HR/Legal/Marketing):** Require simplified "Trello-like" views with powerful workflow backbones.
4.  **Portfolio Managers:** require cross-project "Initiative" tracking and resource capacity planning.

## 2. Core Functional Modules (Feature Parity)

### Issue & Work Item Management
*   **Flexible Entity Model:** Everything is an `Issue` but with strict typing (Epic, Story, Task, Subtask + Custom Types).
*   **Hierarchy:** Infinite nesting supported (Initiative -> Theme -> Epic -> Story -> Task -> Subtask).
*   **Fields:** System fields (Summary, Desc, Assignee, Priority) + Custom Fields (Text, Number, Date, User, Select, Multi-Select, JSON, Lookups).
*   **Rich Text:** Markdown-first WYSIWYG editor with strict schema (ProseMirror based) for consistent rendering and API parsability.

### Agile & Project Management
*   **Boards:** Unified Board View. Toggle between Kanban, Scrum, and List view instantly.
*   **Sprints:** Multi-sprint support (parallel sprints).
*   **Roadmaps:** GANTT-style portfolio views with drag-and-drop dependency management.
*   **Estimation:** Story points, time tracking (original vs remaining), and T-shirt sizing.

### Workflow Engine
*   **State Machine:** Finite State Machine (FSM) core.
*   **visual Editor:** Drag-and-drop workflow builder.
*   **Transition Rules:**
    *   *Conditions:* "User must be assignee", "Field X must be set".
    *   *Validators:* "Check permissions", "Check linked PR status".
    *   *Post-Functions:* "Update field", "Fire webhook", "Trigger automation".

### Search (JQL Equivalent)
*   **AQL (Advanced Query Language):** SQL-like syntax. `project = "ABC" AND status in ("Open", "In Progress") AND assignee = currentUser() ORDER BY rank ASC`.
*   **Performance:** Queries return in < 50ms for 99% of requests.

## 3. ITSM & Service Management

*   **Portal:** Customer-facing, simplified portal. Highly customizable branding.
*   **Request Types:** Maps friendly user requests ("I need a laptop") to complex backend issues.
*   **SLAs:** Measuring "Time to First Response", "Time to Resolution". Calendar-aware (business hours vs 24/7).
*   **Incident Management:** Major Incident workbench, integration with PagerDuty/OpsGenie.
*   **CMDB:** Basic asset management with graph links to Issues.

## 4. Reporting & Analytics

*   **Real-time Dashboards:** Widgets for "Assigned to Me", "Sprint Burndown", "Velocity", "Created vs Resolved".
*   **Custom Report Builder:** Drag-and-drop BI-lite interface.
*   **Export:** PDF, CSV, and "Connect to Excel/Tableau" via OData or specialized API.

## 5. User, Org & Permission Model

*   **Identity:** Users belong to an Organization.
*   **Groups:** Standard groups (Admins, Developers).
*   **Roles:** Project-level roles (Project Admin, Member, Viewer).
*   **Schemes:** Permission Schemes, Notification Schemes, Field Security Schemes (reuse across projects).
*   **Hierarchy:** `User < Group < Project Role < Global Permission`.

## 6. Enterprise NFRs

*   **Security:** SOC2 Type II, GDPR, HIPAA ready.
*   **Auth:** SAML 2.0, OIDC, SCIM for user provisioning.
*   **Scalability:** Tested to 10M+ issues per instance.
*   **Tenancy:** Logical separation (SaaS) with options for "Dedicated Shard" (Enterprise).

## 7. Integrations & Ecosystem

*   **Connect Framework:** iframe-based UI extensions (secure, sandboxed) + Server-side webhooks.
*   **REST API:** Comprehensive, versioned (v1, v2).
*   **GraphQL API:** For high-efficiency frontend fetching.
*   **Marketplace:** Billing engine for 3rd party apps.

## 8. Architecture & Tech Stack

### High-Level Architecture Diagram
```mermaid
graph TD
    UserBrowser[User Browser / Mobile App]
    CDN[Global CDN (Cloudflare)]
    LB[Load Balancers]
    
    subgraph "Frontend Layer"
        SPA[Next.js App (SSR/CSR)]
    end
    
    subgraph "API Gateway Layer"
        Gateway[API Gateway / GraphQL Federation]
        AuthSvc[Auth Service (OIDC/SAML)]
    end
    
    subgraph "Core Microservices"
        IssueSvc[Issue Service (Go)]
        WorkflowSvc[Workflow Engine (Go)]
        ProjectSvc[Project Service (Node.js)]
        UserSvc[User/Org Service (Go)]
    end
    
    subgraph "Data & specialized Services"
        SearchSvc[Search Service (Rust/Java)]
        AnalyticsSvc[Analytics/Reporting (Python)]
        RealtimeSvc[WebSocket/PubSub (Go)]
    end
    
    subgraph "Storage Layer"
        PrimaryDB[(PostgreSQL - Sharded)]
        SearchIndex[(Elasticsearch / OpenSearch)]
        AnalyticsDB[(ClickHouse)]
        Cache[(Redis Cluster)]
        ObjectStore[(S3 - Attachments)]
    end
    
    UserBrowser --> CDN
    CDN --> LB
    LB --> SPA
    SPA --> Gateway
    Gateway --> AuthSvc
    Gateway --> IssueSvc
    Gateway --> WorkflowSvc
    Gateway --> SearchSvc
    
    IssueSvc --> PrimaryDB
    IssueSvc --> RealtimeSvc
    IssueSvc --> SearchIndex
    
    AnalyticsSvc --> AnalyticsDB
```

### Technology Decisions

*   **Frontend:** **Next.js (React)**.
    *   *Rationale:* pervasive ecosystem, excellent performance optimization (Server Components), great developer experience.
*   **Backend Core:** **Go (Golang)**.
    *   *Rationale:* High concurrency, low memory footprint, strict typing alongside great JSON handling. Perfect for high-throughput generic services like "Issue Tracking".
*   **Database (OLTP):** **PostgreSQL**.
    *   *Rationale:* Rock solid, JSONB support allows "Schema-less" custom fields without NoSQL trade-offs.
*   **Search:** **Elasticsearch** (or OpenSearch).
    *   *Rationale:* Industry standard for full-text search and complex filtering (aggregations).
*   **Analytics:** **ClickHouse**.
    *   *Rationale:* Columnar store for incredibly fast aggregation queries ("Average time in status 'In Progress' for last 6 months").
*   **Real-time:** **NATS** or **Redis Pub/Sub** with WebSocket frontend.

## 9. Data Model & Entities

### Core Schemas (Simplified)

**Global/Tenant Level:**
*   `Organization`: { id, name, plan, settings }
*   `User`: { id, email, password_hash, profile }
*   `Group`: { id, name, members[] }

**Project Level:**
*   `Project`: { id, key (e.g. "PROJ"), name, lead_user_id, workflow_scheme_id }
*   `Issue`: { 
      id, 
      project_id, 
      key ("PROJ-123"), 
      summary, 
      description (ProseMirror JSON), 
      type_id, 
      status_id, 
      assignee_id, 
      priority_id, 
      custom_fields: JSONB, 
      created_at, 
      updated_at 
    }
*   `Comment`: { id, issue_id, body, author_id }
*   `WorkLog`: { id, issue_id, time_spent_seconds, date }

**Configuration:**
*   `Workflow`: { id, nodes (statuses), edges (transitions) }
*   `FieldConfig`: { id, name, type, validation_rules }

## 10. UX / UI Expectations

*   **Keyboard First:** `c` to create, `/` to search, `.` to open command palette.
*   **Optimistic UI:** UI updates immediately, syncs in background. Reverts on error.
*   **Density:** "Compact" vs "Comfortable" modes. Engineers typically prefer high information density.
*   **Command Palette:** `Ctrl+K` for everything. "Assign to John", "Move to Done".

## 11. Migration & Interoperability

*   **Jira Importer:**
    *   Direct API-to-API connector (OAuth to Jira Cloud).
    *   XML Backup parser for Data Center migrations.
*   **CSV Import:** robust mapping tool.

## 12. AI & Automation

*   **Copilot:** "Summarize this thread", "Draft a release note based on these 5 tickets".
*   **Auto-Triage:** Suggest assignee and component based on history.
*   **NLP to JQL:** "Show me all high priority bugs from last week" -> `priority = High AND type = Bug AND created > -1w`.

## 13. Delivery Roadmap

**Phase 1: MVP (Months 1-4)**
*   Core Issue Tracking (Values, Statuses).
*   Basic Workflow (Open -> Done).
*   Projects & Users.
*   Next.js Frontend + Go Backend + Postgres.

**Phase 2: Agile Module (Months 5-8)**
*   Scrum/Kanban Boards.
*   Sprints & Backlogs.
*   Drag and Drop prioritization.

**Phase 3: Deep Integrations & Search (Months 9-12)**
*   Elasticsearch integration.
*   GitHub/GitLab sync.
*   Basic Reporting.

**Phase 4: Enterprise & ITSM (Year 2)**
*   SLA Engine.
*   Service Desk Portal.
*   SSO/SCIM.
