#!/usr/bin/env pwsh

$compress = @{
    Path = 
        "../../../omm/omm_start.exe", 
        "../../../omm/omm.ps1", 
        "../../../omm/ommstd", 
        "../../../undra/undrastd", 
        "../../../undra/undra_start.exe",
        "../../../undra/undra.ps1", 
        "../../../oat/oat.ps1",
        "../../../oat/oat_start.exe"

    DestinationPath = "raw.zip"
}
Compress-Archive @compress