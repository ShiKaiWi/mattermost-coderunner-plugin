# mattermost-coderunner-plugin
It is a mattermost plugin for running codes in markdown posts.

## Installation
### Prerequisites
We recommend **CentOS7** as the host os where mattermost run and you should know that the following guide is for it so some difference may exist if you choose some other linux distribution.

[Docker](https://docs.docker.com/install/) is necessary so you have to install it if your mattermost server doesn't have docker.

Then some language docker images needs downloading on mattermost server including:
1. golang: `docker pull golang`
2. node: `docker pull node`

Another important note is that you have to make sure the `mattermost` user(who that starts the mattermost server) is added to the `docker` user group. You can use such command to detect whether `mattermost` user has been added:
```
getent group docker
# output should be something like:
# docker:x:995:xkwei,mattermost
```

If the `mattermost` user does not belong to `docker` group, use such command to add:
```
sudo usermod -aG docker mattermost
```
And then mattermost server needs restarting.

### Build & Install plugins
We provides two plugin, server-coderunner-plugin & web-coderunner-plugin, and they should be installed into mattermost both.

You can either download the two plugins on the release page or build them from the source code.

If you choose to build them, your have to ensure that your host meets the requirements:
1. golang >= 1.9
2. dep >= 0.5

then just follow the script below:
```
go get github.com/ShiKaiWi/mattermost-coderunner-plugin
cd $GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin
dep ensure
sh scripts/build.sh
```

Then the two plugins (**Targeting on Linux**) will occur under the  `$GOPATH/src/github.com/ShiKaiWi/mattermost-coderunner-plugin`.

And as for how to install the plugin please refer to [install mattermost plugin](https://developers.mattermost.com/extend/plugins/server/hello-world/#installing-the-plugin).
