# Terminal Game Client

UI prototype for a Steam-like client.

![game-thing](https://cloud.githubusercontent.com/assets/1373315/13459640/260c03ca-e0c9-11e5-93f4-2fb9dfe0bb90.gif)

* Make a listbox with selectable items (single select only)
  * Doesn't account for _no_ selection or 0 items
  * Would be nice if items had an associated `interface{}` (tag?)
    * Maybe sticking to just the index is good enough, keep it out of the "view"
  * _Multiple_ list item decorations
    * Not downloaded vs. downloaded vs. downloading would be different colours

"Cloud sync" via symlinking game save dirs in to Dropbox

Cool download progress bars (gauge) or IO meter (sparkline)
