#!/bin/bash

TOKEN=$(curl --silent --header 'GET' "https://auth.docker.io/token?service=registry.docker.io&scope=repository:library/redis:pull" | jq -r '.token')

curl -s --header "Accept: application/vnd.oci.image.manifest.v1+json" --header "Authorization: Bearer ${TOKEN}" 'https://registry-1.docker.io/v2/library/redis/manifests/latest' > redis-manifest.json

# curl --header "Accept: application/vnd.oci.image.manifest.v1+json" --header "Authorization: Bearer ${TOKEN}" 'https://registry-1.docker.io/v2/library/ubuntu/manifests/sha256:77d57fd89366f7d16615794a5b53e124d742404e20f035c22032233f1826bd6a'

# curl --location --header "Accept: application/vnd.oci.image.manifest.v1+json" --header "Authorization: Bearer ${TOKEN}" 'https://registry-1.docker.io/v2/library/ubuntu/blobs/sha256:dafa2b0c44d2cfb0be6721f079092ddf15dc8bc537fb07fe7c3264c15cb2e8e6' --output ubuntu.tar

# curl --silent --location --request 'GET' --header "Authorization: Bearer ${TOKEN}" 'https://registry-1.docker.io/v2/library/ubuntu/blobs/sha256:dafa2b0c44d2cfb0be6721f079092ddf15dc8bc537fb07fe7c3264c15cb2e8e6' --output '1.tar.gz'

# curl --silent --location --request 'GET' --header "Authorization: Bearer ${TOKEN}" 'https://registry-1.docker.io/v2/library/ubuntu/blobs/sha256:80ba5fb4d26ccaa99af4e6314b9d6398e856b913117c7cae74da83b351842a50' --output '2.tar'