#!/usr/bin/env bash
cd /var/www/html/toplog; git pull; go build -o ./toplog; systemctl restart toplog
systemctl stop toplog; cd /var/www/html/toplog; git pull; go build -o ./toplog; ./toplog
