# cmdg-image-render

A high-performance standalone HTML renderer designed for terminal applications. 

This tool is a renderer for [cmdg](https://github.com/ThomasHabets/cmdg) that allows rendering of inline images. It extracts HTML into formatted plain text and identifies the precise terminal coordinates for inline images.

## Features
- **High-Density HTML Rendering:** Converts complex HTML layouts into readable, properly spaced terminal text.
- **Image Metadata Extraction:** Detects inline images and returns their source URLs/CIDs with relative positioning.
- **Performance Optimized:** Uses a custom C-based parser for near-instant execution.
- **JSON Protocol:** Easy integration with any TUI project via a clean JSON I/O interface.

## Installation
Requires Go 1.25+.

```bash
go install github.com/kurokirasama/cmdg-image-render/cmd/cmdg-image-render@latest
```

## Integration with cmdg
To enable HTML rendering and inline images in [cmdg](https://github.com/ThomasHabets/cmdg):
1. Install `cmdg-image-render` using the command above.
2. Ensure the compiled binary is in your system's `PATH`.
3. Run `cmdg` with the image protocol enabled:
   ```bash
   cmdg -image_protocol auto
   ```

## Usage
The tool reads raw HTML from `stdin` and outputs a JSON object to `stdout`.

### Command Line
```bash
echo '<h1>Hello</h1><img src=\"cid:123\">' | cmdg-image-render --width 80
```

### Options
- `--width <int>`: Set the terminal wrapping width (default: 80).
- `--start-index <int>`: Set the starting index for image markers (default: 0).

### Output Format
```json
{
  "rendered_text": "\n\n=== HELLO ===\n\n ##IMG_0_## ",
  "inline_images": [
    {
      "index": 0,
      "source": "cid:123"
    }
  ]
}
```

## License
Apache-2.0
