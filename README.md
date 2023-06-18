![Worktrees Logo](/docs/images/worktrees.png?raw=true)

# Worktrees

Introducing Worktrees - the CLI application that automates the process of setting up yiur dev environment for a new git worktree. With Worktrees, developers can create new worktrees with ease, while saving time and increasing efficiency.

Worktrees allows you to create a `.wt` file in your project, where you can define custom commands to initialize and terminate the environment for your new worktree. Once you have set up your `.wt` file, you can create new worktrees with `wt add {worktree_name} {branch}` and remove them with`wt rm {worktree_name}`. Worktrees is a wrapper around git worktrees, so it seamlessly integrates into your existing workflow.

Worktrees is also useful when you're working on your own branch and someone asks you to review their pull request. With Worktrees, it's easy to create a new environment on the PR's branch for testing without losing your work, or having to stash your changes and move to the PR branch.

Ideal for developers who need to work on multiple environments for each git branch, Worktrees ensures that each environment is truly isolated. With Worktrees, you no longer have to manually set up the environment every time you create a new worktree, saving you valuable time.

## Installation

Setting up Worktrees is easy - simply download the CLI's executable from the releases page, rename it to `wt`, and move it where the your system can find it globally, and you're ready to go.

Worktrees is an open source product, available for free. If you find Worktrees useful, you can support the project by paying what you want.

## Setup
On your local repository, create a new `.wt` like the next example:

```yaml
initCommands:
  - cp $MAIN_WORKTREE_PATH/.env $WORKTREE_PATH/.env
  - mysqlsh root@localhost:33068 --sql -e "CREATE DATABASE $(echo $WORKTREE_NAME);" &&
    sed -i '' "s/^DB_DATABASE=.*/DB_DATABASE=$(echo $WORKTREE_NAME)/" $(echo $WORKTREE_PATH)/.env
  - composer install
  - php artisan migrate --seed
  - npm ci
  - valet link $(echo $WORKTREE_NAME) --secure
  - code .

terminateCommands:
  - valet unlink $(echo $WORKTREE_NAME)
  - mysqlsh root@localhost:33068 --sql -e "DROP DATABASE $(echo $WORKTREE_NAME);"
```

On the condfiguration file you have the next env variables available:

- `WORKTREE_BRANCH`: This is the name of the worktree branch.
- `MAIN_WORKTREE_PATH`: This is the git worktree path of your main repository.
- `WORKTREE_PATH`: This is the path new worktree when you are creating a new one. When you are removing the worktree, it is the path of the worktree related to the branch you selected to remove.
- `WORKTREE_NAME`: This is the name of the git worktree.

### Add command

Run `wt add {worktree_name} {branch} --path {path}` on the root project to add a new worktree and initialize the environment.

### Remove command

Run `wt rm {worktree_name}` from the main project or one of the worktrees to remove the worktree and terminate the environment.
