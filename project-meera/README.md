# Video File Info App

A command-line application that takes a file path and returns detailed video metadata information.

## Overview

Build an app that accepts a video file path, validates it exists and is a supported video format, then extracts and returns comprehensive metadata about the file.

## Requirements

### Core Functionality

1. **Input Handling**
   - Accept a file path (absolute or relative) as a string parameter
   - Support command-line argument parsing
   - Convert relative paths to absolute paths internally

2. **Input Validation & Security**
   - Sanitize input path to prevent directory traversal attacks (`../../../etc/passwd`)
   - Reject non-string inputs and malformed paths
   - Validate path contains no suspicious characters or sequences

3. **File Operations**
   - Verify file exists at the specified path
   - Check file is readable with current permissions
   - Handle cross-platform path resolution (Windows vs Unix-like systems)

4. **Format Detection**
   - Support common video container formats: MP4, WebM, AVI, MOV
   - Detect video format from file headers/magic bytes (not just file extension)
   - Reject files that don't match supported video encoding types

5. **Metadata Extraction**
   - Extract and return structured metadata containing:
     - Original filename
     - Absolute file path
     - File size (in bytes)
     - Duration (in seconds)
     - Video resolution (width x height)
     - Video codec (H.264, H.265, VP8, VP9, etc.)
     - Audio codec (if present)
     - Video bitrate
     - Frame rate

### Error Handling & Exit Codes

- **Exit Code 0**: Success - valid video file processed
- **Exit Code 1**: File not found at specified path
- **Exit Code 2**: Permission denied - cannot read file
- **Exit Code 3**: Unsupported video format
- **Exit Code 4**: Invalid input or malformed path
- **Exit Code 5**: General application error

### Output Requirements

- Return structured output in JSON format
- Include human-readable error messages for debugging
- Log processing steps when verbose flag is enabled
- Handle graceful shutdown on interrupt signals

### Performance & Resource Requirements

- Must not require elevated/root permissions in default mode
- Should complete file analysis within reasonable time (< 5 seconds for typical video files)
- Memory usage should be bounded (don't load entire file into memory)

### Cross-Platform Support

- Work consistently across Windows, macOS, and Linux
- Handle platform-specific path separators correctly
- Use appropriate system calls for file operations on each platform

## Technical Decisions

### Language Choice

- **Selected**: Go
- **Rationale**: Excellent cross-platform support, simple deployment, strong standard library for file operations

### Metadata Extraction Strategy

- **Decision Point**: Parse video metadata directly from binary formats OR use external tool/library
- **Options**:
  - Direct binary parsing (more learning, higher complexity)
  - Library/tool integration like `ffprobe` (faster development, less OS-level learning)

## Usage Examples

```bash
# Basic usage
./video-info /path/to/video.mp4

# With verbose logging
./video-info --verbose /path/to/video.mp4

# Relative path
./video-info ../videos/sample.webm
```

## Expected Output Format

```json
{
  "filename": "sample.mp4",
  "absolute_path": "/home/user/videos/sample.mp4",
  "file_size": 15728640,
  "duration": 120.5,
  "resolution": {
    "width": 1920,
    "height": 1080
  },
  "video_codec": "H.264",
  "audio_codec": "AAC",
  "bitrate": 5000000,
  "frame_rate": 30.0
}
```

## Development Phases

1. **Phase 1**: Basic file operations (exists, readable, path resolution)
2. **Phase 2**: Input validation and error handling
3. **Phase 3**: Video format detection from file headers
4. **Phase 4**: Metadata extraction implementation
5. **Phase 5**: JSON output formatting and CLI polish
6. **Phase 6**: Cross-platform testing and optimization

## Future Extensions

- Filename-based filesystem search functionality
- Support for additional video formats
- Batch processing of multiple files
- Integration with video processing pipelines
