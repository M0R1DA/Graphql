# Application Flow Diagram

This diagram illustrates the complete flow of the Graphigl application, from startup to user interaction.

```mermaid
flowchart TD
    A[Start Application] --> B{main.js: Check JWT Token in localStorage}
    B -->|No Token| C[Auth(): Display Login Form]
    C --> D[User submits credentials]
    D --> E[POST /api/auth/signin]
    E -->|Success| F[Store JWT in localStorage]
    F --> G[HomePage()]
    B -->|Token Exists| G
    G --> H[home.js: fetchdata() from GraphQL API]
    H -->|Success| I[Build UI Components]
    I --> J[Display Dashboard: Header, User Data, XP Graph, Skills Chart]
    J --> K[User can Log Out]
    K --> L[logOut(): Clear localStorage, Reload Page]
    L --> B
    H -->|Failure| M[Redirect to Login]
    M --> C
    E -->|Failure| N[Show Error Message]
    N --> C
```

## Detailed Explanation

### 1. Start Application
The application initializes when the user opens `index.html` in their browser. This HTML file includes:
- Font Awesome CSS for icons
- The main stylesheet (`css/style.css`)
- The main JavaScript module (`js/main.js`)

The `main.js` file contains an event listener for `DOMContentLoaded`, which triggers the `main()` function.

### 2. Token Check
The `main()` function in `main.js` performs the first critical check:
- Retrieves the JWT token from `localStorage` using `localStorage.getItem("JWT_TOKEN")`
- **If no token exists**: Calls the `Auth()` function from `auth.js` to display the login interface
- **If a token exists**: Calls the `HomePage()` function from `home.js` to load the main dashboard

This check ensures that only authenticated users can access the application data.

### 3. Authentication Process
When authentication is required, the `Auth()` function:
- Creates a login form with input fields for username/email and password
- Adds form validation and submission handling
- On form submission, sends a POST request to `/api/auth/signin` with credentials encoded in the `Authorization` header
- **Success scenario**: 
  - Receives a JWT token from the server
  - Stores it in `localStorage`
  - Calls `HomePage()` to transition to the main application
- **Failure scenario**: 
  - Displays an error message (e.g., "Invalid Credentials")
  - Allows the user to retry login

The authentication endpoint is separate from the GraphQL API and uses standard HTTP authentication.

### 4. Home Page Loading
The `HomePage()` function in `home.js` handles the main application logic:
- Calls `fetchdata()` from `api.js` to retrieve user information
- The GraphQL query fetches comprehensive data including:
  - User profile (login, name, avatar, contact details, campus)
  - Audit ratio and current level
  - Total XP amount
  - XP transaction history for the progression graph
  - Skills data for the radar chart
- The query is sent to `https://learn.zone01oujda.ma/api/graphql-engine/v1/graphql`
- Includes the JWT token in the `Authorization` header for authentication
- **Success**: Proceeds to build the UI
- **Failure**: Redirects back to login (likely due to expired/invalid token)

### 5. UI Rendering
Using functions from `component.js`, the application dynamically builds the interface:
- **Header**: Created by `Header()` function, includes user avatar/login and navigation (Home, Profile, XP, Skills, Logout)
- **User Data Section**: Displays profile information, contact details, audit ratio, level, and total XP
- **XP Progression Graph**: Line chart showing XP growth over time using transaction data
- **Skills Radar Chart**: Visual representation of user proficiency in various skills
- **Footer**: Simple footer with copyright information

All components use CSS classes defined in their respective stylesheet files for styling.

### 6. Logout Process
The logout functionality:
- Triggered by clicking the "Log out" link in the header
- Calls the `logOut()` function from `auth.js`
- Clears the JWT token from `localStorage`
- Reloads the page, which restarts the application flow from the token check

This ensures a clean logout and prevents unauthorized access if the user leaves their session open.

## Additional Technical Details

- **State Management**: The application uses `localStorage` for simple client-side persistence
- **Error Handling**: Network failures or invalid tokens automatically redirect to login
- **Security**: All API calls include JWT authentication headers
- **Modularity**: Code is organized into separate modules (auth.js, home.js, api.js, component.js) for maintainability
- **Styling**: Modular CSS files allow for easy customization of different UI components

This comprehensive flow ensures a secure, user-friendly experience for visualizing GraphQL data.