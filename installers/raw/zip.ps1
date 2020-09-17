#!/usr/bin/env pwsh

$binhash = @{}

$binhash.Add("tusk_start.exe", "../../../tusktusk_start.exe")
$binhash.Add("undra_start.exe", "../../../undra/undra_start.exe")
$binhash.Add("oat_start.exe", "../../../oat/oat_start.exe")

$binarr = @()

foreach ($i in $binhash.GetEnumerator()) {
    Move-Item -Path $i.Value -Destination $i.Name
    $binarr += $i.Name
}

$compress = @{
    Path = $binarr

    DestinationPath = "raw.zip"
}

make clean
$compress.Path += "../../../"
$compress.Path += "../../../tuskommstd"
$compress.Path += "../../../undra/undrastd"
$compress.Path += "../../../undra/undra.ps1"
$compress.Path += "../../../tuskomm.ps1"
$compress.Path += "../../../oat/oat.ps1"

Compress-Archive @compress