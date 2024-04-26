# Email Verifier Tool

This is a simple command-line tool which verifies the DNS records associated with a domain name, particularly those relevant to email delivery, such as MX, SPF and DMARC records.

## Requirements

- Go 1.11 or higher
- Internet connection for DNS lookups

## Usage

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/rahulkarthickr/email-verifier-tool.git
    ```

2. Navigate to the project directory:

    ```bash
    cd email-verifier-tool
    ```

3. Compile and run the program:

    ```bash
    go build
    ./email-verifier-tool
    ```

4. Enter the domain name and press `Enter`

## Dependencies

This tool uses the Go standard libraries `bufio` `fmt` `log` `os` `strings` and `net` package for DNS lookups.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch `git checkout -b feature/my-feature`
3. Make your changes.
4. Commit your changes `git commit -am 'Add some feature'`
5. Push to the branch `git push origin feature/my-feature`
6. Create a new Pull Request.
