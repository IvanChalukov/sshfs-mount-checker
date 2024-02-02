
# SSHFS Mount Manager

This project provides a Go application that checks if an SSHFS mount is present and mounts it if not. It supports both password-based and SSH key-based authentication for mounting remote filesystems. After a successful operation, it pings a specified health check URL, making it suitable for automated monitoring setups.

## Features

- Checks for the presence of an SSHFS mount.
- Mounts the filesystem using either password-based or SSH key-based authentication.
- Sends a health check ping to a specified URL upon successful mounting or if already mounted.
- Customizable SSHFS options and port settings.

## Prerequisites

- Go (Golang) installed on your system.
- SSHFS must be installed and accessible in your system's PATH.
- Access to a remote filesystem that supports SSH/SFTP connections.
- (Optional) A Uptime Kuma or similar health monitoring setup to receive the pings.

## Installation

### By cloning repository

Clone this repository to your local machine using:

```bash
git clone https://github.com/IvanChalukov/sshfs-mount-checker.git
```

Navigate into the project directory:

```bash
cd sshfs-mount-checker
```

Build the project:

```bash
go build -o sshfs-mount-manager
```

### By go install

Install binary with following command
```sh
go install github.com/IvanChalukov/sshfs-mount-checker@latest
```

## Usage

The application requires several parameters to run, which can be provided as command-line flags:

- `username`: Username for the SSHFS connection.
- `url`: URL or IP address of the storage machine.
- `remotemountpoint`: Mount point on the storage machine.
- `localmountpoint`: Local mount point for SSHFS.
- `healthurl`: URL to ping after successful operation.

Additionally, the SSHFS password should be provided as an environment variable `SSHFS_PASSWORD` if password-based authentication is used.

Example usage:

```bash
export SSHFS_PASSWORD="your_password_here"
./sshfs-mount-manager -username your_username -url storage.example.com -remotemountpoint /remote/path -localmountpoint /local/mount -healthurl https://uptime.example.com/api/ping
```

## Security Note

For password-based authentication, it's highly recommended to use a secure method to set the `SSHFS_PASSWORD` environment variable, avoiding hardcoding or exposing it in scripts or cron jobs.

## Contributing

Contributions to this project are welcome. Please open an issue or submit a pull request with your changes or suggestions.
