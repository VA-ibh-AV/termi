# termi

`termi` = **terminal + AI**.

Ask in plain English, get back a shell command.

## Install (Linux amd64)

Download the latest binary from GitHub Releases:

```bash
curl -LO https://github.com/VA-ibh-AV/termi/releases/latest/download/termi_linux_amd64.tar.gz
tar -xzf termi_linux_amd64.tar.gz
sudo mv termi /usr/local/bin
```

For macOS/Windows and other architectures, use the matching asset from the Releases page.

## Quick Start

```bash
export OPENAI_KEY="your_openai_api_key"
termi "find and delete node_modules folders recursively"
```

## Commands

For now, there is only one command:

```bash
termi [prompt]
```
