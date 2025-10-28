// import library
import Alpine from 'https://cdn.jsdelivr.net/npm/alpinejs@3.14.8/dist/module.esm.js';
import 'https://cdn.jsdelivr.net/npm/htmx.org@2.0.7/dist/htmx.esm.js';
import 'https://unpkg.com/cally';
// import modules
import site from '/assets/js/modules/site.js';

// Expose to global scope
window.site = site;
window.Alpine = Alpine;
// boot AlpineJS
Alpine.start();
