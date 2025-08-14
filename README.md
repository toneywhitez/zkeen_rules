# Rule Lists

A Go utility for converting V2Ray DAT files (GeoSite and GeoIP) into rule list files suitable for various proxy and firewall configurations. 

This repository includes a GitHub Action (`.github/workflows/create-release.yml`) that automatically fetches and processes DAT files [zkeen-domains](https://github.com/jameszeroX/zkeen-domains) and [zkeen-ip](https://github.com/jameszeroX/zkeen-ip) to generate rule lists each week.

## Features

- **GeoSite Support**: Convert domain-based DAT files to rule lists
- **GeoIP Support**: Convert IP CIDR-based DAT files to rule lists with `_ips` postfix
- **Multiple File Processing**: Process multiple DAT files in a single command
- **URL Support**: Download and process DAT files from URLs
- **Flexible Output**: Customizable output directory
- **Auto-formatting**: Proper rule formatting for different proxy clients

## Automated Releases

This repository includes a GitHub Action (`.github/workflows/create-release.yml`) that automatically fetches and processes DAT files [zkeen-domains](https://github.com/jameszeroX/zkeen-domains) and [zkeen-ip](https://github.com/jameszeroX/zkeen-ip) to generate rule lists each week.

### How it works

The GitHub Action:
1. **Runs weekly**: Automatically executes every Sunday at 00:00 UTC
2. **Manual trigger**: Can also be triggered manually via GitHub's workflow dispatch
3. **Fetches latest data**: Downloads the latest DAT files from:
   - [zkeen-domains](https://github.com/jameszeroX/zkeen-domains) for GeoSite data
   - [zkeen-ip](https://github.com/jameszeroX/zkeen-ip) for GeoIP data
4. **Generates rule lists**: Processes both files and creates organized rule lists
5. **Creates releases**: Automatically creates GitHub releases with all generated `.list` files

### Using the automated releases

You can download the latest generated rule lists from the [Releases page](../../releases). Each release contains:
- Domain rule files (e.g., `google.list`, `cn.list`)
- IP rule files with `_ips` suffix (e.g., `google_ips.list`, `cn_ips.list`)



## Usage

### Single File Processing

#### Process a GeoSite DAT file (domains)
```bash
go run main.go -file geosite.dat -type geosite
```

#### Process a GeoIP DAT file (IP ranges)
```bash
go run main.go -file zkeenip.dat -type geoip
```

#### Process from URL
```bash
go run main.go -url https://example.com/geosite.dat -type geosite
go run main.go -url https://example.com/geoip.dat -type geoip
```

### Multiple File Processing

#### Process both GeoSite and GeoIP files
```bash
go run main.go -geosite geosite.dat -geoip zkeenip.dat
```

#### Process multiple GeoSite files
```bash
go run main.go -geosites "file1.dat,file2.dat,file3.dat"
```

#### Process multiple GeoIP files
```bash
go run main.go -geoips "ip1.dat,ip2.dat,ip3.dat"
```

#### Process multiple URLs
```bash
go run main.go -geosite-urls "https://example.com/site1.dat,https://example.com/site2.dat"
go run main.go -geoip-urls "https://example.com/ip1.dat,https://example.com/ip2.dat"
```

#### Combined processing
```bash
go run main.go \
  -geosites "local_sites.dat" \
  -geoips "local_ips.dat" \
  -geosite-urls "https://example.com/remote_sites.dat" \
  -geoip-urls "https://example.com/remote_ips.dat" \
  -out combined_rules
```

### Command Line Options

| Flag | Description | Example |
|------|-------------|---------|
| `-file` | Path to local .dat file | `-file geosite.dat` |
| `-url` | URL to fetch .dat file from | `-url https://example.com/file.dat` |
| `-type` | Type of .dat file: 'geosite' or 'geoip' | `-type geoip` |
| `-out` | Output directory for files | `-out custom_output` |
| `-geosite` | Path to single GeoSite .dat file | `-geosite sites.dat` |
| `-geoip` | Path to single GeoIP .dat file | `-geoip ips.dat` |
| `-geosites` | Comma-separated paths to GeoSite files | `-geosites "file1.dat,file2.dat"` |
| `-geoips` | Comma-separated paths to GeoIP files | `-geoips "ip1.dat,ip2.dat"` |
| `-geosite-urls` | Comma-separated URLs to GeoSite files | `-geosite-urls "url1,url2"` |
| `-geoip-urls` | Comma-separated URLs to GeoIP files | `-geoip-urls "url1,url2"` |

## Output Format

### GeoSite Files (Domains)
Output files are named using the country code from the DAT file:
- `cn.list`
- `us.list`
- `google.list`

Content format:
```
DOMAIN-SUFFIX,example.com
DOMAIN-KEYWORD,google
DOMAIN,exact.domain.com
DOMAIN-REGEX,.*\.example\.com
```

### GeoIP Files (IP Ranges)
Output files are named with `_ips` postfix:
- `cn_ips.list`
- `us_ips.list`
- `google_ips.list`

Content format:
```
IP-CIDR,8.8.8.0/24
IP-CIDR,1.1.1.0/24
IP-CIDR,192.168.0.0/16
```



## License

This project is open source and available under the [MIT License](LICENSE).


