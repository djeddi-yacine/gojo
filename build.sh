#!/usr/bin/env bash

# Define default values
CGO_ENABLED="0"
ARCH="both"

# Parse command-line arguments
while getopts ":ca:" opt; do
  case $opt in
    c) CGO_ENABLED="1" ;;  # Enable CGO when -c flag is passed
    a) ARCH="$OPTARG"       # Set target architecture based on -a argument
      if [[ ! "$ARCH" =~ ^(amd64|arm64|both)$ ]]; then
        echo "Invalid architecture: '$OPTARG'. Valid options are amd64, arm64, or both."
        exit 1
      fi
      ;;
    \?) echo "Invalid option: -$OPTARG" >&2; exit 1 ;;
  esac
done

# Shift arguments to remove parsed options
shift $((OPTIND-1))

# Function for building the Go binary
build_go() {
  local enable_cgo="$1"
  local arch="$2"
  local output_file="gojo"

  # Determine whether CGO is enabled
  if [[ "$enable_cgo" == "1" ]]; then
    output_file="${output_file}-cgo"
    cgo_flag="-a"
    install_suffix="cgo"
    # Remove netgo tag for CGO-enabled build
    tags=""
  else
    install_suffix="nocgo"
    cgo_flag=""
    tags="-tags netgo"
  fi

  # Construct output file name
  output_file="${output_file}-${arch}"

  # Perform the build
  GOOS=linux GOARCH="$arch" CGO_ENABLED="$enable_cgo" \
    go build -v \
      $tags \
      -ldflags="${LDFLAGS[*]}" \
      -gcflags="${GCFLAGS[*]}" \
      -trimpath -mod=readonly -buildmode=pie \
      $cgo_flag -installsuffix "$install_suffix" \
      -o "$output_file" .
}


# Set package name and version information
PACKAGE="github.com/dj-yacine-flutter/gojo"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

# Define common build flags
LDFLAGS_=(
  "-X '${PACKAGE}/version.Version=${VERSION}'"
  "-X '${PACKAGE}/version.CommitHash=${COMMIT_HASH}'"
  "-X '${PACKAGE}/version.BuildTime=${BUILD_TIMESTAMP}'"
  "-w" "-s" "-extldflags '-static'"
)
GCFLAGS_=(
  "-S" "-m"
)

# Clean temporary build files
go clean -x
go clean -cache -x

# Build for specified architectures
case $ARCH in
  "amd64")
    build_go "$CGO_ENABLED" amd64   # Call the build_go function with CGO_ENABLED and arch
    ;;
  "arm64")
    build_go "0" arm64   # Disable CGO for arm64
    ;;
  "both")
    build_go "$CGO_ENABLED" amd64
    build_go "0" amd64
    build_go "0" arm64
    ;;
esac

echo "Build completed!"
