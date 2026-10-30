# Smiles-Sadistic-Adventure

An endless semi-random bullet hell I made in Defold in high school.

## Play in the browser

https://DecibillyJoel.github.io/Smiles-Sadistic-Adventure/

## Watch the preview

https://github.com/user-attachments/assets/4de9e492-813b-44ee-ba36-9ff478500b66

[Watch in a browser page](https://DecibillyJoel.github.io/Smiles-Sadistic-Adventure/preview.html)

## Source code

The game source is in the [`src/`](src/) folder. It is a Defold project. Open `src/` with the Defold editor to build or change the game.

The source folder holds `game.project`, plus the `main/` and `input/` folders. The exported build output and the editor cache are not included.

## Download

Grab a build for your platform from the [Releases page](https://github.com/DecibillyJoel/Smiles-Sadistic-Adventure/releases).

## Rebuilding the export

Defold writes `index.html`, `dmloader.js` and the engine files on every export. Those files carry a few site patches: the dark theme, no focus grab at load, and Tab is not blocked. Copy the new export over the tree, then run the patch script.

    python3 tools/apply-export-patches.py

The script is idempotent. It exits with 1 when a Defold template changed.
