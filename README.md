# Focus

> [!CAUTION]
> Focus is in alpha development stage. It is not yet recommended for production use. Please report any bugs or issues you encounter.

Focus is an open-source, Kanban-style project management tool, emphasizing simplicity and efficiency. The backend is written in go and the frontend is in svelte.

The official source code repository is at [git.bhasher.com/Bhasher/focus](https://git.bhasher.com/Bhasher/focus). It runs the CI pipeline and hosts the Docker images and releases. However, a mirror is available at [github.com/Bhasher/focus](https://github.com/Bhasher/focus).

## Features

- Kanban-style boards
- RESTful API

### Planned Features

- Markdown support
- Real-time collaboration

## KISS Principles

Adhering to [KISS principles](https://en.wikipedia.org/wiki/KISS_principle), Focus boasts a minimalist, efficient codebase, concentrating on essential features. This streamlined approach minimizes complexity, enhancing maintainability and ease of deployment across different architectures.

Go was chosen for the backend because of its simplicity, speed, and cross-platform support. Svelte was chosen for the frontend because it can be compiled to static HTML, CSS, and JavaScript, which can easily be served by any web server.

Certain features are intentionally excluded, left to specialized tools. These include:

- Authentication
- HTTPS
- Notifications
- Mobile App

HTTPS and authentication have been successfully tested with [traefik](https://traefik.io/) and [authelia](https://www.authelia.com/) middleware but it should work with most reverse proxies and authentication middleware as it is platform-agnostic.

Focus offers a comprehensive RESTful API for extensions and integrations.

## Installation

### Docker

The easiest way to run Focus is with Docker. There is three images available:

- `git.bhasher.com/Bhasher/focus-frontend` for the svelte UI frontend
- `git.bhasher.com/Bhasher/focus-backend` for the go backend
- `git.bhasher.com/Bhasher/focus` for both in one image

Each image has a `latest` tag and a `vX.Y.Z` tag for each release. The `latest` tag is updated with each release.

Example `docker-compose.yaml` is available [here](docker-compose.yaml).

### Desktop App

Focus is available as a desktop app using [Tauri](https://tauri.app). Automatic releases are only available for Linux as binaries, AppImage, and DEB packages. Download the latest release from the [releases page](https://git.bhasher.com/Bhasher/focus/releases).

For other platforms, you can build the desktop app yourself using `npm run tauri build` in the `frontend` directory.

> [!CAUTION]
> The desktop app requires the backend to be running. You can specify the backend URL via the `PUBLIC_BACKEND_URL` environment variable. The default is `http://localhost:3000`.

```sh
PUBLIC_BACKEND_URL=http://localhost:3000 ./focus
```

#### AUR

```sh
yay -S focus-desktop-bin
```

### Debug Mode

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
