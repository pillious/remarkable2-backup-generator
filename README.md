# remarkable2-backup-generator

<i><small>Realistically this tool can be used to backup any files.</small></i>

## Overview

The backup generator incrementally generates backups using rsync (with the --link-dest flag).
If you are interested in _only_ backing up the notebooks on your Remarkable2, the notebooks are stored at `/home/root/.local/share/remarkable/xochitl`. The directory to copy is set using the `-src` flag.

The first time a backup is made, the entire directory set by `-src` is copied from your rm2, so this may take a while.
The name of the backup is always set to the current timestamp string in ISO 8601 format.
Once the backup is completed, a `.latest_backup` file is created to store the name of your newest backup.
On subsequent backups, the file listed in `.latest_backup` is used as the --link-dest value for rsync by default.

To automate the entire backup process, make sure to set up passwordless SSH into your Remarkable ([Resources](#resources)).

Intersted in backing up your notebooks as PDFs? Check out [RM2 PDF downloader](https://github.com/pillious/remarkable2-pdf-downloader)

## Requirements

1. Have Golang installed.
2. Have the rsync tool available in your terminal (use WSL if you're using Windows). 
3. Make sure you can SSH into your Remarkable ([Guide](https://remarkable.jms1.info/info/ssh.html)).

## Tool Usage

### Compile the tool
`go install ./`

The executable will the compiled to `${GOPATH}/bin/` or `go/bin/`

### Example usage
`./remarkable2-backup-generator -l -v -src="root@192.168.x.x:/home/root/.local/share/remarkable/xochitl" -backupsDir="/home/user/rm2-backups/"`

### Tool help
`./remarkable2-backup-generator -h`

## Backups Structure

The location of the backups is set using the `-backupsDir` flag.

```
.
├── .latest_backup
│
├── logs/
│   └──  2024-07-19T01:53:05Z.logs
│   └──  2024-07-22T01:39:27Z.logs
│
├── 2024-07-19T01:53:05Z/
│   └──  xochitl/           
│        └──  ...
├── 2024-07-22T01:39:27Z
│   └──  xochitl/
│        └──  ...
└── ...
```

## Acknowledgement

[Passwordless SSH Setup](https://remarkable.jms1.info/info/ssh.html)

[Rsync Info](https://remarkable.jms1.info/info/backups.html)
