# Gin CLI - Scaffold Your Gin Application

Gin CLI is a straightforward command-line tool that helps you quickly set up [Gin Web Framework](https://gin-gonic.com/) applications with a predefined project structure and essential components. Focus on developing your application without worrying about the initial setup!

## Features

- **Quick Application Setup**: Generate a new Gin application with a standard project layout.
- **Component Generation**: Easily create controllers, guards, resources, routes, and services.
- **Database Support**: Configure your app with MySQL, PostgreSQL, SQLite, or MongoDB.
- **User-Friendly Prompts**: Interactive prompts with sensible defaults for a smoother experience.
- **Shorthand Commands**: Utilize command aliases and shorthand flags for efficiency.
- **Safe File Handling**: Prompts before overwriting existing files to prevent data loss.
- **Enhanced Help Command**: Displays commands and flags in a clean, tabular format.

## Installation

### Prerequisites

- **Go**: Make sure you have Go installed (version 1.16 or later). [Download Go](https://golang.org/dl/)
- **Git**: Required if cloning the repository. [Download Git](https://git-scm.com/downloads)

### Install via `curl`

You can quickly install the CLI tool using `curl`. Replace `<version>` with the latest release version.

```bash
# Replace <version> with the actual version, e.g., v1.0.0
curl -LO https://github.com/golang-programming/gincli/releases/download/<version>/gin
chmod +x gin
sudo mv gin /usr/local/bin/
```

### Build from Source

Alternatively, clone the repository and build the tool yourself:

```bash
git clone https://github.com/golang-programming/gincli.git gin
cd gin
go build -o gin
sudo mv gin /usr/local/bin/
```

## Usage

Create and manage your Gin applications using simple commands.

### Available Commands

- `new` (Aliases: `n`, `create`): Create a new Gin application.
- `generate` (Alias: `g`): Generate components like controllers, guards, etc.
- `template` (Alias: `t`): Load application startup templates.

### Command Aliases

Use aliases for quicker command input:

- **new**: `n`, `create`
- **generate**: `g`
- **template**: `t`

#### Generate Subcommand Aliases

- **controller**: `c`
- **guard**: `gd`
- **resource**: `r`
- **route**: `rt`
- **service**: `s`

## Commands Overview

### `new`

Create a new Gin application with a predefined structure.

```bash
gin new [flags]
```

**Flags:**

- `--app-name`, `-a`: Application name (default: `my-gin-app`)
- `--db-type`, `-d`: Database type (`MySQL`, `PostgreSQL`, `SQLite`, `MongoDB`)
- `--db-username`, `-u`: Database username (default: `root`)
- `--db-password`, `-p`: Database password (default: `password`)
- `--yes`, `-y`: Skip prompts and use default values

### `generate`

Generate components for your application.

```bash
gin generate [subcommand] [flags]
```

**Subcommands and Aliases:**

- `controller` (`c`): Generate a new controller.
- `guard` (`gd`): Generate a new guard.
- `resource` (`r`): Generate a new resource.
- `route` (`rt`): Generate a new route.
- `service` (`s`): Generate a new service.

### `template`

Load an application startup template.

```bash
gin template [flags]
```

**Flags:**

- `--template`, `-t`: Template to use (default: `Standard`)
- `--app-name`, `-a`: Application name (default: `my-gin-app`)
- `--yes`, `-y`: Skip prompts and use default values

## Examples

### Create a New Application

```bash
gin new -a myapp
```

Or using aliases:

```bash
gin n -a myapp
```

### Generate a Controller

```bash
gin generate controller user
```

Or using aliases:

```bash
gin g c user
```

### Generate a Resource with Restful Transport

```bash
gin generate resource product -t Restful
```

Or using shorthand:

```bash
gin g r product -t Restful
```

### Load a Template

```bash
gin template -a ecommerce -y
```

Or using aliases:

```bash
gin t -a ecommerce -y
```

## Running the Application

After generating your application:

1. **Navigate to the Project Directory**

   ```bash
   cd myapp
   ```

2. **Run the Application**

   ```bash
   go run *.go
   ```

## Notes

- **Environment Variables**: Ensure your `.env` file is correctly set up.
- **Dependencies**: The CLI runs `go mod tidy` to manage dependencies.
- **Customization**: Feel free to modify templates in the `templates/` directory.

## Support

For issues or suggestions, please open an issue on the [GitHub Repository](https://github.com/golang-programming/gincli/issues) or contact:

📧 [zeshanshakil0@gmail.com](mailto:zeshanshakil0@gmail.com)

---

**Happy Coding!** 🚀