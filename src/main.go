package main

import (
  "encoding/json"
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "strings"
  "bufio"
  "path/filepath"
)

type Response struct{
  Copyright string
  Date string
  Explanation string
  Hdurl string
  Media_type string
  Service_version string
  Title string
  Url string
}

func downloadImage(fileName string, response *http.Response){

  outFile, err := os.Create(fileName)
  
  if err != nil{
    log.Fatal(err)
  }

  defer outFile.Close()  
  
  _, err = io.Copy(outFile, response.Body)
  if err != nil{
    log.Fatal(err)
  }
  
  fmt.Printf("Image successfully downloaded and saved as %s", fileName)

}

func main(){

  file, err := os.Open("./scripts/key.txt")
  if err != nil{
    log.Fatal("Failed to open file: %s", err)
  }

  defer file.Close()

  // Create a scanner to read the file line by line
  scanner := bufio.NewScanner(file)

  lineNumber := 0
  var key string

  // Read lines one by one
  for scanner.Scan() {
    if scanner.Text() == "Line 5 should be a full URL that you can get from the NASA website by signing up for the API key" {
      fmt.Println("Please enter a valid URL with API key from the NASA pictre of the day website on line 5 in the scripts/key.txt file")
      return
    }
    lineNumber++
    if lineNumber == 5 { // Get the 5th line
      key = scanner.Text()
      break // Stop reading after the 5th line
    }
  }

  // Handle possible scanning errors
  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  var url string = strings.TrimSpace(key)
  
  resp, err := http.Get(url)

  if err != nil{
    log.Fatal(err)
  }

  defer resp.Body.Close()

  // Read the response body
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    log.Fatalf("Error reading response body: %v", err)
  }

  // Parse JSON into struct
  var todaysResponse Response 
  if err := json.Unmarshal(body, &todaysResponse); err != nil {
    log.Fatalf("Error unmarshalling JSON: %v", err)
  }

  if todaysResponse.Media_type == "image"{

    imageData, err := http.Get(todaysResponse.Url)
    if err != nil{
      log.Fatal(err)
    }

    var fileName string

    extension := strings.Split(todaysResponse.Url, ".")
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal("Unable to determine user home directory:", err)
    }

    fileName = filepath.Join(homeDir, "Picturess")
    fileName = fileName + todaysResponse.Title + "." + extension[len(extension) - 1]

    downloadImage(fileName, imageData)

    // Print parsed data
    fmt.Printf("Title: %s\nMedia Type: %s\nURL: %s\n", todaysResponse.Title, todaysResponse.Media_type, todaysResponse.Url)

  }

}
