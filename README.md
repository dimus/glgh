# glgh

glgh moves issues from GitLab repository to GitHub repository.

## Usage

### Creating a mirror of a project

Clone the repo from GitLab using the `--mirror` option. This is like
`--bare` but also copies all refs as-is. Useful for a full backup/move.
git clone --mirror git@your-gitlab-site.com:username/repo.git

Change into newly created repo directory

    cd repo

Push to GitHub using the `--mirror` option.  The `--no-verify` option skips any hooks.

    git push --no-verify --mirror git@github.com:username/repo.git

Set push URL to the mirror location

    git remote set-url --push origin git@github.com:username/repo.git

To periodically update the repo on GitHub with what you have in GitLab

    git fetch -p origin
    git push --no-verify --mirror

After doing this, the autolinking of issues, commits, and branches will work.

### Setting up authentication tokens

You will need personal OpenAuth tokens for both GitLab and GitHub.

For GitLab token go to `https://gitlab.com/profile/personal_access_tokens`.
Create a new Access Token with `api` and `read_repository` scopes.

For GitHub token go to `https://github.com/settings/tokens`.
Generate a new token with `repo`, `admin:org`, `user` scope.

Compile and run `glgh -V` first time to generate configuration file at
`~/.config/glgh.yaml`, edit this file according to your needs.

## Running the program

Run `glgh` to read all the tickets from GitLab and cache them on disk.
If you do want to update the cache run `glgh -r`.

This commands will read issues from GitLab repo and recreate them again
at GitHub.
