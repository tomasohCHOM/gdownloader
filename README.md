# gdownloader

A tool to download Google Drive files from the command line (mainly for my personal usage).

## Setup

Enable the [Google Drive API](https://developers.google.com/workspace/drive/api/quickstart/go#enable_the_api). Once you retrieve `credentials.json`, place them inside `~/.gdownloader`. Next, install the tool:

```
go install github.com/tomasohCHOM/gdownloader@latest
```

## Usage

```
gdownloader
```

You will be prompted to authorize access to your Google Drive the first time you run it.

List the available commands for this tool:

```
gdownloader --help
```

---

Developed with 🔥 by Chom.
