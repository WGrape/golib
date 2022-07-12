# How to contribute

## 1、Steps

### (1) create package
You can create a new package directly when you have a great idea, or you would skip this step if it exists already.

### (2) write your function
Write your great idea in a function, such as calculating the distance between two cities.

### (3) test your function
Create a unit test to make sure your function running safely.

## 2、Workflow

### (1) Git Hooks
If you need to use git hooks, please run the following commands.

```bash
chmod +x .github/hooks/prepare-commit-msg
cp .github/hooks/prepare-commit-msg ./.git/hooks
```

### (2) CI/CD
The CI/CD configuration was written in [build.yml](workflows/build.yml), and please click [Github Actions](https://github.com/WGrape/golib/actions) if you need search history builds.
