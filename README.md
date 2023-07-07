# EndExt
EndExt is a .go tool for extracting all the possible endpoints from the JS files

# Idea
When you crawll all the JS files from waybackruls for example, or even collecting the JS files urls from your target website's home source page ..
If the website was using API system and you wanna look for all the endpoints in the JS files, cuz u may find something hidden here or there ..
That's why i made this tool .. I give it the JS files urls .. It graps all the possible endpoints or urls or paths in the submitted JS files for me ..

# Installation
Just need to install go, run:

```
▶ brew install go
▶ git clone https://github.com/SirBugs/endext.git
```

or download from https://go.dev/dl/

# Usage:
```
▶ go run main.go -l js_files_urls.txt


                  ______          ________     __ 	
                 / ____/___  ____/ / ____/  __/ /_	
                / __/ / __ \/ __  / __/ | |/_/ __/	
               / /___/ / / / /_/ / /____>  </ /_  	
              /_____/_/ /_/\__,_/_____/_/|_|\__/  	

            ( * ) EndpointsExtractor Tool By @SirBugs .go Version
            ( * ) For Extracting all possilbe endpoints from Js files 
            ( * ) Version: 1.0.5 (Updated 3.Vrs on 7/7/2023)
            ( * ) Contact: Twitter@SirBagoza, GitHub@SirBugs, Medium@bag0zathev2
            ( * ) Command: go run main.go -l jsurls.txt

            ( ! ) You can use only -u for single URL or -l for .JS file URLs, Not both
            ( ! ) This tool has been received the last 3 updates at once

 ( 1 ) - https://example.com/_home/chunks/preload-helper-xxxxxxxx.js :: (endpoint) _app/
 ( 2 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/dashboard-metadata/bulk
 ( 3 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/applications
 ( 4 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/accounts
 ( 5 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) sign-in
 ( 6 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/sign-out
 ( 7 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/organization/details
 ( 8 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/organization/update
 ( 9 ) - https://example.com/_home/chunks/organization-xxxxxxxx.js :: (endpoint) endpoints/organization/subscribe
 ( 10 ) - https://example.com/_home/chunks/esr-apps-xxxxxxxx.js :: (endpoint) endpoints/express-security-review/application
 ( 11 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) applications
 ( 12 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) applications/new
 ( 13 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) settings
 ( 14 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) integrations
 ( 15 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) grants
 ( 16 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) applications/
 ( 17 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) accounts
 ( 18 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) webhooks
 ( 19 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) quickstart-guides
 ( 20 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) connectivity-api-offering
 ( 21 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) plans
 ( 22 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) users
 ( 23 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) billing
 ( 24 ) - https://example.com/_home/pages/__layout.svelte-xxxxxxxx.js :: (endpoint) experiments/

```

# One Line Command:

```
▶ echo 'target.com' | waybackurls | tee waybackresults.txt; cat waybackresults.txt | grep "\.js" > js_files.txt; go run main.go -l js_files.txt
```

// You can use Gau, HaKrawler, Katana, etc...

# Updates:
- Published (1.0.2)
- removing duplicates (1.0.3)
- RegexGrep with this.fetch(this.url("X") && short the urls filtering functionality (1.0.4)
- flag for single url -u or urls list -l && flag for public the urls -p && flag for output -o (1.0.5)


# Credits:

This tool was written in Golang 1.19.4, Made with all love in Egypt! <3

Twitter@SirBagoza , Github@SirBugs
