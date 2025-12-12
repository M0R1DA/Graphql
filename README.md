# Graphigl

A frontend application for GraphQL data visualization, built with vanilla JavaScript, HTML, and CSS. This project provides an interactive dashboard to visualize user data fetched from a GraphQL API, including XP progression, skills radar chart, and user information.

## Features

- **Authentication**: Secure login using JWT tokens stored in localStorage
- **User Dashboard**: Displays user profile information, audit ratio, level, and total XP
- **XP Progression Graph**: Line chart showing XP gained over time
- **Skills Radar Chart**: Visual representation of user skills proficiency
- **Responsive Design**: Mobile-friendly interface with modular CSS styling

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/M0R1DA/Graphql.git
   cd Graphql
   ```

2. No additional dependencies are required as the project uses vanilla JavaScript and includes Font Awesome via CDN.

## Usage

1. Start the development server:
   ```bash
   npm start
   ```
   or
   ```bash
   python3 -m http.server 8000
   ```

2. Open your browser and navigate to `http://localhost:8000`

3. Log in with your credentials to access the dashboard

## Project Structure

- `index.html` - Main HTML file
- `js/` - JavaScript modules
  - `main.js` - Application entry point
  - `auth.js` - Authentication handling
  - `home.js` - Home page rendering
  - `api.js` - GraphQL API interactions
  - `component.js` - UI component functions
- `css/` - Stylesheets for different components
  - `style.css` - General styles
  - `auth.css` - Login form styles
  - `header.css` - Header styles
  - `footer.css` - Footer styles
  - `sectiondata.css` - User data section styles
  - `xpgraph.css` - XP graph styles
  - `radargraph.css` - Skills radar chart styles

## Application Flow

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

## Technologies Used

- JavaScript (ES6 Modules)
- HTML5
- CSS3
- GraphQL API
- Font Awesome (via CDN)

## API Endpoint

The application fetches data from: `https://learn.zone01oujda.ma/api/graphql-engine/v1/graphql`

## License

ISC

## Author

M0R1DA