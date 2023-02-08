UNAMES=`uname -s`
UNAMEP=`uname -p`

VERSION="$1"
EDITION="$2"

if [ -z "$VERSION" ]; then
    VERSION=`curl -fsSL https://api.github.com/repos/machbase/machbase-neo/releases/latest |grep tag_name | awk '{print $2}' | tr -d '",'`
fi

if [ -z "$VERSION" ]; then
    echo "no version is specified"
    exit 1
fi

case $UNAMES in
    Linux)
        UNAMES="linux"
    ;;
    Darwin)
        UNAMES="darwin"
    ;;
    Windows)
        UNAMES="windows"
    ;;
    *)
        echo "Unsupported OS $UNAMES"
        exit 1
    ;;
esac

case $UNAMEP in
    aarch64)
        UNAMEP="arm64"
    ;;
    arm)
        UNAMEP="arm64"
    ;;
    i386)
        UNAMEP="i386"
    ;;
    x86_64)
        UNAMEP="amd64"
    ;;
esac

case $EDITION in
    edge)
        ;;
    fog)
        ;;
    *)
        if [ "$UNAMES" = "darwin" ]; then
            EDITION="fog"
        elif [ "$UNAMES" = "linux" ]; then
            NP=`nproc`
            if [ $NP -gt 8 ]; then
                EDITION="fog"
            else
                EDITION="edge"
            fi
            echo "Installing '$EDITION' edition, host machine has $NP processors."
        else
            EDITION="edge"
            echo "Installing '$EDITION' edition by default."
        fi
    ;;
esac

# echo $UNAMES $UNAMEP $EDITION $VERSION

FNAME="machbase-neo-$EDITION-$VERSION-$UNAMES-$UNAMEP.zip"

echo "Downloading... $FNAME"

curl -L -o $FNAME \
    "https://github.com/machbase/machbase-neo/releases/download/${VERSION}/${FNAME}" \
    && echo "\n\nDownload complete $FNAME"

