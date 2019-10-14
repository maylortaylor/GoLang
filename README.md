# MaylorTaylor's GoLang / GopherExercises

This is where all of my GoLang exercises and examples will be.

## Setting Up Windows Environment

- Visit [GoLang Download Page](https://golang.org/dl/)
- Download most recent version (as of October 2019 it is v1.13.1)
- Install the tools to the default place C:\Go -- [Instructions](https://golang.org/doc/install#install)
- WINDOWS - make sure PATH in 'Environment Variables' has bene updated. It should now have "C:\Go\bin" at the end of PATH
- Add new 'Environment Varibale'

  - Variable: GOPATH
  - Value: C:\insert path location you want to use GO in

- If you have Windows WSL and use ZSH prompt

  - Add "export GOPATH=\$HOME/go" to ~/.zshrc file

- If you have Windows WSL and use BAHS prompt
  - Add "export GOPATH=\$HOME/go" to ~/.bash_profile file
- Restart computer
