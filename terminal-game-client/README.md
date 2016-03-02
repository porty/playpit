# Terminal Game Client

UI prototype for a Steam-like client.

* Make a listbox with selectable items (single select only)
  * Doesn't account for _no_ selection or 0 items
  * Would be nice if items had an associated `interface{}` (tag?)
    * Maybe sticking to just the index is good enough, keep it out of the "view"
  * _Multiple_ list item decorations
    * Not downloaded vs. downloaded vs. downloading would be different colours

"Cloud sync" via symlinking game save dirs in to Dropbox

Cool download progress bars (gauge) or IO meter (sparkline)
