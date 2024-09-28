Certainly! Here's a streamlined version of your `README.md` that includes the **name**, **logo**, **installation instructions via `curl` and `git clone`**, **support email**, and a **table outlining commands and subcommands**.

---

# Gin CLI

![Gin CLI Logo](https://github.com/golang-programming/gincli/gincli-logo.jpg)

Gin CLI is a powerful command-line tool designed to help you quickly scaffold [Gin Web Framework](https://gin-gonic.com/) applications with a standardized project structure and essential components. Focus on building your application without the hassle of initial setup!

## Installation

### Install via `curl`

**Example for Linux:**

```bash
curl -LO https://github.com/golang-programming/gincli/releases/download/v1.0.0/gin-linux-amd64
chmod +x gin-linux-amd64
sudo mv gin-linux-amd64 /usr/local/bin/gin
gin --version
```

**Example for macOS:**

```bash
curl -LO https://github.com/golang-programming/gincli/releases/download/v1.0.0/gin-darwin-amd64
chmod +x gin-darwin-amd64
sudo mv gin-darwin-amd64 /usr/local/bin/gin
gin --version
```

#### For Windows (Using PowerShell):

```powershell
Invoke-WebRequest -Uri "https://github.com/golang-programming/gincli/releases/download/v1.0.0/gin-windows-amd64.exe" -OutFile "gin.exe"
.\gin.exe --version
```

### Install via Source code

```bash
git clone https://github.com/golang-programming/gincli.git
cd gincli
go build -o gin
sudo mv gin /usr/local/bin/
```


## Commands

| Command    | Alias | Description                                     |
|------------|-------|-------------------------------------------------|
| `new`      | `n`, `create` | Create a new Gin application                     |
| `generate` | `g`   | Generate components like controllers, guards, etc. |
| `template` | `t`   | Load application startup templates               |

### Generate Subcommands

| Subcommand  | Alias | Description                |
|-------------|-------|----------------------------|
| `controller` | `c`   | Generate a new controller    |
| `guard`      | `gd`  | Generate a new guard         |
| `resource`   | `r`   | Generate a new resource      |
| `route`      | `rt`  | Generate a new route         |
| `service`    | `s`   | Generate a new service       |

---

## Quick Start Examples

### Create a New Application

```bash
gin n -a myapp
```

### Generate a Resource with Restful Transport

```bash
gin g r product -t Restful
```

### Load a Template

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

---

## Support

For issues or suggestions, please contact us at:

📧 [zeshanshakil0@gmail.com](mailto:zeshanshakil0@gmail.com)

---

**Happy Coding!** 🚀