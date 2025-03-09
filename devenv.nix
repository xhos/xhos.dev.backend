{
  languages.go.enable = true;

  # Scripts for building and running the application
  scripts = {
    # Build the application
    build.exec = "go build -o bin/api ./cmd/api";

    # Run the application directly (without building first)
    run.exec = "go run ./cmd/api/main.go";

    # Build and then run the application
    start.exec = "go build -o bin/api ./cmd/api && ./bin/api";

    # Run tests
    test.exec = "go test ./...";

    # Clean build artifacts
    clean.exec = "rm -rf bin";

    fmt.exec = "go fmt ./...";
  };

  # Create bin directory on shell entry
  enterShell = ''
    mkdir -p bin
  '';
}
