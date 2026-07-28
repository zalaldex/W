# Repository Wipe Workflow

Adds a manually-triggered GitHub Action that deletes every tracked file in
the repository. No confirmation prompt, no backup — clicking "Run
workflow" wipes it immediately.

## Setup

1. Copy `.github/workflows/wipe-repo.yml` into your repository at that
   exact path.
2. Commit and push it to your default branch.
3. In your repo's Settings → Actions → General, under "Workflow
   permissions," make sure the `GITHUB_TOKEN` has **Read and write
   permissions** — the workflow needs write access to push the deletion
   commit.

## Running it

1. Go to the **Actions** tab → select **Wipe Repository** in the left
   sidebar.
2. Click **Run workflow** → **Run workflow** again to confirm the branch.

That's it — the run starts immediately and deletes every tracked file on
that branch.

## What it does

1. Checks out the repository.
2. Deletes every tracked file (`git rm -rf .`).
3. Commits and pushes the deletion.

## Notes

- This only removes files from the working tree via a commit — git history
  is not rewritten, so old file contents remain in the commit history and
  are recoverable by checking out a prior commit, unless history is later
  rewritten or the repo is deleted.
- It does not touch other branches, issues, PRs, releases, or repo
  settings — only tracked files on the branch you run it against.
- There is no confirmation step. Anyone with permission to run workflows
  in this repo can trigger it with two clicks.
