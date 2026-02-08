# Project Instructions

## Bible Notes Formatting Spec

### Source Format
- Exported from Bible Note app (https://biblenote.ai/)
- Comes as `.txt` files, typically two per sermon:
  - **Summary** file: full structured notes with headings, bullets, scripture refs
  - **Insight** file: numbered insight/reflection paragraphs
- Header lines to strip: `Notes Powered by Bible Note https://biblenote.ai/` and `========` separator lines with section labels like "Summary" / "Insight"

### Target Format: Hugo Markdown

#### Frontmatter
- TOML format (`+++`)
- Title only — no date, slug, tags, or draft fields
- Title is derived from the `#` (H1) heading in the summary file
```
+++
title = "Post Title Here"
+++
```

#### Content Structure (in order)
1. `## Scripture References` — each reference as a separate `*` bullet (source has them semicolon-separated on one line)
2. `## Introduction`
3. `## Key Points / Exposition` — with `###` numbered sub-headings
4. `## Major Lessons & Revelations`
5. `## Practical Application`
6. `## Conclusion & Call to Response`
7. `## Prayer`
8. `## References & Resources`
9. `## Insights` — numbered list merged from the separate Insight file

#### Formatting Rules
- All unordered bullets use `*` (not `-`)
- Sub-bullets indented 2 spaces: `  *`
- No trailing whitespace on lines
- No extra blank lines (single blank line between sections)
- Numbered lists (Insights) use `1.`, `2.`, etc.
- Special characters normalized:
  - Em-dashes (`—`) to `--`
  - Arrows (`→`) to `to`
  - Curly quotes to straight quotes
  - En-dashes in ranges (`–`) to `-`
  - Remove `%` space (`100 %` to `100%`)
  - Ellipsis (`…`) to `...`

#### File Output
- Combined into a single `.md` file (summary + insights merged)
- Saved to `sermon-notes/` directory with kebab-case slug filename
- Example: `sermon-notes/stewards-not-owners-faithful-with-gods-resources.md`
- Destination for Hugo: `content/bible-notes/` (user moves manually or specifies)

#### Reference Example
See `sermon-notes/stewards-not-owners-faithful-with-gods-resources.md` for the canonical formatted output.
