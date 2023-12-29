# Focus

> [!CAUTION]
> Focus is in alpha development stage. Expect breaking changes in every version.

Focus is an open-source, Kanban-style project management tool, emphasizing simplicity and efficiency. The backend is written in go and the frontend is in svelte.

## KISS Principles
Adhering to [KISS principles](https://en.wikipedia.org/wiki/KISS_principle), Focus boasts a minimalist, efficient codebase, concentrating on essential features. This streamlined approach minimizes complexity, enhancing maintainability and ease of deployment across different architectures.

Certain features are intentionally excluded, left to specialized tools. These include:
- Authentication
- HTTPS
- Notifications
- Mobile App

Focus offers a comprehensive RESTful API for extensions and integrations.

## Launching Focus in Debug Mode

To run Focus locally, follow these steps:
1. Clone the repository:
   ```sh
   git clone https://git.bhasher.com/Bhasher/focus
   ```

2. Start the backend:
   ```sh
   cd backend
   go run .
   ```

3. Start the frontend:
   ```sh
   cd frontend
   npm run dev
   ```

4. Access Focus at http://localhost:5000. The API is at http://localhost:3000.

## License

Focus is under the MIT License. For details, see the [LICENSE](license.md) file.