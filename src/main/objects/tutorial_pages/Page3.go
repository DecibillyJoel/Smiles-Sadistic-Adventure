components {
  id: "script"
  component: "/main/scripts/MakeTheHealthGreen.script"
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
  "text: \"These things are, in order from left to right, your health, time and score.\\n"
  "\"\n"
  "  \"The orange bar is your dash indicator.\"\n"
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
  "text: \"This is your info bar.\"\n"
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
  id: "Health"
  type: "sprite"
  data: "tile_set: \"/main/atlases/Atlas.atlas\"\n"
  "default_animation: \"HealthBar\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 115.0
    y: 708.5
    z: 0.8
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Portrait"
  type: "sprite"
  data: "tile_set: \"/main/atlases/Atlas.atlas\"\n"
  "default_animation: \"HealthPortrait\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 360.0
    y: 700.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Score"
  type: "label"
  data: "size {\n"
  "  x: 485.0\n"
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
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"000000000\"\n"
  "font: \"/main/fonts/MyFont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 1550.0
    y: 700.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Time"
  type: "label"
  data: "size {\n"
  "  x: 485.0\n"
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
  "pivot: PIVOT_CENTER\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "line_break: false\n"
  "text: \"00:00.000\"\n"
  "font: \"/main/fonts/MyFont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    x: 960.0
    y: 700.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "Dash"
  type: "sprite"
  data: "tile_set: \"/main/atlases/Atlas.atlas\"\n"
  "default_animation: \"DashBar\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 115.0
    y: 673.5
    z: 0.8
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
