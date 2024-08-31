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
- alpinejs (as a file under static/js/alpine.min.js)

## usage
- install required node modules
```bun install```

- start the vite dev server 
```bun run dev```

- start file watching
```make watch```

