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
▶ go run main.go urls.txt


               /$$$$$$$$                 /$$ /$$$$$$$$             /$$
              | $$_____/                | $$| $$_____/            | $$
              | $$       /$$$$$$$   /$$$$$$$| $$       /$$   /$$ /$$$$$$
              | $$$$$   | $$__  $$ /$$__  $$| $$$$$   |  $$ /$$/|_  $$_/
              | $$__/   | $$  \ $$| $$  | $$| $$__/    \  $$$$/   | $$
              | $$      | $$  | $$| $$  | $$| $$        >$$  $$   | $$ /$$
              | $$$$$$$$| $$  | $$|  $$$$$$$| $$$$$$$$ /$$/\  $$  |  $$$$/
              |________/|__/  |__/ \_______/|________/|__/  \__/   \___/

                       JsValidator Tool By @SirBugs .go Version
                             V: 1.0.1 Made With All Love
                  For Extracting all possilbe endpoints from Js files
                         Twitter@SirBagoza -- GitHub@SirBugs
                           Run : go run main.go jsurls.txt

endpoints/users/password
sign-in
endpoints/sign-out
endpoints/billing/update-billing-info
endpoints/billing/get-account
endpoints/billing/create-account
endpoints/billing/list-subscriptions
endpoints/billing/create-new-subscription-purchase
endpoints/billing/create-one-time-payment
endpoints/billing/get-account
endpoints/billing/create-account
endpoints/billing/list-subscriptions
endpoints/billing/create-new-subscription-purchase
endpoints/billing/create-one-time-payment

```

# One Line Command:

```
▶ echo 'target.com' | waybackurls | tee waybackresults.txt; cat waybackresults.txt | grep "\.js" > js_files.txt; go run main.go js_files.txt
```

// You can use Gau, HaKrawler, Katana, etc...

# Credits
This tool was written in Golang 1.19.4, Made with all love in Egypt! <3

Twitter@SirBagoza , Github@SirBugs
