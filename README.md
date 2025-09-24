# Goath-stack-template

A modern web stack template powered by **Golang**, **Templ**, **TailwindCSS**, **Alpine.js**, and **HTMX**.

This project serves as a starter template that helps you quickly bootstrap your development.
Simply replace `goath-stack` with your own project name, and you’re ready to go.

---

## Features

- 🚀 **Golang + Templ**: Type-safe template engine with enhanced HTML componentization.
- 🎨 **TailwindCSS v4**: A utility-first and elegant CSS framework.
- ⚡ **Alpine.js**: Lightweight JavaScript for interactivity.
- 🔗 **HTMX**: Build interactive frontends with minimal JavaScript.
- 🐳 **Docker Support**: Works for both local development and containerized environments.

---

## Installation

### Install gonew

```bash
go install golang.org/x/tools/cmd/gonew@latest
```

### Create a new project with gonew

```bash
gonew github.com/dachichang/goath-stack-template your-project-name
```

# Project Initialization

1. Replace all occurrences of `goath-stack` in the project with your own project name.
2. Run `cp env.example .env` to create the environment configuration file.

---

# Local Development

## make run / make clean

- `make run` installs the required dependencies and uses **air** for automatic build of the project.
- The generated binary will be placed under the `tmp` directory.
- The default port is `3000`, but you can adjust it in the `.env` file.

## make up / make down

- `make up` starts the entire local development environment using **docker compose**.
- Everything runs fully inside Docker, ensuring a 100% clean environment.
- When you exit the container, all builds and `node_modules` will not be preserved.
- The default port is `8080`, but you can adjust it in the `docker-compose.yaml` file.
- It is recommended to use Docker for development to keep your local system clean.
