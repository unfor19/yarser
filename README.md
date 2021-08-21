# yarser

[![Release](https://github.com/unfor19/yarser/actions/workflows/release.yml/badge.svg)](https://github.com/unfor19/yarser/actions/workflows/release.yml)

A CLI for parsing YAML anchors to plain YAML files.

## Features

- Parses YAML anchors
- Allows usage of `.hidden` nodes by adding the prefix `.` to a root node. All hidden nodes are omitted from the final result

## Getting Started

1. Download the binary file from the releases page, for example [0.0.1rc2](https://github.com/unfor19/yarser/releases/tag/0.0.1rc2)
   - macOS
    ```bash
    YARSER_OS="darwin" && \
    YARSER_VERSION="0.0.1rc2" && \
    wget -O yarser "https://github.com/unfor19/yarser/releases/download/${YARSER_VERSION}/${YARSER_VERSION}_${YARSER_OS}_amd64"
    ```
   - Linux
    ```bash
    YARSER_OS="linux" && \
    YARSER_VERSION="0.0.1rc2" && \
    wget -O yarser "https://github.com/unfor19/yarser/releases/download/${YARSER_VERSION}/${YARSER_VERSION}_${YARSER_OS}_amd64"
    ```
   - Windows - Download the executeable [yarser_0.0.1rc2_windows_amd64.exe](https://github.com/unfor19/yarser/releases/download/0.0.1rc2/yarser_0.0.1rc2_windows_amd64.exe)

   ```
2. Move `yarser` binary to `bin` dir and set permissions to allow execution
   ```
   sudo mv yarser /usr/local/bin/yarser && \
   sudo chmod +x /usr/local/bin/yarser
   ```
   **NOTE**: No need on windows
3. Parse a YAML file once
   ```bash
   SRC_FILE_PATH="tests/github-action/my-action.yml" && \
   DST_FILE_PATH=".my-action-compiled.yml" && \
   yarser parse --watch "$SRC_FILE_PATH" "$DST_FILE_PATH"   
   yarser parse tests/github-action/my-action.yml .test.yml
   # INFO[2021-08-21T19:10:25+03:00] Successfully parsed tests/github-action/my-action.yml to .test.yml
   # Check the file .test.yml
   ```
4. Parse on save automatically while editing the source file by adding the `--watch` flag
   ```bash
   SRC_FILE_PATH="tests/github-action/my-action.yml" && \
   DST_FILE_PATH=".my-action-compiled.yml" && \
   yarser parse --watch "$SRC_FILE_PATH" "$DST_FILE_PATH"
   
   # INFO[2021-08-21T19:13:47+03:00] Watching for changes in tests/github-action/my-action.yml ... 
   # INFO[2021-08-21T19:13:53+03:00] Successfully parsed tests/github-action/my-action.yml to .test.yml
   
   # Keep on editing and saving the source file, and it will automatically parse it
   ```

## Authors

Created and maintained by [Meir Gabay](https://github.com/unfor19)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/unfor19/yarser/blob/master/LICENSE) file for details
