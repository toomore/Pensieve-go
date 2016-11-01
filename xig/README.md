xig
====
To fetch **instagram** user img, content, avatar data.

install
--------

    go install -v github.com/toomore/Pensieve-go/xig

Usage
------

To fetch recently img(12), avatar and content

    xig -name {username}

To fetch **ALL** images data (if user uploaded more, may slow)

    xig -name {username} -all

Fetch folder
------------

```
./{username}
├── avatar
│   ├── {path1}.jpg              // user avatar image
│   └── (...).jpg                // and more ... if put `xig` into cron jobs
├── content
│   ├── {date}_{code}_{id}.json  // json files, for some day `xig` reuse
│   └── {date}_{code}_{id}.txt   // for human readable content
└── img
    ├── {code}_{path1}.jpg       // user uploaded images
    └── (...).jpg                // and more ...
```

Note
-----

* All images will try to fetch original size.
* Private user not work.
* Content's readable date is in `RFC3339` format.
* instagram won't to ban ip, may CDN doesn't check.
* `xig`'s code base are not pretty, I will make it pretty :)

Tips
-----

For crontab, every 1m to fetch

    */1 * * * * cd ~/{some folder}; ({$go_bin_path}/xig -name {username} 2>&1) >> ./{username}.log
