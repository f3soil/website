# f3soil.com Website

## Overview

This site uses [Hugo] to generate the HTML.
Key files:

- `.github/workflow/deploy.yaml`: Defines how the site gets built on GitHub.
- `config.toml`: Main config file for Hugo. Includes the definition of the Locations menu.
- `content/hype`: Content for https://f3soil.com/hype/
- `content/invite`: Content for https://f3soil.com/invite/
- `content/locations/<ao>`: Content for each AO

## Editing

In order to build the site locally you should install [Mise].
It will then handle installing Hugo and enable you to run tasks.

After editing content you will want to view the site locally by running,

```
mise tasks run serve
```

Then you would need to commit and push to GitHub.
GitHub will then automatically build the site; see [GitHub Actions](https://github.com/f3soil/website/actions/workflows/deploy.yml).

Key commands, more as a reference, are:

```
git status
```

```
git add -p
```

```
git commit -v
```

```
git push
```

[Hugo]: https://gohugo.io/
[Mise]: https://mise.jdx.dev/
