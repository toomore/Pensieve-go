xig
====
To fetch **instagram** user img, content, avatar data.

!!!!
-----
Move to [**github.com/toomore/xig**](github.com/toomore/xig)

install
--------

    go get -v -a -u github.com/toomore/Pensieve-go/xig

Usage
------

To fetch recently img(12), avatar and content

    xig {username}

To fetch **ALL** images data (if user uploaded more, may slow)

    xig -a {username}

Print recently data

    xig -i {username}

```
+----------------------------------------------------+
Code: https://www.instagram.com/p/{code}
Date: {date} IsVideo: {true|false}
Caption: {caption}
DisplaySrc: {url}
```

To find some deleted content

    xig -f {username}

Some users turn to private account, using `-u` to login user account for fetch
private data. (required setting environment variables in `IGUSER`, `IGPASS`)

    xig -u {username}

Fetch folder
-------------

```
./{username}
├── profile
│   └── {username}_{hash}.txt    // user profile, biography
├── avatar
│   ├── {username}_{hash}.jpg    // user avatar image
│   └── (...).jpg                // and more ... if put `xig` into cron jobs
├── content
│   ├── {date}_{code}_{id}.json  // json files, for some day `xig` reuse
│   └── {date}_{code}_{id}.txt   // for human readable content
└── img
    ├── {code}_{hash}.jpg        // user uploaded images
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

    */1 * * * * cd ~/{some folder}; ({$go_bin_path}/xig {username} 2>&1) >> ./{username}.log
