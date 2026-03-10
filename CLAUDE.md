# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

Personal learning notes repository — markdown files, code samples, and summaries from books, Udemy courses, YouTube videos, AWS/Azure certifications, C++ conference talks, and other technical resources.

## Repository Structure

- **Top-level `.md` files** (`AWS_Services.md`, `terms.md`, `databases_system_concepts.md`, `hashicorp_stuff.md`): Standalone reference notes on specific topics.
- **`README.md`**: Central hub with miscellaneous notes (WebSockets, Kubernetes, regex, markdown tips, etc.).
- **Course/book directories** (e.g., `udemy_mongodb_complete/`, `c++ Weekly/`, `Effective CPP/`): Each contains a `README.md` or topic-named `.md` files with notes, plus occasionally code samples.
- **`CPP Confrences/`**: Notes organized by conference (CppCon, CppNow, CoreC++) and year.

## Conventions

- **Spell checking**: `cspell.json` files exist both at the root and in subdirectories. Add technical terms to the `ignoreWords` array in the nearest `cspell.json`. Dictionaries in use: `softwareTerms`, `latex`.
- **Markdown style**: `markdown-style.css` at the root is referenced by some files via `<link rel="stylesheet">` for local rendering.
- **Code fences**: Use language identifiers (`cpp`, `python`, `sh`, `bash`, `golang`, `js`, `x86asm`).
- **Collapsible sections**: Long topic sections use `<details>`/`<summary>` tags. In-progress sections use `<!-- <details> -->` with `//TODO: add Summary` as a placeholder.
- **Inline C++ identifiers**: Use `<cpp>...</cpp>` for types, keywords, and standard library names inline in prose.
- **Line breaks in prose**: Use `\` at the end of a line for a soft break within a paragraph.
- **Emphasis**: `_italic_` for light emphasis, `**bold**` for strong.
- **Progress tracking**: Use `- [x]` / `- [ ]` checklists at the top of a series file to track which episodes/chapters are done.
- **Direct quotes**: Use `>` blockquotes for verbatim quotes from talks, books, or slides.
- **LaTeX**: Math blocks use `$$...$$` with `\begin{align*}`.

## Adding New Notes

- Place course/book notes in a dedicated subdirectory with a `README.md` or descriptively named `.md` file.
- Add a `cspell.json` to the subdirectory if the content has domain-specific terms not in the root ignore list.
- Link new major topics from the root `README.md` if they warrant discovery from the top level.
