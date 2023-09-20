path=$(cd "$(dirname "$0")";pwd)
mkdir $path/logs
echo $path
cd $path/../Geth
kill $(lsof -i:9545 -t)
yarn
nohup yarn build > $path/logs/node.log 2>&1 &   

sleep 5
cd $path
yarn 
yarn build
yarn hardhat deploy --network l1

# tail -f -n 100 $path/logs/node.log