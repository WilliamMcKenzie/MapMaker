let holding = false

let visited = new Set()

let moving = false
let last_mouse = { x: 0, y: 0 }

function init_draw() {
    app.canvas.addEventListener("mousedown", (e) => {
        holding = true
        visited = new Set()
        last_mouse = { x: e.clientX, y: e.clientY }
    })
    app.canvas.addEventListener("mousemove", (e) => {
        if (holding && selected_tool == 'move') {
            const dx = e.clientX - last_mouse.x
            const dy = e.clientY - last_mouse.y
            app.stage.position.x += dx
            app.stage.position.y += dy
            last_mouse = { x: e.clientX, y: e.clientY }
        } else if (holding) {
            const [x,y] = localize(e.offsetX, e.offsetY)
            
            const key = `${x},${y}`

            if (!visited.has(key)) {
                visited.add(key)
                
                if (object_data[selected_tile]) {
                    add_object(x,y,selected_tile)
                } else {
                    const size = brush_size(selected_tool)
                    for (let dx = 0; dx < size; dx++) {
                        for (let dy = 0; dy < size; dy++) {
                            add_tile(x + dx, y + dy, selected_tile)
                        }
                    }
                    render_tiles()
                }
            }
        }
    })
    document.addEventListener("mouseup", (e) => {
        holding = false
    })
    
    function scroll(event) {
        const delta = event.deltaY < 0 ? 1 : -1
        const oldScale = app.stage.scale.x
        let newScale = oldScale + delta
        newScale = Math.min(64, newScale)
        newScale = Math.max(0.7, newScale)

        const centerX = app.renderer.width / 2
        const centerY = app.renderer.height / 2
        const worldCenter = {
            x: (centerX - app.stage.position.x) / oldScale,
            y: (centerY - app.stage.position.y) / oldScale
        }
        app.stage.scale.set(newScale)
        app.stage.position.set(
            centerX - worldCenter.x * newScale,
            centerY - worldCenter.y * newScale,
        )
    }
    window.addEventListener("wheel", scroll)
}

function brush_size(tool) {
    if (tool === "brush_1") return 1
    if (tool === "brush_2") return 2
    if (tool === "brush_3") return 3
    return 1
}

function localize(_x, _y) {
    let x = _x / app.stage.scale._x
    x -= app.stage.x / app.stage.scale._x
    x /= SIZE
    x = Math.floor(x)

    let y = _y / app.stage.scale._y
    y -= app.stage.y / app.stage.scale._y
    y /= SIZE
    y = Math.floor(y)

    return [x,y]
}
