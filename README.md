![URLShorty Logo](./assets/main-logo.png)

# URLShorty

Check out URLShorty [here](https://url-short-vubq.onrender.com/)

## Overview

Welcome to URLShorty! This project is a URL shortener built using Go, HTML, and SQLite, allowing users to shorten long URLs for easy sharing. This is a fun little side-project in which I explore networking basics using HTTP as well as exploring the potential of Go as an efficient backend in web applications.

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Tech Stack](#tech-stack)
- [Setup Instructions](#setup-instructions)
- [Deployment](#deployment)
- [How It Works](#how-it-works)
  - [Shortening Algorithm](#shortening-algorithm)
  - [Concurrent Database Access](#concurrent-database-access)
  - [User Interface](#user-interface)
- [Contributions](#contributions)
- [License](#license)

## Architecture Overview

URLShorty's architecture involves a Go backend, `Gorilla Mux` + `HTTP` for routing, `SQLite` for database storage, and `HTML` + `CSS` for the frontend. Mutexes are used to handle concurrent access to the database.

## Tech Stack

- **Backend**:

  - Go (Golang)
  - Gorilla Mux (Router)
  - SQLite3 (Database)

- **Frontend**:

  - HTML
  - CSS

- **Deployment**:
  - Render

## Setup Instructions

1. Clone the repository:

   ```bash
   git clone https://github.com/khairalanam/urlshorty.git
   ```

2. Navigate to the project directory:

   ```cmd
   cd urlshorty
   ```

3. Build and run the Go application:

   ```cmd
   go run .
   ```

4. Access the interface at: http://localhost:3000.

## Deployment

URLShorty is deployed on Render for seamless hosting and accessibility.

## How It Works

### Shortening Algorithm

URLs are shortened using a simple randomised algorithm, guaranteeing short and collision-resistant identifiers.

### Concurrent Database Access

To handle concurrent reads and writes to the SQLite3 database, URLShorty utilizes mutexes. This ensures that only one operation can access the database at a time, preventing race conditions and ensuring data consistency.

### User Interface

The frontend provides a simple and user-friendly interface, allowing users to input long URLs and receive shortened links.

## Contributions

Contributions to URLShorty are welcome! You can fork this repository and change the project to your liking!

## License

URLShorty is open-source software released under the MIT License. See the [LICENSE](LICENSE) file for details.

Feel free to connect with me on [LinkedIn](www.linkedin.com/in/khair-alanam) or [Twitter](https://twitter.com/khair_alanam).
