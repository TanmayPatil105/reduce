## Command line URL shortener

# How to test

``` bash
$ make 
```

## Installation
```bash
$ go install github.com/TanmayPatil105/reduce@latest
```

## Usage
- `-c` to copy reduced url to clipboard
- `-b` to copy long url from clipboard
- `-q` to generate QR code for reduced url
- `-v` to show version
- `-d` to disable ascii art
- `-h` to display help message


# Example

```bash
$ reduce -v
reduced version 0.0.1
```
```bash
$ reduce

                    __                        __     
   _____ ___   ____/ /__  __ _____ ___   ____/ /     
  / ___// _ \ / __  // / / // ___// _ \ / __  /    
 / /   /  __// /_/ // /_/ // /__ /  __// /_/ /     
/_/    \___/ \__ _/ \__ _/ \___/ \___/ \__ _/  


Enter the URL : https://github.com/TanmayPatil105/reduce                  

╭ URL ─────────────────────────────╮
│                                  │
│                                  │
│     https://reduced.to/lcnw9     │
│                                  │
│                                  │
╰──────────────────────────────────╯
```

```bash
$ reduce -b

                    __                        __     
   _____ ___   ____/ /__  __ _____ ___   ____/ /     
  / ___// _ \ / __  // / / // ___// _ \ / __  /    
 / /   /  __// /_/ // /_/ // /__ /  __// /_/ /     
/_/    \___/ \__ _/ \__ _/ \___/ \___/ \__ _/  


╭ URL ─────────────────────────────╮
│                                  │
│                                  │
│     https://reduced.to/8hqeb     │
│                                  │
│                                  │
╰──────────────────────────────────╯
```
