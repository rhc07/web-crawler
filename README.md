
# Web-crawler
A classic web-crawler that uses Golang to output HTML properties

## Introduction

> A [_web crawler_](https://en.wikipedia.org/wiki/Web_crawler) is the portion
> of a search engine that scans web pages looking for links and then follows
> them. It would normally store the data it finds into some datastore where it
> can be useful.

## Tasks

1. Parse all the `href` attributes within anchor elements in the retrieved page
1. Output a list of all of the `href` attributes found
1. Refactor the existing code in to a simple CLI which can take a URL to parse
   as an argument

**Note** that the use of the standard library (including `golang.org/x/...`
packages) is strongly encouraged. 3rd-party community packages should be
avoided if possible (apart from testing related packages).

You can run the code using:

```bash
make run
```
