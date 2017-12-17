#!/bin/bash

# github_api_token should be available

# OWNER=dtchanpura
# REPO_NAME=deployment-agent
TAG=$(git describe --abbrev=0)

function print_usage () {
    echo "  Usage: $0 -cu -t v1.0.0 -o username -r deployment-agent"
    echo
    echo "  -c    --create-release    for creating a new release for given tag"
    echo "  -u    --upload-artifacts  for uploading the generated artifacts to release"
    echo "  -b    --build             for build the project by running build.sh"
    echo "  -t    --tag  tag_name     to specify the tag name for release"
    echo "  -o    --owner owner_name  to specify repository owner name"
    echo "  -r    --repo repo_name    to specify repository name"
    echo "  -h    --help              for displaying this help"
    echo
    exit 0
}

if [[ $# -eq 0 ]]; then
    print_usage
fi

if [[ "$github_api_token" == "" ]]; then
    echo "Environment variable \`github_api_token' is not available."
    exit 1
fi

while [[ $# -gt 0 ]]
do
    key="$1"

    case $key in
        -t|--tag)
            TAG="$2"
            shift 2
            ;;
        -c|--create-release)
            CREATE_RELEASE=YES
            shift # past value
            ;;
        -u|--upload-artifacts)
            UPLOAD_ARTIFACTS=YES
            shift
            ;;
        -b|--build)
            BUILD=YES
            shift
            ;;
        -o|-owner)
            OWNER="$2"
            shift 2
            ;;
        -r|--repo)
            REPO_NAME="$2"
            shift 2
            ;;
        -h|--help)
            print_usage
            shift # past argument
            ;;
    esac
done
#
# if [ "$TAG" == "" ]; then
#     TAG=$(git describe --abbrev=0)
#     RELEASE_ID=$(curl -XGET ${API_URL}/tags/${TAG} | jq .id)
# fi

API_URL="https://api.github.com/repos/${OWNER}/${REPO_NAME}/releases"
UPLOADS_URL="https://uploads.github.com/repos/${OWNER}/${REPO_NAME}/releases"


function create_release () {
    # tag_name=$(git describe --abbrev=0)
    echo "Creating release..."
    tag_name=${TAG}
    title=$(git tag -n10 ${tag_name} | sed 's/^[ ]*//g' | sed "s/${tag_name}[ ]*//" | head -1)
    body=$(git tag -n10 ${tag_name} | sed 's/^[ ]*//g' | sed "s/${tag_name}[ ]*//" | sed 1d)
    data="{\"tag_name\":\"${tag_name}\",\"name\":\"${title}\",\"body\":\"${body//$'\n'/\\n}\"}"
    RELEASE_ID=$(curl -XPOST -H "Authorization: token ${github_api_token}" "${API_URL}" -d"${data}" | jq .id)
    echo "Release ID: ${RELEASE_ID}"
}

function upload_artifact () {
    FILE=$1
    echo "Uploading artifact: ${FILE}"
    curl -H "Authorization: token ${github_api_token}" \
        -H "Content-Type: application/tar+gzip" \
        --data-binary @"${FILE}" \
        "${UPLOADS_URL}/${RELEASE_ID}/assets?name=$(basename $FILE)" > /dev/null
}

if [[ "$CREATE_RELEASE" == "YES" ]]; then
    create_release
else
    RELEASE_ID=$(curl -sXGET ${API_URL}/tags/${TAG} | jq .id)
fi

if [[ "$BUILD" == "YES" ]]; then
    ./build.sh
fi

if [[ "$UPLOAD_ARTIFACTS" == "YES" ]]; then
    for file in $(ls -1 bin/*.tar.gz); do
        upload_artifact $file
    done
fi
