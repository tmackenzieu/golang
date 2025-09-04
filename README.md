### Run the application
With the config files in place, you should be able to run the application from your favorite IDE without any extra configuration.

Use the following command to run the application:
```shell
go run ./...
```

### Other Considerations

#### Policy Agent

The current local configuration allows the Policy Agent client to initiate correctly without
crashing the app, but it will always disallow responses.


## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...
