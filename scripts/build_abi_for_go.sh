#!/bin/bash

ARTIFACTS_DIR="./contract/artifacts/contracts"
OUTPUT_DIR="./backend/abigen"

mkdir -p $OUTPUT_DIR
mkdir -p $OUTPUT_DIR/abi

for artifact in $(find $ARTIFACTS_DIR -name "*.json" ! -name "*.dbg.json"); do
	abi=$(cat $artifact | jq '.abi')
	if [ -z "$abi" ] || [ "$abi" == "[]" ]; then
		echo -e "Empty or non-existent ABI for $artifact, skipping...\n"
		continue
	fi

	filename=$(basename -- "$artifact")
	contract_name="${filename%.*}"

	pascal_case=$contract_name
	snake_case=$(echo $contract_name | awk '{gsub(/([A-Z])/,"_&"); print tolower($0)}' | sed 's/^_//')

	echo $abi >$OUTPUT_DIR/abi/$pascal_case.abi
	echo "ABI saved for $contract_name at: $OUTPUT_DIR/abi/$pascal_case.abi"

	command="abigen --abi=$OUTPUT_DIR/abi/$pascal_case.abi --out=$OUTPUT_DIR/${snake_case}.go --pkg=abigen --type=$pascal_case"
	echo "$command"
	eval "$command"

	echo -e "Go binding generated for $pascal_case at: $OUTPUT_DIR/${snake_case}.go\n"
done
