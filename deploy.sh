#!/bin/sh

sudo -u www-data git pull

echo "build"

go build -buildvcs=false -o ai-chat-go-app .

echo "success"

sudo supervisorctl restart ai-chat-go

#sudo supervisorctl status
#sudo supervisorctl reread
#sudo supervisorctl update