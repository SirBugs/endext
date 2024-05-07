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

Notice that you have to download the `regex.tmp` file and put into it your custom regex depending on the JS File you facing!

# Default Regex Patterns:
```
"\?(.*?)"
"\/(.*?)"
'\/(.*?)'
`\/(.*?)`
this\.fetch\(this\.url\("([^"]+)"\)
```

# Usage:
```
▶ go run main.go -l js_files_urls.txt

[ 0 ] https://dashboard.target.com/assets/index-8d4703d5.js : "?o.credentials="
[ 1 ] https://dashboard.target.com/assets/index-8d4703d5.js : "?id="
[ 2 ] https://dashboard.target.com/assets/index-8d4703d5.js : "?ddforward="
[ 3 ] https://dashboard.target.com/assets/index-8d4703d5.js : "?i="
[ 4 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/api/v2/"
[ 5 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/target-0b66e97b.png"
[ 6 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/welcome-f0600742.png"
[ 7 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/welcome@2x-c3e5437e.png"
[ 8 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/screenshot-063bd7f1.png"
[ 9 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/screenshot@2x-0c6069f9.png"
[ 10 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/ACHAuthorizationTemplate-b7991176.pdf"
[ 11 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/CommercialDepositAccountAgreement_March2023-34034815.pdf"
[ 12 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/target-Prohibited-Categories-7a050116.pdf"
[ 13 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/csv-a3fb6a60.svg"
[ 14 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/gif-2e1908c8.svg"
[ 15 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/html-2dcba250.svg"
[ 16 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/jpeg-2ad1d5df.svg"
[ 17 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/jpg-4e59c322.svg"
[ 18 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/json-3f816b3b.svg"
[ 19 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/pdf-e4bef26d.svg"
[ 20 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/png-478fed1d.svg"
[ 21 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/tiff-c76a207d.svg"
[ 22 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/txt-803993e5.svg"
[ 23 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/unknown-6debfec0.svg"
[ 24 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/zip-e303efb1.svg"
[ 25 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/assets/pattern-5243e697.svg"
[ 26 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/app"
[ 27 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/api-keys"
[ 28 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/bank-accounts"
[ 29 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/counterparties"
[ 30 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/entities/person"
[ 31 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/entities/business"
[ 32 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/entities"
[ 33 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/events"
[ 34 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/events/webhook"
[ 35 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/invites"
[ 36 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/invites/resend_email"
[ 37 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/plaid/auth/oauth2/temp-code"
[ 38 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/simulate/receive-ach-credit"
[ 39 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/simulate/receive-ach-debit"
[ 40 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/simulate/transfers/ach/settle"
[ 41 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/simulate/receive-wire"
[ 42 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/simulate/receive-international-wire"
[ 43 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/simulate/transfers/wire/settle"
[ 44 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers/ach"
[ 45 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers/ach/stop-payments"
[ 46 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers/ach/returns"
[ 47 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers/book"
[ 48 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers/checks/issue"
[ 49 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers"
[ 50 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/transfers/wire"
[ 51 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/dashboard-users"
[ 52 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/password_reset_request"
[ 53 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/login"
[ 54 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/dashboard-users/settings"
[ 55 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/mfa/verify"
[ 56 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/logout"
[ 57 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/self"
[ 58 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/dashboard-users/password"
[ 59 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/mfa/setup"
[ 60 ] https://dashboard.target.com/assets/index-8d4703d5.js : "/dashboard/mfa/resend"
```

# One Line Command:
```
▶ echo 'target.com' | waybackurls | grep "\.js" > js_files.txt; go run main.go -l js_files.txt
```

// You can use Gau, HaKrawler, Katana, etc...

# Options:
```
  -l string
    	.txt File containing JavaScript file URLs
  -o string
    	Output To Save Endpoints
  -s	Silence Bitch
  -u string
    	Signle JavaScript File Direct URL
```

# Updates:
- (1.0.2) :: Published 
- (1.0.3) :: removing duplicates
- (1.0.4) :: RegexGrep with this.fetch(this.url("X") && short the urls filtering functionality
- (1.0.5) :: flag for single url -u or urls list -l && flag for public the urls -p && flag for output -o
- (1.0.6) :: revampig the whole structure, adding silence flag and making it faster


# Credits:

This tool was written in Golang, With all love in Egypt! <3

Twitter@SirBagoza , Github@SirBugs
