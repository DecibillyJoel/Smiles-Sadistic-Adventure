#!/usr/bin/env python3
"""Re-apply the site patches to the committed Defold HTML5 export.

Defold rewrites index.html, dmloader.js and the engine glue files on every
export. Those files live in this repository, so each new export drops the
changes below. Run this script in the repository root after you copy a new
export into the tree:

    python3 tools/apply-export-patches.py

The script is idempotent. It prints one line per patch. It exits with 1 when a
patch cannot be applied, because a changed Defold template must fail loudly.
preview.html is hand-written and is not touched by an export, so it is not
patched here.
"""

from pathlib import Path
import sys

ROOT = Path(__file__).resolve().parents[1]

EMBED_DETECT = """\t<script type='text/javascript'>
\tif (window.self !== window.top) {
\t\tdocument.documentElement.className += ' embedded';
\t}
\t</script>
"""

THEME = """\tbody { background-color: #080b09; }
\t.canvas-app-container { background: #080b09; }
\t.buttons-background { background-color: #080b09; }
\t.canvas-app-progress { background-color: #080b09; }
\t.canvas-app-progress-bar { background-color: #39ff14; }
\t.button { color: #39ff14; }
\t.link { color: #5d8a63; }
\ta { color: #39ff14; }
\thtml.embedded .button { display: none; }
"""

PATCHES = [
    {
        "path": "index.html",
        "name": "detect an embedded frame",
        "old": "\t<style type='text/css'>",
        "new": EMBED_DETECT + "\t<style type='text/css'>",
        "marker": "document.documentElement.className += ' embedded';",
    },
    {
        "path": "index.html",
        "name": "dark theme that matches the portfolio",
        "old": "\t</style>\n</head>",
        "new": THEME + "\t</style>\n</head>",
        "marker": "html.embedded .button { display: none; }",
    },
    {
        "path": "index.html",
        "name": "canvas reached with Tab or a click",
        "old": 'tabindex="1"',
        "new": 'tabindex="0"',
        "marker": 'tabindex="0"',
    },
    {
        "path": "index.html",
        "name": "green fullscreen icon",
        "old": "fill='%23006fff'",
        "new": "fill='%2339ff14'",
        "marker": "fill='%2339ff14'",
    },
    {
        "path": "dmloader.js",
        "name": "no focus grab at load",
        "old": "        canvas.focus();\n        CanvasInput.onFocusIn();\n",
        "new": "        // Do not grab focus at load. Focus arrives on a click or a Tab.\n",
        "marker": "Do not grab focus at load",
    },
]

for engine in ("SmilesSadisticAdventure_wasm.js", "SmilesSadisticAdventure_asmjs.js"):
    PATCHES.append(
        {
            "path": engine,
            "name": "Tab is not blocked",
            "old": "else if(event.keyCode===8||event.keyCode===9||event.keyCode===13){event.preventDefault()}",
            "new": "else if(event.keyCode===8||event.keyCode===13){event.preventDefault()}",
            "marker": "event.keyCode===8||event.keyCode===13",
        }
    )


def main():
    failures = 0
    for patch in PATCHES:
        path = ROOT / patch["path"]
        if not path.exists():
            print(f"FAIL  {patch['path']}: file is missing")
            failures += 1
            continue
        text = path.read_text(encoding="utf-8")
        if patch["marker"] in text:
            print(f"skip  {patch['path']}: {patch['name']}")
            continue
        count = text.count(patch["old"])
        if count != 1:
            print(f"FAIL  {patch['path']}: {patch['name']} — found {count} matches, expected 1")
            failures += 1
            continue
        path.write_text(text.replace(patch["old"], patch["new"]), encoding="utf-8")
        print(f"apply {patch['path']}: {patch['name']}")
    if failures:
        print(f"\n{failures} patch(es) failed. A Defold template probably changed.")
        return 1
    print("\nAll patches are in place.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
