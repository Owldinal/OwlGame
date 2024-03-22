#!/bin/bash

ARTIFACTS_DIR="./contract/artifacts/contracts"
OUTPUT_DIR="./indexer/abi"

for artifact in $(find $ARTIFACTS_DIR -name "*.json" ! -name "*.dbg.json"); do
	abi=$(cat $artifact | jq '.abi')

	filename=$(basename -- "$artifact")
	contract_name="${filename%.*}"

	pascal_case=$contract_name
	snake_case=$(echo $contract_name | awk '{gsub(/([A-Z])/,"_&"); print tolower($0)}' | sed 's/^_//')

	echo $abi >$OUTPUT_DIR/$pascal_case.abi

	echo "ABI saved for $contract_name at: $OUTPUT_DIR/$pascal_case.abi"

	echo "abigen --abi=$OUTPUT_DIR/$pascal_case.abi --out=$OUTPUT_DIR/${snake_case}.go --pkg=abi --type=$pascal_case"
	abigen --abi=$OUTPUT_DIR/$pascal_case.abi --out=$OUTPUT_DIR/${snake_case}.go --pkg=abi --type=$pascal_case

	echo "Go binding generated for $pascal_case at: $OUTPUT_DIR/${snake_case}.go"
done
