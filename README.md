# yarser

[![Release](https://github.com/unfor19/yarser/actions/workflows/release.yml/badge.svg)](https://github.com/unfor19/yarser/actions/workflows/release.yml)

A CLI for parsing YAML anchors to plain YAML files.

## Features

- Parses YAML anchors
- Allows usage of `.hidden` nodes by adding the prefix `.` to a root node. All hidden nodes are omitted from the final result

In my tests, I use [tests/github-action/my-action.yml](https://github.com/unfor19/yarser/blob/master/tests/github-action/my-action.yml) as the source file, and the compiled result is [tests/github-action/my-action-compiled.yml](https://github.com/unfor19/yarser/blob/master/tests/github-action/my-action-compiled.yml). Go ahead and check how it looks like after parsing (compiling) the source YAML file.

## Getting Started

1. Download the binary file from the releases page, for example [0.0.1rc2](https://github.com/unfor19/yarser/releases/tag/0.0.1rc2)
   - macOS
    ```bash
    YARSER_OS="darwin" && \
    YARSER_ARCH="amd64" && \
    YARSER_VERSION="0.0.1rc2" && \
    wget -O yarser "https://github.com/unfor19/yarser/releases/download/${YARSER_VERSION}/yarser_${YARSER_VERSION}_${YARSER_OS}_${YARSER_ARCH}"
    ```
   - Linux
    ```bash
    YARSER_OS="linux" && \
    YARSER_ARCH="amd64" && \
    YARSER_VERSION="0.0.1rc2" && \
    wget -O yarser "https://github.com/unfor19/yarser/releases/download/${YARSER_VERSION}/yarser_${YARSER_VERSION}_${YARSER_OS}_${YARSER_ARCH}"
    ```
   - Windows WSL2
    ```bash
    YARSER_OS="linux" && \
    YARSER_ARCH="386" && \    
    YARSER_VERSION="0.0.1rc2" && \
    wget -O yarser "https://github.com/unfor19/yarser/releases/download/${YARSER_VERSION}/yarser_${YARSER_VERSION}_${YARSER_OS}_${YARSER_ARCH}"   
    ```
2. Move `yarser` binary to `bin` dir and set permissions to allow execution
   ```bash
   chmod +x yarser
   sudo mv yarser "/usr/local/bin/yarser"
   ```
3. Parse a YAML file once - here's the YAML file with anchors that I'm testing with [my-action.yml](tests/github-action/my-action.yml)
   ```bash
   SRC_FILE_PATH="tests/github-action/my-action.yml" && \
   DST_FILE_PATH=".my-action-compiled.yml" && \
   yarser parse --watch "$SRC_FILE_PATH" "$DST_FILE_PATH"
   # INFO[2021-08-21T19:10:25+03:00] Successfully parsed tests/github-action/my-action.yml to .my-action-compiled.yml
   # Check the file .my-action-compiled.yml
   ```
4. Parse on save while editing the source file by adding the `--watch` flag
   ```bash
   SRC_FILE_PATH="tests/github-action/my-action.yml" && \
   DST_FILE_PATH=".my-action-compiled.yml" && \
   yarser parse --watch "$SRC_FILE_PATH" "$DST_FILE_PATH"
   
   # INFO[2021-08-21T19:13:47+03:00] Watching for changes in tests/github-action/my-action.yml ... 
   # INFO[2021-08-21T19:13:53+03:00] Successfully parsed tests/github-action/my-action.yml to .my-action-compiled.yml
   
   # Keep on editing and saving the source file, and it will automatically parse it
   ```

## Authors

Created and maintained by [Meir Gabay](https://github.com/unfor19)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/unfor19/yarser/blob/master/LICENSE) file for details
