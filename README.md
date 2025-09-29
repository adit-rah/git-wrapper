# git-wrapper

A lightweight wrapper around git that simplifies common workflows like creating feature branches, committing changes, submitting pull requests, and merging branches.

## Installation

Clone the repo and add the bin/ directory to your PATH:
```
git clone https://github.com/YOURNAME/git-wrapper.git
cd git-wrapper
echo "export PATH=$PATH:$(pwd)/bin" >> ~/.bashrc
source ~/.bashrc
```
Now you can run the wrapper anywhere as:

`gw`

## Usage

Create a new branch and commit
`gw create update my-feature`

Make additional commits
`gw modify`

Submit a pull request
`gw submit`

Merge a branch into its parent
`gw fold`

Roadmap

Add versioned releases
Provide Homebrew install support
Expand available git workflows
