---
layout: default
title: At The Movies
permalink: /at-the-movies
---

{% for tag in site.tags %}
  {% assign t = tag | first %}
  {% assign posts = tag | last %}
{% endfor %}

# Posts

<ul>
{% for post in posts %}
  {% if post.tags contains "atthemovies" %}
  <li>
    <a href="{{ post.url }}">{{ post.title }}</a>
    <span class="date">{{ post.date | date: "%B %-d, %Y"  }}</span>
  </li>
  {% endif %}
{% endfor %}
</ul>

---

This is a four week series, not on this movie there will be others but if you are interested in checking it out,
check out [Lake Pointe Church - At The Movies](https://lakepointe.church/at-the-movies/){:target="_blank"}

<a href="https://lakepointe.church/at-the-movies/" target="_blank"><img src="/images/at_the_movies.jpg" width="400" alt="At The Movies: Lake Pointe Church"></a>