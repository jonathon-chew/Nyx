# 📸 Nyx (Go)

This is a simple Go script that fetches NASA's **Astronomy Picture of the Day (APOD)** and downloads it to your local machine. It uses the APOD API to retrieve the metadata and image URL, and saves the image with its title as the filename.

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
   git clone https://github.com/your-username/nasa-apod-downloader.git
   cd nasa-apod-downloader
   ```

2. Store your NASA API URL (with the API key) on the 5th line of a text file located at: ../scripts/key.txt

```text
Example contents of ../scripts/key.txt:

- line 1
- line 2
- line 3
- line 4
- https://api.nasa.gov/planetary/apod?api_key=YOUR_API_KEY
```

3. Run the script:

    go run main.go

## 📂 Output

If the media type of the APOD is an image, it will be downloaded and saved with its title as the filename (e.g. A Journey Through the Milky Way.jpg).

The script also prints basic metadata like title, media type, and URL to the console.

## 🧠 Notes

Currently, the script only handles media of type "image". If the APOD is a video (e.g. a YouTube link), it will skip downloading.

Make sure the APOD title is a valid filename. If needed, sanitize the title in the code before using it as a filename.

## 📜 License

This project is licensed under the MIT License. See the LICENSE file for details.
