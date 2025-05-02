# Define parameters
$extension = "mp4"
$root = "F:\整理\dance"

function Process-Files {
    param (
        [string]$dir,
        [string]$transpose_value
    )

    # Get all files with specified extension in the directory
    Get-ChildItem -Path $dir -Filter "*.$extension" | ForEach-Object {
        $file = $_.FullName
        $basename_without_ext = [System.IO.Path]::GetFileNameWithoutExtension($file)
        $aftername = "${basename_without_ext}_rotate.mp4"
        $outputPath = Join-Path $dir $aftername

        Write-Host "文件名： $basename_without_ext"
        Write-Host "扩展名： $extension"

        # Execute ffmpeg command
        ffmpeg -i $file -vf "transpose=$transpose_value" -c:v h264_nvenc -preset medium -c:a copy -map_chapters -1 $outputPath

        if ($LASTEXITCODE -eq 0) {
            Write-Host "旧文件名: $($_.Name) 新文件名: $aftername"
            Write-Host "done"
            Remove-Item $file
        } else {
            Write-Host "ffmpeg command failed for file: $file"
        }
        Write-Host "done"
    }
}

# Process files in toRight and toLeft directories
Process-Files -dir "$root\toRight" -transpose_value 1
Process-Files -dir "$root\toLeft" -transpose_value 2