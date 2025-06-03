// Go WASM Service Worker
//

// Initialize the service worker
self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open("static-v1").then((cache) => {
      return cache.addAll([
        "/",
        "/index.html",
        "/static/css/tailwind.css",
        "/static/js/app.js",
        "/static/icons/icon-192-fs8.png",
      ]);
    }),
  );
});

// Fetch resources from the cache or network
self.addEventListener("fetch", (event) => {
  event.respondWith(
    caches.match(event.request).then((response) => {
      return response || fetch(event.request);
    }),
  );
});
