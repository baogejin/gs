ulimit -c unlimited

dir=$gs/bin
rm -rf $gs/log/*
GOTRACEBACK=crash $dir/server -node=gateway 1>/dev/null 2>&1 &
GOTRACEBACK=crash $dir/server -node=logic 1>/dev/null 2>&1 &

ps -ef | grep "server -node"
