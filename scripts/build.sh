ORIG_PATH=$(pwd)
WORKDIR=$GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin/
cd $WORKDIR

echo "building server-coderunner-plugin..."
sh scripts/build-server-plugin.sh

cd $WORKDIR

echo "building web-coderunner-plugin..."
sh scripts/build-web-plugin.sh

cd $ORIG_PATH