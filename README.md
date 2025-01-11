# URL Shortener 🌐🔗

A simple and user-friendly URL shortener written in Go, supporting flexible URL redirection via YAML, JSON, and database-driven path mappings.

---

## Features ✨

- **Default MapHandler**: Basic in-memory URL mapping.
- **YAMLHandler**: Load URL mappings from a YAML file.
- **JSONHandler**: Load URL mappings from a JSON file.
- **DatabaseHandler**: Extendable to use a database (like BoltDB or SQL) for URL mappings.

---

## How to Run 🚀

### Prerequisites

- Go installed on your system. If you don't have it yet, [Download Go](https://golang.org/dl/).

### Steps

1. **Clone the repository**:

    ```bash
    git clone https://github.com/<your-username>/urlshort.git
    cd urlshort
    ```

2. **Build the project**:

    ```bash
    go build
    ```

3. **Run the application with the desired configuration**:

    - **Default MapHandler**:

      ```bash
      ./urlshort
      ```

    - **YAMLHandler**:

      ```bash
      ./urlshort -yaml urls.yaml
      ```

    - **JSONHandler**:

      ```bash
      ./urlshort -json urls.json
      ```

---

## Example Configuration Files 📄

### Example YAML File (`urls.yaml`)

```yaml
- path: /example
  url: https://example.com
- path: /github
  url: https://github.com
