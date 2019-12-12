cd $PSScriptRoot

function error{
    param([string]$Error)
    Write-Output "ERROR: $(get-date -Format u) $Error"
}

if (-Not (Test-Path env:GOPATH)) {
    error "GOPATH is not defined, please set GOPATH and retry"
    Exit 1
}

docker build -t gobuilder .
Write-Output "start build ..."
docker run --rm --name=gobuilder -v $env:GOPATH:/opt/gopath gobuilder