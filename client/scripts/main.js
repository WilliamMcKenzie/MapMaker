const SIZE = 8
const app = new PIXI.Application();
app.stage.scale = 8;

let textures = {}

async function load_asset(path) {
    const texture = await PIXI.Assets.load("/assets/" + path)
    texture.source.scaleMode = "nearest"
    texture.source.addressMode = "repeat"
    return texture
}

async function init() {
    await app.init({
        background: "#1f1f1f",
        width: window.innerWidth,
        height: window.innerHeight,
        useContextAlpha: false,
        antialias: true,
        autoDensity: true,
        resolution: 1,
    })
    document.body.appendChild(app.canvas)

    app.stage.scale = 4
    app.canvas.style.imageRendering = "pixelated";
    app.canvas.style.imageRendering = "crisp-edges";
    
    await init_tiles()
    await init_objects()
    init_draw()
}
