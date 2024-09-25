# Gin CLI - Scaffold Your Gin Application

![Gin CLI Logo](path/to/logo.png) <!-- Optional: Add a logo image -->

Gin CLI is a powerful command-line tool designed to help you quickly set up [Gin Web Framework](https://gin-gonic.com/) applications with a predefined project structure and essential components. Focus on building your application without worrying about the initial setup!

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
    - [Available Commands](#available-commands)
    - [Command Aliases](#command-aliases)
4. [Commands Overview](#commands-overview)
    - [`new`](#new)
    - [`generate`](#generate)
        - [`generate controller`](#generate-controller)
        - [`generate guard`](#generate-guard)
        - [`generate resource`](#generate-resource)
        - [`generate route`](#generate-route)
        - [`generate service`](#generate-service)
    - [`template`](#template)
5. [Examples](#examples)
6. [Running the Application](#running-the-application)
7. [Notes](#notes)
8. [Support](#support)

---

## Features

- **Quick Scaffold**: Generate a new Gin application with a standard project layout.
- **Modular Generation**: Create controllers, guards, resources, routes, and services effortlessly.
- **Database Configuration**: Supports MySQL, PostgreSQL, SQLite, and MongoDB.
- **User-Friendly**: Interactive prompts with default values for an enhanced user experience.
- **Colorful Logs**: Distinctive colors for informational, success, warning, and error messages.
- **Shorthand Flags**: Use single-letter flags for faster command inputs.
- **Safe File Generation**: Prompts before overwriting existing files to prevent data loss.
- **Enhanced Help**: Displays commands and flags in a neatly formatted table similar to NestJS CLI.

## Installation

### Prerequisites

- **Go**: Ensure you have Go installed (version 1.16 or later). [Download Go](https://golang.org/dl/)
- **Git**: Required to clone the repository. [Download Git](https://git-scm.com/downloads)

### Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/golang-programming/gincli.git
   cd gincli
   ```

2. **Build the CLI Tool**

   ```bash
   go build -o gin
   ```

3. **Move the Executable to Your PATH**

   ```bash
   mv gin /usr/local/bin/
   ```

   *Note: You might need `sudo` permissions depending on your system.*

4. **Verify Installation**

   ```bash
   gin --version
   ```

   You should see output indicating the installed version of Gin CLI.

## Usage

Generate and manage your Gin applications seamlessly using the following commands.

### Available Commands

| Command       | Description                                           |
|---------------|-------------------------------------------------------|
| `new`         | Create a new Gin application with a project structure |
| `generate`    | Generate resources, controllers, guards, routes, and services |
| `template`    | Load application startup templates                    |

### Command Aliases

Aliases allow you to use shorter or alternative names for commands.

| Command    | Aliases        |
|------------|----------------|
| `new`      | `n`, `create`  |
| `generate` | `g`            |
| `template` | `t`            |

#### Subcommand Aliases for `generate`

| Subcommand       | Aliases |
|------------------|---------|
| `generate controller` | `c` |
| `generate guard`      | `gd` |
| `generate resource`   | `r` |
| `generate route`      | `rt` |
| `generate service`    | `s` |

*Note: Subcommand aliases are prefixed with `generate`. For example, `c` is an alias for `generate controller`.*

---

## Commands Overview

### `new`

**Description**: Create a new Gin application with a predefined project structure.

**Aliases**: `n`, `create`

**Usage**:

```bash
gin new [flags]
```

**Flags**:

| Flag                      | Shorthand | Description                                   | Default                                   |
|---------------------------|-----------|-----------------------------------------------|-------------------------------------------|
| `--app-name`              | `-a`      | Name of your application                      | `my-gin-app`                              |
| `--db-type`               | `-d`      | Database type (`MySQL`, `PostgreSQL`, `SQLite`, `MongoDB`) | `MySQL`                              |
| `--db-connection-string`  | `-c`      | Custom database connection string             | *(None)*                                  |
| `--db-host`               | `-H`      | Database host                                 | `localhost`                               |
| `--db-name`               | `-n`      | Database name                                 | `default`                                 |
| `--db-username`           | `-u`      | Database username                             | `root`                                    |
| `--db-password`           | `-p`      | Database password                             | `password`                                |
| `--db-port`               | `-P`      | Database port (default varies by DB type)     | `3306` (MySQL), `5432` (PostgreSQL)       |
| `--yes`                   | `-y`      | Skip all prompts and use default values       | `false`                                   |

---

### `generate`

**Description**: Generate various components such as controllers, guards, resources, routes, and services to build a modular application structure.

**Aliases**: `g`

**Usage**:

```bash
gin generate [subcommand] [flags]
```

**Available Subcommands**:

| Subcommand       | Description                  |
|------------------|------------------------------|
| `controller`     | Generate a new controller    |
| `guard`          | Generate a new guard         |
| `resource`       | Generate a new resource      |
| `route`          | Generate a new route         |
| `service`        | Generate a new service       |

#### Subcommand Flags

Each subcommand may have its own specific flags. Refer to individual subcommand documentation for details.

---

#### `generate controller`

**Description**: Generate a new controller for handling HTTP requests.

**Aliases**: `c`

**Usage**:

```bash
gin generate controller <name> [path] [flags]
```

**Arguments**:

- `<name>`: Name of the controller.
- `[path]`: (Optional) Custom path where the controller should be generated.

**Example**:

```bash
gin g c user
```

---

#### `generate guard`

**Description**: Generate a new guard for implementing authentication or authorization logic.

**Aliases**: `gd`

**Usage**:

```bash
gin generate guard <name> [path] [flags]
```

**Arguments**:

- `<name>`: Name of the guard.
- `[path]`: (Optional) Custom path where the guard should be generated.

**Example**:

```bash
gin g gd auth
```

---

#### `generate resource`

**Description**: Generate a new resource with predefined components like controllers, DTOs, entities, and services.

**Aliases**: `r`

**Usage**:

```bash
gin generate resource <name> [path] [flags]
```

**Arguments**:

- `<name>`: Name of the resource.
- `[path]`: (Optional) Custom path where the resource should be generated.

**Flags**:

| Flag          | Shorthand | Description                        | Default    |
|---------------|-----------|------------------------------------|------------|
| `--transport` | `-t`      | Transport layer (`Restful`, `WebSockets`) | `Restful` |

**Example**:

```bash
gin g r product -t Restful
```

---

#### `generate route`

**Description**: Generate a new route for your application.

**Aliases**: `rt`

**Usage**:

```bash
gin generate route <name> [path] [flags]
```

**Arguments**:

- `<name>`: Name of the route.
- `[path]`: (Optional) Custom path where the route should be generated.

**Example**:

```bash
gin g rt user
```

---

#### `generate service`

**Description**: Generate a new service layer for business logic.

**Aliases**: `s`

**Usage**:

```bash
gin generate service <name> [path] [flags]
```

**Arguments**:

- `<name>`: Name of the service.
- `[path]`: (Optional) Custom path where the service should be generated.

**Example**:

```bash
gin g s payment
```

---

### `template`

**Description**: Load application startup templates to set up different project configurations.

**Aliases**: `t`

**Usage**:

```bash
gin template [flags]
```

**Flags**:

| Flag                      | Shorthand | Description                                   | Default        |
|---------------------------|-----------|-----------------------------------------------|----------------|
| `--template`              | `-t`      | Template to use (`Standard`)                  | `Standard`     |
| `--app-name`              | `-a`      | Name of your application                      | `my-gin-app`   |
| `--db-type`               | `-d`      | Database type (`MySQL`, `PostgreSQL`)         | `MySQL`        |
| `--db-connection-string`  | `-c`      | Custom database connection string             | *(None)*       |
| `--db-host`               | `-H`      | Database host                                 | `localhost`    |
| `--db-name`               | `-n`      | Database name                                 | `default`      |
| `--db-username`           | `-u`      | Database username                             | `root`         |
| `--db-password`           | `-p`      | Database password                             | `password`     |
| `--db-port`               | `-P`      | Database port (default varies by DB type)     | `3306` or `5432`|
| `--yes`                   | `-y`      | Skip all prompts and use default values       | `false`        |

**Example**:

```bash
gin t -a ecommerce -d PostgreSQL -y
```

---

## Examples

### 1. **Create a New Application with Defaults**

Generates a new Gin application named `my-gin-app` with MySQL as the default database.

```bash
gin new
```

Or using aliases:

```bash
gin n
```

### 2. **Create a New Application with a Custom Name**

Generates a new Gin application named `shop-api`.

```bash
gin new --app-name shop-api
```

Or using shorthand flags:

```bash
gin n -a shop-api
```

### 3. **Specify Database Type and Credentials**

Generates a new Gin application with PostgreSQL as the database.

```bash
gin new --db-type PostgreSQL --db-username admin --db-password secret
```

Or using shorthand flags:

```bash
gin n -d PostgreSQL -u admin -p secret
```

### 4. **Use a Custom Database Connection String**

Generates a new Gin application using a custom database connection string.

```bash
gin new --db-connection-string "user:pass@tcp(localhost:3306)/dbname"
```

Or using shorthand flags:

```bash
gin n -c "user:pass@tcp(localhost:3306)/dbname"
```

### 5. **Skip All Prompts**

Generates a new Gin application using all default values without any interactive prompts.

```bash
gin new --yes
```

Or using shorthand flags:

```bash
gin n -y
```

### 6. **Generate a New Controller**

Creates a new controller named `user`.

```bash
gin generate controller user
```

Or using aliases:

```bash
gin g c user
```

### 7. **Generate a New Guard**

Creates a new guard named `auth`.

```bash
gin generate guard auth
```

Or using aliases:

```bash
gin g gd auth
```

### 8. **Generate a New Resource**

Creates a new resource named `product` with Restful transport.

```bash
gin generate resource product --transport Restful
```

Or using shorthand flags:

```bash
gin g r product -t Restful
```

### 9. **Load a Template**

Loads the standard template to set up the project structure.

```bash
gin template
```

Or using aliases:

```bash
gin t
```

## Running the Application

After generating your application, navigate to the project directory and run your Gin server.

1. **Navigate to the Project Directory**

   ```bash
   cd my-gin-app
   ```

2. **Run the Application**

   ```bash
   go run *.go
   ```

   *Alternatively, you can build the application and run the binary:*

   ```bash
   go build -o app
   ./app
   ```

## Notes

- **Environment Variables**: Ensure that all necessary environment variables are set in the `.env` file. The CLI tool automatically loads environment variables during project setup.
- **Dependencies**: The CLI tool runs `go mod tidy` to manage dependencies. Ensure you have an active internet connection during this process.
- **Templates**: You can customize templates located in the `templates/` directory to fit your project's specific needs.
- **Extensibility**: The CLI is designed to be modular. Feel free to add new commands or extend existing ones as your project evolves.

## Support

If you encounter any issues or have suggestions for improvements, feel free to open an issue on the [GitHub Repository](https://github.com/golang-programming/gincli/issues).

For direct support, you can reach out via email:

📧 [zeshanshakil0@gmail.com](mailto:zeshanshakil0@gmail.com)

---

## Contribution

Contributions are welcome! If you'd like to contribute to Gin CLI, please fork the repository and submit a pull request with your changes. Ensure that your code follows the project's coding standards and includes appropriate documentation.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

**Happy Coding!** 🚀

---