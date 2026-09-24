<!---
  Copyright 2024 OCode

  SPDX-License-Identifier: Apache-2.0
-->

# LogSeeker: Your Ultimate Log File Analysis Tool

LogSeeker is a powerful application designed to simplify and streamline the analysis of log files. Whether you're a system administrator, developer, or IT professional, LogSeeker provides the tools you need to quickly and efficiently interpret log data.

## Installation

To install LogSeeker, follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/yourusername/LogSeeker.git
```

2. Navigate to the project directory:

```sh
cd LogSeeker
```

3. Build the application:

```sh
go build -v .
```

## Usage

To start using LogSeeker, run the following command:

```sh
./log-seeker /path/to/logfile
```

For custom log formats, pass a regex pattern to the `analyze`, `analyze-type-log`, or `analyze-date-log` commands:

```sh
./log-seeker analyze /path/to/logfile --pattern '^(\\S+) (\\w+) (\\S+) (.*?) (\\{.*\\})$'
```

For standard presets, use `--format`:

```sh
./log-seeker analyze /path/to/logfile --format bracketed
```

You cannot use `--pattern` and `--format` together.

Available standard formats:
- `bracketed`: `[datetime] [level] [source] [message] [metadata]`
- `csv`: `datetime,level,source,message,metadata`
- `pipe`: `datetime|level|source|message|metadata`
- `kv`: `datetime=<value> level=<value> source=<value> message="<value>" metadata=<value>`

CSV examples:

Example CSV log lines:

```text
2026-09-24T10:00:00Z,INFO,api,request completed,{"status":200}
2026-09-24T10:05:00Z,ERROR,worker,job failed,{"job_id":42}
2026-09-24T10:10:00Z,WARNING,scheduler,retrying task,{"attempt":2}
```

Analyze all CSV logs:

```sh
./log-seeker analyze /path/to/logfile.csv --format csv
```

Filter CSV logs by level:

```sh
./log-seeker analyze-type-log /path/to/logfile.csv ERROR --format csv
```

Filter CSV logs by date range (RFC3339):

```sh
./log-seeker analyze-date-log /path/to/logfile.csv 2026-09-24T10:00:00Z 2026-09-24T11:00:00Z --format csv
```

The pattern must include exactly 5 capture groups in this order:
1. datetime
2. level
3. source
4. message
5. metadata

Default built-in pattern:

```regex
\[(.*?)\] \[(.*?)\] \[(.*?)\] \[(.*?)\] \[(.*?)\]
```

## Contributing

We welcome contributions! Please read our [contributing guidelines](CONTRIBUTING.md) for more details.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
