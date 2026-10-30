embedded_components {
  id: "Description"
  type: "label"
  data: "size {\n"
  "  x: 1800.0\n"
  "  y: 75.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "scale {\n"
  "  x: 1.0\n"
  "  y: 1.0\n"
  "  z: 1.0\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 1.0\n"
  "  y: 1.0\n"
  "  z: 1.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_N\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: true\n"
  "text: \"These guys shoot lasers, which deal 60 damage per second.\\n"
  "\"\n"
  "  \"These lasers have unlimited range, so pay attention to where these enemies are aiming!\"\n"
  "font: \"/main/fonts/MyFont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 960.0
    y: 550.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "MenuBlackout"
  type: "sprite"
  data: "tile_set: \"/main/atlases/Atlas.atlas\"\n"
  "default_animation: \"MenuBlackout\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 960.0
    y: 540.0
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
  id: "Name"
  type: "label"
  data: "size {\n"
  "  x: 485.0\n"
  "  y: 75.0\n"
  "  z: 0.0\n"
  "  w: 0.0\n"
  "}\n"
  "scale {\n"
  "  x: 1.2\n"
  "  y: 1.2\n"
  "  z: 1.2\n"
  "  w: 0.0\n"
  "}\n"
  "color {\n"
  "  x: 1.0\n"
  "  y: 1.0\n"
  "  z: 1.0\n"
  "  w: 1.0\n"
  "}\n"
  "outline {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "shadow {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "  w: 1.0\n"
  "}\n"
  "leading: 1.0\n"
  "tracking: 0.0\n"
  "pivot: PIVOT_N\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"Laser Smile Catchers\"\n"
  "font: \"/main/fonts/MyFont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 960.0
    y: 980.0
    z: 1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Laser"
  type: "sprite"
  data: "tile_set: \"/main/atlases/Atlas.atlas\"\n"
  "default_animation: \"EnemyLaserRight\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 910.0
    y: 700.0
    z: 0.4
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
  "default_animation: \"LaserBeam\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 2234.0
    y: 700.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
