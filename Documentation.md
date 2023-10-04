# My Go Web Forum Application

This is a web forum application written in Go that allows users to communicate, post comments, and more.

## Table of Contents

- [Structure](#structure)
- [Installation](#installation)
- [Usage](#usage)
- [Dockerization](#dockerization)
- [Contributing](#contributing)
- [License](#license)

## Structure

Package Structure
The Go code is organised into packages and files with meaningful names. The following directories are used to group related functionality.

```
forumapp/
    ├── main.go
    ├── app/
    │   ├── app.go
    │   ├── auth.go
    │   ├── database.go
    │   ├── handlers.go
    │   └── models.go
    └── tests/
        ├── app_test.go
        └── database_test.go
```

## Installation

To install this application, follow these steps:

1. Clone the repository: `git clone https://github.com/yourusername/my-golang-app.git`
2. Build the Go application: `go build -o main .`
3. Run the application: `./main`

## Usage

- Register as a new user.
- Log in to access the forum features.
- Create posts and comments.
- Like and dislike posts and comments.
- Filter posts by categories, created posts, and liked posts.

## Dockerization

You can also run this application using Docker. Follow these steps:

1. Build the Docker image: `docker build -t my-golang-app .`
2. Run the Docker container: `docker run -p 8080:8080 my-golang-app`

The application will be available at http://localhost:8080.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

```

```
