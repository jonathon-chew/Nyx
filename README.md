# 🌃 Nyx (Go)

This is a simple Go script that fetches NASA's **Astronomy Picture of the Day (APOD)** and downloads it to your local machine. It uses the APOD API to retrieve the metadata and image URL, then saves the image with its title as the filename.

## 🚀 Features

- Fetches the APOD metadata using the NASA API
- Downloads the image if the media type is `"image"`
- Saves the image with a filename based on the APOD title
- Simple and clean command-line tool

## 🛠️ Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.16+ recommended)
- A NASA API key (get one for free at [api.nasa.gov](https://api.nasa.gov/))

## 📁 Setup

1. Clone this repository:

   ```bash
   git clone https://github.com/jonathon-chew/Nyx.git
   cd Nyx 
   ```

2. Store your NASA API key in a local, untracked file or environment variable before running the app.

```text
Example format:

https://api.nasa.gov/planetary/apod?api_key=YOUR_API_KEY
```

3. Run the script:

    `go run main.go`

## 📂 Output

If the media type of the APOD is an image, it will be downloaded and saved with its title as the filename, for example `A Journey Through the Milky Way.jpg`.

The script also prints basic metadata like title, media type, and URL to the console.

## 🧠 Notes

Currently, the script only handles media of type "image". If the APOD is a video (e.g. a YouTube link), it will skip downloading.

Make sure the APOD title is a valid filename. If needed, sanitize the title in the code before using it as a filename.

The current script will not run until a valid NASA API key is available.

## 📜 License

This project is licensed under the MIT License. See the LICENSE file for details.
