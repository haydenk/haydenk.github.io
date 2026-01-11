+++
title = '{{ getenv "POST_TITLE" | default .File.ContentBaseName }}'
slug = '{{ getenv "POST_SLUG" | default .File.ContentBaseName }}'
date = {{ getenv "POST_DATE" | default .Date }}
draft = false
tags = []
+++
