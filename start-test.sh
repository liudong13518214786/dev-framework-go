#！/bin/bash
echo "start update backend code"
echo "update..."
git pull
echo "update backend code success!"
echo "start update frontend code"
cd /Users/liudong/Documents/vmshare/kennyBlog/
echo "update..."
git pull
echo "update frontend code success!"
echo "start run docker-compose"
cd /Users/liudong/go/src/dev-framework-go/
docker-compose -f docker-compose-prod.yml up -d
each "server run success! http://127.0.0.1:8100"