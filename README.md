# PocketBase Dashboard

PocketBase Dashboard is a self-hosted solution that allows you to manage and use PocketBase for personal use. With PocketBase Dashboard, you can have full control over your data and applications in a convenient and user-friendly way.

## Features

- **Self-Hosting**: Host PocketBase Dashboard on your own server, ensuring privacy and security.
- **User-Friendly Interface**: Easy-to-use dashboard for managing your PocketBase instances.

## Getting Started

Follow the steps below to set up PocketBase Dashboard using Docker.

### Prerequisites

Make sure you have Docker installed on your system. If not, you can download and install it from the [official Docker website](https://www.docker.com/get-started).

### Self-Hosting Guide

1. Create the pocketbase-dashboard docker network

   ```bash
   docker network create lazar-static
   ```

2. Create a `docker-compose.yml` file with the following content:

   ```yaml
   version: '3.8'
   services:
     lazar-dash:
       image: monsieurlazar/pocketbase-dashboard
       container_name: lazar-dash
       environment:
         - ORIGIN=https://pocket.example.com/
         - DEFAULT_PASSWORD=example //defaults to password
       volumes:
         - /var/run/docker.sock:/var/run/docker.sock
         - /home/pocketbase/metadata:/data
       networks:
         - lazar-static
         - lazar-network
       restart: always

   networks:
     lazar-static:
       external: true
     lazar-network:
       external: true
   ```

```bash
   docker run -d -p 8081:80 -e ORIGIN=http://localhost:8081 -e DEFAULT_PASSWORD=example --name lazar-dash -v /var/run/docker.sock:/var/run/docker.sock -v /home/pocketbase/metadata:/data --network=lazar-static monsieurlazar/pocketbase-dashboard
```

3. Start the PocketBase Dashboard container using Docker Compose:

   ```bash
   docker-compose up -d
   ```

   This will pull the necessary Docker image and start the PocketBase Dashboard container in the background.

4. Access PocketBase Dashboard in your web browser by navigating to `http://your-server-ip:port` (replace `your-server-ip` and `port` with your server's IP address and the port you specified in the `docker-compose.yml` file).

5. Log in using the default credentials:

   - **Password:** password

## Roadmap

Our future plans for PocketBase Dashboard include:

- **Improved User Management**: Enhance user roles and permissions management features.
- **More Secury Auth Options**: Improve the security of the system with more secure auth alternatives.
- **API Support**: Provide an api with auth keys for creating projects programatically.

We welcome contributions and feedback from the community to help us improve PocketBase Dashboard. Feel free to open issues and submit pull requests on our [GitHub repository](https://github.com/monsieurlazar/pocketbase-dashboard)!

---

**Note:** Please ensure that you follow best practices for security and server management while self-hosting PocketBase Dashboard.