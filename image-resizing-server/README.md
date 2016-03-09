# Image Resizing Server

Request an image like
http://localhost:8080/from-net?url=http://i.imgur.com/wGVhLkj.jpg&percent=20
and get it back at 20% the size.

Use `/from-net-with-cache` for using a cache instead.
The cache is in memory, currently set at 20M.
It caches source images - it might be smart to cache output images too/instead.
