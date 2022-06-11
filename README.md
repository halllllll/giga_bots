# KANSHI
:eye:
## USAGE

discordのwebhookurlを想定してるよ

expected called by `cron`. so check right for privileges.

due to `cron`, its too bother for env vals.
anyway load `.bash_profile` in `botsCron.sh`. here is my case(lazy example).

```
# botsCron.sh
source /home/user/.bash_profile
cd /path/to/giga_bots
go run main.go >> /path/to/giga_bots/bots.log 2>&1

```

(of course) permissions of the `cron` user also be properly matched, like `/usr/local/go/bin/go` or other paths.


my `crontab` config is here(very straightforward, boring and simple)

```
# crontab
*/5 * * * 1-5  bash /path/to/giga_bots/botsCron.sh
```

## NOTE
caz using [chromedp](https://github.com/chromedp/chromedp) package, required `google-chrome`.