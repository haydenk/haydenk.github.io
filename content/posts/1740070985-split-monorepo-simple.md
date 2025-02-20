# A Simple Guide to Split a Monorepo into a Separate Repository

Assume we have monorepo at `git@github.com/haydenk/monorepo.git`

With directory contents:
```
packageA/
packageB/
packageC/
```

_I did say simple_

---

1. Create the new repository. `git@github.com/haydenk/packageA.git`
2. Clone the repository into an appropriate name for the new repository.
   * `git clone git@github.com/haydenk/monorepo.git packageA/`
3. Replace the origin remote with the new repository.
   * `git remote remove origin`
   * `git remote add origin git@github.com/haydenk/packageA.git`
4. Rewrite the repository to only the subdirectory or subdirectories you want in the new repository. With either master or develop branch or whatever branch you want to use.
   * `git filter-branch --prune-empty --subdirectory-filter packageA -- develop`
   * I chose develop because I will create a release branch to push this to master or you can simply merge into master.
5. Push the rewrites
   * `git push --all`
   * `git push --tags`
