handle_ctrl_c() {
    echo "Stopping surrealdb"
    docker stop go-chat-db
    exit 0
}

trap handle_ctrl_c SIGINT

surrealParams=""
surrealCommand="docker run --user root --name go-chat-db --rm -d -p 8082:8000 -v $HOME/surreal:/data/surreal surrealdb/surrealdb"
if ! command -v docker &> /dev/null; then
    echo "docker not found, trying to run surrealdb locally"
    echo "WARNING: You need to stop surrealdb manually on exit"
    surrealCommand="surreal"
fi

if [ ! -d "$HOME/surreal" ]; then
    surrealParams="--user root --pass root"
fi
$surrealCommand start $surrealParams -- file:/data/surreal/database.db &

if [ $? -ne 0 ]; then
    echo "Failed to start surrealdb"
    exit 1
fi

runCommand="go run ."
if [ "$1" != "-n" ]; then
    runCommand="watchexec -r -e go -- $runCommand"
fi

$runCommand
