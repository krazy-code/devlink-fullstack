---
applyTo: '**'
---

## Project Context

Devlink is a developer community social media website. It is a full stack project, front-end code is inside client folder built with Vite, React, and Typescript, and back-end service is inside server folder built with golang, go fiber and using databases postgres or supabase depends on env config database_name.

## Frontend (./client) Coding Standards

- All new code should adhere to default vite eslint.
- Use PascalCase for pages file name and use kebab-case for other file name.
- Write clear and concise comments for complex logic.

## Backend (./server) Coding Standards

- Use snake_case for all file name.
- Write clear and concise comments for complex logic.
- Write the route of the services inside controllers folder and file (use template like developer system)

## Code Review Focus

- Prioritize security vulnerabilities and performance optimizations.
- Check for adherence to architectural patterns and design principles.
