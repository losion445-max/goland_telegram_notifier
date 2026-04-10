# tg_send

A tiny utility to send Telegram notifications from the shell.

## Setup

1. Build: ```bash go build -o notify```
2. Config: ```bash ~/.config/tg_send/conf.toml```

```toml
[telegram]
token = "your_bot_token"
chat_id = "your_chat_id"
```

### Setup system wide

```bash
mkdir -p ~/.local/bin
mv tg_send ~/.local/bin/
```

## Usage

```bash
./your_gentoo_kernel_compile_script.sh && tg_notify "Kernel compiled"
```
