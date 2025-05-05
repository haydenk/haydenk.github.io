+++
title = 'sync_s3_objects.sh'
date = 2025-05-05T14:38:56Z
type = "snippet"
+++

This assumes you have already fetched the list of s3 objects with `aws s3api list-objects-v2` and saved the json result to files.json

```sh
jsonFile="$HOME/files.json"
sourceBucket="your-source-bucket-name"
targetBucket="your-target-bucket-name"

function keyExists {
    local searchKey=$1;
    aws s3api list-objects-v2 --bucket $targetBucket --query "length(Contents[?ends_with(Key, \`$searchKey\`)])"
}

for object in `jq -cr '.[]' $jsonFile`;
do
    key=$(echo $object | jq -cr '.Key')
    searchFilename=$(basename $key)
    hasKey=$(keyExists $searchFilename)

    # echo $searchFilename

    if [[ $hasKey -gt 0 ]]; then
        continue
    fi

    lastModified="$(echo $object | jq -cr '.LastModified')"
    copySource="$sourceBucket/$key"
    tarketKey="usage/$(date -d "$lastModified" '+year=%Y/month=%m/day=%d')/$searchFilename"

    echo "$copySource -> $targetBucket/$tarketKey"
    aws s3api copy-object --copy-source $copySource --key $tarketKey --bucket $targetBucket | jq .
done
```
