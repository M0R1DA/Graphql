┌─────────────────────────────────────────────────────────────────────────────┐
│                               FRONTEND (Client)                             │
│  HTML / CSS / JS                                                            │
│  (index.html, main.js, auth.js, home.js, api.js, component.js)              │
│                                                                             │
│   ┌──────────────────────────────────────────────────────────────────────┐  │
│   │                            Browser UI Flow                          │  │
│   │                                                                      │  │
│   │  • User loads app → main.js → Check JWT token                       │  │
│   │  • No token → auth.js → Display login form                          │  │
│   │  • User logs in → POST /api/auth/signin                             │  │
│   │  • Token stored → home.js → Fetch GraphQL data                      │  │
│   │  • Display dashboard → Header, User Data, XP Graph, Skills Chart    │  │
│   │  • User logs out → Clear localStorage, reload                       │  │
│   │                                                                      │  │
│   └──────────────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────────────┘

                                      │
                                      ▼  (HTTP Requests)

┌─────────────────────────────────────────────────────────────────────────────┐
│                              GRAPHQL API BACKEND                            │
│                    (https://learn.zone01oujda.ma/api/graphql-engine/v1/graphql) │
├─────────────────────────────────────────────────────────────────────────────┤
│                                 ENDPOINTS                                    │
│                                                                              │
│   • /api/auth/signin    → JWT Authentication                                │
│   • /api/graphql-engine/v1/graphql → GraphQL Query Endpoint                 │
│                                                                              │
│   GraphQL Schema includes:                                                   │
│   - User profile data (login, name, avatar, contact)                        │
│   - Audit ratio, level, total XP                                            │
│   - XP transactions (for progression graph)                                 │
│   - Skills data (for radar chart)                                           │
└─────────────────────────────────────────────────────────────────────────────┘

                                      │
                                      ▼

┌─────────────────────────────────────────────────────────────────────────────┐
│                            AUTHENTICATION LAYER                              │
│                                                                              │
│   • JWT Token Validation → Required for GraphQL queries                    │
│   • Session Management → localStorage in browser                           │
│   • Logout → Clear token and reload                                        │
└─────────────────────────────────────────────────────────────────────────────┘

                                      │
                                      ▼

┌─────────────────────────────────────────────────────────────────────────────┐
│                         APPLICATION BUSINESS LOGIC                           │
│                       (JavaScript Modules)                                  │
│                                                                              │
│   main.js                                                                   │
│     • Entry point, token check, route to auth or home                      │
│                                                                              │
│   auth.js                                                                   │
│     • Login form handling, API authentication                              │
│                                                                              │
│   home.js                                                                   │
│     • Dashboard rendering, component assembly                              │
│                                                                              │
│   api.js                                                                    │
│     • GraphQL query construction and execution                            │
│                                                                              │
│   component.js                                                              │
│     • UI component creation (Header, UserData, XPGraph, SkillsChart)       │
└─────────────────────────────────────────────────────────────────────────────┘

                                      │
                                      ▼

┌─────────────────────────────────────────────────────────────────────────────┐
│                              DATA VISUALIZATION                              │
│                                                                              │
│   • User Data Section → Profile information display                         │
│   • XP Progression Graph → Line chart of XP over time                      │
│   • Skills Radar Chart → Proficiency visualization                         │
│   • Header Component → Navigation and user info                            │
│   • Responsive Design → CSS modules for each component                     │
└─────────────────────────────────────────────────────────────────────────────┘


# **GRAPHQL DATA FETCHING FLOW**

sequenceDiagram
    participant User as User
    participant Frontend as Frontend App
    participant AuthAPI as Auth API
    participant GraphQL as GraphQL API
    participant DB as Backend Database

    User->>Frontend: Load application
    Frontend->>Frontend: Check localStorage for JWT
    alt No JWT Token
        Frontend->>AuthAPI: POST /api/auth/signin (credentials)
        AuthAPI->>DB: Validate user credentials
        DB-->>AuthAPI: User data + JWT token
        AuthAPI-->>Frontend: JWT token
        Frontend->>Frontend: Store JWT in localStorage
    end
    Frontend->>GraphQL: POST /graphql (query + JWT header)
    GraphQL->>DB: Execute GraphQL query
    DB-->>GraphQL: User data, XP, skills
    GraphQL-->>Frontend: JSON response
    Frontend->>Frontend: Render dashboard with charts