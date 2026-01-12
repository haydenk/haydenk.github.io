+++
title = 'Hugo Archtypes are useless'
slug = 'hugo-archtypes-are-useless'
date = 2026-01-11 19:37:19
draft = false
tags = ['hugo', 'blogging', 'development']
+++

Okay, they're not _totally_ useless but they might as well be. I opted to use a shell script in a mise task to manually create the file
with the information I needed.

Specifically, what I figured was an obvious thing was to run the command something like, `hugo new content/posts/1767247204-sample.md --tile "Sample Post" --date "2026-01-01 00:00:04"`, otherise Hugo will insert a titlized version of the filename and the current date for the new post. This probably seems like niche issue but it's important to me
and it's such an obvious option if you needed to back post something or perhaps you wrote it out on paper on one day and wanted it posted as that day not the current day the markdown was created.


```toml
+++
title = '{{  getenv "POST_TITLE" | default (replace .File.ContentBaseName "-" " " | title) }}'
slug = '{{ ..File.ContentBaseName }}'
date = {{ getenv "POST_DATE" | .Date }}
draft = false
+++
```

Then you would create your post with: `POST_TITLE="Sample Post" POST_DATE="2026-01-01 00:00:04" hugo new content/posts/1767247204-sample.md`

The only problem is that you cannot just pass environment variables to the hugo template without updating the securit config. 

```toml
[security]
  [security.funcs]
    getenv = ['^HUGO_', '^CI$', '^POST_']
```

I know it's a small thing and perhaps "not that big of a deal" but if you're that willing to relent on security then what else are you willing to relent with and it just seems hacky and not every elegant.