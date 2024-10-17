# http-script-agent
A containerized HTTP Server written in Go which executes shell commands on requests and returns the shell command's output.

### Example
**Request:**
```
curl -u user:pass http://localhost:8080/cmd/hello
```

**Internal shell call:**
```
/app/your-command.sh hello
```

## Usage
```
docker run -d \
    -p 8080:8080 \
    -e PORT=8080 \
    -e SHELL_COMMAND=/app/demo.sh \
    -e USERNAME=user \
    -e PASSWORD=pass \
    ghcr.io/virtualzone/http-script-agent:latest
```

This launches the container on port 8080, executes ```/app/demo.sh``` on incoming command requests and enforces basic authentication.