## GOAT  
![Banner](thumb.png)
this projects is a starting point for developing static web apps using a-h/templ

### Tech Stack
- Go ( Templ for logic behind templating) 
- Alpinejs (javascript framework)
- Tailwindcss (css framework)
- vite (optional) for hot reloading


## dependencies
- alpinejs (as a file under /public/static/js/alpine.min.js)
- bun (or yarn or pnpm) used to install/run vite and tailwindcss
- go 
- templ cmd (using go install github.com/a-h/templ/cmd/templ@latest)
- make 
- entr


## usage
- clone the repo
```git clone https://github.com/morethancoder/goat.git```

- initalize the project
```go mod init yourprojectname && go mod tidy```

- change the import in main.go 
``` "morethancoder/goat/components" -> "yourprojectname/components" ```

- install required node modules
```bun install```

- start the vite dev server 
```make vite```

- start file watching & compiling
```make watch```

- star the repo
