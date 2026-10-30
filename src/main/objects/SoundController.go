components {
  id: "script"
  component: "/main/scripts/SoundController.script"
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
  id: "MusicLoop"
  type: "sound"
  data: "sound: \"/main/sounds/Music Loop.ogg\"\n"
  "looping: 1\n"
  "group: \"music\"\n"
  "gain: 0.9\n"
  "pan: 0.0\n"
  "speed: 1.0\n"
  ""
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
