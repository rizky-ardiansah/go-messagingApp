# Go Messaging App

Real-time messaging application built with Go, Fiber, and WebSocket for instant communication.


- **Live Demo**: http://chat.rizkyardiansah.online/
## 📋 Documentation

- **Technical Specifications**: [View Technical Specs](https://drive.google.com/file/d/1r7C4tkL_dGUFzJ6gDcrrQnpr4_cqPWOW/view?usp=sharing)


## 🔧 Setup & Installation

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- MySQL Database
- MongoDB

### Local Development

1. **Clone the repository**

   ```bash
   git clone https://github.com/rizky-ardiansah/go-messagingApp
   cd go-messagingApp
   ```

2. **Setup environment variables**

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. **Run the application**

   ```bash
   go mod tidy
   go run main.go
   ```

4. **Access the application**
   - API: `http://localhost:4000`
   - WebSocket: `ws://localhost:8080`

### Docker Deployment

```bash
docker build -t messaging-app .
docker run -p 4000:4000 -p 8080:8080 messaging-app
```


## 📊 Monitoring

The application includes monitoring capabilities:

- **APM**: Elastic Application Performance Monitoring
- **Logs**: Structured logging with rotation
- **Health Checks**: Built-in health check endpoints