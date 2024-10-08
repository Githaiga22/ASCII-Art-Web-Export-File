# ASCII Art Web Export File

## Overview

The ASCII Art Web Export File project is an extension of the ASCII Art Web application. This project allows users to generate ASCII art from text input on a web interface and export the generated ASCII art as a downloadable file. The project focuses on implementing file export functionality with proper HTTP headers and permissions, ensuring the file can be exported in the desired format with the correct settings.

## Objectives

- **File Export:** Ensure that users can export the ASCII art generated by the web application in at least one export format of your choice (e.g., `.txt`).
- **File Permissions:** The exported file must have the correct read and write permissions for the user.
- **HTTP Headers:** Utilize appropriate HTTP headers (`Content-Type`, `Content-Length`, `Content-Disposition`) when sending the file as part of the HTTP response.
- **Error Handling:** Implement proper error handling for the web server, ensuring a smooth user experience.
- **Good Practices:** Ensure the code follows best practices in Go programming and web development.

## Features

- **ASCII Art Generation:** Users can enter text on the website to generate ASCII art in various styles.
- **Export to File:** A new HTTP endpoint allows users to export the generated ASCII art as a downloadable file with appropriate HTTP headers.
- **Responsive Design:** The web interface is designed to be responsive and user-friendly.
- **Error Handling:** Comprehensive error handling is implemented to manage user input and server issues.

## Implementation Details

- **Web Server:** The web server is implemented in Go, using only standard Go packages.
- **Export Format:** The application supports exporting the ASCII art as a `.txt` file.
- **HTTP Headers:** The server correctly sets `Content-Type`, `Content-Length`, and `Content-Disposition` headers for file export.
- **Permissions:** The exported file is created with the correct read and write permissions for the user.

## Getting Started

### Prerequisites

- Go installed on your system.
- A working Go environment.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Githaiga22/ASCII-Art-Web-Export-File.git
2. Navigate to the project directory:
```bash
cd ascii-art-web-export-file
```
3. Run the web server
```bash
 go run main.go
```
4. Open a web browser and navigate to `http://localhost:8080` to access the page

5. Enter the text you want to convert into ASCII art.
6. Select the art style and click "Generate".
7. Once the ASCII art is generated, click the "Export" button to download the file.
## Authors

- allkamau
- mmofat

## Github Profiles

- GitHub Profile: [Githaiga22](https://github.com/Githaiga22)
- Github Profile: [mmofat](https://github.com/mokwathedeveloper/)