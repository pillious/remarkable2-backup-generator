# remarkable2-backup-generator

## Overview

The backup generator incrementally generates backups using rsync (with the --link-dest flag).
If you are only interested in backing up the notebooks on your Remarkable2, the notebooks are stored at `/home/root/.local/share/remarkable/xochitl`. The directory to copy is set using the `-src` flag.

The first time a backup is made, the entire directory set by `-src` is copied from your rm2, so this may take a while.
The name of the backup is always set to the current timestamp string in ISO 8601 format.
Once the backup is completed, a `.latest_backup` file is created to store the name of your newest backup.
On subsequent backups, the file listed in `.latest_backup` is used as the --link-dest value for rsync by default.

To automate the entire backup process, make sure to set up passwordless SSH into your Remarkable ([Resources](#resources)).

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

## Tool Usage

### Compile the tool
`go install ./`

The executable will the compiled to `${GOPATH}/bin/` or `go/bin/`

### Example usage
`./remarkable2-backup-generator -l -v -src="root@192.168.1.11:/home/root/.local/share/remarkable/xochitl" -backupsDir="/home/user/rm2-backups/"`

### Tool help
`./remarkable2-backup-generator -h`

## Acknowledgement

[Passwordless SSH Setup](https://remarkable.jms1.info/info/ssh.html)

[Rsync Info](https://remarkable.jms1.info/info/backups.html)