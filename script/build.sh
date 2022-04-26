#!/usr/bin/env bash

package=github.com/Ekisa-Team/ekisa-chatbots-cli

package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" "windows/386")

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has ocurred. Aborting the script execution...'
        exit 1
    fi
done
