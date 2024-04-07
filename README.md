# AI Command Line Interface (aicli)

`aicli` is a Go CLI application that leverages the power of OpenAI's GPT-3.5 model to provide assistance with Linux terminal commands.

## Installation

1. Ensure you have Go installed on your system.
2. Set your OpenAI API token as an environment variable named `OPENAI_TOKEN`.
3. Install the OpenAI Go client library by running `go get github.com/openai/openai-go/v1`.
4. Clone this repository:

    ```bash
    git clone https://github.com/rezNetDevOps/aicli.git
    ```

5. Navigate to the directory containing the cloned repository.
6. Build the executable:

    ```bash
    go build
    ```

## Usage

To use `aicli`, follow these steps:

1. Execute the built binary:

    ```bash
    ./aicli
    ```

2. Enter your query when prompted. For example:

    ```bash
    Enter your query: list of hidden files in a directory
    ```

3. Press Enter, and `aicli` will provide you with a terminal command along with a brief one-line explanation.

## Example

Here's an example of how `aicli` can assist you:

```bash
Enter your query: list of hidden files in a directory
```

Response:

```bash
ls -a -l
List all files in the directory, including hidden files, with detailed information.
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests on [GitHub](https://github.com/rezNetDevOps/aicli).

## License

This project is licensed under the [MIT License](LICENSE).
