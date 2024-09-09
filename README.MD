## GOAT  
![Banner](thumb.png)
this projects is a starting point for developing static web apps using a-h/templ

### Tech Stack
- Go (for logic behind templates)
- Alpinejs (javascript framework)
- Templ (html template engine)
- Tailwindcss (css framework)
- vite (optional) for hot reloading


## dependencies
- bun (or yarn or pnpm) used to install vite
- go 
- templ (using go install github.com/a-h/templ/cmd/templ@latest)
- make 
- entr
- alpinejs (as a file under /public/static/js/alpine.min.js)



## usage
- clone the repo
```git clone https://github.com/morethancoder/goat.git```

- initalize the project
```go mod init yourprojectname```

- change the import in main.go 
``` "morethancoder/goat/components" -> "yourprojectname/components" ```

- install required node modules
```bun install```

- start the vite dev server 
```make vite```

- start file watching & compiling
```make watch```

