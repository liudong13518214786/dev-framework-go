#！/bin/bash
echo "start update backend code"
echo "update..."
git pull
echo "update backend code success!"
echo "start update frontend code"
cd ../kennyBlog/
echo "update..."
git pull
echo "update frontend code success!"
echo "start run docker-compose"
cd ../dev-framework-go/
docker-compose -f docker-compose-prod.yml up -d
echo "server run success! http://47.93.19.60"