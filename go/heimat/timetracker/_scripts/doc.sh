
#!/bin/sh

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

cd $DIR/..

echo running at http://localhost:6060
godoc -http=:6060 -play=true