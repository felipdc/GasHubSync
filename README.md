# GasHubSync

GasHubSync is an open-source project that synchronizes Google App Script projects with a GitHub repository. This project is written in Go.

## Overview

GasHubSync provides a solution for developers who want to maintain version control on their Google Apps Script projects. It allows a Google Apps Script project to be synced with a GitHub repository. Any changes to the Google Apps Script project files are automatically committed to the GitHub repository.

## Project Status

This project is currently in development. We encourage you to contribute! Please feel free to submit issues and pull requests.

## How to use

To use GasHubSync, you need to provide your Google App Script ID and the URL of your GitHub repository. 

**Usage:**

```shell
go run . <google-app-script-id> <github-repo-url>
```

Replace `<google-app-script-id>` with the ID of your Google Apps Script project, and `<github-repo-url>` with the URL of your GitHub repository.

**Example**

```shell
go run . 1fYGuyTrcLSSiwhjhfjhskj  https://github.com/username/repo.git
```

## Configuration

You can configure the polling interval (how often GasHubSync checks for changes in your Google Apps Script project) by adjusting the value in the `config.json` file.

**config.json**

```json
{
    "polling_interval_seconds": 5
}
```

In the example above, GasHubSync will check for changes every 5 seconds. Adjust the polling_interval_seconds value to your preference.
