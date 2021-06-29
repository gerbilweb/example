package src

import "github.com/gerbilweb/gerbil"

var _ gerbil.Component = (*App)(nil)

type App struct{}

func (component *App) Render() string {
	return `
	<header></header>
	<div class="app">
		<div class="heading">
			<h1 class="title">gerbil.go</h1>
			<h3 class="sub-title">A wasm framework in GO</h3>
		</div>

		<img class="icon" src="media/icon.png" alt="icon" />
		<hr />
		
		<p>
			Edit <code>app.go</code> to change html markup and <code>layout.go</code> to change component layout
		</p>
	</div>
	<footer></footer>
	`
}
