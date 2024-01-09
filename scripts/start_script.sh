#!/bin/bash

# Set the maximum log size in bytes (1GB = 1073741824 bytes)
MAX_LOG_SIZE=1073741824

# Set the path to your Go binary
BINARY_PATH="/home/shardeumcoredev/rpc_local_data/go-service-validator/service_validator"

# Set the log file path
LOG_FILE="/home/shardeumcoredev/rpc_local_data/go-service-validator/main.log"

# Export the environment variables from the file
export CONFIG_FILE_PATH="/home/shardeumcoredev/rpc_local_data/go-service-validator/config.json"
export APP_SQLITE_DB_PATH="/home/shardeumcoredev/rpc_local_data/server/db/shardeum.sqlite"

# Start the Go binary in the background and redirect output to the log file
nohup $BINARY_PATH >$LOG_FILE 2>&1 &

# Create a logrotate configuration file
LOGROTATE_CONF="/etc/logrotate.d/go_binary_logs"
echo "$LOG_FILE {
    size $MAX_LOG_SIZE
    create 0644 root root
    rotate 10
    missingok
    notifempty
    compress
    delaycompress
    copytruncate
}" | sudo tee $LOGROTATE_CONF >/dev/null

# Set proper permissions for the logrotate configuration file
sudo chown root:root $LOGROTATE_CONF
sudo chmod 644 $LOGROTATE_CONF

# Run logrotate manually to rotate the log immediately
sudo logrotate -f $LOGROTATE_CONF
