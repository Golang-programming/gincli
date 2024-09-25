# Gin CLI - Scaffold Your Gin Application

This CLI tool helps you quickly set up [Gin Web Framework](https://gin-gonic.com/) application with a predefined project structure and components. so you can focus on building your application.

## Features

- **Quick Scaffold**: Generate a new Gin application with a standard project layout.
- **Database Configuration**: Supports MySQL, PostgreSQL, SQLite, and MongoDB.
- **User-Friendly**: Includes prompts with default values better UX.

## Installation

1. **Install in your machine**

   ```bash
   git clone https://github.com/golang-programming/gincli.git
   cd gincli
   go build -o gin
   mv gin /usr/local/bin/
   ```

## Usage

Generate a new Gin application by running:

```bash
gin new [flags]
```

### Available Flags

- `--app-name`: Name of your application (default: `my-gin-app`).
- `--db-type`: Database type (`1`: MySQL, `2`: PostgreSQL, `3`: SQLite, `4`: MongoDB).
- `--db-connection-string`: Custom database connection string.
- `--db-host`: Database host (default: `localhost`).
- `--db-name`: Database name (default: `default`).
- `--db-username`: Database username (default: `root`).
- `--db-password`: Database password (default: `password`).
- `--db-port`: Database port (default varies by DB type).
- `-y`, `--yes`: Skip prompts and use default values.

### Examples

1. **Create a New Application with Defaults**

   ```bash
   gin new
   ```

2. **Create a New Application with Custom Name**

   ```bash
   gin new --app-name myapp
   ```

3. **Specify Database Type and Credentials**

   ```bash
   gin new --db-type 2 --db-username user --db-password pass
   ```

4. **Use a Custom Database Connection String**

   ```bash
   gin new --db-connection-string "user:pass@tcp(localhost:3306)/dbname"
   ```

5. **Skip All Prompts**

   ```bash
   gin new --yes
   ```

## Running the Application

1. **Navigate to the Project Directory**

   ```bash
   cd my-gin-app
   ```

2. **Run the Application**

   ```bash
   go run *.go
   ```

## Notes

- Ensure you have [Go](https://golang.org/dl/) installed (version 1.16 or later).
- The CLI automatically runs `go mod tidy` to manage dependencies.
- Check that all environment variables are correctly set in the `.env` file.
- If you encounter any issue, Feel free to open issue
- For support mail me on [email](zeshanshakil0@gmail.com)
