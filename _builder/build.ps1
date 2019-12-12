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
$gopath = ($env:GOPATH).replace('\','/')
Write-Output "start build ..."
docker run --rm --name=gobuilder -v ${gopath}:/opt/gopath gobuilder `
    env GOOS=windows GOARCH=amd64 go build github.com/Schneereich/merge-intervals