#!/bin/bash

set -e

BINDING_DIR=./pkg/bindings
JSON_DIR=./out

# Install latest abigen
echo "Installing latest abigen..."
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

function create_binding {
    contract_name=$1

    mkdir -p $BINDING_DIR/v1/${contract_name}
    mkdir -p $BINDING_DIR/v2/${contract_name}

    contract_json_path="${JSON_DIR}/${contract_name}.sol/${contract_name}.json"

    binding_out_dir="${BINDING_DIR}/v1/${contract_name}"
    mkdir -p $binding_out_dir || true

    cat $contract_json_path | jq -r '.abi' > $binding_out_dir/tmp.abi

    abigen \
        --abi=$binding_out_dir/tmp.abi \
        --pkg="${contract_name}" \
        --out=$BINDING_DIR/v1/$contract_name/binding.go \
        > /dev/null 2>&1

    if [[ $? == "1" ]];
    then
        echo "Failed to generate v1 binding for $contract_json_path"
    fi

    abigen \
        --v2 \
        --abi=$binding_out_dir/tmp.abi \
        --pkg="${contract_name}" \
        --out=$BINDING_DIR/v2/$contract_name/binding.go \
        > /dev/null 2>&1

    if [[ $? == "1" ]];
    then
        echo "Failed to generate v2 binding for $contract_json_path"
    fi
    rm $binding_out_dir/tmp.abi
}

# Array of contracts to generate bindings for
contracts=(
    "AppController"
    "ComputeAVSRegistrar"
    "ComputeOperator"
    "IReleaseManager"
    "IPermissionController"
)

forge b --no-metadata

for contract_name in "${contracts[@]}"; do
    create_binding $contract_name
done
