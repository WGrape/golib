# How to contribute

## 1、Steps

### (1) fork project
First fork this project and switch to your [feat branch](./BRANCH_CONVENTION.md).

### (2) create package
You can create a new package directly when you have a great idea, or you would skip this step if it exists already.

### (3) write your idea
Write your great idea in a function, such as calculating the distance between two cities.

### (4) test your idea
Create a unit test to make sure your function running safely.

### (5) commit and push
After the above steps is all completed, you can [commit](./COMMIT_CONVENTION.md) the code and make a [PR](./PULL_REQUEST_TEMPLATE.md) to the repository.

## 2、Workflow

### (1) Git Hooks
If you need to use git hooks, please run the following commands.

```bash
chmod +x .github/hooks/prepare-commit-msg
cp .github/hooks/prepare-commit-msg ./.git/hooks
```

### (2) CI/CD
The CI/CD configuration was written in [build.yml](workflows/build.yml), and please click [Github Actions](https://github.com/WGrape/golib/actions) if you need search history builds.
