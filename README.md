# godacity

`godacity` is a Go library designed to interact with and control **Audacity** via its built-in Pipe Server (mod-script-pipe). This allows developers to automate audio editing, label management, and project manipulation directly from Go code.

> **ALPHA STATE NOTICE**
> This library is currently in an **Alpha** state. The API is subject to breaking changes, and not all Audacity commands are currently implemented. Use with caution in production environments.

## Prerequisites

Before using this library, you must enable the Pipe Server in Audacity:

1. Open Audacity.
2. Go to **Preferences > Modules**.
3. Set **mod-script-pipe** to **Enabled**.
4. Restart Audacity.

## Installation

```bash
go get github.com/DYankee/godacity
```

## Quick Start

```go
package main

import (
    "log"
    "time"
    "github.com/DYankee/godacity"
)

func main() {
    // Configure connection
    cfg := &godacity.Config{
        AutoStart:    true,          // Attempt to start Audacity if not running
        StartTimeout: 10 * time.Second,
    }

    // Initialize the client
    app, err := godacity.NewAudacity(cfg)
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer app.ClosePipes()

    // Example: Import a file and add a label
    err = app.IO.ImportAudio("/path/to/song.wav")
    if err != nil {
        log.Fatal(err)
    }

    err = app.Labels.Add()
    if err != nil {
        log.Fatal(err)
    }
}
```

## Available Services & Functions

The library is organized into specialized services accessible via the main `Audacity` instance.

### `IO` Service (`app.IO`)

Handles importing and exporting files.

- `ImportAudio(filePath string)`: Imports an audio file into the current project.
- `ExportAudio(destination string, fileName string)`: Exports the selected audio as a 2-channel file.

### `Tracks` Service (`app.Tracks`)

Manages track manipulation and selections.

- `SplitNew()`: Splits the current selection into a new track.
- `SelectRegion(startTime, endTime float64)`: Selects a specific time range in the timeline.

### `Labels` Service (`app.Labels`)

Manages label tracks and individual labels.

- `Add()`: Creates a new label at the current cursor position.
- `Set(id int, text string)`: Updates the text of a specific label by its ID.

### `Clips` Service (`app.Clips`)

Interrogates audio clips within tracks.

- `GetClips()`: Returns a slice of `ClipInfo` containing metadata (start, end, track, name) for all clips in the project.

### Core/Low-Level

- `ExecCommand(command string)`: Sends a raw command string to Audacity. Use this if a specific feature isn't yet wrapped in a service.

## Platform Support

- **Windows**: Uses Named Pipes (`\\.\pipe\ToSrvPipe`).
- **Linux/macOS**: Uses Unix Domain Pipes in `/tmp/`.

## Contributing

Contributions are welcome and if you find a command that isn't implemented, feel free to open a PR adding it to the relevant service. Pull requests will be reviewed when I have time. I work and go to school full time so if I haven't looked at your request yet please be patient.

## License

MIT License

Copyright (c) 2026 Zachary Geary

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
