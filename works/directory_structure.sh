# 创建merchants模块
mkdir -p merchants/{internal,api}
cd merchants && go mod init merchants && cd ..

# 创建blacklist模块
mkdir -p blacklists/{internal,api}
cd blacklists && go mod init blacklists && cd ..

# 创建system模块
mkdir -p systems/{internal,api}
cd systems && go mod init systems && cd .. 

# 创建command-client模块
mkdir -p command-client/{configs}
touch command-client/main.go
cd command-client && go mod init command-client && cd .. 

# 创建command-server模块
mkdir -p command-server/{configs}
touch command-server/main.go
cd command-server && go mod init command-server && cd .. 

# 创建pkg模块
mkdir -p pkgs/{utils,config,database,logger,monitor}
cd pkgs && go mod init pkgs && cd .. 

# 创建go.work文件
touch go.work
echo "go 1.23" > go.work
echo "use (" >> go.work
echo "    ./merchants" >> go.work
echo "    ./blacklists" >> go.work
echo "    ./systems" >> go.work
echo "    ./command-client" >> go.work
echo "    ./command-server" >> go.work
echo "    ./pkgs" >> go.work
echo ")" >> go.work

