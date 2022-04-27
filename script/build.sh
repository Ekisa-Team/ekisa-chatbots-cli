#!/usr/bin/env bash

package=github.com/Ekisa-Team/ekisa-chatbots-cli

platforms=("windows/amd64" "windows/386")

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=ec'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    path=bin/$output_name
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $path $package
    if [ $? -ne 0 ]; then
        echo 'An error has ocurred. Aborting the script execution...'
        exit 1
    fi
done
