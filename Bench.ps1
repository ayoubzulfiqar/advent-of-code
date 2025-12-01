# PowerShell Script to Measure Execution Time of Different Language Implementations
# Save this as: Measure-Performance.ps1

# Configuration - Set your folder paths here
$folders = @{
    "Go"      = "D:\Projects\advent-of-code\2025\Go\Day1"
    "Python"  = "D:\Projects\advent-of-code\2025\Python\Day1" 
    "Dart"    = "D:\Projects\advent-of-code\2025\Dart\Day1"
}

# Files to test (part_1 and part_2 in each language)
$testFiles = @("part_1", "part_2")

# Results storage
$results = @()

# ANSI color codes for better output
$colors = @{
    "Reset" = "`e[0m"
    "Red" = "`e[31m"
    "Green" = "`e[32m"
    "Yellow" = "`e[33m"
    "Blue" = "`e[34m"
    "Magenta" = "`e[35m"
    "Cyan" = "`e[36m"
    "White" = "`e[37m"
    "Bold" = "`e[1m"
}

# Function to print colored output
function Write-Color {
    param(
        [string]$Message,
        [string]$Color = "White",
        [switch]$NoNewLine
    )
    
    $colorCode = $colors[$Color]
    $resetCode = $colors["Reset"]
    
    if ($NoNewLine) {
        Write-Host "$colorCode$Message$resetCode" -NoNewline
    } else {
        Write-Host "$colorCode$Message$resetCode"
    }
}

# Function to compile Dart if needed
function Compile-Dart {
    param([string]$filePath)
    
    $dartFile = "$filePath.dart"
    $compiledFile = "$filePath.exe"
    
    if (Test-Path $dartFile) {
        Write-Color "  Compiling Dart file: $dartFile" "Cyan"
        dart compile exe $dartFile -o $compiledFile
        return $compiledFile
    }
    return $null
}

# Function to compile Go if needed
function Compile-Go {
    param([string]$filePath)
    
    $goFile = "$filePath.go"
    $compiledFile = "$filePath.exe"
    
    if (Test-Path $goFile) {
        Write-Color "  Compiling Go file: $goFile" "Cyan"
        go build -o $compiledFile $goFile
        return $compiledFile
    }
    return $null
}

# Function to measure execution time
function Measure-ExecutionTime {
    param(
        [string]$Language,
        [string]$TestName,
        [string]$Command,
        [string]$WorkingDirectory
    )
    
    Write-Color "  Running: $Command" "Yellow"
    
    try {
        # Create a stopwatch to measure execution time
        $stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
        
        # Run the command and capture output
        $process = Start-Process -FilePath "powershell.exe" `
                                 -ArgumentList "-Command `"$Command`"" `
                                 -WorkingDirectory $WorkingDirectory `
                                 -NoNewWindow `
                                 -PassThru `
                                 -Wait `
                                 -RedirectStandardOutput "output_$Language.txt" `
                                 -RedirectStandardError "error_$Language.txt"
        
        $stopwatch.Stop()
        
        # Read and display output
        $output = Get-Content "output_$Language.txt" -ErrorAction SilentlyContinue
        $errorOutput = Get-Content "error_$Language.txt" -ErrorAction SilentlyContinue
        
        if ($errorOutput) {
            Write-Color "  Error: $errorOutput" "Red"
            return $null
        }
        
        if ($output) {
            Write-Color "  Output: $($output -join ' ')" "Green"
        }
        
        # Cleanup temporary files
        Remove-Item "output_$Language.txt", "error_$Language.txt" -ErrorAction SilentlyContinue
        
        return [PSCustomObject]@{
            Language = $Language
            TestName = $TestName
            ExecutionTime = $stopwatch.Elapsed.TotalMilliseconds
            Output = $output -join ' '
        }
        
    } catch {
        Write-Color "  Failed to execute: $($_.Exception.Message)" "Red"
        return $null
    }
}

# Main execution
Clear-Host
Write-Color "=== Language Performance Benchmark ===" "Bold" "Cyan"
Write-Color "Testing implementations in: $($folders.Keys -join ', ')" "White"
Write-Color "=" * 50 "White"

# Check if input.txt exists in each folder and copy if not
foreach ($lang in $folders.Keys) {
    $folderPath = $folders[$lang]
    $inputFile = Join-Path $folderPath "input.txt"
    
    if (-not (Test-Path $inputFile)) {
        Write-Color "Warning: input.txt not found in $lang folder ($folderPath)" "Yellow"
        Write-Color "Please copy input.txt to: $folderPath" "Yellow"
        Write-Color "=" * 50 "White"
    }
}

# Test each language and each part
foreach ($test in $testFiles) {
    Write-Color "`n=== Testing $test ===" "Bold" "Magenta"
    Write-Color "-" * 40 "White"
    
    foreach ($lang in $folders.Keys) {
        $folderPath = $folders[$lang]
        $fullTestPath = Join-Path $folderPath $test
        
        Write-Color "`n[$lang]" "Bold" "Blue"
        
        $result = $null
        
        switch ($lang) {
            "Go" {
                # Compile Go
                $compiledExe = Compile-Go $fullTestPath
                if ($compiledExe -and (Test-Path $compiledExe)) {
                    $command = ".\$test.exe"
                    $result = Measure-ExecutionTime -Language $lang -TestName $test -Command $command -WorkingDirectory $folderPath
                    
                    # Cleanup compiled executable
                    Remove-Item $compiledExe -ErrorAction SilentlyContinue
                } else {
                    Write-Color "  Go file not found or failed to compile: $fullTestPath.go" "Red"
                }
            }
            
            "Python" {
                $pythonFile = "$fullTestPath.py"
                if (Test-Path $pythonFile) {
                    $command = "python `"$pythonFile`""
                    $result = Measure-ExecutionTime -Language $lang -TestName $test -Command $command -WorkingDirectory $folderPath
                } else {
                    Write-Color "  Python file not found: $pythonFile" "Red"
                }
            }
            
            "Dart" {
                # Compile Dart to exe first
                $compiledExe = Compile-Dart $fullTestPath
                if ($compiledExe -and (Test-Path $compiledExe)) {
                    $command = ".\$test.exe"
                    $result = Measure-ExecutionTime -Language $lang -TestName $test -Command $command -WorkingDirectory $folderPath
                    
                    # Cleanup compiled executable
                    Remove-Item $compiledExe -ErrorAction SilentlyContinue
                } else {
                    # Try running with dart interpreter
                    $dartFile = "$fullTestPath.dart"
                    if (Test-Path $dartFile) {
                        $command = "dart run `"$dartFile`""
                        $result = Measure-ExecutionTime -Language $lang -TestName $test -Command $command -WorkingDirectory $folderPath
                    } else {
                        Write-Color "  Dart file not found: $dartFile" "Red"
                    }
                }
            }
        }
        
        if ($result) {
            $results += $result
            Write-Color "  Time: $($result.ExecutionTime.ToString('F3')) ms" "Green"
        }
    }
}

# Display summary
Write-Color "`n=== PERFORMANCE SUMMARY ===" "Bold" "Cyan"
Write-Color "=" * 50 "White"

if ($results.Count -eq 0) {
    Write-Color "No results to display. Check if files exist and can be executed." "Red"
    exit 1
}

# Group results by test
foreach ($test in $testFiles) {
    $testResults = $results | Where-Object { $_.TestName -eq $test }
    
    if ($testResults.Count -gt 0) {
        Write-Color "`n[$test Results]" "Bold" "Magenta"
        
        # Sort by fastest time
        $sortedResults = $testResults | Sort-Object ExecutionTime
        
        # Display results
        $rank = 1
        foreach ($result in $sortedResults) {
            $color = switch ($rank) {
                1 { "Green" }
                2 { "Yellow" }
                3 { "White" }
                default { "White" }
            }
            
            $medal = switch ($rank) {
                1 { "🥇 " }
                2 { "🥈 " }
                3 { "🥉 " }
                default { "   " }
            }
            
            Write-Color "  $rank. $medal$($result.Language): $($result.ExecutionTime.ToString('F3')) ms" $color
            $rank++
        }
        
        # Calculate speed differences
        $fastest = $sortedResults[0]
        $slowest = $sortedResults[-1]
        
        if ($sortedResults.Count -gt 1) {
            $speedup = $slowest.ExecutionTime / $fastest.ExecutionTime
            Write-Color "  Fastest ($($fastest.Language)) is $($speedup.ToString('F2'))x faster than slowest ($($slowest.Language))" "Cyan"
        }
    }
}

# Overall champion
$overallResults = $results | Group-Object Language | ForEach-Object {
    [PSCustomObject]@{
        Language = $_.Name
        AvgTime = ($_.Group.ExecutionTime | Measure-Object -Average).Average
        TotalTime = ($_.Group.ExecutionTime | Measure-Object -Sum).Sum
    }
} | Sort-Object AvgTime

Write-Color "`n=== OVERALL CHAMPION ===" "Bold" "Cyan"
Write-Color "=" * 50 "White"

if ($overallResults.Count -gt 0) {
    Write-Color "🥇 OVERALL FASTEST: $($overallResults[0].Language)" "Bold" "Green"
    Write-Color "  Average time: $($overallResults[0].AvgTime.ToString('F3')) ms" "Green"
    
    Write-Color "`nRankings by average execution time:" "White"
    $rank = 1
    foreach ($result in $overallResults) {
        $medal = switch ($rank) {
            1 { "🥇 " }
            2 { "🥈 " }
            3 { "🥉 " }
            default { "   " }
        }
        
        Write-Color "  $rank. $medal$($result.Language): $($result.AvgTime.ToString('F3')) ms (total: $($result.TotalTime.ToString('F3')) ms)" "White"
        $rank++
    }
}

# Save results to CSV for further analysis
$timestamp = Get-Date -Format "yyyyMMdd_HHmmss"
$csvPath = "performance_results_$timestamp.csv"
$results | Export-Csv -Path $csvPath -NoTypeInformation

Write-Color "`nResults saved to: $csvPath" "Cyan"
Write-Color "Benchmark completed! 🏁" "Bold" "Green"