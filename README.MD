# Golang Telegram Notifier

A tiny utility to send Telegram notifications from the shell.

## Setup

1. Build: ``` go build -o tg_notify```
2. Config: ``` ~/.config/tg_send/conf.toml```

```toml
[telegram]
token = "your_bot_token"
chat_id = "your_chat_id"
```

### Setup system wide

```bash
mkdir -p ~/go/bin
mv tg_notify ~/.go/bin/
```

## Usage

```bash
./your_gentoo_kernel_compile_script.sh && tg_notify "Kernel compiled"
```
