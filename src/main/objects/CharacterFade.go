components {
  id: "script"
  component: "/main/scripts/CharacterFadeController.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/atlases/Atlas.atlas\"\n"
  "default_animation: \"Character\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.4
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
