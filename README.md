# Web Crawler

A simple, concurrent web crawler written in Go. This program starts at a given URL, downloads the content of the web page, saves it to a specified directory, and extracts valid URLs from the page. It then recursively crawls the extracted URLs and repeats the process. The program ensures that URLs are not visited multiple times by using a thread-safe data structure.

## Installation

To install the web crawler, make sure you have [Go installed](https://golang.org/doc/install) on your system. 
Then, to clone the repository and install webcrawler run the following commands:

```bash
go get github.com/oliverasistemas/webcrawler
cd webcrawler
go install
```

## Usage

To run the web crawler, execute the following command:

```bash
webcrawler <starting-url> <destination-directory>
```

Replace <starting-url> with the URL you want to start crawling from and <destination-directory> with the directory where you want to save the downloaded content.

For example, to start crawling from https://example.com and save the content in a directory named output:
```bash
webcrawler https://example.com output
```

## Dependencies
The web crawler uses two custom packages:

- crawler/pagedownloader: Responsible for downloading the content of a URL and saving it to a specified directory.
- crawler/url_extractor: Responsible for extracting valid URLs from the downloaded content.

## Limitations
- This web crawler is a basic implementation and is not optimized for large-scale web crawling.
- It does not handle JavaScript rendering or dynamic content.
- There is no support for robots.txt or other crawling rules.
- Support resume functionality by checking the destination directory for downloaded pages and skip downloading and processing where not necessary.